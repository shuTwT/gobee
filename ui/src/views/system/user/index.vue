<script lang="ts" setup>
import { NButton, NIcon,NDataTable, type DataTableColumns, NTag, NPopconfirm } from 'naive-ui'
import { Pencil,RefreshOutline,TrashOutline } from '@vicons/ionicons5'
import * as userApi from '@/api/system/user'
import { addDialog } from '@/components/dialog'
import EditForm from './editForm.vue'

const showModal = ref(false)
const editFormRef = ref<any>(null)
const currentUserId = ref<number | undefined>(undefined)

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

const dataList = ref<any>([])
const loading = ref(false)

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
    key: 'phone_number',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
    {
    title: '角色',
    key: 'role_id',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
    render: (row) => {
      return h(NTag,{type:'primary'},()=>row.role?.name || '无')
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
                openEditDialog('编辑',row)
              },
            },
            {
              icon: () => h(NIcon, {}, () => h(Pencil)),
              default: () => '编辑',
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
              default: () => '确定删除该用户吗？',
            },
          ),
        ],
      )
    },
  },
]

const openEditDialog = (title = '新增', row?: any) => {
  currentUserId.value = row?.id
  addDialog({
    title: `${title}用户`,
    props: {
      formInline: {
        id: row?.id || undefined,
        email: row?.email || '',
        name: row?.name || '',
        phone_number: row?.phone_number || '',
        password: '',
        role_id: row?.role_id || undefined,
      }
    },
    contentRenderer: ({ options }) => h(EditForm, { ref: editFormRef, formInline: options.props!.formInline }),
    beforeSure: async (done) => {
      try {
        const data = await editFormRef.value?.getData()
        
        if (currentUserId.value) {
          await userApi.updateUser(currentUserId.value, data)
          window.$message?.success('更新成功')
        } else {
          await userApi.createUser(data)
          window.$message?.success('创建成功')
        }
        
        onSearch()
        done()
      } catch (error) {
        console.error('提交失败:', error)
        window.$message?.error('提交失败')
      }
    },
  })
}

const onSearch = () => {
  loading.value = true
  userApi.getUserPage({
    page: pagination.page,
    page_size: pagination.pageSize,
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
    await userApi.deleteUser(id)
    window.$message?.success('删除成功')
    onSearch()
  } catch (error) {
    console.error('删除失败:', error)
    window.$message?.error('删除失败')
  }
}

onMounted(() => {
  onSearch()
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
          <n-button type="primary" style="margin-right: 12px" @click="openEditDialog('新增')"> <i class="bi bi-plus"></i> 添加用户 </n-button>
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
      />
    </n-card>
  </div>
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
