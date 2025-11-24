<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import * as settingApi from '@/api/system/setting'

const message = useMessage()

const siteFormRef = ref<FormInst | null>(null)
const siteLoading = ref(false)

const defaultForm = {
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
}
// 站点设置表单
const siteForm = ref({
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


// 保存站点设置
const saveSiteSettings = async () => {
  siteLoading.value = true
  try {
    await settingApi.saveSettings('site',siteForm.value)
    await new Promise((resolve) => setTimeout(resolve, 1000))
    onSearch()
    message.success('站点设置保存成功')
  } catch {
    message.error('站点设置保存失败')
  } finally {
    siteLoading.value = false
  }
}

const onSearch = async()=>{
  const res = await settingApi.getSettingsMap('site')
  siteForm.value = Object.assign({},defaultForm,res.data)
}

onMounted(()=>{
  onSearch()
})
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
