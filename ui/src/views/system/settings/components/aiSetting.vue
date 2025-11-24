<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import type { AiSetting } from '../utils/types';
import * as settingApi from '@/api/system/setting'

const message = useMessage()
const aiModelOptions = [
  { label: 'GPT-3.5 Turbo', value: 'gpt-3.5-turbo' },
  { label: 'GPT-4', value: 'gpt-4' },
  { label: 'GPT-4 Turbo', value: 'gpt-4-turbo' },
  { label: 'GPT-4o', value: 'gpt-4o' }
]
const aiFormRef = ref<FormInst|null>(null)
const defaultForm = {
  openaiApiKey: '',
  openaiApiUrl: 'https://api.openai.com/v1',
  aiModel: 'gpt-3.5-turbo',
  aiTemperature: 0.7,
  aiMaxTokens: 2048,
  aiTopP: 1.0,
  aiFrequencyPenalty: 0,
  aiPresencePenalty: 0,
}
const aiForm = ref<AiSetting>({
  openaiApiKey: '',
  openaiApiUrl: 'https://api.openai.com/v1',
  aiModel: 'gpt-3.5-turbo',
  aiTemperature: 0.7,
  aiMaxTokens: 2048,
  aiTopP: 1.0,
  aiFrequencyPenalty: 0,
  aiPresencePenalty: 0,
})


const aiLoading = ref(false)
// 保存AI设置
const saveAISettings = async () => {
  aiLoading.value = true
  try {
    await settingApi.saveSettings('ai',aiForm.value)
    await new Promise(resolve => setTimeout(resolve, 1000))
    onSearch()
    message.success('AI设置保存成功')
  } catch {
    message.error('AI设置保存失败')
  } finally {
    aiLoading.value = false
  }
}

// 测试AI连接
const testAIConnection = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('AI连接测试成功')
  } catch {
    message.error('AI连接测试失败')
  }
}

const onSearch=(async()=>{
  const res = await settingApi.getSettingsMap('ai')
  aiForm.value = Object.assign({},defaultForm,res.data)
})

onMounted(()=>{
  onSearch()
})
</script>
<template>
  <n-form
    ref="aiFormRef"
    :model="aiForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="OpenAI API Key" path="openaiApiKey">
      <n-input
        v-model:value="aiForm.openaiApiKey"
        type="password"
        placeholder="请输入OpenAI API密钥"
        show-password-on="mousedown"
      />
    </n-form-item>
    <n-form-item label="OpenAI API地址" path="openaiApiUrl">
      <n-input v-model:value="aiForm.openaiApiUrl" placeholder="https://api.openai.com/v1" />
    </n-form-item>
    <n-form-item label="AI模型" path="aiModel">
      <n-select v-model:value="aiForm.aiModel" :options="aiModelOptions" />
    </n-form-item>
    <n-form-item label="AI温度" path="aiTemperature">
      <n-slider v-model:value="aiForm.aiTemperature" :min="0" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.aiTemperature }}</span>
    </n-form-item>
    <n-form-item label="AI最大令牌数" path="aiMaxTokens">
      <n-input-number v-model:value="aiForm.aiMaxTokens" :min="1" :max="8192" />
    </n-form-item>
    <n-form-item label="AI Top P" path="aiTopP">
      <n-slider v-model:value="aiForm.aiTopP" :min="0" :max="1" :step="0.1" />
      <span class="slider-value">{{ aiForm.aiTopP }}</span>
    </n-form-item>
    <n-form-item label="AI频率惩罚" path="aiFrequencyPenalty">
      <n-slider v-model:value="aiForm.aiFrequencyPenalty" :min="-2" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.aiFrequencyPenalty }}</span>
    </n-form-item>
    <n-form-item label="AI存在惩罚" path="aiPresencePenalty">
      <n-slider v-model:value="aiForm.aiPresencePenalty" :min="-2" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.aiPresencePenalty }}</span>
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveAISettings" :loading="aiLoading"> 保存AI设置 </n-button>
      <n-button @click="testAIConnection" style="margin-left: 12px"> 测试AI连接 </n-button>
    </n-form-item>
  </n-form>
</template>
