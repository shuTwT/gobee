1.项目使用框架版本:gofiber v2.52.9,ent orm v0.14.5
2.包名应为小写。例如，使用 `apihandler` 而不是 `api_handler` 或 `apiHandler`
3.避免使用 `util`, `common`, `base` 等过于宽泛的名称。如果需要，请创建更具描述性的包，例如 `strutil`。
4.ui文件夹下为前端项目。vue 组件采用大驼峰命名，ts 文件采用小驼峰命名。

5.在 `pkg/domain/model/error.go` 中定义业务错误，方便统一处理。`services` 层应该处理和包装来自底层的错误，并返回给 `handlers` 层。

6. 所有数据库操作都应通过 `ent` ORM 进行。 对于需要多个数据库操作的业务逻辑，使用事务来保证数据一致性。

7. 使用 `slog` 或 `zap` 等库进行结构化日志记录。日志应包含时间戳、日志级别、消息以及相关的上下文信息（如 request ID, user ID）
