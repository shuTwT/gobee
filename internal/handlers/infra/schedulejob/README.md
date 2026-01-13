# ScheduleJob Handler

定时任务管理模块的处理器层，负责处理HTTP请求和响应。

## 功能

- 创建定时任务
- 查询定时任务分页列表
- 查询单个定时任务详情
- 更新定时任务配置
- 删除定时任务

## 接口

### 创建定时任务
- 路径: `POST /api/v1/infrastructure/schedule-job/create`
- 请求体:
  ```json
  {
    "name": "任务名称",
    "type": "cron",
    "expression": "0 0 * * *",
    "description": "任务描述",
    "enabled": true,
    "execution_type": "http",
    "http_method": "POST",
    "http_url": "http://example.com/api/endpoint",
    "http_headers": {
      "Authorization": "Bearer token"
    },
    "http_body": "{}",
    "http_timeout": 30,
    "max_retries": 3,
    "failure_notification": false
  }
  ```
- 响应: 创建成功的定时任务详情

### 查询定时任务分页列表
- 路径: `GET /api/v1/infrastructure/schedule-job/page`
- 参数:
  - page: 页码
  - size: 每页数量
- 响应: 分页的定时任务列表

### 查询单个定时任务
- 路径: `GET /api/v1/infrastructure/schedule-job/query/:id`
- 参数:
  - id: 定时任务ID
- 响应: 定时任务详情

### 更新定时任务
- 路径: `PUT /api/v1/infrastructure/schedule-job/update/:id`
- 参数:
  - id: 定时任务ID
- 请求体: 支持部分更新，所有字段都是可选的
- 响应: 更新后的定时任务详情

### 删除定时任务
- 路径: `DELETE /api/v1/infrastructure/schedule-job/delete/:id`
- 参数:
  - id: 定时任务ID
- 响应: 删除成功提示

## 错误处理

所有接口都包含完善的错误处理和日志记录：

- 请求参数解析失败：返回 400 错误
- 数据库操作失败：返回 500 错误
- 所有错误都会记录详细的日志信息，包括错误上下文

## 日志记录

使用结构化日志记录关键操作：

- 创建任务：记录任务ID和名称
- 更新任务：记录任务ID和名称
- 删除任务：记录任务ID
- 查询失败：记录错误详情
- 操作失败：记录错误详情和上下文信息
