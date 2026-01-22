# 说说

说说功能允许用户发布和分享动态，支持文字、图片等内容。

## 功能特性

- 发布说说：支持文字和图片
- 查看说说列表：查看自己和他人的说说
- 点赞和评论：对说说进行点赞和评论
- 删除说说：删除自己发布的说说

## API 接口

### 获取说说列表

```http
GET /api/v1/social/moment/page
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认 1 |
| page_size | int | 否 | 每页数量，默认 10 |
| user_id | int | 否 | 用户 ID |

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
        "user_id": 1,
        "user_name": "用户名",
        "user_avatar": "avatar.jpg",
        "content": "今天天气真好！",
        "images": ["image1.jpg", "image2.jpg"],
        "like_count": 10,
        "comment_count": 5,
        "is_liked": false,
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

### 发布说说

```http
POST /api/v1/social/moment/create
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| content | string | 是 | 说说内容 |
| images | array | 否 | 图片列表 |

**请求示例：**

```json
{
  "content": "今天天气真好！",
  "images": ["image1.jpg", "image2.jpg"]
}
```

### 删除说说

```http
DELETE /api/v1/social/moment/delete/:id
```

### 点赞说说

```http
POST /api/v1/social/moment/like/:id
```

### 评论说说

```http
POST /api/v1/social/moment/comment/:id
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| content | string | 是 | 评论内容 |

## 使用示例

### 发布说说

```javascript
import { momentApi } from '@/api/social/moment'

const publishMoment = async () => {
  const data = {
    content: '今天天气真好！',
    images: ['image1.jpg', 'image2.jpg']
  }
  
  const result = await momentApi.create(data)
  console.log(result)
}
```

### 获取说说列表

```javascript
import { momentApi } from '@/api/social/moment'

const getMoments = async () => {
  const params = {
    page: 1,
    page_size: 10
  }
  
  const result = await momentApi.page(params)
  console.log(result)
}
```

### 点赞说说

```javascript
import { momentApi } from '@/api/social/moment'

const likeMoment = async () => {
  const result = await momentApi.like(1)
  console.log(result)
}
```

### 评论说说

```javascript
import { momentApi } from '@/api/social/moment'

const commentMoment = async () => {
  const data = {
    content: '我也这么觉得！'
  }
  
  const result = await momentApi.comment(1, data)
  console.log(result)
}
```

## 最佳实践

### 内容规范

- 遵守社区规范
- 发布积极向上的内容
- 避免发布敏感内容

### 图片使用

- 使用高质量的图片
- 控制图片数量（建议不超过 9 张）
- 注意图片大小

### 互动建议

- 积极点赞和评论
- 尊重他人观点
- 避免恶意评论

## 相关文档

- [社交模块](index.md) - 社交模块总览
- [朋友圈](friend.md) - 朋友圈文档
- [社交登录](login.md) - 社交登录文档
