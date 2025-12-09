<script lang="ts" setup>
import { Language } from '@vicons/ionicons5'
const currentLanguage = ref('zh-CN')
const languageOptions = [
  { label: '简体中文', key: 'zh-CN' },
  { label: 'English', key: 'en-US' },
]

const handleLanguageChange = (key: string) => {
  currentLanguage.value = key
  // 这里可以添加语言切换逻辑
}
</script>
<template>
  <div
    class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4"
  >
    <!-- 语言切换 -->
    <div class="flex justify-end mb-4 absolute top-4 right-4">
      <n-dropdown trigger="click" :options="languageOptions" @select="handleLanguageChange">
        <n-button text>
          <template #icon>
            <n-icon><Language /></n-icon>
          </template>
          {{ currentLanguage === 'zh-CN' ? '简体中文' : 'English' }}
        </n-button>
      </n-dropdown>
    </div>
    <div class="w-full max-w-md">
      <!-- Logo -->
      <div class="text-center mb-8">
        <div
          class="w-20 h-20 mx-auto mb-4 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center"
        >
          <span class="text-white text-2xl font-bold">G</span>
        </div>
        <h1 class="text-3xl font-bold text-gray-800 mb-2">GoBee</h1>
        <p class="text-gray-600">欢迎来到我们的平台</p>
      </div>

      <!-- 登录卡片 -->
      <n-card class="shadow-xl rounded-2xl">
        <div class="p-6">
          <slot/>
        </div>
      </n-card>

      <!-- 底部信息 -->
      <div class="text-center mt-6 text-gray-500 text-sm">
        <p>© 2024 GoBee. All rights reserved.</p>
      </div>
    </div>
  </div>
</template>
<style scoped>
/* 自定义样式 */
:deep(.n-card) {
  border-radius: 1rem;
}

:deep(.n-button--primary) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

:deep(.n-button--primary:hover) {
  background: linear-gradient(135deg, #5a6fd8 0%, #6a4190 100%);
}
</style>
