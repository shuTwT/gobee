# AI 模块

AI 模块是 Hoshikuzu 的核心功能之一，提供强大的 AI 能力，包括 AI 摘要、AI 播客、AI 写作、AI 知识库等功能。

## 功能特性

- **AI 摘要**：自动生成文章摘要，快速了解文章内容
- **AI 播客**：将文本内容转换为音频播客
- **AI 写作**：智能辅助写作，提供写作建议和内容生成
- **AI 知识库**：构建和管理知识库，智能检索和问答

## 快速开始

### 1. 配置 AI 参数

在 `config.yaml` 中配置 AI 参数：

```yaml
ai:
  api_key: your_ai_api_key
  model: gpt-4
  max_tokens: 2000
```

### 2. 使用 AI 功能

```javascript
import { aiApi } from '@/api/ai'

const generateSummary = async () => {
  const data = {
    content: '文章内容...'
  }
  
  const result = await aiApi.summary(data)
  console.log(result)
}
```

## AI 功能

### AI 摘要

自动生成文章摘要，帮助用户快速了解文章内容。

### AI 播客

将文本内容转换为音频播客，提供更好的阅读体验。

### AI 写作

智能辅助写作，提供写作建议和内容生成。

### AI 知识库

构建和管理知识库，智能检索和问答。

## API 接口

### AI 摘要

```http
POST /api/v1/ai/summary
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| content | string | 是 | 文章内容 |
| max_length | int | 否 | 摘要最大长度 |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "summary": "文章摘要..."
  }
}
```

### AI 播客

```http
POST /api/v1/ai/podcast
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| content | string | 是 | 文章内容 |
| voice | string | 否 | 语音类型 |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "audio_url": "https://your-domain.com/audio/xxx.mp3"
  }
}
```

### AI 写作

```http
POST /api/v1/ai/writing
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| prompt | string | 是 | 写作提示 |
| type | string | 否 | 写作类型 |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "content": "生成的内容..."
  }
}
```

### AI 知识库

```http
POST /api/v1/ai/knowledge
```

**请求参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| question | string | 是 | 问题 |
| context | string | 否 | 上下文 |

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "answer": "回答..."
  }
}
```

## 使用示例

### 生成文章摘要

```javascript
import { aiApi } from '@/api/ai'

const generateSummary = async () => {
  const data = {
    content: 'Hoshikuzu 是一款高性能、高稳定性、易扩展的内容管理系统...',
    max_length: 100
  }
  
  const result = await aiApi.summary(data)
  console.log(result.summary)
}
```

### 生成播客

```javascript
import { aiApi } from '@/api/ai'

const generatePodcast = async () => {
  const data = {
    content: 'Hoshikuzu 是一款高性能、高稳定性、易扩展的内容管理系统...',
    voice: 'male'
  }
  
  const result = await aiApi.podcast(data)
  console.log(result.audio_url)
}
```

### AI 写作

```javascript
import { aiApi } from '@/api/ai'

const aiWriting = async () => {
  const data = {
    prompt: '写一篇关于 Hoshikuzu 的介绍文章',
    type: 'article'
  }
  
  const result = await aiApi.writing(data)
  console.log(result.content)
}
```

### 知识库问答

```javascript
import { aiApi } from '@/api/ai'

const knowledgeQuery = async () => {
  const data = {
    question: 'Hoshikuzu 支持哪些数据库？'
  }
  
  const result = await aiApi.knowledge(data)
  console.log(result.answer)
}
```

## 最佳实践

### AI 摘要

1. **内容长度**：建议文章内容在 500 字以上
2. **摘要长度**：根据需求设置合适的摘要长度
3. **内容质量**：高质量的内容能生成更好的摘要

### AI 播客

1. **语音选择**：根据内容类型选择合适的语音
2. **音频质量**：使用高质量的音频格式
3. **文件大小**：注意音频文件大小，避免过大

### AI 写作

1. **提示词**：使用清晰明确的提示词
2. **内容类型**：指定内容类型以获得更好的结果
3. **人工审核**：AI 生成的内容需要人工审核

### AI 知识库

1. **知识库构建**：定期更新知识库内容
2. **问题质量**：使用清晰明确的问题
3. **上下文信息**：提供足够的上下文信息

## 相关文档

- [AI 摘要](summary.md) - AI 摘要详细文档
- [AI 播客](podcast.md) - AI 播客详细文档
- [AI 写作](writing.md) - AI 写作详细文档
- [AI 知识库](knowledge.md) - AI 知识库详细文档
- [内容管理模块](../content/article.md) - 内容管理模块文档

## AI 聊天（新增）

AI 聊天通过受认证的服务端代理调用 OpenAI 兼容的 Chat Completions 流式接口。浏览器不直连供应商，聊天会话和消息按登录用户持久化。

### 启用配置

启动服务前必须设置服务端环境变量 `AI_CONFIG_ENCRYPTION_KEY`，值为 Base64 编码的 32 字节 AES-256 密钥。缺失或无法解密时，AI 配置和聊天接口会关闭，不使用默认密钥或旧配置兜底。

超级管理员在“系统设置 → AI 设置”填写完整配置：

- `base_url`：必须是以 `/v1` 结尾的 HTTP(S) 地址
- `model`：供应商支持的模型 ID
- `api_key`：每次保存重新输入，服务端使用 AES-256-GCM 加密保存且不回显
- 温度、Top P、惩罚参数和最大输出 Token

只有数据库真实角色码为 `superAdmin` 的用户可以读取、保存、测试配置和刷新模型列表；所有已登录用户可以使用自己的聊天会话。旧通用设置中的 `ai` 和 `openai_*` 配置不会被读取或迁移。

### 聊天接口

- `GET/POST /api/v1/ai/chat/sessions`
- `GET /api/v1/ai/chat/sessions/{id}/messages`
- `DELETE /api/v1/ai/chat/sessions/{id}`
- `DELETE /api/v1/ai/chat/sessions/{id}/messages`
- `POST /api/v1/ai/chat/sessions/{id}/stream`

流接口返回 `text/event-stream`，事件为 `delta`、`done` 和 `error`。前端使用携带 Authorization 请求头的 `fetch` 和 `ReadableStream` 解析，不使用 `EventSource`。每次请求提交会话的全部已保存历史，不自动截断、摘要、重试或模型兜底；停止生成时未完成的助手消息不会入库。
