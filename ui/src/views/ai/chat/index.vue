<script setup lang="ts">
import type { ScrollbarInst } from 'naive-ui'
import * as aiApi from '@/api/ai'
import type { ChatMessage as Message, ChatSession } from '@/api/ai'

const chatSessions = ref<ChatSession[]>([])
const currentSessionId = ref<number | null>(null)
const messages = ref<Message[]>([])
const inputContent = ref('')
const loading = ref(false)
const messageListRef = ref<ScrollbarInst | null>(null)
let abortController: AbortController | null = null

const currentSession = computed(() =>
  chatSessions.value.find(session => session.id === currentSessionId.value),
)

const scrollToBottom = () => {
  nextTick(() => {
    messageListRef.value?.scrollTo({ top: 999999, behavior: 'smooth' })
  })
}

const refreshSessions = async () => {
  const res = await aiApi.listSessions()
  chatSessions.value = res.data
}

const loadMessages = async (sessionId: number) => {
  const res = await aiApi.listMessages(sessionId)
  currentSessionId.value = sessionId
  messages.value = res.data
  scrollToBottom()
}

const createNewSession = async () => {
  if (loading.value) return
  try {
    const res = await aiApi.createSession()
    chatSessions.value.unshift(res.data)
    currentSessionId.value = res.data.id
    messages.value = []
    scrollToBottom()
  } catch {
    window.$message.error('新建对话失败')
  }
}

const loadChat = async () => {
  try {
    await refreshSessions()
    if (chatSessions.value.length === 0) {
      await createNewSession()
      return
    }
    await loadMessages(chatSessions.value[0].id)
  } catch {
    window.$message.error('加载聊天记录失败')
  }
}

const selectSession = async (sessionId: number) => {
  if (loading.value || sessionId === currentSessionId.value) return
  try {
    await loadMessages(sessionId)
  } catch {
    window.$message.error('加载对话消息失败')
  }
}

const deleteSession = async (sessionId: number, event: Event) => {
  event.stopPropagation()
  if (loading.value) return
  try {
    await aiApi.deleteSession(sessionId)
    const remaining = chatSessions.value.filter(session => session.id !== sessionId)
    chatSessions.value = remaining
    if (currentSessionId.value === sessionId) {
      if (remaining.length > 0) {
        await loadMessages(remaining[0].id)
      } else {
        await createNewSession()
      }
    }
  } catch {
    window.$message.error('删除对话失败')
  }
}

const clearChat = async () => {
  if (loading.value || currentSessionId.value === null) return
  try {
    await aiApi.clearSession(currentSessionId.value)
    messages.value = []
    const session = chatSessions.value.find(item => item.id === currentSessionId.value)
    if (session) {
      session.title = '新对话'
      session.updated_at = Date.now()
    }
    scrollToBottom()
  } catch {
    window.$message.error('清空对话失败')
  }
}

const stopGeneration = () => {
  abortController?.abort()
}

const sendMessage = async () => {
  const content = inputContent.value.trim()
  const sessionId = currentSessionId.value
  if (!content || sessionId === null || loading.value) return

  const userMessage: Message = {
    id: -Date.now(),
    role: 'user',
    content,
    created_at: Date.now(),
  }
  const assistantMessage: Message = {
    id: -Date.now() - 1,
    role: 'assistant',
    content: '',
    created_at: Date.now(),
  }
  messages.value.push(userMessage, assistantMessage)
  inputContent.value = ''
  loading.value = true
  abortController = new AbortController()
  scrollToBottom()

  try {
    await aiApi.streamChat(
      sessionId,
      content,
      (event, data) => {
        if (event === 'delta') {
          assistantMessage.content += data.content || ''
          scrollToBottom()
        } else if (event === 'done' && data.message_id) {
          assistantMessage.id = data.message_id
        } else if (event === 'error') {
          throw new Error(data.message || 'AI 回复失败')
        }
      },
      abortController.signal,
    )
    await loadMessages(sessionId)
    await refreshSessions()
  } catch (error) {
    await loadMessages(sessionId).catch(() => undefined)
    if (!(error instanceof DOMException && error.name === 'AbortError')) {
      window.$message.error(error instanceof Error ? error.message : 'AI 回复失败')
    }
  } finally {
    loading.value = false
    abortController = null
    scrollToBottom()
  }
}

const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    sendMessage()
  }
}

const formatTime = (timestamp: number) => new Date(timestamp).toLocaleTimeString()

onMounted(loadChat)
onBeforeUnmount(() => abortController?.abort())
</script>

