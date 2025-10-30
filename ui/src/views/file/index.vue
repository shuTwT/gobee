<script lang="ts" setup>
import { NButton,NIcon } from 'naive-ui'
import { RefreshOutline,Eye } from '@vicons/ionicons5'
import type { DataTableColumns } from 'naive-ui'
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

const dataList = ref([
  {
    id:"1",
    type:"图片",
    size:"1.2MB",
    createdTime:"2023-10-26"
  }
])
const loading = ref(false)

// 表格列定义
const columns: DataTableColumns<any> = [
  {
    title:"编号",
    key:"id",
    width:180,
    ellipsis:{
      tooltip:true
    }
  },
    {
    title:"类型",
    key:"type",
    width:180,
    ellipsis:{
      tooltip:true
    }
  },
    {
    title:"大小",
    key:"size",
    width:180,
    ellipsis:{
      tooltip:true
    }
  },
    {
    title:"上传时间",
    key:"createdTime",
    width:180,
    ellipsis:{
      tooltip:true
    }
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
          icon: () => h(NIcon, {}, () => h(Eye)),
          default: () => '编辑',
        },
      )
    },
  },
]
</script>
<template>
  <div class="container-fluid p-6">
    <n-card title="文件管理" class="file-card">
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

    <!-- 文件上传 Modal -->
    <n-modal>
      <div
        class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800"
      >
        <div class="mt-3 text-center">
          <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white">上传文件</h3>
          <div class="mt-2 px-7 py-3">
            <input
              type="file"
              id="fileInput"
              class="block w-full text-sm text-gray-500 dark:text-gray-400 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100 dark:file:bg-blue-900/30 dark:file:text-blue-400 dark:hover:file:bg-blue-800"
            />
          </div>
          <div class="items-center px-4 py-3">
            <button
              id="doUploadBtn"
              class="px-4 py-2 bg-blue-600 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
            >
              上传
            </button>
            <button
              id="closeUploadModalBtn"
              class="mt-3 px-4 py-2 bg-gray-200 text-gray-800 text-base font-medium rounded-md w-full shadow-sm hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-300 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white dark:hover:bg-gray-600"
            >
              取消
            </button>
          </div>
        </div>
      </div>
    </n-modal>
  </div>
</template>
<style scoped>
.file-card {
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
