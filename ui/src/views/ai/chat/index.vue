<script setup lang="ts">
import type { ScrollbarInst } from 'naive-ui'
import type { Message, ChatSession } from './utils/types'

const chatSessions = ref<ChatSession[]>([
  {
    id: '1',
    title: '新对话',
    messages: [
      {
        id: '1',
        role: 'assistant',
        content: '你好！我是 AI 助手，有什么可以帮助你的吗？',
        timestamp: Date.now()
      }
    ],
    createdAt: Date.now(),
    updatedAt: Date.now()
  }
])

const currentSessionId = ref('1')
const inputContent = ref('')
const loading = ref(false)
const messageListRef = ref<ScrollbarInst | null>(null)

const models = ref([
  { label: 'GPT-4', value: 'gpt-4' },
  { label: 'GPT-3.5 Turbo', value: 'gpt-3.5-turbo' },
  { label: 'Claude 3', value: 'claude-3' },
  { label: 'Gemini Pro', value: 'gemini-pro' }
])

const selectedModel = ref('gpt-4')

const currentSession = computed(() => {
  return chatSessions.value.find(session => session.id === currentSessionId.value)
})

const messages = computed(() => {
  return currentSession.value?.messages || []
})

const createNewSession = () => {
  const newSession: ChatSession = {
    id: Date.now().toString(),
    title: '新对话',
    messages: [
      {
        id: Date.now().toString(),
        role: 'assistant',
        content: '你好！我是 AI 助手，有什么可以帮助你的吗？',
        timestamp: Date.now()
      }
    ],
    createdAt: Date.now(),
    updatedAt: Date.now()
  }
  chatSessions.value.unshift(newSession)
  currentSessionId.value = newSession.id
}

const selectSession = (sessionId: string) => {
  currentSessionId.value = sessionId
  nextTick(() => {
    scrollToBottom()
  })
}

const deleteSession = (sessionId: string, event: Event) => {
  event.stopPropagation()
  if (chatSessions.value.length === 1) {
    window.$message.warning('至少保留一个对话')
    return
  }
  const index = chatSessions.value.findIndex(session => session.id === sessionId)
  chatSessions.value.splice(index, 1)
  if (currentSessionId.value === sessionId) {
    currentSessionId.value = chatSessions.value[0].id
  }
}

const updateSessionTitle = (sessionId: string, firstMessage: string) => {
  const session = chatSessions.value.find(s => s.id === sessionId)
  if (session) {
    session.title = firstMessage.slice(0, 20) + (firstMessage.length > 20 ? '...' : '')
    session.updatedAt = Date.now()
  }
}

const sendMessage = async () => {
  if (!inputContent.value.trim() || loading.value) return

  const session = chatSessions.value.find(s => s.id === currentSessionId.value)
  if (!session) return

  const userMessage: Message = {
    id: Date.now().toString(),
    role: 'user',
    content: inputContent.value.trim(),
    timestamp: Date.now()
  }

  session.messages.push(userMessage)

  if (session.messages.length === 2) {
    updateSessionTitle(session.id, userMessage.content)
  }

  session.updatedAt = Date.now()
  inputContent.value = ''
  loading.value = true

  scrollToBottom()

  setTimeout(() => {
    const aiMessage: Message = {
      id: (Date.now() + 1).toString(),
      role: 'assistant',
      content: `这是一个模拟的 AI 回复。当前使用的模型是：${models.value.find(m => m.value === selectedModel.value)?.label}。目前后端接口尚未接入，请稍后再试。`,
      timestamp: Date.now()
    }
    session.messages.push(aiMessage)
    session.updatedAt = Date.now()
    loading.value = false
    scrollToBottom()
  }, 1000)
}

const scrollToBottom = () => {
  nextTick(() => {
    messageListRef.value?.scrollTo({ top: 999999, behavior: 'smooth' })
  })
}

const clearChat = () => {
  const session = chatSessions.value.find(s => s.id === currentSessionId.value)
  if (session) {
    session.messages = [
      {
        id: Date.now().toString(),
        role: 'assistant',
        content: '你好！我是 AI 助手，有什么可以帮助你的吗？',
        timestamp: Date.now()
      }
    ]
    session.updatedAt = Date.now()
  }
}

const handleKeyDown = (e: KeyboardEvent) => {
  if (e.key === 'Enter' && !e.shiftKey) {
    e.preventDefault()
    sendMessage()
  }
}

onMounted(() => {
  scrollToBottom()
})
</script>

