import { http } from '@/utils/http'
import { getToken } from '@/utils/auth'
import { BASE_URL, type ApiResponse } from '@/api/utils'

const AI_BASE_URL = `${BASE_URL}/v1/ai`

export interface AIModelItem {
  id: number
  model_name: string
  display_name?: string
  is_enabled: boolean
  sort: number
}

export interface AIProvider {
  id: number
  name: string
  provider_type: string
  base_url: string
  api_key_configured: boolean
  temperature: number
  max_tokens: number
  top_p: number
  frequency_penalty: number
  presence_penalty: number
  is_default: boolean
  is_enabled: boolean
  sort: number
  remark: string
  created_at: number
  updated_at: number
  models: AIModelItem[]
}

export interface AIProviderInput {
  name: string
  provider_type: string
  base_url: string
  api_key: string
  temperature: number
  max_tokens: number
  top_p: number
  frequency_penalty: number
  presence_penalty: number
  is_default: boolean
  is_enabled: boolean
  sort: number
  remark: string
}

export interface AIModelInput {
  model_name: string
  display_name: string
  is_enabled: boolean
  sort: number
}

export interface ChatSession {
  id: number
  title: string
  created_at: number
  updated_at: number
}

export interface ChatMessage {
  id: number
  role: 'user' | 'assistant'
  content: string
  model?: string
  created_at: number
}

export interface AIStreamEvent {
  content?: string
  message_id?: number
  code?: string
  message?: string
}

export const listProviders = () => {
  return http.request<ApiResponse<AIProvider[]>>('get', `${AI_BASE_URL}/providers/list`)
}

export const createProvider = (data: AIProviderInput) => {
  return http.request<ApiResponse<null>>('post', `${AI_BASE_URL}/providers/create`, { data })
}

export const updateProvider = (id: number, data: AIProviderInput) => {
  return http.request<ApiResponse<null>>('put', `${AI_BASE_URL}/providers/update/${id}`, { data })
}

export const deleteProvider = (id: number) => {
  return http.request<ApiResponse<null>>('delete', `${AI_BASE_URL}/providers/delete/${id}`)
}

export const testProvider = (data: Pick<AIProviderInput, 'base_url' | 'api_key'>) => {
  return http.request<ApiResponse<null>>('post', `${AI_BASE_URL}/providers/test`, { data })
}

export const syncProviderModels = (providerId: number) => {
  return http.request<ApiResponse<null>>('post', `${AI_BASE_URL}/providers/${providerId}/models/sync`)
}

export const createProviderModel = (providerId: number, data: AIModelInput) => {
  return http.request<ApiResponse<null>>('post', `${AI_BASE_URL}/providers/${providerId}/models/create`, { data })
}

export const updateProviderModel = (providerId: number, modelId: number, data: AIModelInput) => {
  return http.request<ApiResponse<null>>('put', `${AI_BASE_URL}/providers/${providerId}/models/update/${modelId}`, { data })
}

export const deleteProviderModel = (providerId: number, modelId: number) => {
  return http.request<ApiResponse<null>>('delete', `${AI_BASE_URL}/providers/${providerId}/models/delete/${modelId}`)
}

export const listSessions = () => {
  return http.request<ApiResponse<ChatSession[]>>('get', `${AI_BASE_URL}/chat/sessions`)
}

export const createSession = () => {
  return http.request<ApiResponse<ChatSession>>('post', `${AI_BASE_URL}/chat/sessions`)
}

export const listMessages = (sessionId: number) => {
  return http.request<ApiResponse<ChatMessage[]>>('get', `${AI_BASE_URL}/chat/sessions/${sessionId}/messages`)
}

export const deleteSession = (sessionId: number) => {
  return http.request<ApiResponse<null>>('delete', `${AI_BASE_URL}/chat/sessions/${sessionId}`)
}

export const clearSession = (sessionId: number) => {
  return http.request<ApiResponse<null>>('delete', `${AI_BASE_URL}/chat/sessions/${sessionId}/messages`)
}

type StreamEventName = 'delta' | 'done' | 'error'

function parseStreamBlock(block: string): { event: StreamEventName; data: AIStreamEvent } | null {
  let event: StreamEventName = 'delta'
  const dataLines: string[] = []
  for (const line of block.split(/\r?\n/)) {
    if (line.startsWith('event:')) {
      const value = line.slice('event:'.length).trim()
      if (value === 'delta' || value === 'done' || value === 'error') event = value
    } else if (line.startsWith('data:')) {
      dataLines.push(line.slice('data:'.length).trimStart())
    }
  }
  if (dataLines.length === 0) return null
  try {
    return { event, data: JSON.parse(dataLines.join('\n')) as AIStreamEvent }
  } catch {
    throw new Error('无法解析 AI 流响应')
  }
}

async function responseError(response: Response): Promise<Error> {
  try {
    const body = (await response.json()) as { msg?: string; message?: string }
    return new Error(body.msg || body.message || `AI 请求失败（${response.status}）`)
  } catch {
    return new Error(`AI 请求失败（${response.status}）`)
  }
}

export async function streamChat(
  sessionId: number,
  content: string,
  onEvent: (event: StreamEventName, data: AIStreamEvent) => void,
  signal: AbortSignal,
): Promise<void> {
  const token = getToken()?.accessToken
  const response = await fetch(`${AI_BASE_URL}/chat/sessions/${sessionId}/stream`, {
    method: 'POST',
    headers: {
      Accept: 'text/event-stream',
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    body: JSON.stringify({ content }),
    signal,
  })
  if (!response.ok) throw await responseError(response)
  if (!response.body) throw new Error('AI 服务未返回流响应')

  const reader = response.body.getReader()
  const decoder = new TextDecoder()
  let buffer = ''
  const consume = (final = false) => {
    if (final) buffer += decoder.decode()
    const blocks = buffer.split(/\r?\n\r?\n/)
    buffer = blocks.pop() || ''
    for (const block of blocks) {
      const parsed = parseStreamBlock(block)
      if (parsed) onEvent(parsed.event, parsed.data)
    }
  }

  try {
    while (true) {
      const { value, done } = await reader.read()
      if (done) break
      buffer += decoder.decode(value, { stream: true })
      consume()
    }
    consume(true)
  } finally {
    await reader.cancel().catch(() => undefined)
    reader.releaseLock()
  }
}
