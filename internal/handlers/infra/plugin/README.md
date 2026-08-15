# Plugin Handler

插件管理模块的处理器层，负责处理 HTTP 请求和响应，业务逻辑委托给服务层 `PluginService`。

## 功能

- 上传安装插件（zip 包）
- 查询插件分页列表
- 查询单个插件详情
- 删除插件
- 启动 / 停止 / 重启插件

## 接口

### 创建插件
- 路径: `POST /api/v1/plugin/create`
- 请求体: `multipart/form-data`，字段 `file` 为插件 zip 包（含 `plugin-config.yaml` 与可执行文件）
- 响应: 创建成功的插件详情

### 查询插件分页列表
- 路径: `GET /api/v1/plugin/page`
- 参数:
  - page: 页码
  - size: 每页数量
  - name: 插件名称（模糊）
  - key: 插件标识（模糊）
  - status: 状态（stopped/running/error/loading）
  - enabled: 是否启用
  - auto_start: 是否自动启动
- 响应: 分页的插件列表

### 查询单个插件
- 路径: `GET /api/v1/plugin/query/:id`
- 参数:
  - id: 插件ID
- 响应: 插件详情

### 删除插件
- 路径: `DELETE /api/v1/plugin/delete/:id`
- 参数:
  - id: 插件ID
- 响应: 删除成功提示（运行中的插件会先被停止）

### 启动插件
- 路径: `POST /api/v1/plugin/:id/start`
- 参数:
  - id: 插件ID
- 响应: 启动成功提示（异步拉起插件进程，状态先置 loading，成功后 running，失败置 error）

### 停止插件
- 路径: `POST /api/v1/plugin/:id/stop`
- 参数:
  - id: 插件ID
- 响应: 停止成功提示

### 重启插件
- 路径: `POST /api/v1/plugin/:id/restart`
- 参数:
  - id: 插件ID
- 响应: 重启成功提示

## 错误处理

- 请求参数解析失败：返回 400 错误
- 插件 ID 无效：返回 400 错误
- 服务层操作失败：返回 500 错误，包含具体原因（如依赖插件未运行、二进制不存在等）

## 日志记录

使用结构化日志记录关键操作：

- 创建插件：记录插件 ID 和 key
- 启动/停止/重启/删除：记录插件 ID
- 查询失败：记录错误详情和上下文信息
