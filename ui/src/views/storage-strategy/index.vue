<script setup lang="ts">
import { h, ref, computed, onMounted } from 'vue'
import type { FormInst } from 'naive-ui'
import {
  NCard,
  NDataTable,
  NButton,
  NModal,
  NForm,
  NFormItem,
  NInput,
  NSelect,
  NSpace,
  NSwitch,
  useMessage,
  NIcon,
  useDialog,
} from 'naive-ui'
import { Pencil } from '@vicons/ionicons5'
import type { DataTableColumns } from 'naive-ui'
import {
  getStorageStrategyList,
  createStorageStrategy,
  updateStorageStrategy,
  deleteStorageStrategy,
  setDefaultStorageStrategy,
  type StorageStrategy,
  type CreateStorageStrategyParams,
} from '@/api/storage'

const message = useMessage()
const dialog = useDialog()

// 表格数据
const loading = ref(false)
const data = ref<StorageStrategy[]>([])
const columns: DataTableColumns<StorageStrategy> = [
  {
    title: '策略名称',
    key: 'name',
  },
  {
    title: '存储类型',
    key: 'type',
    render(row) {
      const typeMap = {
        local: '本地存储',
        aliyun: '阿里云OSS',
        tencent: '腾讯云COS',
        qiniu: '七牛云',
      }
      return typeMap[row.type] || row.type
    },
  },
  {
    title: '访问域名',
    key: 'domain',
  },
  {
    title: '是否默认',
    key: 'is_default',
    render(row) {
      return h(NSwitch, {
        value: row.is_default,
        disabled: row.is_default,
        onUpdateValue: () => handleSetDefault(row),
      })
    },
  },
  {
    title: '操作',
    key: 'actions',
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
                type: 'primary',
                onClick: () => handleEdit(row),
              },
              {
                icon: () => h(NIcon, null, { default: () => h(Pencil) }),
                default: () => '编辑',
              },
            ),
            h(
              NButton,
              {
                size: 'small',
                quaternary: true,
                type: 'error',
                disabled: row.is_default,
                onClick: () => handleDelete(row),
              },
              { default: () => '删除' },
            ),
          ],
        },
      )
    },
  },
]

// 存储类型选项
const storageTypes = [
  { label: '本地存储', value: 'local' as const },
  { label: '阿里云OSS', value: 'aliyun' as const },
  { label: '腾讯云COS', value: 'tencent' as const },
  { label: '七牛云', value: 'qiniu' as const },
]

// 编辑弹窗相关
const showModal = ref(false)
const modalTitle = ref('')
const formRef = ref<FormInst | null>(null)
const formValue = ref<Omit<CreateStorageStrategyParams, 'type'> & { id: number; type: StorageStrategy['type'] }>({
  id: 0,
  name: '',
  type: 'local',
  domain: '',
  access_key: '',
  secret_key: '',
  bucket: '',
  region: '',
  is_default: false,
})

// 根据存储类型显示不同的表单项
const showAccessKey = computed(() => formValue.value.type !== 'local')
const showBucket = computed(() => formValue.value.type !== 'local')
const showRegion = computed(() => ['aliyun', 'tencent'].includes(formValue.value.type))

// 获取存储策略列表
const fetchStorageStrategies = async () => {
  loading.value = true
  try {
    const res = await getStorageStrategyList()
    data.value = res.data
  } catch (error) {
    message.error('获取存储策略列表失败：' + (error as Error).message)
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  modalTitle.value = '添加存储策略'
  formValue.value = {
    id: 0,
    name: '',
    type: 'local',
    domain: '',
    access_key: '',
    secret_key: '',
    bucket: '',
    region: '',
    is_default: false,
  }
  showModal.value = true
}

const handleEdit = (row: StorageStrategy) => {
  modalTitle.value = '编辑存储策略'
  formValue.value = {
    id: row.id,
    name: row.name,
    type: row.type,
    domain: row.domain,
    access_key: row.access_key || '',
    secret_key: row.secret_key || '',
    bucket: row.bucket || '',
    region: row.region || '',
    is_default: row.is_default,
  }
  showModal.value = true
}

const handleDelete = async (row: StorageStrategy) => {
  dialog.warning({
    title: '确认删除',
    content: '确定要删除该存储策略吗？删除后无法恢复！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: async () => {
      try {
        await deleteStorageStrategy(row.id)
        message.success('删除成功喵~')
        fetchStorageStrategies()
      } catch (error) {
        message.error('删除失败：' + (error as Error).message)
      }
    },
  })
}

