<script setup lang="ts">
import { NTag,NDataTable } from 'naive-ui'
import * as roleApi from '@/api/system/role'
import type { DataTableColumns } from 'naive-ui'
import { RefreshOutline } from '@vicons/ionicons5'

const dataList = ref<any[]>([])
const loading = ref(false)

// 分页配置
const pagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  total: 0,
  pageSizes: [10, 20, 50, 100],
  onChange: (page: number) => {
    pagination.page = page
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.pageSize = pageSize
    pagination.page = 1
  },
})

// 表格列定义
const columns: DataTableColumns<any> = [
  {
    title: '角色 ID',
    key: 'id',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '角色名称',
    key: 'name',
    width: 200,
  },
  {
    title:"默认角色",
    key:"is_default",
    width: 180,
    render: (row: any) => {
      return row.is_default ? h(NTag, { type: 'success' }, ()=>'是') : h(NTag, { type: 'info' }, ()=>'否')
    }
  },
  {
    title: '角色描述',
    key: 'description',
    ellipsis: true,
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
  }
]

const onSearch = ()=>{
  roleApi.getRolePage().then(res=>{
    dataList.value = res.data.records
  })
}

onMounted(()=>{
  onSearch()
})
</script>
<template>
  <div class="container-fluid p-6">
    <n-card title="角色管理" class="role-card">
      <!-- 头部操作栏 -->
       <div class="header-section">
        <div class="search-section">

        </div>
        <div class="action-section">
          <n-button type="primary" style="margin-right: 12px">新增角色</n-button>
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
      />
    </n-card>
  </div>
</template>
<style scoped>
.role-card{
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
