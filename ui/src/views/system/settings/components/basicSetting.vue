<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import type { SettingsProps } from '../utils/types';

const props = defineProps<{
  settings:SettingsProps
}>()
const emit = defineEmits(["refresh"])
const message = useMessage()
// 选项数据
const languageOptions = [
  { label: '简体中文', value: 'zh-CN' },
  { label: '繁體中文', value: 'zh-TW' },
  { label: 'English', value: 'en-US' },
  { label: '日本語', value: 'ja-JP' }
]

const timezoneOptions = [
  { label: '北京时间 (UTC+8)', value: 'Asia/Shanghai' },
  { label: '东京时间 (UTC+9)', value: 'Asia/Tokyo' },
  { label: '纽约时间 (UTC-5)', value: 'America/New_York' },
  { label: '伦敦时间 (UTC+0)', value: 'Europe/London' }
]

const formRef = ref<FormInst | null>(null)

const basicForm = reactive({
  siteName: '',
  siteDescription: '',
  siteLogo: '',
  siteFavicon: '',
  keywords: '',
  author: '',
  language: 'zh-CN',
  timezone: 'Asia/Shanghai',
  dateFormat: 'YYYY-MM-DD',
  timeFormat: 'HH:mm:ss'
})

watch(()=>props.settings,(newSettings)=>{
  basicForm.siteName = newSettings.siteName ?? ''
  basicForm.siteDescription = newSettings.siteDescription ?? ''
  basicForm.siteLogo = newSettings.siteLogo ?? ''
  basicForm.siteFavicon = newSettings.siteFavicon ?? ''
  basicForm.author = newSettings.author ?? ''
  basicForm.language = newSettings.language ?? 'zh-CN'
  basicForm.timezone = newSettings.timezone ?? 'Asia/Shanghai'
  basicForm.dateFormat = newSettings.dateFormat ?? 'YYYY-MM-DD'
  basicForm.timeFormat = newSettings.timeFormat ?? 'HH:mm:ss'
})

const basicLoading = ref(false)

// 保存基本设置
const saveBasicSettings = async () => {
  basicLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    emit('refresh')
    message.success('基本设置保存成功')
  } catch {
    message.error('基本设置保存失败')
  } finally {
    basicLoading.value = false
  }
}
</script>
<template>
  <n-form
    ref="formRef"
    :model="basicForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="网站名称" path="siteName">
      <n-input v-model:value="basicForm.siteName" placeholder="请输入网站名称" />
    </n-form-item>
    <n-form-item label="网站描述" path="siteDescription">
      <n-input
        v-model:value="basicForm.siteDescription"
        type="textarea"
        placeholder="请输入网站描述"
      />
    </n-form-item>
    <n-form-item label="网站Logo" path="siteLogo">
      <n-input v-model:value="basicForm.siteLogo" placeholder="请输入Logo地址" />
    </n-form-item>
    <n-form-item label="网站图标" path="siteFavicon">
      <n-input v-model:value="basicForm.siteFavicon" placeholder="请输入图标地址" />
    </n-form-item>
    <n-form-item label="关键词" path="keywords">
      <n-input v-model:value="basicForm.keywords" placeholder="请输入关键词，用逗号分隔" />
    </n-form-item>
    <n-form-item label="作者" path="author">
      <n-input v-model:value="basicForm.author" placeholder="请输入作者名称" />
    </n-form-item>
    <n-form-item label="语言" path="language">
      <n-select v-model:value="basicForm.language" :options="languageOptions" />
    </n-form-item>
    <n-form-item label="时区" path="timezone">
      <n-select v-model:value="basicForm.timezone" :options="timezoneOptions" />
    </n-form-item>
    <n-form-item label="日期格式" path="dateFormat">
      <n-input v-model:value="basicForm.dateFormat" placeholder="例如: YYYY-MM-DD" />
    </n-form-item>
    <n-form-item label="时间格式" path="timeFormat">
      <n-input v-model:value="basicForm.timeFormat" placeholder="例如: HH:mm:ss" />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveBasicSettings" :loading="basicLoading">
        保存基本设置
      </n-button>
    </n-form-item>
  </n-form>
</template>
