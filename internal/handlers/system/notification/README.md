# Notification Handler

通知管理模块的处理器层，负责处理HTTP请求和响应。

## 功能

- 查询通知分页列表（支持按已读/未读状态过滤）
- 查询通知详情
- 删除通知
- 批量标记通知为已读

## 接口

### 查询通知分页列表
- 路径: `GET /api/v1/notifications/page`
- 参数:
  - page: 页码
  - size: 每页数量
  - is_read: 是否已读（可选）

### 查询通知详情
- 路径: `GET /api/v1/notifications/query/:id`
- 参数:
  - id: 通知ID

### 删除通知
- 路径: `DELETE /api/v1/notifications/delete/:id`
- 参数:
  - id: 通知ID

### 批量标记为已读
- 路径: `POST /api/v1/notifications/batch/read`
- 参数:
  - ids: 通知ID列表
