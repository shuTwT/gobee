# 社交登录

社交登录功能支持用户使用第三方社交平台账号登录系统。

## 功能特性

- 微信登录：支持微信扫码登录和授权登录
- QQ 登录：支持 QQ 授权登录
- 微博登录：支持微博授权登录
- 绑定账号：支持绑定多个社交账号

## 配置说明

在 `config.yaml` 中配置社交登录参数：

```yaml
social:
  wechat:
    app_id: your_wechat_app_id
    app_secret: your_wechat_app_secret
    redirect_uri: your_redirect_uri
  qq:
    app_id: your_qq_app_id
    app_secret: your_qq_app_secret
    redirect_uri: your_redirect_uri
  weibo:
    app_id: your_weibo_app_id
    app_secret: your_weibo_app_secret
    redirect_uri: your_redirect_uri
```

## API 接口

### 获取登录授权链接

```http
GET /api/v1/social/login/url
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| provider | string | 是 | 登录方式 (wechat/qq/weibo) |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "auth_url": "https://open.weixin.qq.com/connect/qrconnect?xxx"
  }
}
```

### 社交登录

```http
POST /api/v1/social/login
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| provider | string | 是 | 登录方式 (wechat/qq/weibo) |
| code | string | 是 | 授权码 |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "jwt_token",
    "user": {
      "id": 1,
      "username": "用户名",
      "avatar": "avatar.jpg"
    }
  }
}
```

### 绑定社交账号

```http
POST /api/v1/social/bind
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| provider | string | 是 | 登录方式 (wechat/qq/weibo) |
| code | string | 是 | 授权码 |

### 解绑社交账号

```http
POST /api/v1/social/unbind
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| provider | string | 是 | 登录方式 (wechat/qq/weibo) |

## 使用示例

### 获取微信登录授权链接

```javascript
import { loginApi } from '@/api/social/login'

const getWechatAuthUrl = async () => {
  const params = {
    provider: 'wechat'
  }
  
  const result = await loginApi.getAuthUrl(params)
  window.location.href = result.auth_url
}
```

### 微信登录

```javascript
import { loginApi } from '@/api/social/login'

const wechatLogin = async (code) => {
  const data = {
    provider: 'wechat',
    code: code
  }
  
  const result = await loginApi.social(data)
  console.log(result)
}
```

### 绑定 QQ 账号

```javascript
import { loginApi } from '@/api/social/login'

const bindQq = async (code) => {
  const data = {
    provider: 'qq',
    code: code
  }
  
  const result = await loginApi.bind(data)
  console.log(result)
}
```

### 解绑微博账号

```javascript
import { loginApi } from '@/api/social/login'

const unbindWeibo = async () => {
  const data = {
    provider: 'weibo'
  }
  
  const result = await loginApi.unbind(data)
  console.log(result)
}
```

## 登录流程

### 微信登录流程

1. 获取微信登录授权链接
2. 用户扫码授权
3. 获取授权码
4. 使用授权码登录
5. 返回用户信息和 token

### QQ 登录流程

1. 获取 QQ 登录授权链接
2. 用户点击授权
3. 获取授权码
4. 使用授权码登录
5. 返回用户信息和 token

### 微博登录流程

1. 获取微博登录授权链接
2. 用户点击授权
3. 获取授权码
4. 使用授权码登录
5. 返回用户信息和 token

## 最佳实践

### 安全建议

- 使用 HTTPS 协议
- 验证授权码的有效性
- 设置合理的 token 过期时间
- 定期更新 app_secret

### 用户体验

- 提供清晰的登录指引
- 优化授权流程
- 提供登录状态提示
- 支持记住登录状态

### 错误处理

- 处理授权失败的情况
- 提供友好的错误提示
- 记录登录日志
- 支持重试机制

## 相关文档

- [社交模块](index.md) - 社交模块总览
- [说说](moment.md) - 说说文档
- [朋友圈](friend.md) - 朋友圈文档
