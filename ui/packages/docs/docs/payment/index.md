# 支付模块

支付模块提供了完整的支付功能，支持多种支付方式，包括支付宝、微信支付和易支付。

## 功能特性

- **多支付方式**：支持支付宝、微信支付、易支付
- **统一接口**：提供统一的支付接口，方便接入
- **订单管理**：完整的订单创建、查询、退款流程
- **支付回调**：支持异步通知和同步跳转
- **安全可靠**：完善的签名验证和错误处理机制

## 支付方式

### 支付宝

支持支付宝扫码支付、网页支付、移动支付等多种支付方式。

### 微信支付

支持微信扫码支付、公众号支付、小程序支付等多种支付方式。

### 易支付

支持易支付平台接入，提供统一的支付接口。

## 快速开始

### 1. 配置支付参数

在 `config.yaml` 中配置支付参数：

```yaml
payment:
  alipay:
    app_id: your_alipay_app_id
    private_key: your_alipay_private_key
    public_key: your_alipay_public_key
  wechat:
    app_id: your_wechat_app_id
    mch_id: your_wechat_mch_id
    api_key: your_wechat_api_key
  epay:
    url: your_epay_url
    pid: your_epay_pid
    key: your_epay_key
```

### 2. 创建支付订单

```javascript
import { paymentApi } from '@/api/payment'

const createPayment = async () => {
  const data = {
    order_id: 'ORDER_20240101',
    amount: 100,
    subject: '商品名称',
    type: 'alipay'
  }
  
  const result = await paymentApi.create(data)
  console.log(result)
}
```

### 3. 处理支付回调

支付成功后，系统会自动处理支付回调，更新订单状态。

## API 接口

### 创建支付订单

```http
POST /api/v1/payment/create
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| order_id | string | 是 | 订单号 |
| amount | decimal | 是 | 支付金额 |
| subject | string | 是 | 支付标题 |
| type | string | 是 | 支付类型 (alipay/wechat/epay) |
| return_url | string | 否 | 同步返回地址 |
| notify_url | string | 否 | 异步通知地址 |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "order_id": "ORDER_20240101",
    "payment_url": "https://qr.alipay.com/xxx",
    "qr_code": "data:image/png;base64,..."
  }
}
```

### 查询支付订单

```http
GET /api/v1/payment/query/:order_id
```

### 申请退款

```http
POST /api/v1/payment/refund
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| order_id | string | 是 | 订单号 |
| refund_amount | decimal | 是 | 退款金额 |
| reason | string | 否 | 退款原因 |

## 支付流程

### 支付宝支付流程

1. 创建支付订单
2. 获取支付二维码或跳转链接
3. 用户扫码或点击链接完成支付
4. 支付宝异步通知支付结果
5. 系统更新订单状态

### 微信支付流程

1. 创建支付订单
2. 获取支付二维码或跳转链接
3. 用户扫码或点击链接完成支付
4. 微信异步通知支付结果
5. 系统更新订单状态

### 易支付流程

1. 创建支付订单
2. 获取支付二维码或跳转链接
3. 用户扫码或点击链接完成支付
4. 易支付异步通知支付结果
5. 系统更新订单状态

## 安全建议

1. **签名验证**：所有支付回调必须验证签名
2. **订单校验**：处理回调时校验订单金额和状态
3. **重复通知**：处理重复的支付回调通知
4. **日志记录**：记录所有支付相关日志
5. **错误处理**：完善的错误处理机制

## 相关文档

- [支付宝支付](alipay.md) - 支付宝支付详细文档
- [微信支付](wechat.md) - 微信支付详细文档
- [易支付](epay.md) - 易支付详细文档
- [商城模块](../mall/index.md) - 商城模块文档
