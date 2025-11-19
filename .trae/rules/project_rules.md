# 项目开发规范

## 1. 包命名规范 (Go)

- **简洁明了**: 包名应为小写，简短且具有描述性。
- **不使用下划线或驼峰命名**: 例如，使用 `apihandler` 而不是 `api_handler` 或 `apiHandler`。
- **与目录名一致**: 包名应与其所在目录的名称相匹配。
- **避免使用通用名称**: 避免使用 `util`, `common`, `base` 等过于宽泛的名称。如果需要，请创建更具描述性的包，例如 `strutil`。

## 2. 模块命名规范

### Go

- **文件名**: 使用小写蛇形命名法 (snake_case)，例如 `user_handler.go`。
- **接口**: 接口名称以 `er` 结尾，例如 `type Reader interface { Read(p []byte) (n int, err error) }`。
- **结构体**: 使用驼峰命名法 (CamelCase)，并根据其用途添加后缀，例如 `UserHandler`, `UserService`。

### Vue.js

- **组件文件名**: 使用大驼峰命名法 (PascalCase)，例如 `MyComponent.vue`。
- **TypeScript 文件**: 使用小驼峰命名法 (camelCase)，例如 `userService.ts`。

## 3. 代码结构

### Go

- **`cmd/`**: 应用程序的主入口。
- **`internal/`**: 项目内部的私有代码，外部无法导入。
    - **`handlers/`**: HTTP 请求处理器。
    - **`services/`**: 业务逻辑。
    - **`database/`**: 数据库初始化和连接。
    - **`middleware/`**: 中间件。
    - **`router/`**: 路由定义。
- **`pkg/`**: 可以被外部应用程序引用的代码。
- **`ent/`**: `ent` ORM 生成的代码和 schema 定义。
- **`config/`**: 配置文件。

### Vue.js (`ui/`)

- **`src/`**: 主要源代码。
    - **`api/`**: API 请求。
    - **`assets/`**: 静态资源，如 CSS 和图片。
    - **`components/`**: 可重用组件。
    - **`layout/`**: 页面布局。
    - **`router/`**: 路由配置。
    - **`stores/`**: Pinia 状态管理。
    - **`utils/`**: 工具函数。
    - **`views/`**: 页面级组件。

## 4. 文件结构

- **保持内聚性**: 相关的功能应该放在同一个包或目录中。
- **按功能组织**: 在 `internal/handlers` 和 `internal/services` 中，按功能模块创建子目录。

## 5. 错误处理

### Go

- **总是处理错误**: 不要忽略函数返回的 `error`。
- **错误信息要清晰**: 错误信息应该提供足够的上下文。
- **自定义错误类型**: 在 `pkg/domain/model/error.go` 中定义业务错误，方便统一处理。
- **在业务逻辑层处理错误**: `services` 层应该处理和包装来自底层的错误，并返回给 `handlers` 层。

### Vue.js

- **API 请求错误**: 在 `utils/http.ts` 中统一处理 API 请求错误，例如 401, 403, 404, 500 等。
- **组件内部错误**: 使用 `try...catch` 或 Vue 的 `onErrorCaptured` 生命周期钩子来捕获组件内部的错误。

## 6. HTTP 处理

- **统一的响应结构**: 在 `pkg/domain/model/success.go` 和 `pkg/domain/model/error.go` 中定义统一的成功和失败响应结构。
- **RESTful API**: 遵循 RESTful 设计原则。
- **参数校验**: 在 `handlers` 层对传入的参数进行校验。

## 7. 数据库操作

- **使用 `ent`**: 所有数据库操作都应通过 `ent` ORM 进行。
- **事务**: 对于需要多个数据库操作的业务逻辑，使用事务来保证数据一致性。
- **避免裸露 SQL**: 尽可能使用 `ent` 的 API，避免手写 SQL。

## 8. 安全规范

- **认证与授权**: 使用中间件进行用户认证和权限校验。
- **防止 SQL 注入**: `ent` 默认可以防止 SQL 注入，但仍需注意。
- **XSS 防护**: Vue.js 默认会对输出进行转义，但对于 `v-html` 等指令需要特别小心。
- **CSRF 防护**: 确保您的框架或库已启用 CSRF 防护。
- **不硬编码敏感信息**: 不要将密码、API 密钥等硬编码在代码中，应使用环境变量或配置文件。

## 9. 日志规范

- **使用结构化日志**: 使用 `slog` 或 `zap` 等库进行结构化日志记录。
- **日志级别**:
    - **Debug**: 用于开发过程中的调试信息。
    - **Info**: 记录关键操作和流程信息。
    - **Warn**: 记录可预见的、但需要注意的问题。
    - **Error**: 记录影响功能正常运行的错误。
- **日志内容**: 日志应包含时间戳、日志级别、消息以及相关的上下文信息（如 request ID, user ID）。