const handleSetDefault = async (row: StorageStrategy) => {
  try {
    await setDefaultStorageStrategy(row.id)
    message.success('设置默认策略成功喵~')
    fetchStorageStrategies()
  } catch (error) {
    message.error('设置默认策略失败：' + (error as Error).message)
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return
  try {
    await formRef.value.validate()

    if (formValue.value.id) {
      await updateStorageStrategy({
        id: formValue.value.id,
        name: formValue.value.name,
        type: formValue.value.type,
        domain: formValue.value.domain,
        access_key: formValue.value.access_key,
        secret_key: formValue.value.secret_key,
        bucket: formValue.value.bucket,
        region: formValue.value.region,
        is_default: formValue.value.is_default,
      })
    } else {
      await createStorageStrategy({
        name: formValue.value.name,
        type: formValue.value.type,
        domain: formValue.value.domain,
        access_key: formValue.value.access_key,
        secret_key: formValue.value.secret_key,
        bucket: formValue.value.bucket,
        region: formValue.value.region,
        is_default: formValue.value.is_default,
      })
    }
    showModal.value = false
    message.success('保存成功喵~')
    fetchStorageStrategies()
  } catch (error) {
    if (error instanceof Error) {
      message.error('保存失败：' + error.message)
    }
  }
}

onMounted(() => {
  fetchStorageStrategies()
})
</script>

<template>
  <div class="storage-strategy">
    <n-card>
      <template #header>
        <div class="header-section">
          <div class="title">存储策略管理</div>
          <n-button type="primary" @click="handleAdd">添加策略</n-button>
        </div>
      </template>
      <n-data-table
        :loading="loading"
        :columns="columns"
        :data="data"
        :pagination="{ pageSize: 10 }"
      />
    </n-card>

    <!-- 编辑弹窗 -->
    <n-modal
      v-model:show="showModal"
      :title="modalTitle"
      preset="dialog"
      @positive-click="handleSubmit"
    >
      <n-form
        ref="formRef"
        :model="formValue"
        label-placement="left"
        label-width="100"
        require-mark-placement="right-hanging"
      >
        <n-form-item label="策略名称" path="name" :rule="{ required: true }">
          <n-input v-model:value="formValue.name" placeholder="请输入策略名称" />
        </n-form-item>
        <n-form-item label="存储类型" path="type" :rule="{ required: true }">
          <n-select
            v-model:value="formValue.type"
            :options="storageTypes"
            placeholder="请选择存储类型"
          />
        </n-form-item>
        <n-form-item label="访问域名" path="domain" :rule="{ required: true }">
          <n-input v-model:value="formValue.domain" placeholder="请输入访问域名" />
        </n-form-item>
        <template v-if="showAccessKey">
          <n-form-item
            label="Access Key"
            path="access_key"
            :rule="{ required: true }"
          >
            <n-input
              v-model:value="formValue.access_key"
              placeholder="请输入Access Key"
            />
          </n-form-item>
          <n-form-item
            label="Secret Key"
            path="secret_key"
            :rule="{ required: true }"
          >
            <n-input
              v-model:value="formValue.secret_key"
              type="password"
              show-password-on="click"
              placeholder="请输入Secret Key"
            />
          </n-form-item>
        </template>
        <template v-if="showBucket">
          <n-form-item label="Bucket" path="bucket" :rule="{ required: true }">
            <n-input v-model:value="formValue.bucket" placeholder="请输入Bucket" />
          </n-form-item>
        </template>
        <template v-if="showRegion">
          <n-form-item label="地域" path="region" :rule="{ required: true }">
            <n-input v-model:value="formValue.region" placeholder="请输入地域" />
          </n-form-item>
        </template>
        <n-form-item label="设为默认">
          <n-switch v-model:value="formValue.is_default" />
        </n-form-item>
      </n-form>
    </n-modal>
  </div>
</template>

<style scoped>
.storage-strategy {
  max-width: 1200px;
  margin: 0 auto;
  padding: 24px;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title {
  font-size: 16px;
  font-weight: 500;
}
</style>
