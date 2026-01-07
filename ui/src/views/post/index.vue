<script setup lang="ts">
import {
  NCard,
  NButton,
  NInput,
  NIcon,
  NSelect,
  NDatePicker,
  NDataTable,
  NTag,
  NDropdown,
  useMessage,
  type DataTableColumns,
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
import { useRouter } from 'vue-router'
import * as postApi from '@/api/post'
import { usePostHook } from './utils/hook'
import dayjs from 'dayjs'

const message = useMessage()
const dialog = useDialog()
const router = useRouter()
const { settingPost,publishPost,unpublishPost } = usePostHook()

// 搜索和筛选
const searchKeyword = ref('')
const filterStatus = ref(null)
const dateRange = ref<[number, number] | null>(null)

// 表格加载状态
const loading = ref(false)

// 分页配置
const pagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 30, 40],
  onChange: (page: number) => {
    pagination.page = page
    onSearch()
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.pageSize = pageSize
    pagination.page = 1
    onSearch()
  },
})

// 文章状态选项
const statusOptions = [
  { label: '草稿', value: 'draft' },
  { label: '已发布', value: 'published' },
  { label: '已归档', value: 'archived' },
]

// 模拟文章数据
const dataList = ref([
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
const columns: DataTableColumns<any> = [
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
        undefined: { text: '未知', type: 'error' },
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
    key: 'created_at',
    width: 180,
    render:(row)=>dayjs(row.created_at).format("YYYY-MM-DD HH:mm:ss")
  },
  {
    title: '更新时间',
    key: 'updated_at',
    width: 180,
    render:(row)=>dayjs(row.updated_at).format("YYYY-MM-DD HH:mm:ss")
  },
  {
    title: '操作',
    key: 'actions',
    width: 100,
    fixed: 'right',
    render: (row: any) => {
      const options: DropdownMixedOption[] = [
        {
          label: '编辑',
          key: 'edit',
          icon: () => h(NIcon, null, { default: () => h(CreateOutline) }),
        },
        {
          label: row.status === 'draft' ? '发布' : '取消发布',
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
          type: 'danger',
          icon: () => h(NIcon, null, { default: () => h(TrashOutline) }),
          disabled: row.status === 'published',
          props: {
            class: 'n-dropdown-option-body--danger',
          },
        },
      ]

      const handleSelect = (key: string) => {
        switch (key) {
          case 'edit':
            editPost(row)
            break
          case 'publish':
            if(row.status==='publish'){
               publishPost(row).then(()=>{
                onSearch()
               })
            } else{
              unpublishPost(row).then(()=>{
                onSearch()
              })
            }

            break
          case 'setting':
            handleSettingPost(row)
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
          options: options,
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

// 创建文章
const createPost = () => {
  postApi
    .createPost({
      title: '未命名的文章',
      content: '<p>此处是文章内容</p>',
    })
    .then((res) => {
      if (res.code === 200) {
        message.success('创建成功')
        router.push({
          name: 'PostEditor',
          query: {
            id: res.data.id,
          },
        })
      }
    })
}

// 编辑文章
const editPost = (post: any) => {
  router.push({
    name: 'PostEditor',
    query: {
      id: post.id,
    },
  })
}


const handleSettingPost = (row: any) => {
  settingPost(row).then(()=>{
    onSearch()
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
      dataList.value.unshift(newPost)
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
      const index = dataList.value.findIndex((item) => item.id === row.id)
      dataList.value.splice(index, 1)
    },
  })
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
  onSearch().then(() => {
    setTimeout(() => {
      loading.value = false
    }, 1000)
  })
}

const onSearch = async () => {
  const res = await postApi.getPostList({
    page: pagination.page,
    page_size: pagination.pageSize,
    title: searchKeyword.value,
    status: filterStatus.value,
    start_date: dateRange.value?.[0],
    end_date: dateRange.value?.[1],
  })
  console.log(res)
  dataList.value = res.data
}
onMounted(() => {
  onSearch()
})
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
        :data="dataList"
        :loading="loading"
        :pagination="pagination"
        remote
        :row-key="(row) => row.id"
      />
    </n-card>
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
