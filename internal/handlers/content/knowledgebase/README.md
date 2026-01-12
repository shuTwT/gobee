# 知识库处理器模块

## 功能说明

知识库处理器模块提供知识库相关的HTTP请求处理，包括知识库的创建、查询、更新和删除等操作。

## 主要功能

### KnowledgeBaseHandler
- CreateKnowledgeBase：创建新的知识库
- UpdateKnowledgeBase：更新指定知识库的信息
- GetKnowledgeBasePage：分页查询知识库列表
- GetKnowledgeBase：查询指定知识库的详细信息
- DeleteKnowledgeBase：删除指定的知识库
- GetKnowledgeBaseList：获取知识库列表

## API 接口

### 创建知识库
- 路径：POST /api/v1/knowledgebase/create
- 请求体：KnowledgeBaseCreateReq
- 响应：成功消息

### 更新知识库
- 路径：PUT /api/v1/knowledgebase/update/:id
- 路径参数：id（知识库ID）
- 请求体：KnowledgeBaseUpdateReq
- 响应：成功消息

### 分页查询知识库
- 路径：GET /api/v1/knowledgebase/page
- 查询参数：page（页码）、size（每页数量）、name（名称，可选）、model_provider（模型供应商，可选）
- 响应：分页结果（包含总数和记录列表）

### 查询知识库详情
- 路径：GET /api/v1/knowledgebase/query/:id
- 路径参数：id（知识库ID）
- 响应：知识库详细信息

### 删除知识库
- 路径：DELETE /api/v1/knowledgebase/delete/:id
- 路径参数：id（知识库ID）
- 响应：成功消息

### 获取知识库列表
- 路径：GET /api/v1/knowledgebase/list
- 响应：知识库列表

## 数据验证

处理器层会对请求数据进行验证：
- 路径参数（如id）必须为有效整数
- 请求体必须符合JSON格式
- 必填字段不能为空
- 字段值必须在指定范围内

## 错误处理

所有错误都返回统一的错误格式：
- 状态码：400（请求参数错误）或其他HTTP状态码
- 错误消息：包含详细的错误信息

## 使用说明

在路由中注册处理器：

```go
knowledgeBaseApi := router.Group("/knowledgebase")
{
    knowledgeBaseApi.Post("/create", handlerMap.KnowledgeBaseHandler.CreateKnowledgeBase)
    knowledgeBaseApi.Put("/update/:id", handlerMap.KnowledgeBaseHandler.UpdateKnowledgeBase)
    knowledgeBaseApi.Get("/page", handlerMap.KnowledgeBaseHandler.GetKnowledgeBasePage)
    knowledgeBaseApi.Get("/query/:id", handlerMap.KnowledgeBaseHandler.GetKnowledgeBase)
    knowledgeBaseApi.Delete("/delete/:id", handlerMap.KnowledgeBaseHandler.DeleteKnowledgeBase)
    knowledgeBaseApi.Get("/list", handlerMap.KnowledgeBaseHandler.GetKnowledgeBaseList)
}
```

## 性能要求

- 单条记录CRUD操作响应时间不超过200ms
- 分页查询接口在10万级数据量下响应时间不超过500ms
