# 优惠券

优惠券是商城模块的重要组成部分，提供完整的优惠券管理能力。

## 功能特性

- 优惠券创建：创建不同类型的优惠券
- 优惠券发放：向用户发放优惠券
- 优惠券核销：用户使用优惠券
- 优惠券统计：优惠券使用统计

## API 接口

### 获取优惠券列表

```http
GET /api/v1/mall/coupon/list
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| type | string | 否 | 优惠券类型 (discount/amount) |
| status | string | 否 | 状态 (active/inactive) |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": 1,
        "name": "新用户优惠券",
        "type": "discount",
        "value": 10,
        "min_amount": 100,
        "total": 1000,
        "used": 100,
        "status": "active",
        "start_time": "2024-01-01T00:00:00Z",
        "end_time": "2024-12-31T23:59:59Z"
      }
    ]
  }
}
```

### 创建优惠券

```http
POST /api/v1/mall/coupon/create
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 是 | 优惠券名称 |
| type | string | 是 | 优惠券类型 (discount/amount) |
| value | decimal | 是 | 优惠券值 |
| min_amount | decimal | 否 | 最小使用金额 |
| total | int | 是 | 总数量 |
| start_time | datetime | 是 | 开始时间 |
| end_time | datetime | 是 | 结束时间 |

**请求示例：**

```json
{
  "name": "新用户优惠券",
  "type": "discount",
  "value": 10,
  "min_amount": 100,
  "total": 1000,
  "start_time": "2024-01-01T00:00:00Z",
  "end_time": "2024-12-31T23:59:59Z"
}
```

### 更新优惠券

```http
PUT /api/v1/mall/coupon/update/:id
```

**请求参数：** 同创建优惠券

### 删除优惠券

```http
DELETE /api/v1/mall/coupon/delete/:id
```

### 领取优惠券

```http
POST /api/v1/mall/coupon/receive
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| coupon_id | int | 是 | 优惠券 ID |

### 获取我的优惠券

```http
GET /api/v1/mall/coupon/my
```

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": 1,
        "coupon_id": 1,
        "coupon_name": "新用户优惠券",
        "type": "discount",
        "value": 10,
        "min_amount": 100,
        "status": "unused",
        "start_time": "2024-01-01T00:00:00Z",
        "end_time": "2024-12-31T23:59:59Z"
      }
    ]
  }
}
```

## 使用示例

### 创建优惠券

```javascript
import { couponApi } from '@/api/mall/coupon'

const createCoupon = async () => {
  const data = {
    name: '新用户优惠券',
    type: 'discount',
    value: 10,
    min_amount: 100,
    total: 1000,
    start_time: '2024-01-01T00:00:00Z',
    end_time: '2024-12-31T23:59:59Z'
  }
  
  const result = await couponApi.create(data)
  console.log(result)
}
```

### 领取优惠券

```javascript
import { couponApi } from '@/api/mall/coupon'

const receiveCoupon = async () => {
  const data = {
    coupon_id: 1
  }
  
  const result = await couponApi.receive(data)
  console.log(result)
}
```

### 获取我的优惠券

```javascript
import { couponApi } from '@/api/mall/coupon'

const getMyCoupons = async () => {
  const result = await couponApi.getMyCoupons()
  console.log(result)
}
```

## 优惠券类型

### 折扣券 (discount)

- 按折扣比例减免
- value 为折扣百分比（如 10 表示 9 折）
- 适用于百分比优惠场景

### 满减券 (amount)

- 按固定金额减免
- value 为减免金额
- 需要设置 min_amount（最小使用金额）

## 最佳实践

### 优惠券设计

- 优惠券名称要吸引人
- 优惠券值要有吸引力
- 设置合理的使用条件

### 优惠券发放

- 定期发放优惠券
- 针对不同用户发放不同优惠券
- 结合活动发放优惠券

### 优惠券核销

- 优化核销流程
- 提供核销提示
- 记录核销数据

### 优惠券统计

- 定期分析优惠券数据
- 优化优惠券策略
- 提高优惠券使用率

## 相关文档

- [商城模块](index.md) - 商城模块总览
- [商品](product.md) - 商品管理文档
- [会员](member.md) - 会员系统文档
