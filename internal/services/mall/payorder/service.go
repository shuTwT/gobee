package payorder

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/payorder"
	"github.com/shuTwT/hoshikuzu/ent/postpurchase"
	"github.com/shuTwT/hoshikuzu/internal/infra/pay/epay"
	setting_service "github.com/shuTwT/hoshikuzu/internal/services/system/setting"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
)

type PayOrderService interface {
	ListPayOrderPage(ctx context.Context, req *model.PageQuery) ([]*ent.PayOrder, int, error)
	SubmitPayOrder(ctx context.Context, userID int, req *model.PayOrderSubmitReq) (*model.PayOrderSubmitResp, error)
	HandleNotify(ctx context.Context, params map[string]string) error
	SyncOrderStatus(ctx context.Context, orderID int) (*model.PayOrderStatusResp, error)
	GetTodayStats(ctx context.Context) (*model.PayOrderTodayStats, error)
}

type PayOrderServiceImpl struct {
	db             *ent.Client
	settingService setting_service.SettingService
}

func NewPayOrderServiceImpl(db *ent.Client, settingService setting_service.SettingService) PayOrderService {
	return &PayOrderServiceImpl{db: db, settingService: settingService}
}

// paymentSettings 对应系统设置中 key='payment' 的 JSON 配置，仅声明后端消费的易支付字段。
type paymentSettings struct {
	EnableEpay      bool   `json:"enableEpay"`
	EpayApiUrl      string `json:"epayApiUrl"`
	EpayMerchantId  string `json:"epayMerchantId"`
	EpayMerchantKey string `json:"epayMerchantKey"`
	EpayNotifyUrl   string `json:"epayNotifyUrl"`
	EpayReturnUrl   string `json:"epayReturnUrl"`
}

// loadEpayConfig 从系统设置读取易支付配置，每次下单按需读取以保证设置修改即时生效。
func (s *PayOrderServiceImpl) loadEpayConfig(ctx context.Context) (epay.Config, bool, error) {
	setting, err := s.settingService.GetSettingByKey(ctx, "payment")
	if err != nil {
		if ent.IsNotFound(err) {
			return epay.Config{}, false, fmt.Errorf("支付配置未初始化，请先在系统设置-支付设置中配置")
		}
		return epay.Config{}, false, err
	}

	var ps paymentSettings
	if err := json.Unmarshal([]byte(setting.Value), &ps); err != nil {
		return epay.Config{}, false, fmt.Errorf("解析支付配置失败: %w", err)
	}

	cfg := epay.Config{
		MchID:     ps.EpayMerchantId,
		Key:       ps.EpayMerchantKey,
		APIURL:    ps.EpayApiUrl,
		NotifyURL: ps.EpayNotifyUrl,
		ReturnURL: ps.EpayReturnUrl,
	}
	return cfg, ps.EnableEpay, nil
}

// 查询支付订单列表
func (s *PayOrderServiceImpl) ListPayOrderPage(ctx context.Context, req *model.PageQuery) ([]*ent.PayOrder, int, error) {
	orders, err := s.db.PayOrder.Query().
		Limit(req.Size).
		Offset((req.Page - 1) * req.Size).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.db.PayOrder.Query().
		Order(ent.Desc(payorder.FieldID)).
		Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return orders, count, nil
}

func (s *PayOrderServiceImpl) SubmitPayOrder(ctx context.Context, userID int, req *model.PayOrderSubmitReq) (*model.PayOrderSubmitResp, error) {
	// 金额单位为分，epay 需要元字符串
	moneyYuan := strconv.FormatFloat(float64(req.Money)/100, 'f', 2, 64)

	// 创建订单并落库业务字段
	orderCreate := s.db.PayOrder.Create().
		SetUserID(userID).
		SetOrderType(req.OrderType).
		SetChannelType(req.ChannelType).
		SetSubject(req.Name).
		SetBody(req.Name).
		SetOrderPrice(req.Money).
		SetPrice(req.Money).
		SetState("1")
	switch req.OrderType {
	case model.PayOrderTypePost:
		orderCreate = orderCreate.SetPostID(req.PostId)
	case model.PayOrderTypeProduct:
		orderCreate = orderCreate.SetProductID(req.ProductId)
	}
	order, err := orderCreate.Save(ctx)
	if err != nil {
		return nil, err
	}

	// 生成商户订单号
	orderNo := time.Now().Format("20060102150405") + fmt.Sprintf("%09d", order.ID)
	order, err = s.db.PayOrder.UpdateOneID(order.ID).SetOutTradeNo(orderNo).Save(ctx)
	if err != nil {
		return nil, err
	}

	// 从系统设置读取易支付配置（每次下单按需读取，保证设置修改即时生效）
	cfg, enabled, err := s.loadEpayConfig(ctx)
	if err != nil {
		return nil, err
	}
	if !enabled {
		return nil, fmt.Errorf("易支付未启用")
	}

	// 调用易支付接口下单
	returnURL := req.ReturnUrl
	params := epay.V1PayRequestParams{
		PID:        cfg.MchID,
		Type:       req.ChannelType,
		OutTradeNo: orderNo,
		Name:       req.Name,
		Money:      moneyYuan,
		ReturnURL:  &returnURL,
	}
	resp, err := epay.NewV1Client(cfg).CreateOrder(params)
	if err != nil {
		// 下单失败，记录错误信息
		_ = s.db.PayOrder.UpdateOneID(order.ID).SetState("3").SetErrorMsg(err.Error()).Exec(ctx)
		return nil, err
	}

	// 存储支付链接与易支付订单号
	payURL := ""
	if resp.Payurl != nil {
		payURL = *resp.Payurl
	}
	_, err = s.db.PayOrder.UpdateOneID(order.ID).
		SetPayURL(payURL).
		SetOrderID(resp.TradeNO).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.PayOrderSubmitResp{
		OrderID:    order.ID,
		OutTradeNo: orderNo,
		PayURL:     payURL,
		TradeNO:    resp.TradeNO,
	}, nil
}

