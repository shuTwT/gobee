# Plugin Service

插件管理模块的服务层，负责插件 CRUD 与上传安装，插件进程的加载、启停与心跳管理委托给基础设施层的 `PluginManager`（`internal/infra/plugin`）。

## 功能

- 插件分页列表 / 详情查询
- 插件上传安装：解析 zip 包内的 `plugin-config.yaml`、校验依赖、解压到 `./data/plugins/{key}`、落库
- 删除插件（运行中先停止）
- 启动 / 停止 / 重启 / 自动启动（委托 `PluginManager`）
- 插件注册 / 心跳 / 心跳超时检查（委托 `PluginManager`，对应公开接口 `/api/v1/public/plugin/register`、`/heartbeat`，仅 debug 模式）

## 分层说明

- 服务层（本包）只做 CRUD 和上传安装，不直接管理插件进程
- 基础设施层 `PluginManager` 持有 go-plugin 客户端、运行实例缓存与心跳注册表，负责：
  - 进程加载（gRPC 握手、Dispense `plugin_store`、Init 注入配置）
  - 启停/重启/自动启动与状态机（stopped/loading/running/error）
  - 插件注册、心跳、30 秒超时检测（`StartHeartbeatChecker` 由 `cmd/server/main.go` 启动）

## 插件包格式

上传的 zip 包必须包含：

- `plugin-config.yaml`（根级）：`key`、`name`、`version`、`magic_cookie_value` 必填；`protocol_version`、`magic_cookie_key`、`description`、`dependencies`、`config`、`auto_start` 可选
- 一个可执行文件（无扩展名或文件名含 `bin`）

插件二进制必须实现 `pkg/plugin/shared` 的 `PluginStore` gRPC 协议（使用 `StoreGRPCPlugin` 与 `GRPCServer.SetImpl`），否则启动握手失败进入 error 状态。

## 使用

```go
pluginManager := plugin_infra.NewPluginManager(db)
pluginService := NewPluginServiceImpl(db, pluginManager)
```
