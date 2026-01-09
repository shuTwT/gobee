package epay

const (
	//支付宝
	PayTypeAlipay PayType = "alipay"
	//微信支付
	PayTypeWechat PayType = "wxpay"
	//QQ钱包
	PayTypeQQ PayType = "qqpay"
	//网银支付
	PayTypeBank PayType = "bank"
)

// 配置信息
type Config struct {
	MchID     string // 商户ID
	Key       string // 商户密钥
	APIURL    string // 支付接口地址
	NotifyURL string // 回调通知地址
	ReturnURL string // 跳转返回地址
}

type BasePayResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type EPayClient interface {
	// 创建支付订单
	CreateOrder(params PayRequestParams) (PayResponse, error)
	// 查询商户信息
	QueryMerchant(pid int, key string) error
	// v1查询结算记录 deprecated
	QuerySettle(pid int, key string) error
	// 查询支付订单
	QueryOrder(pid int, key string, trade_no string, out_trade_no string) error
	// 批量查询支付订单
	QueryOrderBatch(pid int, key string, limit int, page int) error
	// v1提交订单退款
	RefundOrder(pid int, key string, trade_no int, refund_no int, money string) error
}
