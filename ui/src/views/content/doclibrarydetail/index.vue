<script setup lang="ts">
import { ref, onMounted, h } from 'vue'
import { useRoute } from 'vue-router'
import {
  NButton,
  NSpace,
  NModal,
  NInput,
  NSelect,
  NForm,
  NFormItem,
  NCard,
  NDropdown,
  useMessage,
  useDialog,
  type TreeOption,
  type DropdownOption,
} from 'naive-ui'
import * as doclibraryApi from '@/api/content/doclibrary'
import * as doclibrarydetailApi from '@/api/content/doclibrarydetail'
import LeftPanel from './components/LeftPanel.vue'
import RightPanel from './components/RightPanel.vue'

const route = useRoute()
const message = useMessage()
const dialog = useDialog()

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

interface DocLibraryDetailFormData {
  library_id: number
  title: string
  version: string
  path: string
  url: string
  language: string
}

const loading = ref(false)
const docLibraries = ref<any[]>([])
const selectedLibraryId = ref<number | null>(null)
const selectedNodeKey = ref<number | null>(null)

const showModal = ref(false)
const isEdit = ref(false)
const isSettings = ref(false)
const currentId = ref<number | null>(null)
const formData = ref<DocLibraryDetailFormData>({
  library_id: 0,
  title: '',
  version: '',
  path: '',
  url: '',
  language: 'zh',
})

const languageOptions = [
  { label: '中文', value: 'zh' },
  { label: '英文', value: 'en' },
  { label: '日文', value: 'ja' },
  { label: '韩文', value: 'ko' },
]

const contextMenuOptions: DropdownOption[] = [
  {
    label: '新建文档',
    key: 'create',
    icon: () => h('svg', { width: 16, height: 16, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2 }, [
      h('line', { x1: 12, y1: 5, x2: 12, y2: 19 }),
      h('line', { x1: 5, y1: 12, x2: 19, y2: 12 })
    ])
  },
  {
    label: '设置',
    key: 'settings',
    icon: () => h('svg', { width: 16, height: 16, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2 }, [
      h('circle', { cx: 12, cy: 12, r: 3 }),
      h('path', { d: 'M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z' })
    ])
  },
  {
    label: '删除',
    key: 'delete',
    icon: () => h('svg', { width: 16, height: 16, viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': 2 }, [
      h('polyline', { points: '3 6 5 6 21 6' }),
      h('path', { d: 'M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2' })
    ])
  }
]

const contextMenuPosition = ref({ x: 0, y: 0 })
const showContextMenu = ref(false)
const contextMenuNode = ref<TreeOption | null>(null)

const fetchDocLibraries = async () => {
  try {
    const res = await doclibraryApi.getDocLibraryList()
    docLibraries.value = res.data
    
    const docLibraryIdFromRoute = route.query.docLibraryId as string
    
    if (docLibraryIdFromRoute) {
      const libraryId = parseInt(docLibraryIdFromRoute)
      if (docLibraries.value.some(lib => lib.id === libraryId)) {
        selectedLibraryId.value = libraryId
      } else {
        message.warning('指定的文档库不存在')
        if (docLibraries.value.length > 0) {
          selectedLibraryId.value = docLibraries.value[0].id
        }
      }
    } else if (docLibraries.value.length > 0 && !selectedLibraryId.value) {
      selectedLibraryId.value = docLibraries.value[0].id
    }
  } catch (error) {
    console.error('获取文档库列表失败:', error)
    message.error('获取文档库列表失败')
  }
}

const handleCreate = () => {
  if (!selectedLibraryId.value) {
    message.error('请先选择文档库')
    return
  }
  isEdit.value = false
  isSettings.value = false
  currentId.value = null
  formData.value = {
    library_id: selectedLibraryId.value,
    title: '',
    version: '',
    path: '',
    url: '',
    language: 'zh',
  }
  showModal.value = true
}

const handleCreateFromContext = () => {
  if (!selectedLibraryId.value) {
    message.error('请先选择文档库')
    return
  }
  isEdit.value = false
  isSettings.value = false
  currentId.value = null
  formData.value = {
    library_id: selectedLibraryId.value,
    title: '',
    version: '',
    path: '',
    url: '',
    language: 'zh',
  }
  showContextMenu.value = false
  showModal.value = true
}

const handleSettingsFromContext = () => {
  if (!contextMenuNode.value) return
  const data = contextMenuNode.value.data as DocLibraryDetail
  isEdit.value = true
  isSettings.value = true
  currentId.value = data.id
  formData.value = {
    library_id: data.library_id,
    title: data.title,
    version: data.version,
    path: data.path,
    url: data.url,
    language: data.language,
  }
  showContextMenu.value = false
  showModal.value = true
}

const handleDeleteFromContext = async () => {
  if (!contextMenuNode.value) return
  const data = contextMenuNode.value.data as DocLibraryDetail
  showContextMenu.value = false
  
  dialog.warning({
    title: '确认删除',
    content: `确定要删除文档"${data.title}"吗？`,
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: async () => {
      try {
        await doclibrarydetailApi.deleteDocLibraryDetail(data.id)
        message.success('删除成功')
        if (selectedNodeKey.value === data.id) {
          selectedNodeKey.value = null
        }
        fetchDocLibraries()
      } catch (error) {
        console.error('删除失败:', error)
        message.error('删除失败')
      }
    }
  })
}

