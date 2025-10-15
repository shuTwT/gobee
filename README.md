# GoBee - 现代化内容管理系统

GoBee 是一个基于 Go 语言和 Fiber 框架开发的现代化内容管理系统（CMS），支持内容模型管理、用户管理、支付系统等功能。

## 🌟 主要特性

- **现代化架构**: 基于 Go 1.25+ 和 Fiber 框架
- **数据库支持**: 使用 Ent ORM，支持 SQLite、MySQL、PostgreSQL
- **前后端分离**: 后端提供 API 服务，前端使用 Vue 3 + TypeScript
- **用户认证**: JWT 认证系统，支持邮箱和手机号登录
- **内容管理**: 灵活的内容模型系统，支持自定义字段
- **支付系统**: 集成支付渠道管理和订单处理
- **文件管理**: 支持文件上传和管理，集成 AWS S3
- **相册管理**: 图片相册功能
- **评论系统**: 内容评论管理
- **响应式设计**: 基于 Layui 的管理界面

## 🚀 快速开始

### 环境要求

- Go 1.25 或更高版本
- Node.js 18+ (用于前端开发)
- SQLite (默认数据库)

### 安装依赖

```bash
# 安装 Go 依赖
go mod download

# 安装前端依赖
cd ui && npm install
```

### 数据库初始化

```bash
# 生成 Ent 代码
go generate ./ent

# 运行数据库迁移
# 系统会自动创建所需的表结构
```

### 运行项目

#### 开发模式

```bash
# 使用 Air 热重载 (推荐)
# Windows
air -c .air.windows.toml

# macOS/Linux
air -c .air.mac.toml

# 或者直接运行
make run
```

#### 前端开发

```bash
cd ui
npm run dev
```

### 访问应用

- 前端页面: http://localhost:3000
- 管理后台: http://localhost:3000/console
- API 文档: http://localhost:3000/api/v1

## 📁 项目结构

```
gobee/
├── assets/              # 静态资源
├── config/              # 配置文件
├── ent/                  # Ent ORM 生成的代码
├── internal/            # 内部包
│   ├── database/        # 数据库连接
│   ├── handlers/        # HTTP 处理器
│   ├── middleware/      # 中间件
│   ├── router/          # 路由配置
│   └── services/        # 业务逻辑
├── pkg/                 # 公共包
├── public/              # 静态文件
├── ui/                  # Vue 3 前端项目
├── views/               # Go 模板文件
├── main.go              # 应用入口
├── go.mod               # Go 模块文件
└── Makefile             # 构建脚本
```

## 🔧 配置说明

### 环境变量

```bash
# 数据库连接
DATABASE_URL=file:ent?mode=memory&cache=shared&_fk=1

# 服务端口
PORT=3000

# 运行环境
STAGE=dev  # dev, prod

# JWT 密钥
SECRET=your-secret-key
```

### 数据库配置

系统默认使用 SQLite 内存数据库，支持以下数据库：
- SQLite: `file:ent?mode=memory&cache=shared&_fk=1`
- MySQL: `mysql://user:password@localhost/database`
- PostgreSQL: `postgres://user:password@localhost/database`

## 🎯 功能模块

### 用户管理
- 用户注册、登录、认证
- JWT Token 管理
- 用户信息管理

### 内容管理
- 内容模型定义
- 动态内容管理
- 内容分类和标签

### 支付系统
- 支付渠道管理
- 订单处理
- 支付状态跟踪

### 文件管理
- 文件上传下载
- AWS S3 集成
- 文件分类管理

### 相册功能
- 图片上传
- 相册分类
- 图片展示

## 🔒 安全特性

- JWT 认证
- CORS 保护
- 输入验证
- 密码加密存储
- 环境变量管理

## 🛠️ 开发指南

### 添加新的数据模型

```bash
# 创建新的 Ent schema
ent init ModelName

# 生成代码
go generate ./ent
```

### API 开发

所有 API 路由都在 `internal/router/router.go` 中定义，遵循 RESTful 设计原则。

### 前端开发

前端使用 Vue 3 + TypeScript + Vite 构建，代码位于 `ui/` 目录。

## 📦 部署

### Docker 部署（推荐）

```bash
# 构建镜像
docker build -t gobee .

# 运行容器
docker run -p 3000:3000 -e DATABASE_URL=your-db-url gobee
```

### 传统部署

```bash
# 构建二进制文件
make build

# 运行应用
./main
```

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 📝 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🆘 支持

如果你遇到问题或有建议，请通过以下方式联系我们：
- 提交 Issue
- 发送邮件
- 加入讨论组

## 🙏 致谢

感谢以下开源项目的支持：
- [Fiber](https://github.com/gofiber/fiber) - Web 框架
- [Ent](https://entgo.io/) - ORM 框架
- [Vue.js](https://vuejs.org/) - 前端框架
- [Layui](https://layui.dev/) - UI 框架