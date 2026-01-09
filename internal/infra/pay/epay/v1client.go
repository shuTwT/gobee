package epay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// 支付请求参数
type V1PayRequestParams struct {
	PID        string  // 商户ID
	Type       string  // 支付方式
	OutTradeNo string  // 商户订单号
	NotifyURL  *string // 异步回调地址
	ReturnURL  *string // 跳转通知地址
	Name       string  // 商品名称
	Money      string  // 商品金额
	ClientIP   string  // 客户端IP
	Device     *string // 设备类型
	Param      *string // 业务扩展参数
	Sign       string  // 签名字符串
	SignType   *string // 签名类型，默认MD5
}

// 支付结果返回响应
type V1CreateOrderResp struct {
	BasePayResult
	TradeNO   string  `json:"trade_no"`            // 易支付订单号
	Payurl    *string `json:"payurl,omitempty"`    // 支付URL
	Qrcode    *string `json:"qrcode,omitempty"`    // 二维码URL
	Urlscheme *string `json:"urlscheme,omitempty"` // 小程序跳转地址
}

// 支付结果通知
type V1PayResultNotify struct {
	PID         string `json:"pid"`
	TradeNO     string `json:"trade_no"`     // 易支付订单号
	OutTradeNo  string `json:"out_trade_no"` // 商户系统内部的订单号
	Type        string `json:"type"`         // 支付方式
	Name        string `json:"name"`         // 商品名称
	Money       string `json:"money"`        // 商品金额
	TradeStatus string `json:"trade_status"`
	Sign        string `json:"sign"`
	SignType    string `json:"sign_type"` //签名类型，默认为md5
}

// 查询支付订单响应
type V1QueryOrder struct {
	TradeNO    string  `json:"trade_no"`     // 易支付订单号
	OutTradeNo string  `json:"out_trade_no"` // 商户系统内部的订单号
	ApiTradeNO string  `json:"api_trade_no"` // 支付宝微信等接口方订单号
	Type       string  `json:"type"`         // 支付方式
	PID        int     `json:"pid"`          //商户 ID
	Addtime    string  `json:"addtime"`      // 订单创建时间
	Endtime    string  `json:"endtime"`      // 订单支付时间
	Name       string  `json:"name"`         // 商品名称
	Money      string  `json:"money"`        // 商品金额
	Status     int     `json:"status"`
	Param      *string `json:"param,omitempty"`
	Buyer      *string `json:"buyer,omitempty"`
}

// 查询商户信息响应
type V1QueryMerchantResp struct {
	BasePayResult
	PID          int    `json:"pid"`
	Key          string `json:"key"`           // 商户密钥
	Active       int    `json:"active"`        // 商户状态
	Money        string `json:"money"`         // 商户余额
	Type         int    `json:"type"`          // 结算方式
	Account      string `json:"account"`       // 结算账号
	Username     string `json:"username"`      //结算姓名
	Orders       int    `json:"orders"`        // 订单总数
	OrderToday   int    `json:"order_today"`   // 今日订单
	OrderLastday int    `json:"order_lastday"` // 昨日订单
}

// 查询结算记录
type V1QuerySettleResp struct {
	BasePayResult
	Data []interface{} `json:"data"`
}

// 查询支付订单响应
type V1QueryOrderResp struct {
	BasePayResult
	V1QueryOrder
}

// 批量查询支付订单响应
type V1BatchQueryOrderResp struct {
	BasePayResult
	Data []V1QueryOrder `json:"data"`
}

// 提交订单退款响应
type V1RefundOrderResp struct {
	BasePayResult
}

type V1Client struct {
	config Config
}

func NewV1Client(config Config) *V1Client {
	return &V1Client{config: config}
}

