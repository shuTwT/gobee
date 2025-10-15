# 🐝 GoBee - 现代化内容管理系统

> 🌟 一个基于 Go 语言开发的超现代化 CMS 系统，让你的内容管理变得轻松愉快！

## 📋 项目概述

GoBee 是一个基于 **Go 1.25+** 和 **Fiber 框架** 开发的现代化内容管理系统（CMS），专为开发者打造的高效、灵活、易扩展的内容管理解决方案。支持内容模型管理、用户系统、支付集成等丰富功能，前后端分离架构让你的开发体验飞起来！🚀

### ✨ 核心亮点

- 🏗️ **现代化架构**: 基于 Go 1.25+ 和 Fiber 高性能 Web 框架
- 💾 **多数据库支持**: Ent ORM 支持 SQLite、MySQL、PostgreSQL
- 🎨 **前后端分离**: 后端提供 RESTful API，前端使用 Vue 3 + TypeScript
- 🔐 **完整用户系统**: JWT 认证，支持邮箱/手机号登录
- 📝 **灵活内容管理**: 自定义内容模型，动态字段扩展
- 💰 **支付系统集成**: 多渠道支付管理和订单处理
- 📁 **文件管理**: 支持本地和 AWS S3 云存储
- 🖼️ **相册功能**: 图片管理和相册分类
- 💬 **评论系统**: 内容互动和评论管理

## 🚀 安装指南

### 📦 环境要求

在开始之前，请确保你的开发环境满足以下要求：

| 依赖项 | 最低版本 | 推荐版本 | 备注 |
|--------|----------|----------|------|
| Go | 1.25 | 1.25+ | 后端运行环境 |
| Node.js | 18.0.0 | 18+ | 前端开发环境 |
| SQLite | 3.x | 最新版 | 默认数据库 |
| Git | 任意版本 | 最新版 | 代码管理 |

### 🔧 快速安装

#### 1. 克隆项目代码

```bash
git clone https://github.com/your-username/gobee.git
cd gobee
```

#### 2. 安装后端依赖

```bash
# 下载 Go 依赖包
go mod download

# 验证依赖完整性
go mod verify
```

#### 3. 安装前端依赖

```bash
# 进入前端目录
cd ui

# 安装 npm 依赖
npm install

# 返回项目根目录
cd ..
```

#### 4. 数据库初始化

```bash
# 生成 Ent ORM 代码
go generate ./ent

# 数据库迁移（首次运行会自动创建表结构）
go run main.go
```

## 📖 使用说明

### 🏃‍♂️ 开发模式运行

#### 后端开发（推荐热重载）

```bash
# Windows 系统
air -c .air.windows.toml

# macOS/Linux 系统
air -c .air.mac.toml

# 或者使用 Makefile
make run
```

#### 前端开发

```bash
# 进入前端目录
cd ui

# 启动开发服务器
npm run dev

# 前端开发服务器运行在 http://localhost:5173
```

### 🌐 访问应用

当项目成功启动后，你可以通过以下地址访问：

- 🏠 **前端页面**: http://localhost:3000
- 🔧 **管理后台**: http://localhost:3000/console
- 📚 **API 文档**: http://localhost:3000/api/v1

### ⚙️ 环境配置

#### 基础环境变量

```bash
# 数据库连接配置
DATABASE_URL=file:ent?mode=memory&cache=shared&_fk=1

# 服务端口（默认：3000）
PORT=3000

# 运行环境（dev/prod）
STAGE=dev

# JWT 密钥（生产环境务必修改）
SECRET=your-super-secret-jwt-key-here
```

#### 数据库配置示例

| 数据库类型 | 连接字符串示例 |
|------------|----------------|
| SQLite | `file:ent?mode=memory&cache=shared&_fk=1` |
| MySQL | `mysql://user:password@localhost:3306/gobee` |
| PostgreSQL | `postgres://user:password@localhost:5432/gobee` |

### 🎯 功能模块详解

#### 👥 用户管理模块
- ✅ 用户注册与登录
- ✅ JWT Token 认证
- ✅ 用户信息管理
- ✅ 权限控制

#### 📝 内容管理模块
- ✅ 动态内容模型定义
- ✅ 自定义字段扩展
- ✅ 内容分类和标签
- ✅ 内容发布与审核

