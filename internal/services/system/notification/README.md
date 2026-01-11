# Notification Service

通知管理模块的服务层，负责业务逻辑处理和数据库操作。

## 功能

- 查询通知分页列表（支持按已读/未读状态过滤）
- 查询通知详情
- 删除通知
- 批量标记通知为已读

## 方法

### QueryNotificationPage
查询通知分页列表，支持按已读/未读状态过滤。

### QueryNotification
根据ID查询单个通知的详细信息。

### DeleteNotification
根据ID删除通知。

### BatchMarkAsRead
批量将通知标记为已读。