// 生成签名
func (c *V1Client) generateSign(params map[string]string, key string) string {
	// 1. 收集所有需要签名的 k=v 字符串（排除 sign 和空值）
	var pairs []string
	for k, v := range params {
		if k == "sign" || v == "" {
			continue
		}
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}

	// 2. 按 key 排序
	sort.Strings(pairs)

	// 3. 拼接为 a=b&c=d&e=f
	signStr := strings.Join(pairs, "&")

	// 4. 直接拼接 key（前面无 &）
	signStr += key

	log.Printf("签名前字符串: %s", signStr)

	// 5. MD5 加密
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStr))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// 创建支付订单
func (c *V1Client) CreateOrder(params V1PayRequestParams) (*V1CreateOrderResp, error) {

	paramsMap := map[string]string{
		"pid":          params.PID,
		"type":         params.Type,
		"out_trade_no": params.OutTradeNo,
		"notify_url":   c.config.NotifyURL,
		"return_url":   c.config.ReturnURL,
		"name":         params.Name,
		"money":        params.Money,
		"clientIp":     params.ClientIP,
		"device":       "PC",
	}

	// 生成签名
	paramsMap["sign"] = c.generateSign(paramsMap, c.config.Key)
	paramsMap["sign_type"] = "MD5"

	// 发送请求 - 使用转换函数将map转换为url.Values
	log.Printf("易支付请求地址: %v", c.config.APIURL+"/mapi.php")
	data := toURLValues(paramsMap)
	log.Printf("请求参数: %v", data)
	resp, err := http.PostForm(c.config.APIURL+"/mapi.php", data)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	log.Printf("响应内容: %s", string(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var payResp V1CreateOrderResp
	if err := json.Unmarshal(bodyBytes, &payResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if payResp.Code != 0 {
		return nil, fmt.Errorf("接口返回错误: %s", payResp.Msg)
	}

	return &payResp, nil

}

// 查询商户信息
func (c *V1Client) QueryMerchant(pid int, key string) (*V1QueryMerchantResp, error) {
	paramsMap := map[string]string{
		"act":  "query",
		"pid":  strconv.Itoa(pid),
		"type": key,
	}

	// 发送请求 - 使用转换函数将map转换为url.Values
	url := c.config.APIURL + "/api.php" + "?" + toURLValues(paramsMap).Encode()
	log.Printf("易支付请求地址: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	log.Printf("响应内容: %s", string(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var payResp V1QueryMerchantResp
	if err := json.Unmarshal(bodyBytes, &payResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if payResp.Code != 0 {
		return nil, fmt.Errorf("接口返回错误: %s", payResp.Msg)
	}

	return &payResp, nil
}

// v1查询结算记录 deprecated
func (c *V1Client) QuerySettle(pid int, key string) (*V1QuerySettleResp, error) {
	paramsMap := map[string]string{
		"act":  "settle",
		"pid":  strconv.Itoa(pid),
		"type": key,
	}

	// 发送请求 - 使用转换函数将map转换为url.Values
	url := c.config.APIURL + "/api.php" + "?" + toURLValues(paramsMap).Encode()
	log.Printf("易支付请求地址: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	log.Printf("响应内容: %s", string(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var payResp V1QuerySettleResp
	if err := json.Unmarshal(bodyBytes, &payResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if payResp.Code != 0 {
		return nil, fmt.Errorf("接口返回错误: %s", payResp.Msg)
	}

	return &payResp, nil
}

// 查询支付订单
func (c *V1Client) QueryOrder(pid int, key string, trade_no string, out_trade_no string) (*V1QueryOrderResp, error) {
	paramsMap := map[string]string{
		"act":          "order",
		"pid":          strconv.Itoa(pid),
		"trade_no":     trade_no,
		"out_trade_no": out_trade_no,
	}

	// 发送请求 - 使用转换函数将map转换为url.Values
	url := c.config.APIURL + "/api.php" + "?" + toURLValues(paramsMap).Encode()
	log.Printf("易支付请求地址: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	log.Printf("响应内容: %s", string(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var payResp V1QueryOrderResp
	if err := json.Unmarshal(bodyBytes, &payResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if payResp.Code != 0 {
		return nil, fmt.Errorf("接口返回错误: %s", payResp.Msg)
	}

	return &payResp, nil
}

// 批量查询支付订单
func (c *V1Client) QueryOrderBatch(pid int, key string, limit int, page int) (*V1BatchQueryOrderResp, error) {
	paramsMap := map[string]string{
		"act":   "orders",
		"pid":   strconv.Itoa(pid),
		"key":   key,
		"limit": fmt.Sprintf("%d", limit),
		"page":  fmt.Sprintf("%d", page),
	}

	// 发送请求 - 使用转换函数将map转换为url.Values
	url := c.config.APIURL + "/api.php" + "?" + toURLValues(paramsMap).Encode()
	log.Printf("易支付请求地址: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	log.Printf("响应内容: %s", string(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var payResp V1BatchQueryOrderResp
	if err := json.Unmarshal(bodyBytes, &payResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if payResp.Code != 0 {
		return nil, fmt.Errorf("接口返回错误: %s", payResp.Msg)
	}

	return &payResp, nil
}

// v1提交订单退款
func (c *V1Client) RefundOrder(pid int, key string, trade_no string, refund_no string, money string) (*V1RefundOrderResp, error) {
	paramsMap := map[string]string{
		"act":       "refund",
		"pid":       strconv.Itoa(pid),
		"key":       key,
		"trade_no":  trade_no,
		"refund_no": refund_no,
		"money":     money,
	}

	// 发送请求 - 使用转换函数将map转换为url.Values
	url := c.config.APIURL + "/api.php" + "?" + toURLValues(paramsMap).Encode()
	log.Printf("易支付请求地址: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	log.Printf("响应内容: %s", string(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var payResp V1RefundOrderResp
	if err := json.Unmarshal(bodyBytes, &payResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if payResp.Code != 0 {
		return nil, fmt.Errorf("接口返回错误: %s", payResp.Msg)
	}

	return &payResp, nil
}
