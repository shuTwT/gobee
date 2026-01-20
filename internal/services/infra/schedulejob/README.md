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
- **type**: 任务类型（cron, interval）
- **expression**: 调度表达式（cron表达式或时间间隔）
- **description**: 任务描述（可选）
- **enabled**: 是否启用
- **next_run_time**: 下次执行时间
- **last_run_time**: 上次执行时间
- **job_name**: 内部任务名称
- **max_retries**: 最大重试次数
- **failure_notification**: 失败是否通知

## 方法

### CreateScheduleJob
创建新的定时任务。

**参数验证**：
- 任务名称不能为空
- 任务类型不能为空（必须为 cron, interval 之一）
- 调度表达式不能为空
- 内部任务名称不能为空

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
使用时间间隔定义执行频率，例如：`24h` 表示每24小时执行一次。

## 内部任务说明

定时任务通过内部任务名称（job_name）来指定要执行的任务。系统启动时会从数据库加载所有启用的定时任务，并根据任务名称找到对应的任务实现并注册到gocron调度器中。

### 可用的内部任务

- **friendCircle**: 朋友圈爬虫任务，用于抓取友链的RSS更新

### 添加新的内部任务

1. 在 `internal/job/` 目录下创建新的任务包
2. 实现任务结构体，满足 `schedule_model.DurationJob` 或 `schedule_model.CronJob` 接口
3. 在 `internal/schedule/init.go` 的 `InitializeSchedule` 函数中将任务添加到 `jobCache` 中
4. 在前端表单的 `jobNameOptions` 中添加对应的选项
