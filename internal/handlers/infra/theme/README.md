# 主题管理 Handler（infra/theme）

主题管理的 HTTP 接口层，负责参数解析与响应组装，业务逻辑在 `services/infra/theme`。

## 接口一览

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/theme/upload` | 上传主题 zip 包，返回临时 `file_path` |
| POST | `/api/v1/theme/create` | 创建主题（internal/static 传 `file_path`，external 传外部地址） |
| GET | `/api/v1/theme/page` | 分页查询主题列表 |
| GET | `/api/v1/theme/query/:id` | 查询单个主题 |
| DELETE | `/api/v1/theme/delete/:id` | 删除主题（已启用主题禁止删除） |
| POST | `/api/v1/theme/:id/enable` | 启用主题（全局互斥，自动禁用其他主题） |
| POST | `/api/v1/theme/:id/disable` | 禁用主题 |
| GET | `/api/v1/theme/:id/setting` | 获取主题设置表单定义与当前配置值 |
| POST | `/api/v1/theme/:id/setting/save` | 保存主题配置值（body 为 `{"values": {...}}`） |

## 主题设置接口说明

`GET /:id/setting` 返回 `{forms, values}`：`forms` 来自主题包 `setting.yaml` 的实时解析（外部主题返回空数组），`values` 为 setting 表已存值并回填 schema 默认值。

`POST /:id/setting/save` 保存时按当前 schema 过滤未知字段并校验必填项，值以 JSON 存入 setting 表（key 为 `theme:{主题name}`）。
