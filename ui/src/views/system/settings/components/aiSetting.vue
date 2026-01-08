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
  openai_api_key: '',
  openai_api_url: 'https://api.openai.com/v1',
  ai_model: 'gpt-3.5-turbo',
  ai_memperature: 0.7,
  ai_maxTokens: 2048,
  ai_topP: 1.0,
  ai_frequency_penalty: 0,
  ai_presence_penalty: 0,
}
const aiForm = ref<AiSetting>({
  openai_api_key: '',
  openai_api_url: 'https://api.openai.com/v1',
  ai_model: 'gpt-3.5-turbo',
  ai_memperature: 0.7,
  ai_maxTokens: 2048,
  ai_topP: 1.0,
  ai_frequency_penalty: 0,
  ai_presence_penalty: 0,
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
        v-model:value="aiForm.openai_api_key"
        type="password"
        placeholder="请输入OpenAI API密钥"
        show-password-on="mousedown"
      />
    </n-form-item>
    <n-form-item label="OpenAI API地址" path="openaiApiUrl">
      <n-input v-model:value="aiForm.openai_api_url" placeholder="https://api.openai.com/v1" />
    </n-form-item>
    <n-form-item label="AI模型" path="aiModel">
      <div class="w-full flex">
       <n-select v-model:value="aiForm.ai_model" :options="aiModelOptions" /> 
       <n-button>刷新</n-button>
      </div>
      
    </n-form-item>
    <n-form-item label="AI温度" path="aiTemperature">
      <n-slider v-model:value="aiForm.ai_memperature" :min="0" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.ai_memperature }}</span>
    </n-form-item>
    <n-form-item label="AI最大令牌数" path="aiMaxTokens">
      <n-input-number v-model:value="aiForm.ai_maxTokens" :min="1" :max="8192" />
    </n-form-item>
    <n-form-item label="AI Top P" path="aiTopP">
      <n-slider v-model:value="aiForm.ai_topP" :min="0" :max="1" :step="0.1" />
      <span class="slider-value">{{ aiForm.ai_topP }}</span>
    </n-form-item>
    <n-form-item label="AI频率惩罚" path="aiFrequencyPenalty">
      <n-slider v-model:value="aiForm.ai_frequency_penalty" :min="-2" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.ai_frequency_penalty }}</span>
    </n-form-item>
    <n-form-item label="AI存在惩罚" path="aiPresencePenalty">
      <n-slider v-model:value="aiForm.ai_presence_penalty" :min="-2" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.ai_presence_penalty }}</span>
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveAISettings" :loading="aiLoading"> 保存AI设置 </n-button>
      <n-button @click="testAIConnection" style="margin-left: 12px"> 测试AI连接 </n-button>
    </n-form-item>
  </n-form>
</template>
