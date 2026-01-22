# 会员

会员系统是商城模块的重要组成部分，提供完整的会员管理能力。

## 功能特性

- 会员等级管理：创建和管理会员等级
- 会员权益管理：设置不同等级的会员权益
- 积分系统：会员积分获取和使用
- 会员统计：会员数据统计和分析

## API 接口

### 获取会员等级列表

```http
GET /api/v1/mall/member/level/list
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
        "name": "普通会员",
        "discount": 1.0,
        "points": 0,
        "privileges": [],
        "sort": 1
      },
      {
        "id": 2,
        "name": "黄金会员",
        "discount": 0.9,
        "points": 1000,
        "privileges": ["免运费", "专属客服"],
        "sort": 2
      }
    ]
  }
}
```

### 创建会员等级

```http
POST /api/v1/mall/member/level/create
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 是 | 会员等级名称 |
| discount | decimal | 是 | 折扣比例 (0-1) |
| points | int | 是 | 所需积分 |
| privileges | array | 否 | 会员权益列表 |
| sort | int | 否 | 排序 |

**请求示例：**

```json
{
  "name": "黄金会员",
  "discount": 0.9,
  "points": 1000,
  "privileges": ["免运费", "专属客服"],
  "sort": 2
}
```

### 更新会员等级

```http
PUT /api/v1/mall/member/level/update/:id
```

**请求参数：** 同创建会员等级

### 删除会员等级

```http
DELETE /api/v1/mall/member/level/delete/:id
```

### 获取会员信息

```http
GET /api/v1/mall/member/info
```

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "user_id": 1,
    "level_id": 2,
    "level_name": "黄金会员",
    "points": 1500,
    "total_points": 5000,
    "discount": 0.9,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### 增加积分

```http
POST /api/v1/mall/member/points/add
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| user_id | int | 是 | 用户 ID |
| points | int | 是 | 积分数量 |
| reason | string | 否 | 原因 |

### 扣除积分

```http
POST /api/v1/mall/member/points/deduct
```

**请求参数：** 同增加积分

## 使用示例

### 创建会员等级

```javascript
import { memberApi } from '@/api/mall/member'

const createMemberLevel = async () => {
  const data = {
    name: '黄金会员',
    discount: 0.9,
    points: 1000,
    privileges: ['免运费', '专属客服'],
    sort: 2
  }
  
  const result = await memberApi.createLevel(data)
  console.log(result)
}
```

### 获取会员信息

```javascript
import { memberApi } from '@/api/mall/member'

const getMemberInfo = async () => {
  const result = await memberApi.getInfo()
  console.log(result)
}
```

### 增加积分

```javascript
import { memberApi } from '@/api/mall/member'

const addPoints = async () => {
  const data = {
    user_id: 1,
    points: 100,
    reason: '购物奖励'
  }
  
  const result = await memberApi.addPoints(data)
  console.log(result)
}
```

## 最佳实践

### 会员等级设计

- 合理设置会员等级数量（建议 3-5 个）
- 等级名称要简洁明了
- 积分要求要合理递增

### 会员权益设置

- 权益要有吸引力
- 不同等级要有明显差异
- 定期更新权益内容

### 积分系统

- 积分获取规则要清晰
- 积分使用场景要丰富
- 定期举办积分活动

### 会员运营

- 定期分析会员数据
- 优化会员等级体系
- 提供个性化服务

## 相关文档

- [商城模块](index.md) - 商城模块总览
- [商品](product.md) - 商品管理文档
- [优惠券](coupon.md) - 优惠券文档
