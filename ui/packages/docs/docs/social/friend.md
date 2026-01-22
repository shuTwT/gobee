# 朋友圈

朋友圈功能允许用户与好友分享生活动态，支持设置可见范围。

## 功能特性

- 发布朋友圈：支持文字和图片
- 查看朋友圈动态：查看好友的朋友圈
- 点赞和评论：对朋友圈进行点赞和评论
- 设置可见范围：控制朋友圈的可见范围

## API 接口

### 获取朋友圈列表

```http
GET /api/v1/social/friend/page
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认 1 |
| page_size | int | 否 | 每页数量，默认 10 |

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
        "content": "分享我的生活",
        "images": ["image1.jpg"],
        "like_count": 10,
        "comment_count": 5,
        "is_liked": false,
        "visibility": "friends",
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

### 发布朋友圈

```http
POST /api/v1/social/friend/create
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| content | string | 是 | 朋友圈内容 |
| images | array | 否 | 图片列表 |
| visibility | string | 否 | 可见范围 (public/friends/private)，默认 friends |

**请求示例：**

```json
{
  "content": "分享我的生活",
  "images": ["image1.jpg"],
  "visibility": "friends"
}
```

### 删除朋友圈

```http
DELETE /api/v1/social/friend/delete/:id
```

### 点赞朋友圈

```http
POST /api/v1/social/friend/like/:id
```

### 评论朋友圈

```http
POST /api/v1/social/friend/comment/:id
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| content | string | 是 | 评论内容 |

## 使用示例

### 发布朋友圈

```javascript
import { friendApi } from '@/api/social/friend'

const publishFriend = async () => {
  const data = {
    content: '分享我的生活',
    images: ['image1.jpg'],
    visibility: 'friends'
  }
  
  const result = await friendApi.create(data)
  console.log(result)
}
```

### 获取朋友圈列表

```javascript
import { friendApi } from '@/api/social/friend'

const getFriends = async () => {
  const params = {
    page: 1,
    page_size: 10
  }
  
  const result = await friendApi.page(params)
  console.log(result)
}
```

### 点赞朋友圈

```javascript
import { friendApi } from '@/api/social/friend'

const likeFriend = async () => {
  const result = await friendApi.like(1)
  console.log(result)
}
```

### 评论朋友圈

```javascript
import { friendApi } from '@/api/social/friend'

const commentFriend = async () => {
  const data = {
    content: '很棒！'
  }
  
  const result = await friendApi.comment(1, data)
  console.log(result)
}
```

## 可见范围

### public

- 所有人可见
- 适合公开分享的内容

### friends

- 仅好友可见
- 适合与好友分享的内容

### private

- 仅自己可见
- 适合私密内容

## 最佳实践

### 内容规范

- 遵守社区规范
- 发布积极向上的内容
- 避免发布敏感内容

### 隐私保护

- 合理设置可见范围
- 注意保护个人隐私
- 避免泄露他人信息

### 互动建议

- 积极点赞和评论
- 尊重他人观点
- 避免恶意评论

## 相关文档

- [社交模块](index.md) - 社交模块总览
- [说说](moment.md) - 说说文档
- [社交登录](login.md) - 社交登录文档
