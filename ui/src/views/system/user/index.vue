<script lang="ts" setup>
import { NButton, NIcon, type DataTableColumns } from 'naive-ui'
import { Pencil,RefreshOutline } from '@vicons/ionicons5'
import * as userApi from '@/api/system/user'
import { addDialog } from '@/components/dialog'
import EditForm from './editForm.vue'


const showModal = ref(false)
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
    title: '用户 ID',
    key: 'id',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '邮箱',
    key: 'email',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '用户名',
    key: 'name',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '手机号',
    key: 'phoneNumber',
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
          onClick: (row) => {
            openEditDialog('编辑',row)
          },
        },
        {
          icon: () => h(NIcon, {}, () => h(Pencil)),
          default: () => '编辑',
        },
      )
    },
  },
]

const openEditDialog=(title="新增",row?:any)=>{
  const formRef = ref<any>(null)
  addDialog({
    title:`${title}用户`,
    props:{
      formInline:{
        id:row?.id || undefined,
        email:row?.email || '',
        name:row?.name || '',
        phone_number:row?.phone_number || '',
        password:row?.password || '',
      }
    },
    contentRenderer:({ options })=> h(EditForm,{ref:formRef,formInline:options.props!.formInline}),
  })
}

onMounted(()=>{
  userApi.getUserList().then(res=>{
    dataList.value = res.data || []
  })
})
</script>
<template>
  <div class="container-fluid p-6">
    <n-card title="用户管理" class="user-card">
      <!-- 头部操作栏 -->
      <div class="header-section">
        <div class="search-section">

        </div>
        <div class="action-section">
          <n-button type="primary"  style="margin-right: 12px" @click="openEditDialog('新增')"> <i class="bi bi-plus"></i> 添加用户 </n-button>
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

  <!-- 添加/编辑用户模态框 -->
  <n-modal v-model:show="showModal" preset="dialog" :closable="true">
    <template #header>
      <h5 class="modal-title" >添加用户</h5>
    </template>
    <template #default>
      <n-form>
        <n-form-item label="邮箱" path="email">
          <n-input/>
        </n-form-item>
        <n-form-item label="用户名" path="username">
          <n-input/>
        </n-form-item>
        <n-form-item label="手机号" path="phonenumber">
          <n-input/>
        </n-form-item>
        <n-form-item label="密码" path="password">
          <n-input/>
        </n-form-item>
      </n-form>
    </template>
    <template #action>
      <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
      <button type="button" class="btn btn-primary" onclick="saveUser()">保存</button>
    </template>
  </n-modal>
</template>
<style scoped>
.user-card{
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