const handleNodeSelect = (nodeId: number) => {
  selectedNodeKey.value = nodeId
}

const handleContextMenu = (node: TreeOption, position: { x: number, y: number }) => {
  contextMenuNode.value = node
  contextMenuPosition.value = position
  showContextMenu.value = true
}

const handleEdit = (node: TreeOption) => {
  const data = node.data as DocLibraryDetail
  isEdit.value = true
  isSettings.value = false
  currentId.value = data.id
  formData.value = {
    library_id: data.library_id,
    title: data.title,
    version: data.version,
    path: data.path,
    url: data.url,
    language: data.language,
  }
  showModal.value = true
}

const handleDelete = async (node: TreeOption) => {
  const data = node.data as DocLibraryDetail
  try {
    await doclibrarydetailApi.deleteDocLibraryDetail(data.id)
    message.success('删除成功')
    if (selectedNodeKey.value === data.id) {
      selectedNodeKey.value = null
    }
    fetchDocLibraries()
  } catch (error) {
    console.error('删除失败:', error)
    message.error('删除失败')
  }
}

const handleSubmit = async () => {
  if (!formData.value.title.trim()) {
    message.error('请输入文档标题')
    return
  }

  try {
    if (isEdit.value && currentId.value) {
      await doclibrarydetailApi.updateDocLibraryDetail(currentId.value, formData.value)
      message.success(isSettings.value ? '设置已保存' : '更新成功')
    } else {
      await doclibrarydetailApi.createDocLibraryDetail(formData.value)
      message.success('创建成功')
    }
    showModal.value = false
    fetchDocLibraries()
  } catch (error) {
    console.error('操作失败:', error)
    message.error(isEdit.value ? '更新失败' : '创建失败')
  }
}

const handleContextMenuSelect = (key: string) => {
  switch (key) {
    case 'create':
      handleCreateFromContext()
      break
    case 'settings':
      handleSettingsFromContext()
      break
    case 'delete':
      handleDeleteFromContext()
      break
  }
}

onMounted(() => {
  fetchDocLibraries()
})
</script>

<template>
  <div class="doclibrarydetail-container">
    <NCard :bordered="false" class="doclibrarydetail-card">
      <div class="doclibrarydetail-header">
        <h1 class="page-title">文档详情管理</h1>
      </div>

      <div class="content-layout">
        <LeftPanel
          :selected-library-id="selectedLibraryId"
          :doc-libraries="docLibraries"
          @update:selected-library-id="(value) => selectedLibraryId = value"
          @node-select="handleNodeSelect"
          @create="handleCreate"
          @context-menu="handleContextMenu"
        />

        <RightPanel :selected-node-id="selectedNodeKey" />
      </div>
    </NCard>

    <NDropdown
      placement="bottom-start"
      trigger="manual"
      :x="contextMenuPosition.x"
      :y="contextMenuPosition.y"
      :options="contextMenuOptions"
      :show="showContextMenu"
      :show-arrow="false"
      @clickoutside="showContextMenu = false"
      @select="handleContextMenuSelect"
    />

    <NModal
      v-model:show="showModal"
      preset="card"
      :title="isSettings ? '文档设置' : (isEdit ? '编辑文档' : '新建文档')"
      style="width: 600px"
      :mask-closable="false"
    >
      <NForm :model="formData" label-placement="left" label-width="100px">
        <NFormItem label="文档标题" required>
          <NInput v-model:value="formData.title" placeholder="请输入文档标题" />
        </NFormItem>
        <NFormItem label="版本">
          <NInput v-model:value="formData.version" placeholder="请输入文档版本" />
        </NFormItem>
        <NFormItem label="文档路径">
          <NInput v-model:value="formData.path" placeholder="请输入文档路径" />
        </NFormItem>
        <NFormItem label="文档URL">
          <NInput v-model:value="formData.url" placeholder="请输入文档URL" />
        </NFormItem>
        <NFormItem label="语言">
          <NSelect
            v-model:value="formData.language"
            :options="languageOptions"
            placeholder="请选择文档语言"
          />
        </NFormItem>
      </NForm>

      <template #footer>
        <NSpace justify="end">
          <NButton @click="showModal = false">取消</NButton>
          <NButton type="primary" @click="handleSubmit">
            {{ isSettings ? '保存' : (isEdit ? '保存' : '创建') }}
          </NButton>
        </NSpace>
      </template>
    </NModal>
  </div>
</template>

<style scoped>
.doclibrarydetail-container {
  padding: 24px;
  background-color: #f5f7fa;
  min-height: 100vh;
}

.doclibrarydetail-card {
  border-radius: 8px;
  min-height: calc(100vh - 48px);
}

.doclibrarydetail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #1a1a1a;
  margin: 0;
}

.content-layout {
  display: flex;
  gap: 24px;
  height: calc(100vh - 200px);
  min-height: 600px;
}

@media (max-width: 1024px) {
  .content-layout {
    flex-direction: column;
    height: auto;
    min-height: auto;
  }
}

@media (max-width: 768px) {
  .doclibrarydetail-container {
    padding: 16px;
  }
  
  .doclibrarydetail-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
}

@media (max-width: 480px) {
  .page-title {
    font-size: 20px;
  }
}
</style>
