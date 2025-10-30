<script setup lang="ts">
import { ref, computed, reactive } from 'vue'
import {
  NCard,
  NButton,
  NInput,
  NIcon,
  NSelect,
  NDatePicker,
  NDataTable,
  NModal,
  NForm,
  NFormItem,
  NDynamicTags,
  NSpace,
  NTag,
  NDropdown,
  useMessage,
  type DataTableColumns,
  useModal,
} from 'naive-ui'
import {
  SearchOutline,
  AddOutline,
  RefreshOutline,
  CreateOutline,
  TrashOutline,
  SendOutline,
  SettingsOutline,
  ShareSocialOutline,
  DownloadOutline,
  CopyOutline,
  DuplicateOutline,
} from '@vicons/ionicons5'
import { h } from 'vue'
import type { DropdownMixedOption } from 'naive-ui/es/dropdown/src/interface'
import SettingForm from './settingForm.vue'
import { addDialog } from '@/components/dialog'

const message = useMessage()
const dialog = useDialog()
const modal = useModal()



// 搜索和筛选
const searchKeyword = ref('')
const filterStatus = ref(null)
const dateRange = ref<[number,number]|null>(null)

// 表格加载状态
const loading = ref(false)
const saving = ref(false)

// 编辑弹窗控制
const showEditModal = ref(false)

// 分页配置
const pagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 30, 40],
  onChange: (page: number) => {
    pagination.page = page
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.pageSize = pageSize
    pagination.page = 1
  },
})

// 文章状态选项
const statusOptions = [
  { label: '草稿', value: 'draft' },
  { label: '已发布', value: 'published' },
  { label: '已归档', value: 'archived' },
]

// 表单校验规则
const rules = {
  title: {
    required: true,
    message: '请输入文章标题',
    trigger: ['blur', 'input'],
  },
  content: {
    required: true,
    message: '请输入文章内容',
    trigger: ['blur', 'input'],
  },
}

// 表单数据
const postForm = reactive({
  id: '',
  title: '',
  summary: '',
  content: '',
  tags: [],
  status: 'draft',
})

// 模拟文章数据
const postList = ref([
  {
    id: '1',
    title: '示例文章1',
    summary: '这是一篇示例文章的摘要内容...',
    author: '管理员',
    status: 'published',
    views: 100,
    comments: 5,
    createdAt: '2024-01-15 10:00:00',
    updatedAt: '2024-01-15 10:00:00',
  },
  {
    id: '2',
    title: '示例文章2',
    summary: '这是另一篇示例文章的摘要内容...',
    author: '管理员',
    status: 'draft',
    views: 0,
    comments: 0,
    createdAt: '2024-01-15 11:00:00',
    updatedAt: '2024-01-15 11:00:00',
  },
])

// 表格列配置
const columns :DataTableColumns<any>= [
  {
    title: '标题',
    key: 'title',
    width: 200,
  },
  {
    title: '摘要',
    key: 'summary',
    ellipsis: true,
  },
  {
    title: '作者',
    key: 'author',
    width: 100,
  },
  {
    title: '状态',
    key: 'status',
    width: 100,
    render: (row: any) => {
      const statusMap: any = {
        draft: { text: '草稿', type: 'default' },
        published: { text: '已发布', type: 'success' },
        archived: { text: '已归档', type: 'warning' },
      }
      const status = statusMap[row.status]
      return h(NTag, { type: status.type, size: 'small' }, { default: () => status.text })
    },
  },
  {
    title: '浏览量',
    key: 'views',
    width: 100,
  },
  {
    title: '评论数',
    key: 'comments',
    width: 100,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 180,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    width: 180,
  },
  {
    title: '操作',
    key: 'actions',
    width: 100,
    fixed: 'right',
    render: (row: any) => {
      const options:DropdownMixedOption[] = [
        {
          label: '编辑',
          key: 'edit',
          icon: () => h(NIcon, null, { default: () => h(CreateOutline) }),
        },
        {
          label: row.status === 'draft'?'发布':'取消发布',
          key: 'publish',
          icon: () => h(NIcon, null, { default: () => h(SendOutline) }),
        },
        {
          label: '设置',
          key: 'setting',
          icon: () => h(NIcon, null, { default: () => h(SettingsOutline) }),
        },
        {
          label: '分享',
          key: 'share',
          icon: () => h(NIcon, null, { default: () => h(ShareSocialOutline) }),
        },
        {
          label: '导出',
          key: 'export',
          icon: () => h(NIcon, null, { default: () => h(DownloadOutline) }),
        },
        {
          label: '复制内容',
          key: 'copy',
          icon: () => h(NIcon, null, { default: () => h(CopyOutline) }),
        },
        {
          label: '克隆',
          key: 'clone',
          icon: () => h(NIcon, null, { default: () => h(DuplicateOutline) }),
        },
        {
          type: 'divider',
          key: 'd1',
        },
        {
          label: '删除',
          key: 'delete',
          type:'danger',
          icon: () => h(NIcon, null, { default: () => h(TrashOutline) }),
          disabled: row.status === 'published',
          props:{
            class:'n-dropdown-option-body--danger'
          }
        },
      ]

      const handleSelect = (key: string) => {
        switch (key) {
          case 'edit':
            editPost(row)
            break
          case 'publish':
            publishPost(row)
            break
          case 'setting':
            settingPost(row)
            break
          case 'share':
            sharePost(row)
            break
          case 'export':
            exportPost(row)
            break
          case 'copy':
            copyPostContent(row)
            break
          case 'clone':
            clonePost(row)
            break
          case 'delete':
            deletePost(row)
            break
        }
      }

      return h(
        NDropdown,
        {
          trigger: 'hover',
          options:options,
          onSelect: handleSelect,
        },
        {
          default: () =>
            h(
              NButton,
              {
                size: 'small',
                quaternary: true,
              },
              {
                default: () => '操作',
              },
            ),
        },
      )
    },
  },
]

