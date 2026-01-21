<script setup lang="ts">
import { h, ref, reactive } from 'vue'
import { NCard, NDataTable, NButton, NSpace, NTag, NIcon, NPopconfirm, useMessage } from 'naive-ui'
import { RefreshOutline, AddOutline, TrashOutline, CheckmarkCircleOutline } from '@vicons/ionicons5'
import type { DataTableColumns } from 'naive-ui'
import {
  getLicensePage,
  deleteLicense,
  verifyLicense,
  type License,
} from '@/api/infra/license'
import { addDialog } from '@/components/dialog'
import FormComponent from './form.vue'

const message = useMessage()

const loading = ref(false)
const data = ref<License[]>([])

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

const columns: DataTableColumns<License> = [
  {
    title: 'ID',
    key: 'id',
    width: 80,
  },
  {
    title: '机器码',
    key: 'machine_code',
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '授权密钥',
    key: 'license_key',
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '客户名称',
    key: 'customer_name',
  },
  {
    title: '过期日期',
    key: 'expire_date',
  },
  {
    title: '状态',
    key: 'status',
    render(row) {
      const statusMap = {
        1: { type: 'success', label: '有效' },
        2: { type: 'warning', label: '过期' },
        3: { type: 'error', label: '禁用' },
      }
      const status = statusMap[row.status] || { type: 'default', label: '未知' }
      return h(NTag, { type: status.type as any }, { default: () => status.label })
    },
  },
  {
    title: '创建时间',
    key: 'created_at',
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    render(row) {
      return h(
        NSpace,
        {},
        {
          default: () => [
            h(
              NButton,
              {
                size: 'small',
                quaternary: true,
                type: 'info',
                onClick: () => handleVerify(row),
              },
              {
                icon: () => h(NIcon, null, { default: () => h(CheckmarkCircleOutline) }),
                default: () => '验证',
              },
            ),
            h(
              NButton,
              {
                size: 'small',
                quaternary: true,
                type: 'primary',
                onClick: () => handleEdit(row),
              },
              {
                default: () => '编辑',
              },
            ),
            h(
              NPopconfirm,
              {
                onPositiveClick: () => handleDelete(row),
              },
              {
                trigger: () =>
                  h(
                    NButton,
                    {
                      size: 'small',
                      quaternary: true,
                      type: 'error',
                    },
                    {
                      icon: () => h(NIcon, null, { default: () => h(TrashOutline) }),
                      default: () => '删除',
                    },
                  ),
                default: () => '确定要删除该授权吗？',
              },
            ),
          ],
        },
      )
    },
  },
]

const onSearch = async () => {
  loading.value = true
  try {
    const res = await getLicensePage({ page: pagination.page, size: pagination.pageSize })
    data.value = res.data.records
    pagination.total = res.data.total
  } catch (error) {
    message.error('获取授权列表失败：' + (error as Error).message)
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  addDialog({
    title: '添加授权',
    contentRenderer: ({ options }) =>
      h(FormComponent, { ...options.props, onSuccess: () => onSearch() }),
  })
}

const handleEdit = (row: License) => {
  addDialog({
    title: '编辑授权',
    contentRenderer: ({ options }) =>
      h(FormComponent, { ...options.props, license: row, onSuccess: () => onSearch() }),
  })
}

const handleDelete = async (row: License) => {
  try {
    await deleteLicense(row.id)
    message.success('授权删除成功')
    onSearch()
  } catch (error) {
    message.error('授权删除失败：' + (error as Error).message)
  }
}

const handleVerify = async (row: License) => {
  try {
    const res = await verifyLicense({ machine_code: row.machine_code })
    if (res.data.valid) {
      message.success(`授权验证成功：${res.data.message}，客户：${res.data.customer_name}，过期时间：${res.data.expire_date}`)
    } else {
      message.warning(`授权验证失败：${res.data.message}`)
    }
  } catch (error) {
    message.error('授权验证失败：' + (error as Error).message)
  }
}

onMounted(() => {
  onSearch()
})
</script>

<template>
  <NCard title="授权管理">
    <template #header-extra>
      <NSpace>
        <NButton type="primary" @click="handleAdd">
          <template #icon>
            <NIcon><AddOutline /></NIcon>
          </template>
          添加授权
        </NButton>
        <NButton @click="onSearch">
          <template #icon>
            <NIcon><RefreshOutline /></NIcon>
          </template>
          刷新
        </NButton>
      </NSpace>
    </template>

    <NDataTable
      :columns="columns"
      :data="data"
      :loading="loading"
      :pagination="pagination"
      :bordered="false"
      striped
    />
  </NCard>
</template>
