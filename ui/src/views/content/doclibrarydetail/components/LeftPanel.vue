<script setup lang="ts">
import { ref, h, watch } from 'vue'
import {
  NTree,
  NSpin,
  NEmpty,
  NTag,
  NButton,
  NSelect,
  NCard,
  type TreeOption,
} from 'naive-ui'
import * as doclibraryApi from '@/api/content/doclibrary'
import * as doclibrarydetailApi from '@/api/content/doclibrarydetail'

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
  selectedLibraryId: number | null
  docLibraries: any[]
}

interface Emits {
  (e: 'update:selectedLibraryId', value: number | null): void
  (e: 'nodeSelect', nodeId: number): void
  (e: 'create'): void
  (e: 'contextMenu', node: TreeOption, position: { x: number, y: number }): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const treeData = ref<TreeOption[]>([])
const treeLoading = ref(false)
const selectedNodeKey = ref<number | null>(null)

const fetchTree = async () => {
  if (!props.selectedLibraryId) return
  
  treeLoading.value = true
  try {
    const res = await doclibrarydetailApi.getDocLibraryDetailTree({
      library_id: props.selectedLibraryId,
    })
    treeData.value = transformToTreeOptions(res.data)
  } catch (error) {
    console.error('获取文档树失败:', error)
  } finally {
    treeLoading.value = false
  }
}

const transformToTreeOptions = (data: DocLibraryDetail[]): TreeOption[] => {
  return data.map((item) => ({
    key: item.id,
    label: item.title,
    version: item.version,
    children: item.children ? transformToTreeOptions(item.children) : undefined,
    data: item,
  }))
}

const handleNodeSelect = (keys: Array<string | number>) => {
  if (keys.length === 0) return
  const key = keys[0] as number
  if (selectedNodeKey.value === key) return
  selectedNodeKey.value = key
  emit('nodeSelect', key)
}

const handleCreate = () => {
  emit('create')
}

const renderLabel = ({ option }: { option: TreeOption }) => {
  return h('div', { 
    class: 'tree-node',
    onContextmenu: (e: MouseEvent) => {
      e.preventDefault()
      e.stopPropagation()
      emit('contextMenu', option, { x: e.clientX, y: e.clientY })
    }
  }, [
    h('span', { class: 'tree-node-title' }, option.label as string),
    option.version
      ? h(NTag, { size: 'small', type: 'info', style: 'margin-left: 8px' }, { default: () => option.version as string })
      : null,
  ])
}

watch(() => props.selectedLibraryId, () => {
  fetchTree()
}, { immediate: true })
</script>

<template>
  <div class="left-panel">
    <div class="panel-header">
      <span class="panel-title">文档结构</span>
      <NSelect
        :value="props.selectedLibraryId"
        :options="docLibraries.map(lib => ({ label: lib.name, value: lib.id }))"
        placeholder="请选择文档库"
        style="width: 100%; margin-top: 8px"
        @update:value="(value) => emit('update:selectedLibraryId', value)"
      />
    </div>
    <div class="panel-actions">
      <NButton type="primary" size="small" @click="handleCreate">
        <template #icon>
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="12" y1="5" x2="12" y2="19"></line>
            <line x1="5" y1="12" x2="19" y2="12"></line>
          </svg>
        </template>
        新建文档
      </NButton>
    </div>
    <NSpin :show="treeLoading">
      <div v-if="treeData.length > 0" class="tree-container">
        <NTree
          :data="treeData"
          :render-label="renderLabel"
          block-line
          expand-on-click
          selectable
          :cancelable="false"
          :selected-keys="selectedNodeKey ? [selectedNodeKey] : []"
          @update:selected-keys="handleNodeSelect"
        />
      </div>
      <NEmpty v-else description="暂无文档数据" />
    </NSpin>
  </div>
</template>

<style scoped>
.left-panel {
  flex: 0 0 350px;
  display: flex;
  flex-direction: column;
  border-right: 1px solid #e5e7eb;
  padding-right: 24px;
  overflow: hidden;
}

.panel-header {
  padding: 12px 0;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 12px;
}

.panel-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a1a;
}

.panel-actions {
  padding: 8px 0;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 12px;
}

.tree-container {
  flex: 1;
  overflow-y: auto;
  padding: 8px 0;
}

.tree-node {
  display: flex;
  align-items: center;
  width: 100%;
}

.tree-node-title {
  font-weight: 500;
}

@media (max-width: 1024px) {
  .left-panel {
    flex: none;
    width: 100%;
    border-right: none;
    border-bottom: 1px solid #e5e7eb;
    padding-right: 0;
    padding-bottom: 24px;
    max-height: 400px;
  }
}

@media (max-width: 768px) {
  .left-panel {
    max-height: 300px;
  }
}
</style>
