工作流程
系统启动
    ↓
初始化 ScheduleManager
    ↓
添加任务到缓存（friendCircle）
    ↓
从数据库加载启用的任务
    ↓
添加到 gocron 调度器
    ↓
启动调度器
    ↓
任务按配置自动执行

用户操作（添加/更新/删除）
    ↓
Service层处理
    ↓
数据库操作
    ↓
ScheduleManager 同步调度器
    ↓
实时生效（无需重启）