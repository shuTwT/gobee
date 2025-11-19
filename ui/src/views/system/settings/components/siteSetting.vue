<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import type { SettingsProps } from '../utils/types';
import * as settingApi from '@/api/system/setting'

const props = defineProps<{
  settings:SettingsProps
}>()
const emit = defineEmits(["refresh"])
const message = useMessage()

const siteFormRef = ref<FormInst | null>(null)
const siteLoading = ref(false)

// 站点设置表单
const siteForm = reactive({
  maintenanceMode: false,
  allowRegistration: true,
  emailVerification: true,
  commentModeration: true,
  uploadLimit: 10,
  allowedFileTypes: ['jpg', 'jpeg', 'png', 'gif', 'pdf', 'doc', 'docx'],
  enableCache: true,
  cacheTime: 60,
  enableSSL: false,
  enableCDN: false,
  cdnUrl: '',
})

watch(()=>props.settings,(newSettings)=>{
  siteForm.maintenanceMode = newSettings.maintenanceMode || false
  siteForm.allowRegistration = newSettings.allowRegistration || true
  siteForm.emailVerification = newSettings.emailVerification || true
  siteForm.commentModeration = newSettings.commentModeration || true
  siteForm.uploadLimit = newSettings.uploadLimit || 10
  siteForm.allowedFileTypes = newSettings.allowedFileTypes || ['jpg', 'jpeg', 'png', 'gif', 'pdf', 'doc', 'docx']
  siteForm.enableCache = newSettings.enableCache || true
  siteForm.cacheTime = newSettings.cacheTime || 60
  siteForm.enableSSL = newSettings.enableSSL || false
  siteForm.enableCDN = newSettings.enableCDN || false
  siteForm.cdnUrl = newSettings.cdnUrl || ''
})

// 保存站点设置
const saveSiteSettings = async () => {
  siteLoading.value = true
  try {
    await settingApi.saveSettings(siteForm)
    await new Promise((resolve) => setTimeout(resolve, 1000))
    emit('refresh')
    message.success('站点设置保存成功')
  } catch {
    message.error('站点设置保存失败')
  } finally {
    siteLoading.value = false
  }
}
</script>
<template>
  <n-form
    ref="siteFormRef"
    :model="siteForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="维护模式" path="maintenanceMode">
      <n-switch v-model:value="siteForm.maintenanceMode" />
    </n-form-item>
    <n-form-item label="允许注册" path="allowRegistration">
      <n-switch v-model:value="siteForm.allowRegistration" />
    </n-form-item>
    <n-form-item label="邮箱验证" path="emailVerification">
      <n-switch v-model:value="siteForm.emailVerification" />
    </n-form-item>
    <n-form-item label="评论审核" path="commentModeration">
      <n-switch v-model:value="siteForm.commentModeration" />
    </n-form-item>
    <n-form-item label="文件上传限制" path="uploadLimit">
      <n-input-number v-model:value="siteForm.uploadLimit" :min="1" :max="100">
        <template #suffix>MB</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="允许的文件类型" path="allowedFileTypes">
      <n-dynamic-tags v-model:value="siteForm.allowedFileTypes" />
    </n-form-item>
    <n-form-item label="启用缓存" path="enableCache">
      <n-switch v-model:value="siteForm.enableCache" />
    </n-form-item>
    <n-form-item label="缓存时间" path="cacheTime">
      <n-input-number v-model:value="siteForm.cacheTime" :min="1">
        <template #suffix>分钟</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="启用SSL" path="enableSSL">
      <n-switch v-model:value="siteForm.enableSSL" />
    </n-form-item>
    <n-form-item label="启用CDN" path="enableCDN">
      <n-switch v-model:value="siteForm.enableCDN" />
    </n-form-item>
    <n-form-item label="CDN地址" path="cdnUrl">
      <n-input v-model:value="siteForm.cdnUrl" placeholder="请输入CDN地址" />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveSiteSettings" :loading="siteLoading">
        保存站点设置
      </n-button>
    </n-form-item>
  </n-form>
</template>
