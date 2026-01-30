<template>
  <div class="theme-detail">
    <n-descriptions bordered label-placement="left" :column="1">
      <n-descriptions-item label="主题类型">
        {{ theme.type === 'internal' ? '内部主题' : '外部主题' }}
      </n-descriptions-item>
      <n-descriptions-item label="主题名称">
        {{ theme.name }}
      </n-descriptions-item>
      <n-descriptions-item label="显示名称">
        {{ theme.display_name }}
      </n-descriptions-item>
      <n-descriptions-item label="版本">
        {{ theme.version }}
      </n-descriptions-item>
      <n-descriptions-item label="作者">
        {{ theme.author_name || '-' }}
      </n-descriptions-item>
      <n-descriptions-item label="作者邮箱">
        {{ theme.author_email || '-' }}
      </n-descriptions-item>
      <n-descriptions-item label="描述">
        {{ theme.description || '-' }}
      </n-descriptions-item>
      <n-descriptions-item label="外部URL" v-if="theme.type === 'external'">
        <n-a v-if="theme.external_url" :href="theme.external_url" target="_blank">{{ theme.external_url }}</n-a>
        <span v-else>-</span>
      </n-descriptions-item>
      <n-descriptions-item label="文件路径" v-if="theme.type === 'internal'">
        {{ theme.path || '-' }}
      </n-descriptions-item>
      <n-descriptions-item label="主页">
        <n-a v-if="theme.homepage" :href="theme.homepage" target="_blank">{{ theme.homepage }}</n-a>
        <span v-else>-</span>
      </n-descriptions-item>
      <n-descriptions-item label="仓库">
        <n-a v-if="theme.repo" :href="theme.repo" target="_blank">{{ theme.repo }}</n-a>
        <span v-else>-</span>
      </n-descriptions-item>
      <n-descriptions-item label="状态">
        <n-tag :type="theme.enabled ? 'success' : 'default'">
          {{ theme.enabled ? '当前默认主题 (已启用)' : '未启用' }}
        </n-tag>
      </n-descriptions-item>
    </n-descriptions>
  </div>
</template>

<script setup lang="ts">
import { NDescriptions, NDescriptionsItem, NTag, NA } from 'naive-ui'

const props = defineProps({
  theme: {
    type: Object,
    required: true,
  },
})
</script>

<style scoped>
.theme-detail {
  padding: 20px;
}
</style>