// handlePaySuccess 标记订单为已支付并执行业务履约（文章付费写购买记录），幂等。
func (s *PayOrderServiceImpl) handlePaySuccess(ctx context.Context, order *ent.PayOrder) error {
	if order.State == "2" {
		return nil
	}

	order, err := s.db.PayOrder.UpdateOneID(order.ID).SetState("2").Save(ctx)
	if err != nil {
		return err
	}

	// 文章付费：写购买记录（唯一索引保证不重复）
	if order.OrderType == model.PayOrderTypePost && order.PostID != 0 {
		exist, err := s.db.PostPurchase.Query().
			Where(postpurchase.UserIDEQ(order.UserID), postpurchase.PostIDEQ(order.PostID)).
			Exist(ctx)
		if err != nil {
			return err
		}
		if !exist {
			if _, err := s.db.PostPurchase.Create().
				SetUserID(order.UserID).
				SetPostID(order.PostID).
				SetOrderID(order.ID).
				Save(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

// HandleNotify 处理易支付异步回调：验签 -> 找订单 -> 支付成功则履约。
func (s *PayOrderServiceImpl) HandleNotify(ctx context.Context, params map[string]string) error {
	cfg, _, err := s.loadEpayConfig(ctx)
	if err != nil {
		return err
	}

	sign, ok := params["sign"]
	if !ok {
		return fmt.Errorf("缺少签名参数")
	}
	if !epay.VerifySign(params, cfg.Key, sign) {
		return fmt.Errorf("签名校验失败")
	}

	// 仅处理支付成功的通知
	if params["trade_status"] != "TRADE_SUCCESS" {
		return nil
	}

	outTradeNo := params["out_trade_no"]
	order, err := s.db.PayOrder.Query().
		Where(payorder.OutTradeNoEQ(outTradeNo)).
		Only(ctx)
	if err != nil {
		return err
	}

	return s.handlePaySuccess(ctx, order)
}

// SyncOrderStatus 主动查询易支付订单状态并同步本地（用于补单/后台手动刷新）。
func (s *PayOrderServiceImpl) SyncOrderStatus(ctx context.Context, orderID int) (*model.PayOrderStatusResp, error) {
	order, err := s.db.PayOrder.Get(ctx, orderID)
	if err != nil {
		return nil, err
	}

	// 已支付的订单无需同步
	if order.State == "2" {
		return toPayOrderStatusResp(order), nil
	}

	cfg, enabled, err := s.loadEpayConfig(ctx)
	if err != nil {
		return nil, err
	}
	if !enabled {
		return nil, fmt.Errorf("易支付未启用")
	}

	pid, err := strconv.Atoi(cfg.MchID)
	if err != nil {
		return nil, fmt.Errorf("商户ID无效: %w", err)
	}
	tradeNo := ""
	if order.OrderID != nil {
		tradeNo = *order.OrderID
	}
	outTradeNo := ""
	if order.OutTradeNo != nil {
		outTradeNo = *order.OutTradeNo
	}

	queryResp, err := epay.NewV1Client(cfg).QueryOrder(pid, cfg.Key, tradeNo, outTradeNo)
	if err != nil {
		return nil, err
	}

	// 易支付 status=2 表示已支付
	if queryResp.Status == 2 {
		if err := s.handlePaySuccess(ctx, order); err != nil {
			return nil, err
		}
		order, err = s.db.PayOrder.Get(ctx, orderID)
		if err != nil {
			return nil, err
		}
	}

	return toPayOrderStatusResp(order), nil
}

func toPayOrderStatusResp(o *ent.PayOrder) *model.PayOrderStatusResp {
	outTradeNo := ""
	if o.OutTradeNo != nil {
		outTradeNo = *o.OutTradeNo
	}
	return &model.PayOrderStatusResp{
		ID:         o.ID,
		OutTradeNo: outTradeNo,
		State:      o.State,
		Subject:    o.Subject,
		Price:      o.Price,
	}
}

func (s *PayOrderServiceImpl) GetTodayStats(ctx context.Context) (*model.PayOrderTodayStats, error) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := todayStart.Add(24 * time.Hour)

	orders, err := s.db.PayOrder.Query().
		Where(payorder.CreatedAtGTE(todayStart)).
		Where(payorder.CreatedAtLT(todayEnd)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	total := len(orders)
	var amount float64
	pending := 0
	successCount := 0

	for _, order := range orders {
		if order.State == "2" {
			successCount++
			amount += float64(order.Price) / 100
		} else if order.State == "1" {
			pending++
		}
	}

	var successRate float64
	if total > 0 {
		successRate = float64(successCount) / float64(total) * 100
	}

	return &model.PayOrderTodayStats{
		Total:       total,
		Amount:      amount,
		SuccessRate: successRate,
		Pending:     pending,
	}, nil
}
