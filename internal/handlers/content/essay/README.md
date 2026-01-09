# Essay Handler

## 功能说明

说说(即刻短文)模块的请求处理器，负责处理与说说相关的HTTP请求。

## 接口列表

### 创建说说
- **路径**: `POST /api/v1/essay/create`
- **描述**: 创建一个新说说
- **请求体**: `model.EssayCreateReq`
  - `content`: string - 说说内容
  - `draft`: bool - 是否为草稿
  - `images`: []string - 图片列表

### 更新说说
- **路径**: `PUT /api/v1/essay/update/:id`
- **描述**: 更新指定说说的信息
- **路径参数**: `id` - 说说ID
- **请求体**: `model.EssayUpdateReq`
  - `id`: int - 说说ID
  - `content`: string - 说说内容
  - `draft`: bool - 是否为草稿
  - `images`: []string - 图片列表

### 获取说说分页列表
- **路径**: `GET /api/v1/essay/page`
- **描述**: 获取说说分页列表
- **查询参数**:
  - `page`: int - 页码，默认1
  - `size`: int - 每页数量，默认10

### 删除说说
- **路径**: `DELETE /api/v1/essay/delete/:id`
- **描述**: 删除指定说说
- **路径参数**: `id` - 说说ID

## 使用方法

```go
// 在路由中注册
essayApi := router.Group("/essay")
{
    essayApi.Get("/page", handlerMap.EssayHandler.GetEssayPage)
    essayApi.Post("/create", handlerMap.EssayHandler.CreateEssay)
    essayApi.Put("/update/:id", handlerMap.EssayHandler.UpdateEssay)
    essayApi.Delete("/delete/:id", handlerMap.EssayHandler.DeleteEssay)
}
```
