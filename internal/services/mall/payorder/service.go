package payorder

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/payorder"
	"github.com/shuTwT/hoshikuzu/internal/infra/pay/epay"
	setting_service "github.com/shuTwT/hoshikuzu/internal/services/system/setting"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
)

type PayOrderService interface {
	ListPayOrderPage(ctx context.Context, req *model.PageQuery) ([]*ent.PayOrder, int, error)
	SubmitPayOrder(ctx context.Context, req *model.PayOrderSubmitReq) error
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

func (s *PayOrderServiceImpl) CreatePayOrder(ctx context.Context) (*ent.PayOrder, error) {
	order, err := s.db.PayOrder.Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// 更新支付订单号
func (s *PayOrderServiceImpl) UpdatePayOrderNO(ctx context.Context, order *ent.PayOrder) (*ent.PayOrder, error) {
	id := order.ID

	now := time.Now().Format("20060102150405")
	idStr := fmt.Sprintf("%09d", id)

	orderNo := now + idStr
	return s.db.PayOrder.UpdateOne(order).
		SetOutTradeNo(orderNo).
		Save(ctx)
}

func (s *PayOrderServiceImpl) SubmitPayOrder(ctx context.Context, req *model.PayOrderSubmitReq) error {
	// 先创建一个支付订单
	order, err := s.CreatePayOrder(ctx)
	if err != nil {
		return err
	}
	// 更新支付订单号
	order, err = s.UpdatePayOrderNO(ctx, order)
	if err != nil {
		return err
	}

	// 从系统设置读取易支付配置（每次下单按需读取，保证设置修改即时生效）
	cfg, enabled, err := s.loadEpayConfig(ctx)
	if err != nil {
		return err
	}
	if enabled {
		// 调用易支付接口
		params := epay.V1PayRequestParams{
			PID:        cfg.MchID,
			Type:       req.ChannelType,
			OutTradeNo: *order.OutTradeNo,
			Name:       req.Name,
			Money:      strconv.Itoa(req.Money),
		}
		resp, err := epay.NewV1Client(cfg).CreateOrder(params)
		if err != nil {
			return err
		}
		if resp.Code != 1 {
			return fmt.Errorf("创建支付订单失败: %s", resp.Msg)
		}
	}
	return nil
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
