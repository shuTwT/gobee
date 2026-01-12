# 文档详情服务模块

## 功能说明

文档详情服务模块提供文档详情相关的业务逻辑处理，包括文档的创建、查询、更新、删除和树形结构构建等操作。

## 主要功能

### DocLibraryDetailService
- CreateDocLibraryDetail：创建新的文档详情
- UpdateDocLibraryDetail：更新指定文档详情的信息
- GetDocLibraryDetail：查询指定文档详情的详细信息
- GetDocLibraryDetailPage：分页查询文档详情列表（支持按文档库筛选）
- DeleteDocLibraryDetail：删除指定的文档详情
- GetDocLibraryDetailTree：获取文档树形结构（支持按文档库筛选）

## 树形结构构建

树形结构通过parent_id字段构建，服务层提供GetDocLibraryDetailTree方法获取所有文档，handler层负责构建树形结构。

## 数据访问

所有数据库操作都通过ent ORM进行，确保数据一致性和操作的安全性。

## 错误处理

服务层处理和包装来自底层的错误，并返回给handlers层统一处理。

## 使用说明

在handler中注入DocLibraryDetailService接口，调用相应的方法进行业务操作。

```go
service := doclibrarydetail.NewDocLibraryDetailServiceImpl(client)
detail, err := service.CreateDocLibraryDetail(ctx, req)
```

## 树形结构说明

- 根文档：parent_id为nil的文档
- 子文档：parent_id指向父文档ID的文档
- 树形构建：handler层的buildTree函数负责将扁平数据转换为树形结构
