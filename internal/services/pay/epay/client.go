package epay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 配置信息
type Config struct {
	MchID     string // 商户ID
	Key       string // 商户密钥
	APIURL    string // 支付接口地址
	NotifyURL string // 回调通知地址
	ReturnURL string // 跳转返回地址
}

// 支付请求参数
type PayRequest struct {
	MchID      string // 商户ID
	OutTradeNo string // 商户订单号
	TotalFee   int    // 支付金额(分)
	Body       string // 商品描述
	PayType    string // 支付方式 alipay,wechat
	NotifyURL  string // 回调地址
	ReturnURL  string // 返回地址
	TimeStamp  string // 时间戳
	Sign       string // 签名
}

// 支付响应
type PayResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PayInfo string `json:"pay_info"` // 支付信息，可能是URL或HTML
		OrderNo string `json:"order_no"` // 平台订单号
	} `json:"data"`
}

// 订单查询请求
type QueryRequest struct {
	MchID      string // 商户ID
	OutTradeNo string // 商户订单号
	OrderNo    string // 平台订单号(二选一)
	TimeStamp  string // 时间戳
	Sign       string // 签名
}

// 订单查询响应
type QueryResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OutTradeNo string `json:"out_trade_no"`
		OrderNo    string `json:"order_no"`
		Status     int    `json:"status"` // 0-未支付 1-已支付
		TotalFee   int    `json:"total_fee"`
		PayTime    string `json:"pay_time"`
	} `json:"data"`
}

// 回调通知参数
type NotifyParams struct {
	MchID      string `json:"mch_id"`
	OutTradeNo string `json:"out_trade_no"`
	OrderNo    string `json:"order_no"`
	TotalFee   int    `json:"total_fee"`
	Status     int    `json:"status"`
	Sign       string `json:"sign"`
}

// 客户端
type Client struct {
	config Config
}

// 创建客户端
func NewClient(config Config) *Client {
	return &Client{
		config: config,
	}
}

// 生成签名
func (c *Client) generateSign(params map[string]string) string {
	// 1. 排序参数
	keys := make([]string, 0, len(params))
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接参数
	var signStr strings.Builder
	for _, k := range keys {
		if params[k] != "" {
			signStr.WriteString(fmt.Sprintf("%s=%s&", k, params[k]))
		}
	}

	// 3. 拼接密钥
	signStr.WriteString(fmt.Sprintf("key=%s", c.config.Key))

	// 4. MD5加密
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStr.String()))
	return strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil)))
}

// 将map[string]string转换为url.Values
func toURLValues(params map[string]string) url.Values {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	return values
}

// 创建支付订单
func (c *Client) CreatePay(payType, outTradeNo, body string, totalFee int) (*PayResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"mch_id":       c.config.MchID,
		"out_trade_no": outTradeNo,
		"total_fee":    fmt.Sprintf("%d", totalFee),
		"body":         body,
		"pay_type":     payType,
		"notify_url":   c.config.NotifyURL,
		"return_url":   c.config.ReturnURL,
		"time_stamp":   fmt.Sprintf("%d", time.Now().Unix()),
	}

	// 生成签名
	params["sign"] = c.generateSign(params)

	// 发送请求 - 使用转换函数将map转换为url.Values
	resp, err := http.PostForm(c.config.APIURL+"/pay", toURLValues(params))
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var payResp PayResponse
	if err := json.Unmarshal(bodyBytes, &payResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if payResp.Code != 0 {
		return nil, fmt.Errorf("接口返回错误: %s", payResp.Msg)
	}

	return &payResp, nil
}

// 查询订单状态
func (c *Client) QueryOrder(outTradeNo, orderNo string) (*QueryResponse, error) {
	params := map[string]string{
		"mch_id":       c.config.MchID,
		"out_trade_no": outTradeNo,
		"order_no":     orderNo,
		"time_stamp":   fmt.Sprintf("%d", time.Now().Unix()),
	}

	// 生成签名
	params["sign"] = c.generateSign(params)

	// 构建查询URL
	u, err := url.Parse(c.config.APIURL + "/query")
	if err != nil {
		return nil, err
	}

	// 使用转换函数将map转换为url.Values
	query := toURLValues(params)
	u.RawQuery = query.Encode()

	// 发送请求
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var queryResp QueryResponse
	if err := json.Unmarshal(bodyBytes, &queryResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(bodyBytes))
	}

	if queryResp.Code != 0 {
		return nil, fmt.Errorf("查询失败: %s", queryResp.Msg)
	}

	return &queryResp, nil
}

// 验证回调签名
func (c *Client) VerifyNotify(params NotifyParams) bool {
	// 构建参数map
	paramMap := map[string]string{
		"mch_id":       params.MchID,
		"out_trade_no": params.OutTradeNo,
		"order_no":     params.OrderNo,
		"total_fee":    fmt.Sprintf("%d", params.TotalFee),
		"status":       fmt.Sprintf("%d", params.Status),
	}

	// 生成签名并验证
	sign := c.generateSign(paramMap)
	return sign == params.Sign
}

// 处理回调通知的HTTP处理器
func (c *Client) NotifyHandler(handleFunc func(NotifyParams) bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析请求
		r.ParseForm()

		// 提取参数
		totalFee, _ := strconv.Atoi(r.Form.Get("total_fee"))
		status, _ := strconv.Atoi(r.Form.Get("status"))

		params := NotifyParams{
			MchID:      r.Form.Get("mch_id"),
			OutTradeNo: r.Form.Get("out_trade_no"),
			OrderNo:    r.Form.Get("order_no"),
			TotalFee:   totalFee,
			Status:     status,
			Sign:       r.Form.Get("sign"),
		}

		// 验证签名
		if !c.VerifyNotify(params) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("fail"))
			return
		}

		// 处理业务逻辑
		success := handleFunc(params)

		// 返回结果
		if success {
			w.Write([]byte("success"))
		} else {
			w.Write([]byte("fail"))
		}
	}
}
