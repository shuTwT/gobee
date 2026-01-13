# ScheduleJob Service

定时任务管理模块的服务层，负责业务逻辑处理和数据库操作。

## 功能

- 创建定时任务
- 查询定时任务分页列表
- 查询单个定时任务详情
- 更新定时任务配置
- 删除定时任务

## 数据模型

定时任务包含以下主要字段：

- **name**: 任务名称
- **type**: 任务类型（cron, interval, once）
- **expression**: 调度表达式（cron表达式或时间间隔）
- **description**: 任务描述（可选）
- **enabled**: 是否启用
- **next_run_time**: 下次执行时间
- **last_run_time**: 上次执行时间
- **execution_type**: 执行类型（http, internal, command, mq）
- **http_method**: HTTP方法（GET, POST, PUT, DELETE）
- **http_url**: HTTP URL
- **http_headers**: HTTP请求头
- **http_body**: HTTP请求体
- **http_timeout**: HTTP超时时间（秒）
- **max_retries**: 最大重试次数
- **failure_notification**: 失败是否通知

## 方法

### CreateScheduleJob
创建新的定时任务。

**参数验证**：
- 任务名称不能为空
- 任务类型不能为空（必须为 cron, interval, once 之一）
- 调度表达式不能为空
- 执行类型不能为空（必须为 http, internal, command, mq 之一）

### ListScheduleJobPage
查询定时任务分页列表，按ID降序排列。

### QueryScheduleJob
根据ID查询单个定时任务的详细信息。

### UpdateScheduleJob
更新定时任务配置，支持部分更新。

### DeleteScheduleJob
根据ID删除定时任务。

## 任务类型说明

### Cron 任务
使用cron表达式定义执行时间，例如：`0 0 * * *` 表示每天午夜执行。

### Interval 任务
使用时间间隔定义执行频率，例如：`5m` 表示每5分钟执行一次。

### Once 任务
一次性任务，在指定时间执行一次。

## 执行类型说明

### HTTP 执行
通过HTTP请求触发任务执行，支持自定义请求头、请求体和超时时间。

### Internal 执行
调用内部服务方法执行任务（待实现）。

### Command 执行
执行系统命令或脚本（待实现）。

### MQ 执行
将任务发布到消息队列，由消费者异步处理（待实现）。