// 筛选后的文章列表
const filteredPostList = computed(() => {
  let filtered = postList.value

  // 关键词搜索
  if (searchKeyword.value) {
    const keyword = searchKeyword.value.toLowerCase()
    filtered = filtered.filter(
      (post) =>
        post.title.toLowerCase().includes(keyword) ||
        post.summary.toLowerCase().includes(keyword),
    )
  }

  // 状态筛选
  if (filterStatus.value) {
    filtered = filtered.filter((post) => post.status === filterStatus.value)
  }

  // 日期范围筛选
  if (dateRange.value) {
    const [startDate, endDate] = dateRange.value
    filtered = filtered.filter((post) => {
      const postDate = new Date(post.createdAt).valueOf()
      return postDate >= startDate && postDate <= endDate
    })
  }

  return filtered
})

// 创建文章
const createPost = () => {
  postForm.id = ''
  postForm.title = ''
  postForm.summary = ''
  postForm.content = ''
  postForm.tags = []
  postForm.status = 'draft'
  showEditModal.value = true
}

// 编辑文章
const editPost = (post: any) => {
  postForm.id = post.id
  postForm.title = post.title
  postForm.summary = post.summary
  postForm.content = post.content || ''
  postForm.tags = post.tags || []
  postForm.status = post.status
  showEditModal.value = true
}

// 发布文章
const publishPost = (row: any) => {
  dialog.info({
    title: '确认',
    content: '确定要发布该文章吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      row.status = 'published'
      message.success('发布成功')
    },
  })
}

// 文章设置
const settingPost = (row: any) => {
  // TODO: 实现文章设置功能
  addDialog({
    props:{},
    contentRenderer:()=>h(SettingForm)
  })
}

// 分享文章
const sharePost = (row: any) => {
  // TODO: 实现分享功能
  message.info('分享功能开发中')
}

// 导出文章
const exportPost = (row: any) => {
  // TODO: 实现导出功能
  message.info('导出功能开发中')
}

// 复制文章内容
const copyPostContent = (row: any) => {
  // TODO: 实现复制内容功能
  message.info('复制内容功能开发中')
}

// 克隆文章
const clonePost = (row: any) => {
  dialog.info({
    title: '确认',
    content: '确定要克隆该文章吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      const newPost = { ...row }
      newPost.id = Date.now()
      newPost.title = `${newPost.title} (副本)`
      newPost.status = 'draft'
      postList.value.unshift(newPost)
      message.success('克隆成功')
    },
  })
}

// 删除文章
const deletePost = (row: any) => {
  dialog.warning({
    title: '警告',
    content: '确定要删除该文章吗？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      message.success('删除成功')
      const index = postList.value.findIndex((item) => item.id === row.id)
      postList.value.splice(index, 1)
    },
  })
}

