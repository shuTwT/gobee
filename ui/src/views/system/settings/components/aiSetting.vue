<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import * as aiApi from '@/api/ai'
import type { AiSetting } from '../utils/types'

const message = useMessage()
const aiFormRef = ref<FormInst | null>(null)
const aiLoading = ref(false)
const modelLoading = ref(false)
const configured = ref(false)
const aiModelOptions = ref<{ label: string; value: string }[]>([])

const defaultForm = (): AiSetting => ({
  base_url: 'https://api.openai.com/v1',
  api_key: '',
  model: '',
  temperature: 0.7,
  max_tokens: 2048,
  top_p: 1,
  frequency_penalty: 0,
  presence_penalty: 0,
})

const aiForm = ref<AiSetting>(defaultForm())

const loadConfig = async () => {
  try {
    const res = await aiApi.getConfig()
    const defaults = defaultForm()
    const data = res.data
    aiForm.value = {
      ...defaults,
      base_url: data.base_url || defaults.base_url,
      model: data.model || defaults.model,
      temperature: data.temperature ?? defaults.temperature,
      max_tokens: data.max_tokens || defaults.max_tokens,
      top_p: data.top_p ?? defaults.top_p,
      frequency_penalty: data.frequency_penalty ?? defaults.frequency_penalty,
      presence_penalty: data.presence_penalty ?? defaults.presence_penalty,
      api_key: '',
    }
    configured.value = data.api_key_configured
    if (data.model) aiModelOptions.value = [{ label: data.model, value: data.model }]
  } catch {
    message.error('加载 AI 配置失败')
  }
}

const saveAISettings = async () => {
  if (!aiForm.value.api_key.trim()) {
    message.warning('保存配置时必须重新输入 API Key')
    return
  }
  aiLoading.value = true
  try {
    await aiApi.saveConfig(aiForm.value)
    configured.value = true
    aiForm.value.api_key = ''
    message.success('AI 设置保存成功')
  } catch {
    message.error('AI 设置保存失败')
  } finally {
    aiLoading.value = false
  }
}

const testAIConnection = async () => {
  if (!aiForm.value.api_key.trim()) {
    message.warning('测试连接时请输入 API Key')
    return
  }
  aiLoading.value = true
  try {
    await aiApi.testConfig({
      base_url: aiForm.value.base_url,
      api_key: aiForm.value.api_key,
    })
    message.success('AI 连接测试成功')
  } catch {
    message.error('AI 连接测试失败')
  } finally {
    aiLoading.value = false
  }
}

const refreshModelList = async () => {
  modelLoading.value = true
  try {
    const res = await aiApi.getModels()
    aiModelOptions.value = res.data.map(item => ({ label: item.id, value: item.id }))
    if (aiModelOptions.value.length > 0 && !aiModelOptions.value.some(item => item.value === aiForm.value.model)) {
      aiForm.value.model = aiModelOptions.value[0].value
    }
  } catch {
    message.error('刷新模型列表失败，请先保存配置')
  } finally {
    modelLoading.value = false
  }
}

onMounted(loadConfig)
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
    <n-alert v-if="configured" type="success" class="mb-4">
      当前已配置 AI 服务。出于安全原因，API Key 不会回显；每次保存或测试都需要重新输入。
    </n-alert>
    <n-alert v-else type="warning" class="mb-4">
      尚未配置 AI 服务。请填写完整配置后保存。
    </n-alert>
    <n-form-item label="OpenAI API Key" path="api_key">
      <n-input
        v-model:value="aiForm.api_key"
        type="password"
        placeholder="请输入 API Key（不会回显）"
        show-password-on="mousedown"
      />
    </n-form-item>
    <n-form-item label="OpenAI API 地址" path="base_url">
      <n-input v-model:value="aiForm.base_url" placeholder="https://api.openai.com/v1" />
    </n-form-item>
    <n-form-item label="AI 模型" path="model">
      <div class="w-full flex gap-2">
        <n-select v-model:value="aiForm.model" :options="aiModelOptions" filterable tag />
        <n-button :loading="modelLoading" @click="refreshModelList">刷新模型</n-button>
      </div>
    </n-form-item>
    <n-form-item label="AI 温度" path="temperature">
      <n-slider v-model:value="aiForm.temperature" :min="0" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.temperature }}</span>
    </n-form-item>
    <n-form-item label="AI 最大令牌数" path="max_tokens">
      <n-input-number v-model:value="aiForm.max_tokens" :min="1" :max="8192" />
    </n-form-item>
    <n-form-item label="AI Top P" path="top_p">
      <n-slider v-model:value="aiForm.top_p" :min="0" :max="1" :step="0.1" />
      <span class="slider-value">{{ aiForm.top_p }}</span>
    </n-form-item>
    <n-form-item label="AI 频率惩罚" path="frequency_penalty">
      <n-slider v-model:value="aiForm.frequency_penalty" :min="-2" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.frequency_penalty }}</span>
    </n-form-item>
    <n-form-item label="AI 存在惩罚" path="presence_penalty">
      <n-slider v-model:value="aiForm.presence_penalty" :min="-2" :max="2" :step="0.1" />
      <span class="slider-value">{{ aiForm.presence_penalty }}</span>
    </n-form-item>
    <n-form-item>
      <n-button type="primary" :loading="aiLoading" @click="saveAISettings">保存 AI 设置</n-button>
      <n-button :loading="aiLoading" class="ml-3" @click="testAIConnection">测试 AI 连接</n-button>
    </n-form-item>
  </n-form>
</template>
