<script setup lang="ts">
import { ref, watch } from 'vue'
import {
  NSpin,
  NButton,
  NTag,
  NDivider,
  NEmpty,
  useMessage,
} from 'naive-ui'
import Editor from '@tinymce/tinymce-vue'
import * as doclibrarydetailApi from '@/api/content/doclibrarydetail'
import dayjs from 'dayjs'

interface DocLibraryDetail {
  id: number
  library_id: number
  title: string
  version: string
  content: string
  parent_id: number | null
  path: string
  url: string
  language: string
  created_at: string
  updated_at: string
  children?: DocLibraryDetail[]
}

interface Props {
  selectedNodeId: number | null
}

const message = useMessage()

const props = defineProps<Props>()

const rightPanelData = ref<DocLibraryDetail | null>(null)
const rightPanelLoading = ref(false)
const rightPanelSaving = ref(false)

const initConfig = {
  height: 'calc(100vh - 450px)',
  menubar: true,
  plugins: [
    'advlist', 'autolink', 'lists', 'link', 'image', 'charmap', 'preview',
    'anchor', 'searchreplace', 'visualblocks', 'code', 'fullscreen',
    'insertdatetime', 'media', 'table', 'help', 'wordcount'
  ],
  toolbar: 'undo redo | blocks | ' +
    'bold italic backcolor | alignleft aligncenter ' +
    'alignright alignjustify | bullist numlist outdent indent | ' +
    'removeformat | help',
  content_style: 'body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; font-size: 14px; padding: 20px; }',
}

const languageOptions = [
  { label: '中文', value: 'zh' },
  { label: '英文', value: 'en' },
  { label: '日文', value: 'ja' },
  { label: '韩文', value: 'ko' },
]

const fetchRightPanelData = async (id: number) => {
  rightPanelLoading.value = true
  try {
    const res = await doclibrarydetailApi.getDocLibraryDetail(id)
    rightPanelData.value = res.data
  } catch (error) {
    console.error('获取文档详情失败:', error)
    message.error('获取文档详情失败')
    rightPanelData.value = null
  } finally {
    rightPanelLoading.value = false
  }
}

const handleRightPanelUpdate = async () => {
  if (!rightPanelData.value) return

  rightPanelSaving.value = true
  try {
    await doclibrarydetailApi.updateDocLibraryDetail(rightPanelData.value.id, {
      content: rightPanelData.value.content,
    })
    message.success('保存成功')
  } catch (error) {
    console.error('保存失败:', error)
    message.error('保存失败')
  } finally {
    rightPanelSaving.value = false
  }
}

watch(() => props.selectedNodeId, (newId) => {
  if (newId) {
    fetchRightPanelData(newId)
  } else {
    rightPanelData.value = null
  }
})
</script>

<template>
  <div class="right-panel">
    <div class="panel-header">
      <span class="panel-title">文档详情</span>
      <div v-if="rightPanelData" class="panel-actions">
        <NButton size="small" :loading="rightPanelSaving" @click="handleRightPanelUpdate">
          <template #icon>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path>
              <polyline points="17 21 17 13 7 13 7 21"></polyline>
              <polyline points="7 3 7 8 15 8"></polyline>
            </svg>
          </template>
          保存
        </NButton>
      </div>
    </div>
    <NSpin :show="rightPanelLoading">
      <div v-if="rightPanelData" class="right-panel-content">
        <div class="document-editor">
          <Editor v-model="rightPanelData.content" :init="initConfig" />
        </div>
        <NDivider />
      </div>
      <NEmpty v-else description="请选择左侧文档查看详情" />
    </NSpin>
  </div>
</template>

<style scoped>
.right-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.panel-header {
  padding: 12px 0;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.panel-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
}

.panel-actions {
  display: flex;
  gap: 8px;
}

.right-panel-content {
  flex: 1;
  overflow-y: auto;
  padding: 8px 0;
}

.document-header {
  margin-bottom: 16px;
}

.document-title {
  font-size: 24px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0 0 12px 0;
}

.document-meta {
  display: flex;
  gap: 8px;
}

.document-editor {
  flex: 1;
  overflow: hidden;
  margin-bottom: 16px;
}

.document-editor :deep(.tox-tinymce) {
  border: 1px solid #e5e7eb;
  border-radius: 4px;
}

.document-editor :deep(.tox-editor-header) {
  border-bottom: 1px solid #e5e7eb;
}

.document-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  align-items: center;
}

.info-label {
  font-weight: 500;
  color: #6b7280;
  min-width: 100px;
}

.info-value {
  color: #1a1a1a;
  flex: 1;
}

.link-text {
  color: #3b82f6;
  text-decoration: none;
}

.link-text:hover {
  text-decoration: underline;
}

@media (max-width: 1024px) {
  .right-panel {
    flex: none;
    width: 100%;
    min-height: 400px;
  }
}
</style>