<template>
  <div class="h-[calc(100vh_-_100px)] flex bg-gray-50 dark:bg-gray-900">
    <div class="w-72 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 flex flex-col">
      <div class="p-4 border-b border-gray-200 dark:border-gray-700">
        <n-button type="primary" block :disabled="loading" @click="createNewSession">
          <template #icon>
            <n-icon>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path fill="currentColor" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z" />
              </svg>
            </n-icon>
          </template>
          新建对话
        </n-button>
      </div>

      <n-scrollbar class="flex-1">
        <div v-if="chatSessions.length" class="p-2 space-y-1">
          <div
            v-for="session in chatSessions"
            :key="session.id"
            class="relative group p-3 rounded-lg cursor-pointer transition-colors"
            :class="currentSessionId === session.id
              ? 'bg-blue-50 dark:bg-blue-900/20'
              : 'hover:bg-gray-100 dark:hover:bg-gray-700'"
            @click="selectSession(session.id)"
          >
            <div class="flex items-start justify-between">
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-gray-900 dark:text-white truncate">
                  {{ session.title }}
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
                  {{ formatTime(session.updated_at) }}
                </p>
              </div>
              <n-button
                text
                size="tiny"
                class="opacity-0 group-hover:opacity-100 transition-opacity"
                @click="deleteSession(session.id, $event)"
              >
                <template #icon>
                  <n-icon>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                      <path fill="currentColor" d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z" />
                    </svg>
                  </n-icon>
                </template>
              </n-button>
            </div>
          </div>
        </div>
        <n-empty v-else class="p-6" description="暂无对话" />
      </n-scrollbar>
    </div>

    <div class="flex-1 flex flex-col relative">
      <div class="absolute top-0 left-0 right-0 z-10">
        <n-card class="border-b border-gray-200 dark:border-gray-700">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <n-icon size="24" color="#18a058">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                  <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2z" />
                </svg>
              </n-icon>
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
                {{ currentSession?.title || 'AI 助手' }}
              </h2>
            </div>
            <n-button quaternary :disabled="loading || !currentSessionId" @click="clearChat">
              清空对话
            </n-button>
          </div>
        </n-card>
      </div>

      <div class="absolute top-[72px] left-0 bottom-[132px] right-0">
        <n-scrollbar ref="messageListRef" class="h-full">
          <div class="p-6 space-y-6">
            <n-empty v-if="messages.length === 0" description="开始一段新对话" class="mt-20" />
            <div v-for="message in messages" :key="message.id" class="flex" :class="message.role === 'user' ? 'justify-end' : 'justify-start'">
              <div class="flex items-start space-x-3 max-w-[80%]" :class="message.role === 'user' ? 'flex-row-reverse space-x-reverse' : ''">
                <n-avatar round :size="40" :style="{ backgroundColor: message.role === 'user' ? '#18a058' : '#2080f0' }">
                  {{ message.role === 'user' ? 'U' : 'AI' }}
                </n-avatar>
                <div
                  class="px-4 py-3 rounded-2xl shadow-sm"
                  :class="message.role === 'user'
                    ? 'bg-green-500 text-white rounded-br-md'
                    : 'bg-white dark:bg-gray-800 text-gray-900 dark:text-white rounded-bl-md'"
                >
                  <n-spin v-if="message.role === 'assistant' && !message.content && loading" size="small" />
                  <p v-else class="whitespace-pre-wrap break-words">{{ message.content }}</p>
                  <p class="text-xs mt-2 opacity-70" :class="message.role === 'user' ? 'text-green-100' : 'text-gray-500 dark:text-gray-400'">
                    {{ formatTime(message.created_at) }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </n-scrollbar>
      </div>

      <div class="absolute bottom-0 left-0 right-0">
        <n-card class="border-t border-gray-200 dark:border-gray-700">
          <div class="flex items-end space-x-3">
            <n-input
              v-model:value="inputContent"
              type="textarea"
              placeholder="输入消息...（Enter 发送，Shift+Enter 换行）"
              :autosize="{ minRows: 1, maxRows: 4 }"
              :disabled="loading || !currentSessionId"
              class="flex-1"
              @keydown="handleKeyDown"
            />
            <n-button v-if="loading" type="warning" class="h-[42px] px-6" @click="stopGeneration">
              停止
            </n-button>
            <n-button v-else type="primary" :disabled="!inputContent.trim() || !currentSessionId" class="h-[42px] px-6" @click="sendMessage">
              发送
            </n-button>
          </div>
          <p class="text-xs text-gray-500 dark:text-gray-400 mt-2">
            回复由服务端通过 OpenAI 兼容接口流式生成。
          </p>
        </n-card>
      </div>
    </div>
  </div>
</template>

<style scoped>
:deep(.n-card__content) {
  padding: 16px;
}

:deep(.n-input__textarea-el) {
  resize: none;
}
</style>
