<script setup lang="ts">
import { NTag, NDataTable, NButton, NIcon, NPopconfirm } from 'naive-ui'
import * as roleApi from '@/api/system/role'
import type { DataTableColumns } from 'naive-ui'
import { RefreshOutline, Pencil, TrashOutline } from '@vicons/ionicons5'
import { addDialog } from '@/components/dialog'
import EditForm from './editForm.vue'

const dataList = ref<any[]>([])
const loading = ref(false)
const editFormRef = ref<any>(null)
const currentRoleId = ref<number | undefined>(undefined)

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
    title: '角色代码',
    key: 'code',
    width: 180,
  },
  {
    title: "默认角色",
    key: "is_default",
    width: 120,
    render: (row: any) => {
      return row.is_default ? h(NTag, { type: 'success' }, () => '是') : h(NTag, { type: 'info' }, () => '否')
    }
  },
  {
    title: '角色描述',
    key: 'description',
    ellipsis: {
      tooltip: true,
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
                openEditDialog('编辑', row)
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
              default: () => '确定删除该角色吗？',
            },
          ),
        ],
      )
    },
  },
]

const openEditDialog = (title = '新增', row?: any) => {
  currentRoleId.value = row?.id
  addDialog({
    title: `${title}角色`,
    props: {
      formInline: {
        id: row?.id || undefined,
        name: row?.name || '',
        code: row?.code || '',
        description: row?.description || '',
      }
    },
    contentRenderer: ({ options }) => h(EditForm, { ref: editFormRef, formInline: options.props!.formInline }),
    beforeSure: async (done) => {
      try {
        const data = await editFormRef.value?.getData()

        if (currentRoleId.value) {
          await roleApi.updateRole(currentRoleId.value, data)
          window.$message?.success('更新成功')
        } else {
          await roleApi.createRole(data)
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
  roleApi.getRolePage({
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
    await roleApi.deleteRole(id)
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
    <n-card title="角色管理" class="role-card">
      <!-- 头部操作栏 -->
       <div class="header-section">
        <div class="search-section">

        </div>
        <div class="action-section">
          <n-button type="primary" style="margin-right: 12px" @click="openEditDialog('新增')">新增角色</n-button>
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
