<script setup lang="ts">
import { NButton, NIcon, NDataTable, type DataTableColumns, NTag, NPopconfirm, NDrawer, NDrawerContent, NDescriptions, NDescriptionsItem } from 'naive-ui'
import { RefreshOutline, TrashOutline, EyeOutline, CheckmarkDoneOutline } from '@vicons/ionicons5'
import * as notificationApi from '@/api/system/notification'
import { addDialog } from '@/components/dialog'

const dataList = ref<any[]>([])
const loading = ref(false)
const selectedRowKeys = ref<number[]>([])
const showDetailDrawer = ref(false)
const currentNotification = ref<any>(null)

const pagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  total: 0,
  pageSizes: [10, 20, 50, 100],
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

const searchForm = reactive({
  is_read: null as boolean | null,
})

const columns: DataTableColumns<any> = [
  {
    type: 'selection',
  },
  {
    title: '通知 ID',
    key: 'id',
    width: 100,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '通知标题',
    key: 'title',
    width: 200,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '通知内容',
    key: 'content',
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '发布时间',
    key: 'publish_time',
    width: 180,
    render: (row) => {
      return row.publish_time ? new Date(row.publish_time).toLocaleString() : '-'
    },
  },
  {
    title: '是否已读',
    key: 'is_read',
    width: 100,
    render: (row) => {
      return row.is_read ? h(NTag, { type: 'success' }, () => '已读') : h(NTag, { type: 'warning' }, () => '未读')
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 180,
    render: (row) => {
      return h(
        'div',
        { style: { display: 'flex', gap: '8px' } },
        [
          h(
            NButton,
            {
              size: 'small',
              type: 'primary',
              quaternary: true,
              onClick: () => {
                openDetailDrawer(row)
              },
            },
            {
              icon: () => h(NIcon, {}, () => h(EyeOutline)),
              default: () => '详情',
            },
          ),
          h(
            NPopconfirm,
            {
              onPositiveClick: () => handleDelete(row.id),
            },
            {
              trigger: () =>
                h(
                  NButton,
                  {
                    size: 'small',
                    type: 'error',
                    quaternary: true,
                  },
                  {
                    icon: () => h(NIcon, {}, () => h(TrashOutline)),
                    default: () => '删除',
                  },
                ),
              default: () => '确定删除该通知吗？',
            },
          ),
        ],
      )
    },
  },
]

const openDetailDrawer = (row: any) => {
  currentNotification.value = row
  showDetailDrawer.value = true
}

const onSearch = () => {
  loading.value = true
  notificationApi.getNotificationPage({
    page: pagination.page,
    page_size: pagination.pageSize,
    is_read: searchForm.is_read,
  }).then(res => {
    if (res.code === 200) {
      dataList.value = res.data.records || []
      pagination.total = res.data.total || 0
    }
  }).finally(() => {
    loading.value = false
  })
}

const handleDelete = async (id: number) => {
  try {
    await notificationApi.deleteNotification(id)
    window.$message?.success('删除成功')
    onSearch()
  } catch (error) {
    console.error('删除失败:', error)
    window.$message?.error('删除失败')
  }
}

const handleBatchRead = async () => {
  if (selectedRowKeys.value.length === 0) {
    window.$message?.warning('请选择要标记为已读的通知')
    return
  }
  try {
    await notificationApi.batchMarkAsRead({ ids: selectedRowKeys.value })
    window.$message?.success('批量标记已读成功')
    selectedRowKeys.value = []
    onSearch()
  } catch (error) {
    console.error('批量标记已读失败:', error)
    window.$message?.error('批量标记已读失败')
  }
}

const handleSearchByStatus = (status: boolean | null) => {
  searchForm.is_read = status
  pagination.page = 1
  onSearch()
}

onMounted(() => {
  onSearch()
})
</script>
<template>
  <div class="container-fluid p-6">
    <n-card title="通知管理" class="notification-card">
      <!-- 头部操作栏 -->
      <div class="header-section">
        <div class="search-section">
          <n-button-group>
            <n-button :type="searchForm.is_read === null ? 'primary' : 'default'" @click="handleSearchByStatus(null)">
              全部
            </n-button>
            <n-button :type="searchForm.is_read === false ? 'primary' : 'default'" @click="handleSearchByStatus(false)">
              未读
            </n-button>
            <n-button :type="searchForm.is_read === true ? 'primary' : 'default'" @click="handleSearchByStatus(true)">
              已读
            </n-button>
          </n-button-group>
        </div>
        <div class="action-section">
          <n-button type="success" :disabled="selectedRowKeys.length === 0" @click="handleBatchRead">
            <template #icon>
              <n-icon><checkmark-done-outline /></n-icon>
            </template>
            批量已读
          </n-button>
          <n-button @click="onSearch()">
            <template #icon>
              <n-icon><refresh-outline /></n-icon>
            </template>
            刷新
          </n-button>
        </div>
      </div>

      <n-data-table
        :columns="columns"
        :data="dataList"
        :loading="loading"
        :pagination="pagination"
        :row-key="(row) => row.id"
        :remote="true"
        v-model:checked-row-keys="selectedRowKeys"
      />
    </n-card>

    <!-- 通知详情抽屉 -->
    <n-drawer v-model:show="showDetailDrawer" :width="600">
      <n-drawer-content title="通知详情">
        <n-descriptions v-if="currentNotification" bordered :column="1">
          <n-descriptions-item label="通知ID">
            {{ currentNotification.id }}
          </n-descriptions-item>
          <n-descriptions-item label="通知标题">
            {{ currentNotification.title }}
          </n-descriptions-item>
          <n-descriptions-item label="通知内容">
            {{ currentNotification.content }}
          </n-descriptions-item>
          <n-descriptions-item label="发布时间">
            {{ currentNotification.publish_time ? new Date(currentNotification.publish_time).toLocaleString() : '-' }}
          </n-descriptions-item>
          <n-descriptions-item label="是否已读">
            <n-tag :type="currentNotification.is_read ? 'success' : 'warning'">
              {{ currentNotification.is_read ? '已读' : '未读' }}
            </n-tag>
          </n-descriptions-item>
          <n-descriptions-item label="创建时间">
            {{ new Date(currentNotification.created_at).toLocaleString() }}
          </n-descriptions-item>
        </n-descriptions>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<style scoped>
.notification-card {
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
  gap: 12px;
}
</style>
