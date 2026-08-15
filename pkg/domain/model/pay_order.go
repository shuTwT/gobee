package model

// PayOrderCreateReq represents the request body for creating a payment order.
type PayOrderCreateReq struct {
	ChannelType string `json:"channel_id" validate:"required"`
	OrderID     string `json:"order_id"`
	OutTradeNo  string `json:"out_trade_no" validate:"required"`
	TotalFee    string `json:"total_fee" validate:"required"`
	Subject     string `json:"subject" validate:"required"`
	Body        string `json:"body" validate:"required"`
	NotifyURL   string `json:"notify_url" validate:"required,url"`
	ReturnURL   string `json:"return_url" validate:"required,url"`
	Extra       string `json:"extra"`
	PayURL      string `json:"pay_url,omitempty"`
	State       string `json:"state"`
	ErrorMsg    string `json:"error_msg,omitempty"`
	Raw         string `json:"raw,omitempty"`
}

// PayOrderUpdateReq represents the request body for updating a payment order.
type PayOrderUpdateReq struct {
	ChannelType     string `json:"channel_type"`
	OrderID         string `json:"order_id"`
	OutTradeNo      string `json:"out_trade_no"`
	OrderPrice      int    `json:"order_price"`
	Price           int    `json:"price"`
	ChannelFeePrice int    `json:"channel_fee_price"`
	Subject         string `json:"subject"`
	Body            string `json:"body"`
	NotifyURL       string `json:"notify_url"`
	ReturnURL       string `json:"return_url"`
	Extra           string `json:"extra"`
	PayURL          string `json:"pay_url,omitempty"`
	State           string `json:"state"`
	ErrorMsg        string `json:"error_msg,omitempty"`
	Raw             string `json:"raw,omitempty"`
}

// PayOrderResp represents the response body for a payment order.
type PayOrderResp struct {
	ID         int       `json:"id"`
	CreatedAt  LocalTime `json:"created_at"`
	ChannelID  string    `json:"channel_id"`
	OrderID    string    `json:"order_id"`
	OutTradeNo string    `json:"out_trade_no"`
	TotalFee   string    `json:"total_fee"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	NotifyURL  string    `json:"notify_url"`
	ReturnURL  string    `json:"return_url"`
	Extra      string    `json:"extra"`
	PayURL     *string   `json:"pay_url,omitempty"`
	State      string    `json:"state"`
	ErrorMsg   *string   `json:"error_msg,omitempty"`
}

const (
	PayOrderTypePost     = "1"
	PayOrderTypeProduct  = "2"
	PayOrderTypeRecharge = "3"

	// 支付渠道
	PayChannelBalance = "balance"
)

// PayMethodResp 可用支付方式查询响应
type PayMethodResp struct {
	// 支付宝（易支付或直连支付宝启用）
	Alipay bool `json:"alipay"`
	// 微信支付（易支付或直连微信支付启用）
	Wechat bool `json:"wechat"`
	// 余额支付（钱包余额大于 0）
	Balance bool `json:"balance"`
	// 钱包余额,单位分
	BalanceAmount int `json:"balance_amount"`
}

type PayOrderSubmitReq struct {
	// 渠道类型 1 支付宝 2 微信 3 银联
	ChannelType string `json:"channel_type" validate:"required"`
	// 返回地址
	ReturnUrl string `json:"return_url" validate:"required,url"`
	// 订单类型 1 文章付费 2 商品购买
	OrderType string `json:"order_type" validate:"required"`
	// 商品名称
	Name string `json:"name"`
	// 金额
	Money int `json:"money"`
	// 文章 id，可选
	PostId int `json:"post_id"`
	// 商品 id，可选
	ProductId int `json:"product_id"`
}

type PayOrderTodayStats struct {
	Total       int     `json:"total"`
	Amount      float64 `json:"amount"`
	SuccessRate float64 `json:"success_rate"`
	Pending     int     `json:"pending"`
}

// PayOrderSubmitResp 提交支付订单后返回给前端的信息
type PayOrderSubmitResp struct {
	OrderID    int    `json:"order_id"`     // 系统订单ID
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	PayURL     string `json:"pay_url"`      // 支付链接
	TradeNO    string `json:"trade_no"`     // 易支付订单号
}

// PayOrderStatusResp 订单状态查询响应
type PayOrderStatusResp struct {
	ID         int    `json:"id"`
	OutTradeNo string `json:"out_trade_no"`
	State      string `json:"state"`
	Subject    string `json:"subject"`
	Price      int    `json:"price"`
}

// PayOrderRefundReq 退款请求（金额可选，为空或大于等于订单支付金额时全额退款）
type PayOrderRefundReq struct {
	// 退款金额,单位分,0 表示全额退款
	Amount int `json:"amount"`
}

// PayOrderRefundResp 退款响应
type PayOrderRefundResp struct {
	OrderID      int    `json:"order_id"`
	RefundNo     string `json:"refund_no"`
	RefundAmount int    `json:"refund_amount"`
	State        string `json:"state"`
}

// PayOrderRechargeReq 余额充值请求
type PayOrderRechargeReq struct {
	// 渠道类型 1 支付宝 2 微信 3 银联
	ChannelType string `json:"channel_type" validate:"required"`
	// 返回地址
	ReturnUrl string `json:"return_url" validate:"required,url"`
	// 充值金额,单位分
	Amount int `json:"amount"`
}
