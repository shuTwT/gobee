# 文档详情管理模块

## 功能说明

文档详情管理模块用于管理文档库中的具体文档内容，支持树形结构展示，提供文档的创建、查询、更新和删除功能。

## 主要功能

### 文档详情管理
- 创建文档：支持创建新的文档详情，包括标题、版本、内容、父文档ID等
- 查询文档：支持分页查询文档详情列表
- 更新文档：支持更新文档的详细信息
- 删除文档：支持删除指定的文档
- 树形展示：支持以树形结构展示文档层级关系

### 树形结构
- 通过parent_id字段构建文档的父子关系
- 支持无限层级的文档嵌套
- 树形展示便于查看文档的组织结构

## 数据结构

### DocLibraryDetail
- id：文档ID
- library_id：所属文档库ID
- title：文档标题
- version：文档版本（可选）
- content：文档内容
- parent_id：父文档ID（可选，用于构建树形结构）
- path：文档路径（可选）
- url：文档URL（可选）
- language：文档语言（默认为zh）
- created_at：创建时间
- updated_at：更新时间

## API接口

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 创建文档 | POST | /api/v1/doclibrarydetail/create | 创建新的文档详情 |
| 更新文档 | PUT | /api/v1/doclibrarydetail/update/:id | 更新指定文档详情 |
| 查询文档 | GET | /api/v1/doclibrarydetail/query/:id | 查询指定文档详情 |
| 分页查询 | GET | /api/v1/doclibrarydetail/page | 分页查询文档详情列表 |
| 删除文档 | DELETE | /api/v1/doclibrarydetail/delete/:id | 删除指定文档详情 |
| 获取树形结构 | GET | /api/v1/doclibrarydetail/tree | 获取文档树形结构 |

## 使用说明

1. 在前端页面"文档详情管理"中，首先选择要管理的文档库
2. 点击"新建文档"按钮创建新的文档
3. 填写文档信息，包括标题、版本、内容等
4. 可以通过设置parent_id来指定父文档，构建文档层级关系
5. 文档以树形结构展示，可以展开/折叠查看子文档
6. 可以在树形结构中直接编辑或删除文档

## 树形结构说明

- 根文档：parent_id为null或0的文档
- 子文档：parent_id指向父文档ID的文档
- 树形展示：自动根据parent_id构建文档树，支持无限层级嵌套

## 扩展性

该模块支持多语言文档管理，预留了path和url字段，便于后续扩展文档路径管理和外部链接功能。
