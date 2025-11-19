<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import type { SettingsProps } from '../utils/types';
import * as settingApi from '@/api/system/setting'

const props = defineProps<{
  settings:SettingsProps
}>()
const emit = defineEmits(["refresh"])
const message = useMessage()

const seoFormRef = ref<FormInst | null>(null)
const seoLoading = ref(false)

// SEO设置表单
const seoForm = reactive({
  enableSEO: true,
  siteTitleSuffix: '',
  siteKeywords: '',
  siteDescription: '',
  siteAuthor: '',
  siteCopyright: '',
  siteIcp: '',
  sitePoliceIcp: '',
  analyticsCode: '',
  robotsContent: '',
  autoGenerateSitemap: true,
  urlRewrite: true,
  custom404Content: ''
})

watch(()=>props.settings,(newSettings)=>{
  seoForm.enableSEO = newSettings.enableSEO || true
  seoForm.siteTitleSuffix = newSettings.siteTitleSuffix || ''
  seoForm.siteKeywords = newSettings.siteKeywords || ''
  seoForm.siteDescription = newSettings.siteDescription || ''
  seoForm.siteAuthor = newSettings.siteAuthor || ''
  seoForm.siteCopyright = newSettings.siteCopyright || ''
  seoForm.siteIcp = newSettings.siteIcp || ''
  seoForm.sitePoliceIcp = newSettings.sitePoliceIcp || ''
  seoForm.analyticsCode = newSettings.analyticsCode || ''
  seoForm.robotsContent = newSettings.robotsContent || ''
  seoForm.autoGenerateSitemap = newSettings.autoGenerateSitemap || true
  seoForm.urlRewrite = newSettings.urlRewrite || true
  seoForm.custom404Content = newSettings.custom404Content || ''
})

// 保存SEO设置
const saveSEOSettings = async () => {
  seoLoading.value = true
  try {
    await settingApi.saveSettings(seoForm)
    await new Promise(resolve => setTimeout(resolve, 1000))
    emit('refresh')
    message.success('SEO设置保存成功')
  } catch {
    message.error('SEO设置保存失败')
  } finally {
    seoLoading.value = false
  }
}

// 生成Sitemap
const generateSitemap = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('Sitemap生成成功')
  } catch {
    message.error('Sitemap生成失败')
  }
}
</script>
<template>
  <n-form
    ref="seoFormRef"
    :model="seoForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="启用SEO优化" path="enableSEO">
      <n-switch v-model:value="seoForm.enableSEO" />
    </n-form-item>
    <n-form-item label="网站标题后缀" path="siteTitleSuffix">
      <n-input v-model:value="seoForm.siteTitleSuffix" placeholder="例如: - 我的网站" />
    </n-form-item>
    <n-form-item label="网站关键词" path="siteKeywords">
      <n-input v-model:value="seoForm.siteKeywords" placeholder="请输入网站关键词，用逗号分隔" />
    </n-form-item>
    <n-form-item label="网站描述" path="siteDescription">
      <n-input
        v-model:value="seoForm.siteDescription"
        type="textarea"
        placeholder="请输入网站描述"
        :rows="3"
      />
    </n-form-item>
    <n-form-item label="作者信息" path="siteAuthor">
      <n-input v-model:value="seoForm.siteAuthor" placeholder="请输入网站作者" />
    </n-form-item>
    <n-form-item label="版权信息" path="siteCopyright">
      <n-input v-model:value="seoForm.siteCopyright" placeholder="请输入版权信息" />
    </n-form-item>
    <n-form-item label="ICP备案号" path="siteIcp">
      <n-input v-model:value="seoForm.siteIcp" placeholder="请输入ICP备案号" />
    </n-form-item>
    <n-form-item label="公安备案号" path="sitePoliceIcp">
      <n-input v-model:value="seoForm.sitePoliceIcp" placeholder="请输入公安备案号" />
    </n-form-item>
    <n-form-item label="统计代码" path="analyticsCode">
      <n-input
        v-model:value="seoForm.analyticsCode"
        type="textarea"
        placeholder="请输入百度统计、Google Analytics等统计代码"
        :rows="4"
      />
    </n-form-item>
    <n-form-item label=" robots.txt 内容" path="robotsContent">
      <n-input
        v-model:value="seoForm.robotsContent"
        type="textarea"
        placeholder="请输入robots.txt内容"
        :rows="6"
      />
    </n-form-item>
    <n-form-item label="Sitemap自动生成" path="autoGenerateSitemap">
      <n-switch v-model:value="seoForm.autoGenerateSitemap" />
    </n-form-item>
    <n-form-item label="URL伪静态" path="urlRewrite">
      <n-switch v-model:value="seoForm.urlRewrite" />
    </n-form-item>
    <n-form-item label="404页面自定义内容" path="custom404Content">
      <n-input
        v-model:value="seoForm.custom404Content"
        type="textarea"
        placeholder="请输入自定义404页面内容"
        :rows="4"
      />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveSEOSettings" :loading="seoLoading">
        保存SEO设置
      </n-button>
      <n-button @click="generateSitemap" style="margin-left: 12px"> 生成Sitemap </n-button>
    </n-form-item>
  </n-form>
</template>
