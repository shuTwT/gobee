# 商品

商品管理是商城模块的核心功能，提供完整的商品管理能力。

## 功能特性

- 商品创建与编辑：支持富文本编辑器，支持 Markdown 格式
- 商品分类管理：多级分类支持
- 商品上下架：支持商品上下架管理
- 库存管理：实时库存管理
- 商品搜索：支持关键词搜索和筛选

## API 接口

### 获取商品列表

```http
GET /api/v1/mall/product/page
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认 1 |
| page_size | int | 否 | 每页数量，默认 10 |
| category_id | int | 否 | 分类 ID |
| keyword | string | 否 | 搜索关键词 |
| status | string | 否 | 状态 (on_sale/off_sale) |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 100,
    "page": 1,
    "page_size": 10,
    "items": [
      {
        "id": 1,
        "name": "商品名称",
        "price": 100,
        "stock": 100,
        "description": "商品描述",
        "category_id": 1,
        "category_name": "分类名称",
        "status": "on_sale",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

### 创建商品

```http
POST /api/v1/mall/product/create
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 是 | 商品名称 |
| price | decimal | 是 | 商品价格 |
| stock | int | 是 | 库存数量 |
| description | string | 否 | 商品描述 |
| category_id | int | 否 | 分类 ID |
| images | array | 否 | 商品图片列表 |
| status | string | 否 | 状态 (on_sale/off_sale)，默认 on_sale |

**请求示例：**

```json
{
  "name": "商品名称",
  "price": 100,
  "stock": 100,
  "description": "商品描述",
  "category_id": 1,
  "images": ["image1.jpg", "image2.jpg"],
  "status": "on_sale"
}
```

### 更新商品

```http
PUT /api/v1/mall/product/update/:id
```

**请求参数：** 同创建商品

### 删除商品

```http
DELETE /api/v1/mall/product/delete/:id
```

### 查询商品详情

```http
GET /api/v1/mall/product/query/:id
```

## 商品分类管理

### 获取分类列表

```http
GET /api/v1/mall/product/category/list
```

### 创建分类

```http
POST /api/v1/mall/product/category/create
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| name | string | 是 | 分类名称 |
| parent_id | int | 否 | 父分类 ID |
| description | string | 否 | 分类描述 |
| sort | int | 否 | 排序 |

## 使用示例

### 创建商品

```javascript
import { productApi } from '@/api/mall/product'

const createProduct = async () => {
  const data = {
    name: '商品名称',
    price: 100,
    stock: 100,
    description: '商品描述',
    category_id: 1,
    images: ['image1.jpg', 'image2.jpg']
  }
  
  const result = await productApi.create(data)
  console.log(result)
}
```

### 获取商品列表

```javascript
import { productApi } from '@/api/mall/product'

const getProducts = async () => {
  const params = {
    page: 1,
    page_size: 10,
    status: 'on_sale'
  }
  
  const result = await productApi.page(params)
  console.log(result)
}
```

### 更新商品

```javascript
import { productApi } from '@/api/mall/product'

const updateProduct = async () => {
  const data = {
    id: 1,
    name: '更新后的商品名称',
    price: 120,
    stock: 80
  }
  
  const result = await productApi.update(data.id, data)
  console.log(result)
}
```

## 最佳实践

### 商品命名

- 使用简洁明了的商品名称
- 包含关键信息（品牌、型号等）
- 避免使用过于复杂的名称

### 商品描述

- 详细描述商品特性
- 使用清晰的段落结构
- 添加高质量的图片

### 库存管理

- 定期更新库存数量
- 设置库存预警
- 及时补货

### 商品分类

- 合理设置商品分类
- 避免分类过于复杂
- 定期优化分类结构

## 相关文档

- [商城模块](index.md) - 商城模块总览
- [会员](member.md) - 会员系统文档
- [优惠券](coupon.md) - 优惠券文档
