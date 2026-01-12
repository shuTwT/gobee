# 文档库服务模块

## 功能说明

文档库服务模块提供文档库相关的业务逻辑处理，包括文档库的创建、查询、更新和删除等操作。

## 主要功能

### DocLibraryService
- CreateDocLibrary：创建新的文档库
- UpdateDocLibrary：更新指定文档库的信息
- GetDocLibrary：查询指定文档库的详细信息
- GetDocLibraryPage：分页查询文档库列表
- DeleteDocLibrary：删除指定的文档库
- GetDocLibraryList：获取文档库列表（按名称排序）

## 数据访问

所有数据库操作都通过ent ORM进行，确保数据一致性和操作的安全性。

## 错误处理

服务层处理和包装来自底层的错误，并返回给handlers层统一处理。

## 使用说明

在handler中注入DocLibraryService接口，调用相应的方法进行业务操作。

```go
service := doclibrary.NewDocLibraryServiceImpl(client)
library, err := service.CreateDocLibrary(ctx, req)
```
