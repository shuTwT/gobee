import { http } from '@/utils/http'
import { getToken } from '@/utils/auth'
import { BASE_URL, type ApiResponse } from '@/api/utils'

const AI_BASE_URL = `${BASE_URL}/v1/ai`

export interface AIConfig {
  base_url: string
  model: string
  temperature: number
  max_tokens: number
  top_p: number
  frequency_penalty: number
  presence_penalty: number
  api_key_configured: boolean
}

export interface AIConfigInput {
  base_url: string
  api_key: string
  model: string
  temperature: number
  max_tokens: number
  top_p: number
  frequency_penalty: number
  presence_penalty: number
}

export interface AIModel {
  id: string
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

export const getConfig = () => {
  return http.request<ApiResponse<AIConfig>>('get', `${AI_BASE_URL}/config`)
}

export const saveConfig = (data: AIConfigInput) => {
  return http.request<ApiResponse<null>>('put', `${AI_BASE_URL}/config`, { data })
}

export const testConfig = (data: Pick<AIConfigInput, 'base_url' | 'api_key'>) => {
  return http.request<ApiResponse<null>>('post', `${AI_BASE_URL}/config/test`, { data })
}

export const getModels = () => {
  return http.request<ApiResponse<AIModel[]>>('get', `${AI_BASE_URL}/config/models`)
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
