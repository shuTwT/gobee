<script setup lang="ts">
import { NButton, NIcon } from 'naive-ui'
import { Pencil, RefreshOutline } from '@vicons/ionicons5'
import type { DataTableColumns } from 'naive-ui'
const webhookFormData = ref({
  webhookUrl: '',
  webhookEvents: [],
})

// 分页配置
const pagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 50, 100],
  onChange: (page: number) => {
    pagination.page = page
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.pageSize = pageSize
    pagination.page = 1
  },
})

const dataList = ref([])
const loading = ref(false)

// 表格列定义
const columns: DataTableColumns<any> = [
  {
    title: '编号',
    key: 'id',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: 'url',
    key: 'url',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '事件',
    key: 'event',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 180,
    render: () => {
      return h(
        NButton,
        {
          size: 'small',
          type: 'primary',
          quaternary: true,
        },
        {
          icon: () => h(NIcon, {}, () => h(Pencil)),
          default: () => '编辑',
        },
      )
    },
  },
]

const eventOptions = ref([
  {
    value: 'content.create',
    label: '内容创建',
  },
  {
    value: 'content.update',
    label: '内容更新',
  },
])
</script>
<template>
  <div class="container-fluid p-4">
    <n-card title="webhook 管理" class="webhook-card">
      <!-- 头部操作栏 -->
      <div class="header-section">
        <div class="search-section"></div>
        <div class="action-section">
          <n-button type="primary"  style="margin-right: 12px"> <i class="bi bi-plus"></i> 添加webhook </n-button>
          <n-button>
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
      />
    </n-card>

    <!-- Add/Edit Webhook Modal -->
    <n-modal>
      <template #header> 添加 Webhook </template>
      <n-form>
        <n-form-item label="url">
          <n-input v-model:value="webhookFormData.webhookUrl" placeholder="请输入 Webhook URL" />
        </n-form-item>
        <n-form-item label="事件">
          <n-checkbox-group v-model:value="webhookFormData.webhookEvents" :options="eventOptions" />
        </n-form-item>
      </n-form>
      <template #action>
        <div class="items-center px-4 py-3">
          <n-button type="primary"> 保存 </n-button>
          <n-button type="default"> 取消 </n-button>
        </div>
      </template>
    </n-modal>
  </div>
</template>
<style scoped>
.webhook-card{
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