#### 💰 支付系统模块
- ✅ 多渠道支付管理
- ✅ 订单生命周期管理
- ✅ 支付状态实时跟踪
- ✅ 退款处理

#### 📁 文件管理模块
- ✅ 文件上传下载
- ✅ AWS S3 云存储集成
- ✅ 文件分类管理
- ✅ 图片压缩处理

#### 🖼️ 相册功能模块
- ✅ 图片批量上传
- ✅ 相册分类管理
- ✅ 图片懒加载
- ✅ 响应式图片展示

## 🤝 贡献指南

我们热烈欢迎社区成员参与项目贡献！无论你是开发者、设计师还是文档撰写者，都可以为项目贡献自己的力量。🌟

### 📝 贡献方式

#### 🐛 报告问题
- 在 [Issues](https://github.com/your-username/gobee/issues) 页面搜索是否已存在相关问题
- 创建新 Issue 时，请使用我们提供的模板
- 详细描述问题复现步骤和环境信息

#### 💡 功能建议
- 我们非常欢迎新功能的想法和建议
- 请在 Issue 中详细描述你的功能需求
- 说明该功能对用户的价值和意义

#### 🔧 代码贡献

1. **Fork 项目仓库**
   ```bash
   # 点击 GitHub 上的 Fork 按钮
   # 然后克隆你 Fork 的仓库
   git clone https://github.com/your-username/gobee.git
   ```

2. **创建特性分支**
   ```bash
   git checkout -b feature/amazing-feature
   # 或者修复 bug
   git checkout -b fix/nasty-bug
   ```

3. **进行开发**
   - 遵循项目的代码规范和风格
   - 编写必要的单元测试
   - 更新相关文档

4. **提交更改**
   ```bash
   git add .
   git commit -m '🚀 Add some amazing feature'
   # 或者
   git commit -m '🐛 Fix nasty bug in user authentication'
   ```

5. **推送和创建 PR**
   ```bash
   git push origin feature/amazing-feature
   # 然后在 GitHub 上创建 Pull Request
   ```

### 🎯 贡献准则

- 🎨 **代码风格**: 遵循 Go 官方代码规范，保持代码整洁
- 🧪 **测试覆盖**: 新功能需要包含相应的单元测试
- 📚 **文档更新**: 更新相关的 API 文档和使用说明
- 🔄 **保持同步**: 定期同步主分支，避免冲突

### 🏆 贡献者荣誉

所有贡献者的名字都会被记录在项目的贡献者列表中，你的付出将会被社区铭记！❤️

## 📄 许可证信息

本项目采用 **MIT 许可证** 开源协议，这意味着你可以：

- ✅ **自由使用**: 在个人和商业项目中免费使用
- ✅ **修改源码**: 根据需求修改和定制代码
- ✅ **分发软件**: 重新分发原始或修改后的版本
- ✅ **私人使用**: 在私有项目中使用

### 📋 MIT 许可证摘要

```
MIT License

Copyright (c) 2024 GoBee Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

完整许可证文本请查看 [LICENSE](LICENSE) 文件。

## 🙏 致谢

GoBee 项目的成功离不开以下优秀开源项目的支持，在此表示衷心感谢！

| 项目名称 | 用途 | 许可证 |
|----------|------|--------|
| [Fiber](https://github.com/gofiber/fiber) | Web 框架 | MIT |
| [Ent](https://entgo.io/) | ORM 框架 | Apache 2.0 |
| [Vue.js](https://vuejs.org/) | 前端框架 | MIT |
| [Vite](https://vitejs.dev/) | 构建工具 | MIT |

## 📞 支持与联系

如果你在使用 GoBee 过程中遇到问题，或者有任何建议，欢迎通过以下方式联系我们：

- 🐛 **问题反馈**: [提交 Issue](https://github.com/your-username/gobee/issues)
- 💬 **讨论交流**: [加入 Discussions](https://github.com/your-username/gobee/discussions)
- 📧 **邮件联系**: gobee@example.com
- 🌟 **给项目点星**: 如果项目对你有帮助，欢迎点亮 Star！

---

<div align="center">

**🎉 让 GoBee 成为你内容管理的最佳伙伴！**

Made with ❤️ by the GoBee Community

</div>