// 保存文章
const savePost = async () => {
  saving.value = true
  try {
    // 模拟API调用
    await new Promise((resolve) => setTimeout(resolve, 1000))
    if (postForm.id) {
      // 更新现有文章
      const index = postList.value.findIndex((item) => item.id === postForm.id)
      if (index !== -1) {
        postList.value[index] = {
          ...postList.value[index],
          ...postForm,
          updatedAt: new Date().toLocaleString('zh-CN'),
        }
      }
    } else {
      // 创建新文章
      postList.value.unshift({
        ...postForm,
        id: String(Date.now()),
        author: '管理员',
        views: 0,
        comments: 0,
        createdAt: new Date().toLocaleString('zh-CN'),
        updatedAt: new Date().toLocaleString('zh-CN'),
      })
    }
    message.success('文章保存成功')
    showEditModal.value = false
  } catch {
    message.error('文章保存失败')
  } finally {
    saving.value = false
  }
}

// 处理搜索
const handleSearch = () => {
  pagination.page = 1
}

// 处理筛选变化
const handleFilterChange = () => {
  pagination.page = 1
}

// 刷新数据
const handleRefresh = () => {
  loading.value = true
  // 模拟加载数据
  setTimeout(() => {
    loading.value = false
  }, 1000)
}
</script>

<template>
  <div class="post-page">
    <n-card title="文章管理" class="post-card">
      <!-- 头部操作栏 -->
      <div class="header-section">
        <div class="search-section">
          <n-input
            v-model:value="searchKeyword"
            placeholder="搜索文章标题或内容"
            clearable
            style="width: 200px"
            @keyup.enter="handleSearch"
          >
            <template #prefix>
              <n-icon><search-outline /></n-icon>
            </template>
          </n-input>
          <n-select
            v-model:value="filterStatus"
            placeholder="文章状态"
            :options="statusOptions"
            clearable
            style="width: 150px"
            @update:value="handleFilterChange"
          />
          <n-date-picker
            v-model:value="dateRange"
            type="daterange"
            clearable
            style="width: 240px"
            @update:value="handleFilterChange"
          />
        </div>
        <div class="action-section">
          <n-button type="primary" style="margin-right: 12px" @click="createPost">
            <template #icon>
              <n-icon><add-outline /></n-icon>
            </template>
            新建文章
          </n-button>
          <n-button @click="handleRefresh">
            <template #icon>
              <n-icon><refresh-outline /></n-icon>
            </template>
            刷新
          </n-button>
        </div>
      </div>

      <!-- 文章列表 -->
      <n-data-table
        :columns="columns"
        :data="filteredPostList"
        :loading="loading"
        :pagination="pagination"
        :row-key="(row) => row.id"
      />
    </n-card>

    <!-- 文章编辑弹窗 -->
    <n-modal
      v-model:show="showEditModal"
      preset="card"
      style="width: 900px"
      title="编辑文章"
      :bordered="false"
      size="huge"
      :segmented="true"
    >
      <n-form
        ref="formRef"
        :model="postForm"
        :rules="rules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
        size="medium"
      >
        <n-form-item label="标题" path="title">
          <n-input v-model:value="postForm.title" placeholder="请输入文章标题" />
        </n-form-item>
        <n-form-item label="摘要" path="summary">
          <n-input
            v-model:value="postForm.summary"
            type="textarea"
            placeholder="请输入文章摘要"
            :rows="3"
          />
        </n-form-item>
        <n-form-item label="内容" path="content">
          <n-input
            v-model:value="postForm.content"
            type="textarea"
            placeholder="请输入文章内容"
            :rows="10"
          />
        </n-form-item>
        <n-form-item label="标签" path="tags">
          <n-dynamic-tags v-model:value="postForm.tags" />
        </n-form-item>
        <n-form-item label="状态" path="status">
          <n-select
            v-model:value="postForm.status"
            :options="statusOptions"
            placeholder="请选择文章状态"
          />
        </n-form-item>
      </n-form>
      <template #footer>
        <n-space justify="end">
          <n-button @click="showEditModal = false">取消</n-button>
          <n-button type="primary" :loading="saving" @click="savePost">保存</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<style scoped>
.post-page {
  padding: 16px;
}

.post-card {
  max-width: 1600px;
  margin: 0 auto;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 16px 0;
  border-bottom: 1px solid #f0f0f0;
}

.search-section {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: wrap;
}

.action-section {
  display: flex;
  align-items: center;
}
</style>