<template>
  <div class="h-[calc(100vh_-_100px)] flex bg-gray-50 dark:bg-gray-900">
    <div class="w-72 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 flex flex-col">
      <div class="p-4 border-b border-gray-200 dark:border-gray-700">
        <n-button type="primary" block @click="createNewSession">
          <template #icon>
            <n-icon>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path fill="currentColor" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/>
              </svg>
            </n-icon>
          </template>
          新建对话
        </n-button>
      </div>
      
      <n-scrollbar class="flex-1">
        <div class="p-2 space-y-1">
          <div
            v-for="session in chatSessions"
            :key="session.id"
            @click="selectSession(session.id)"
            class="relative group p-3 rounded-lg cursor-pointer transition-colors"
            :class="currentSessionId === session.id
              ? 'bg-blue-50 dark:bg-blue-900/20'
              : 'hover:bg-gray-100 dark:hover:bg-gray-700'"
          >
            <div class="flex items-start justify-between">
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-gray-900 dark:text-white truncate">
                  {{ session.title }}
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
                  {{ new Date(session.updatedAt).toLocaleString() }}
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
                      <path fill="currentColor" d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"/>
                    </svg>
                  </n-icon>
                </template>
              </n-button>
            </div>
          </div>
        </div>
      </n-scrollbar>
    </div>

    <div class="flex-1 flex flex-col">
      <div class="absolute top-0 left-72 right-0">
        <n-card class="flex-none border-b border-gray-200 dark:border-gray-700">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <n-icon size="24" color="#18a058">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                  <path fill="currentColor"
                    d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zm-1 17.93c-3.95-.49-7-3.85-7-7.93c0-.62.08-1.21.21-1.79L9 15v1c0 1.1.9 2 2 2v1.93zm6.9-2.54c-.26-.81-1-1.39-1.9-1.39h-1v-3c0-.55-.45-1-1-1H8v-2h2c.55 0 1-.45 1-1V7h2c1.1 0 2-.9 2-2v-.41c2.93 1.19 5 4.06 5 7.41c0 2.08-.8 3.97-2.1 5.39z" />
                </svg>
              </n-icon>
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white">AI 助手</h2>
            </div>
            <n-button quaternary @click="clearChat">
              <template #icon>
                <n-icon>
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                    <path fill="currentColor"
                      d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z" />
                  </svg>
                </n-icon>
              </template>
              清空对话
            </n-button>
          </div>
        </n-card>
      </div>

      <div class="absolute top-[72px] left-72 bottom-[120px] right-0">
        <n-scrollbar ref="messageListRef" class="flex-1">
          <div class="p-6 space-y-6">
            <div v-for="message in messages" :key="message.id" class="flex"
              :class="message.role === 'user' ? 'justify-end' : 'justify-start'">
              <div class="flex items-start space-x-3 max-w-[80%]"
                :class="message.role === 'user' ? 'flex-row-reverse space-x-reverse' : ''">
                <n-avatar round :size="40" :style="{
                  backgroundColor: message.role === 'user' ? '#18a058' : '#2080f0'
                }">
                  {{ message.role === 'user' ? 'U' : 'AI' }}
                </n-avatar>
                <div class="px-4 py-3 rounded-2xl shadow-sm" :class="message.role === 'user'
                  ? 'bg-green-500 text-white rounded-br-md'
                  : 'bg-white dark:bg-gray-800 text-gray-900 dark:text-white rounded-bl-md'">
                  <p class="whitespace-pre-wrap break-words">{{ message.content }}</p>
                  <p class="text-xs mt-2 opacity-70"
                    :class="message.role === 'user' ? 'text-green-100' : 'text-gray-500 dark:text-gray-400'">
                    {{ new Date(message.timestamp).toLocaleTimeString() }}
                  </p>
                </div>
              </div>
            </div>

            <div v-if="loading" class="flex justify-start">
              <div class="flex items-start space-x-3 max-w-[80%]">
                <n-avatar round :size="40" :style="{ backgroundColor: '#2080f0' }">
                  AI
                </n-avatar>
                <div class="px-4 py-3 rounded-2xl rounded-bl-md bg-white dark:bg-gray-800 shadow-sm">
                  <n-spin size="small" />
                </div>
              </div>
            </div>
          </div>
        </n-scrollbar>
      </div>

      <div class="absolute bottom-0 left-72 right-0">
        <n-card class="flex-none border-t border-gray-200 dark:border-gray-700">
          <div class="flex items-center space-x-3 mb-3">
            <span class="text-sm text-gray-600 dark:text-gray-400">选择模型：</span>
            <n-select
              v-model:value="selectedModel"
              :options="models"
              class="w-48"
            />
          </div>
          <div class="flex items-end space-x-3">
            <n-input v-model:value="inputContent" type="textarea" placeholder="输入消息... (Enter 发送，Shift+Enter 换行)"
              :autosize="{ minRows: 1, maxRows: 4 }" :disabled="loading" @keydown="handleKeyDown" class="flex-1" />
            <n-button type="primary" :disabled="!inputContent.trim() || loading" :loading="loading" @click="sendMessage"
              class="h-[42px] px-6">
              发送
            </n-button>
          </div>
          <p class="text-xs text-gray-500 dark:text-gray-400 mt-2">
            提示：目前使用模拟数据，后端接口尚未接入
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
