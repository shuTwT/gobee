<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import { apiClient, useApi } from '@/api'
import ImageUpload from '@/components/upload/ImageUpload.vue'

const message = useMessage()

const siteFormRef = ref<FormInst | null>(null)
const siteLoading = ref(false)

// 基本信息（原基本设置，仍存储于 basic 组）
const defaultBasicForm = {
  siteName: '',
  siteDescription: '',
  siteLogo: '',
  siteFavicon: '',
  siteUrl: '',
  keywords: '',
  author: '',
  language: 'zh-CN',
  icpBeian: '',
  gonganBeian: '',
  siteAnnouncement: '',
}
const basicForm = ref({ ...defaultBasicForm })

// 站点设置（存储于 site 组）
const defaultSiteForm = {
  maintenanceMode: false,
  allowRegistration: true,
  emailVerification: true,
  commentModeration: true,
  enableCDN: false,
  cdnUrl: '',
}
const siteForm = ref({ ...defaultSiteForm })

// 保存站点设置（基本信息与站点设置分两组保存，保持 basic 组数据兼容）
const saveSiteSettings = async () => {
  siteLoading.value = true
  try {
    await Promise.all([
      useApi(apiClient.api.v1SettingsJsonSaveCreate, 'basic', basicForm.value),
      useApi(apiClient.api.v1SettingsJsonSaveCreate, 'site', siteForm.value),
    ])
    await new Promise((resolve) => setTimeout(resolve, 1000))
    onSearch()
    message.success('站点设置保存成功')
  } catch {
    message.error('站点设置保存失败')
  } finally {
    siteLoading.value = false
  }
}

const onSearch = async () => {
  const [basicRes, siteRes] = await Promise.all([
    useApi(apiClient.api.v1SettingsJsonDetail, 'basic'),
    useApi(apiClient.api.v1SettingsJsonDetail, 'site'),
  ])
  basicForm.value = { ...defaultBasicForm, ...basicRes.data }
  siteForm.value = { ...defaultSiteForm, ...siteRes.data }
}

onMounted(() => {
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
      <ImageUpload
        :file-list="basicForm.siteLogo ? [basicForm.siteLogo] : []"
        @update:file-list="(list) => (basicForm.siteLogo = list[0] || '')"
      />
    </n-form-item>
    <n-form-item label="网站图标" path="siteFavicon">
      <ImageUpload
        :file-list="basicForm.siteFavicon ? [basicForm.siteFavicon] : []"
        @update:file-list="(list) => (basicForm.siteFavicon = list[0] || '')"
      />
    </n-form-item>
    <n-form-item label="站点地址" path="siteUrl">
      <n-input v-model:value="basicForm.siteUrl" placeholder="例如: https://example.com" />
    </n-form-item>
    <n-form-item label="关键词" path="keywords">
      <n-input v-model:value="basicForm.keywords" placeholder="请输入关键词，用逗号分隔" />
    </n-form-item>
    <n-form-item label="作者" path="author">
      <n-input v-model:value="basicForm.author" placeholder="请输入作者名称" />
    </n-form-item>
    <n-form-item label="首选语言" path="language">
      <n-select
        v-model:value="basicForm.language"
        :options="[
          { label: '简体中文', value: 'zh-CN' },
          { label: '繁體中文', value: 'zh-TW' },
          { label: 'English', value: 'en-US' },
          { label: '日本語', value: 'ja-JP' },
        ]"
      />
    </n-form-item>
    <n-form-item label="ICP备案号" path="icpBeian">
      <n-input v-model:value="basicForm.icpBeian" placeholder="请输入ICP备案号" />
    </n-form-item>
    <n-form-item label="公安备案号" path="gonganBeian">
      <n-input v-model:value="basicForm.gonganBeian" placeholder="请输入公安备案号" />
    </n-form-item>
    <n-form-item label="站点公告（HTML）" path="siteAnnouncement">
      <n-input
        v-model:value="basicForm.siteAnnouncement"
        type="textarea"
        :rows="5"
        placeholder="请输入站点公告，支持 HTML，例如：&lt;p&gt;欢迎访问本站&lt;/p&gt;"
      />
    </n-form-item>
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
