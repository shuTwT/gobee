<script lang="ts" setup>
import { NButton,NIcon } from 'naive-ui'
import { RefreshOutline,Eye } from '@vicons/ionicons5'
import type { DataTableColumns } from 'naive-ui'
import { addDialog } from '@/components/dialog'
import uploadForm from './uploadForm.vue'


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

const openUploadDialog = (title="新增",row?:any)=>{
  addDialog({
    title:"上传文件",
    contentRenderer:()=>h(uploadForm)
  })
}

const onSearch=()=>{

}

onMounted(()=>{
  onSearch()
})
</script>
<template>
  <div class="container-fluid p-6">
    <n-card title="文件管理" class="file-card">
      <!-- 头部操作栏 -->
      <div class="header-section">
        <div class="search-section"></div>
        <div class="action-section">
           <n-button type="primary"  style="margin-right: 12px" @click="openUploadDialog('新增')"> <i class="bi bi-plus"></i> 上传文件 </n-button>
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
