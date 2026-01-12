# 文档库管理模块

## 功能说明

文档库管理模块用于管理需要定期爬取的文档库信息，提供文档库的创建、查询、更新和删除功能。

## 主要功能

### 文档库管理
- 创建文档库：支持创建新的文档库，包括名称、别名、描述、来源和URL
- 查询文档库：支持分页查询文档库列表
- 更新文档库：支持更新文档库的基本信息
- 删除文档库：支持删除指定的文档库

### 文档来源类型
- Git：Git仓库类型的文档库
- OpenAPI：OpenAPI规范的文档库
- LLMs Txt：LLMs文本格式的文档库
- Website：网站类型的文档库

## 数据结构

### DocLibrary
- id：文档库ID
- name：文档库名称
- alias：文档库别名
- description：文档库描述（可选）
- source：文档库来源（git/openapi/llms_txt/website）
- url：文档库URL
- created_at：创建时间
- updated_at：更新时间

## API接口

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 创建文档库 | POST | /api/v1/doclibrary/create | 创建新的文档库 |
| 更新文档库 | PUT | /api/v1/doclibrary/update/:id | 更新指定文档库 |
| 查询文档库 | GET | /api/v1/doclibrary/query/:id | 查询指定文档库详情 |
| 分页查询 | GET | /api/v1/doclibrary/page | 分页查询文档库列表 |
| 删除文档库 | DELETE | /api/v1/doclibrary/delete/:id | 删除指定文档库 |
| 获取列表 | GET | /api/v1/doclibrary/list | 获取文档库列表 |

## 使用说明

1. 在前端页面"文档库管理"中，点击"新建文档库"按钮创建新的文档库
2. 填写文档库信息，包括名称、别名、来源和URL等
3. 点击"保存"按钮创建文档库
4. 可以在列表中查看、编辑或删除文档库

## 扩展性

该模块预留了爬虫接口及数据存储结构，便于后续扩展爬虫功能。
