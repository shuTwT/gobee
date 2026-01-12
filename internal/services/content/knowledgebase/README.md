# 知识库服务模块

## 功能说明

知识库服务模块提供知识库相关的业务逻辑处理，包括知识库的创建、查询、更新和删除等操作。

## 主要功能

### KnowledgeBaseService
- CreateKnowledgeBase：创建新的知识库
- UpdateKnowledgeBase：更新指定知识库的信息
- GetKnowledgeBase：查询指定知识库的详细信息
- GetKnowledgeBasePage：分页查询知识库列表，支持按名称和模型供应商筛选
- DeleteKnowledgeBase：删除指定的知识库
- GetKnowledgeBaseList：获取知识库列表（按名称排序）

## 数据验证

服务层会对输入数据进行以下验证：
- 名称：必填，长度1-100字符
- 模型供应商：必填，支持 openai、anthropic、google、azure、cohere、huggingface、local
- 模型名称：必填，长度1-50字符
- 向量维度：必填，正整数，范围1-10000
- 最大批量处理文档数量：必填，正整数，范围1-1000

## 数据访问

所有数据库操作都通过ent ORM进行，确保数据一致性和操作的安全性。

## 错误处理

服务层处理和包装来自底层的错误，并返回给handlers层统一处理。所有错误都包含详细的错误信息，便于调试和问题追踪。

## 使用说明

在handler中注入KnowledgeBaseService接口，调用相应的方法进行业务操作。

```go
service := knowledgebase.NewKnowledgeBaseServiceImpl(client)
kb, err := service.CreateKnowledgeBase(ctx, req)
```

## 性能优化

- 知识库名称字段已添加唯一索引，确保快速查询
- 分页查询使用 LIMIT 和 OFFSET 优化大数据量场景
- 支持按名称模糊查询和模型供应商精确查询
