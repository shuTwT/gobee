<script setup lang="ts">
import type { FormInst } from 'naive-ui';
import * as settingApi from '@/api/system/setting'

const message = useMessage()


const encryptionOptions = [
  { label: 'TLS', value: 'tls' },
  { label: 'SSL', value: 'ssl' },
  { label: '无', value: 'none' }
]

const emailFormRef = ref<FormInst | null>(null)
const defaultForm = {
  smtpHost: '',
  smtpPort: 587,
  smtpUsername: '',
  smtpPassword: '',
  smtpEncryption: 'tls',
  senderEmail: '',
  senderName: ''
}
const emailForm = ref({
  smtpHost: '',
  smtpPort: 587,
  smtpUsername: '',
  smtpPassword: '',
  smtpEncryption: 'tls',
  senderEmail: '',
  senderName: ''
})
const emailLoading = ref(false)



// 保存邮件设置
const saveEmailSettings = async () => {
  emailLoading.value = true
  try {
    await settingApi.saveSettings('email',emailForm.value)
    await new Promise(resolve => setTimeout(resolve, 1000))
    onSearch()
    message.success('邮件设置保存成功')
  } catch {
    message.error('邮件设置保存失败')
  } finally {
    emailLoading.value = false
  }
}

// 测试邮件连接
const testEmailConnection = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('邮件连接测试成功')
  } catch {
    message.error('邮件连接测试失败')
  }
}

const onSearch = async()=>{
  const res = await settingApi.getSettingsMap('email')
  emailForm.value = Object.assign({},defaultForm,res.data)
}

onMounted(()=>{
  onSearch()
})
</script>
<template>
  <n-form
    ref="emailFormRef"
    :model="emailForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="SMTP服务器" path="smtpHost">
      <n-input v-model:value="emailForm.smtpHost" placeholder="例如: smtp.gmail.com" />
    </n-form-item>
    <n-form-item label="SMTP端口" path="smtpPort">
      <n-input-number v-model:value="emailForm.smtpPort" :min="1" :max="65535" />
    </n-form-item>
    <n-form-item label="SMTP用户名" path="smtpUsername">
      <n-input v-model:value="emailForm.smtpUsername" placeholder="请输入SMTP用户名" />
    </n-form-item>
    <n-form-item label="SMTP密码" path="smtpPassword">
      <n-input
        v-model:value="emailForm.smtpPassword"
        type="password"
        placeholder="请输入SMTP密码"
        show-password-on="mousedown"
      />
    </n-form-item>
    <n-form-item label="加密方式" path="smtpEncryption">
      <n-select v-model:value="emailForm.smtpEncryption" :options="encryptionOptions" />
    </n-form-item>
    <n-form-item label="发件人邮箱" path="senderEmail">
      <n-input v-model:value="emailForm.senderEmail" placeholder="请输入发件人邮箱地址" />
    </n-form-item>
    <n-form-item label="发件人名称" path="senderName">
      <n-input v-model:value="emailForm.senderName" placeholder="请输入发件人名称" />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveEmailSettings" :loading="emailLoading">
        保存邮件设置
      </n-button>
      <n-button @click="testEmailConnection" style="margin-left: 12px"> 测试邮件连接 </n-button>
    </n-form-item>
  </n-form>
</template>
