<template>
  <div class="pay-channel-container">
    <n-card title="支付渠道管理" class="channel-card">
      <!-- 头部操作栏 -->
      <div class="header-section">
        <div class="search-section">
          <n-input
            v-model:value="searchKeyword"
            placeholder="搜索渠道名称或代码"
            clearable
            style="width: 250px"
            @input="handleSearch"
          >
            <template #prefix>
              <n-icon><search-outline /></n-icon>
            </template>
          </n-input>
          <n-select
            v-model:value="filterType"
            placeholder="筛选类型"
            clearable
            style="width: 150px; margin-left: 12px"
            :options="typeOptions"
            @update:value="handleFilterChange"
          />
          <n-select
            v-model:value="filterStatus"
            placeholder="筛选状态"
            clearable
            style="width: 150px; margin-left: 12px"
            :options="statusOptions"
            @update:value="handleFilterChange"
          />
        </div>
        <div class="action-section">
          <n-button type="primary" @click="openAddModal">
            <template #icon>
              <n-icon><add-outline /></n-icon>
            </template>
            添加渠道
          </n-button>
          <n-button @click="handleRefresh" style="margin-left: 12px">
            <template #icon>
              <n-icon><refresh-outline /></n-icon>
            </template>
            刷新
          </n-button>
        </div>
      </div>

      <!-- 渠道列表表格 -->
      <n-data-table
        :columns="columns"
        :data="filteredChannelList"
        :loading="loading"
        :pagination="pagination"
        :row-key="(row) => row.id"
        class="channel-table"
      />
    </n-card>

    <!-- 添加/编辑渠道模态框 -->
    <n-modal v-model:show="showModal" preset="dialog" :title="modalTitle" style="width: 600px">
      <div class="modal-content">
        <n-form
          ref="formRef"
          :model="channelForm"
          :rules="formRules"
          label-placement="left"
          label-width="100px"
        >
          <n-form-item label="渠道名称" path="name">
            <n-input
              v-model:value="channelForm.name"
              placeholder="请输入渠道名称"
              maxlength="50"
              show-count
            />
          </n-form-item>

          <n-form-item label="渠道代码" path="code">
            <n-input
              v-model:value="channelForm.code"
              placeholder="请输入渠道代码"
              maxlength="30"
              show-count
              :disabled="isEditMode"
            />
          </n-form-item>

          <n-form-item label="渠道类型" path="type">
            <n-select
              v-model:value="channelForm.type"
              placeholder="请选择渠道类型"
              :options="typeOptions"
              :disabled="isEditMode"
            />
          </n-form-item>

          <n-form-item label="配置信息" path="config">
            <n-input
              v-model:value="channelForm.config"
              type="textarea"
              placeholder="请输入渠道配置信息（JSON格式）"
              :rows="6"
              @blur="validateConfig"
            />
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-radio-group v-model:value="channelForm.status">
              <n-radio-button value="active">启用</n-radio-button>
              <n-radio-button value="inactive">禁用</n-radio-button>
            </n-radio-group>
          </n-form-item>

          <n-form-item label="描述" path="description">
            <n-input
              v-model:value="channelForm.description"
              type="textarea"
              placeholder="请输入渠道描述"
              :rows="3"
              maxlength="200"
              show-count
            />
          </n-form-item>

          <n-form-item label="排序" path="sortOrder">
            <n-input-number
              v-model:value="channelForm.sortOrder"
              :min="0"
              :max="999"
              placeholder="排序值"
            />
          </n-form-item>
        </n-form>
      </div>
      <template #action>
        <n-space>
          <n-button @click="closeModal">取消</n-button>
          <n-button type="primary" @click="handleSubmit" :loading="saving"> 保存 </n-button>
        </n-space>
      </template>
    </n-modal>

    <!-- 配置详情模态框 -->
    <n-modal v-model:show="showConfigModal" preset="dialog" title="配置详情" style="width: 700px">
      <div class="config-modal-content">
        <n-form label-placement="left" label-width="100px">
          <n-form-item label="渠道名称">
            <n-input :value="currentChannel?.name" readonly />
          </n-form-item>
          <n-form-item label="配置信息">
            <n-code :code="formatConfig" language="json" :word-wrap="true" />
          </n-form-item>
        </n-form>
      </div>
      <template #action>
        <n-button @click="showConfigModal = false">关闭</n-button>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, h } from 'vue'
import { NButton, NIcon, NTag, useMessage } from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import {
  SearchOutline,
  RefreshOutline,
  AddOutline,
  EyeOutline,
  PencilOutline,
  TrashOutline,
  CheckmarkCircleOutline,
  CloseCircleOutline,
} from '@vicons/ionicons5'

const message = useMessage()

// 搜索和筛选
const searchKeyword = ref('')
const filterType = ref<string | null>(null)
const filterStatus = ref<string | null>(null)

// 表格数据
const loading = ref(false)
const channelList = ref<ChannelItem[]>([])

// 模态框状态
const showModal = ref(false)
const showConfigModal = ref(false)
const saving = ref(false)
const isEditMode = ref(false)

// 当前操作的渠道
const currentChannel = ref<ChannelItem | null>(null)

// 表单引用
const formRef = ref<FormInst | null>(null)

// 渠道表单
const channelForm = reactive({
  id: '',
  name: '',
  code: '',
  type: '',
  config: '',
  status: 'active' as 'active' | 'inactive',
  description: '',
  sortOrder: 0,
})

// 渠道类型选项
const typeOptions = [
  { label: '支付宝', value: 'alipay' },
  { label: '微信支付', value: 'wechatpay' },
  { label: '银联支付', value: 'unionpay' },
  { label: 'PayPal', value: 'paypal' },
  { label: 'Stripe', value: 'stripe' },
]

// 状态选项
const statusOptions = [
  { label: '全部', value: null as any },
  { label: '启用', value: 'active' },
  { label: '禁用', value: 'inactive' },
]

// 表单验证规则
const formRules: FormRules = {
  name: {
    required: true,
    message: '请输入渠道名称',
    trigger: 'blur',
  },
  code: {
    required: true,
    message: '请输入渠道代码',
    trigger: 'blur',
  },
  type: {
    required: true,
    message: '请选择渠道类型',
    trigger: 'change',
  },
  config: {
    required: true,
    message: '请输入配置信息',
    trigger: 'blur',
  },
}

// 分页配置
const pagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 50],
  onChange: (page: number) => {
    pagination.page = page
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.pageSize = pageSize
    pagination.page = 1
  },
})

// 渠道数据类型
interface ChannelItem {
  id: string
  name: string
  code: string
  type: string
  typeLabel: string
  config: string
  status: 'active' | 'inactive'
  description: string
  sortOrder: number
  createdAt: string
  updatedAt: string
}

// 模态框标题
const modalTitle = computed(() => (isEditMode.value ? '编辑支付渠道' : '添加支付渠道'))

// 格式化配置信息
const formatConfig = computed(() => {
  if (!currentChannel.value?.config) return '{}'
  try {
    const config = JSON.parse(currentChannel.value.config)
    return JSON.stringify(config, null, 2)
  } catch {
    return currentChannel.value.config
  }
})

// 筛选后的渠道列表
const filteredChannelList = computed(() => {
  let filtered = channelList.value

  // 关键词搜索
  if (searchKeyword.value) {
    const keyword = searchKeyword.value.toLowerCase()
    filtered = filtered.filter(
      (channel) =>
        channel.name.toLowerCase().includes(keyword) ||
        channel.code.toLowerCase().includes(keyword) ||
        channel.description.toLowerCase().includes(keyword),
    )
  }

  // 类型筛选
  if (filterType.value) {
    filtered = filtered.filter((channel) => channel.type === filterType.value)
  }

  // 状态筛选
  if (filterStatus.value) {
    filtered = filtered.filter((channel) => channel.status === filterStatus.value)
  }

  return filtered
})

// 表格列定义
const columns: DataTableColumns<ChannelItem> = [
  {
    title: '渠道名称',
    key: 'name',
    width: 150,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '渠道代码',
    key: 'code',
    width: 120,
  },
  {
    title: '渠道类型',
    key: 'typeLabel',
    width: 120,
    render: (row) => {
      const typeMap: Record<string, string> = {
        alipay: '支付宝',
        wechatpay: '微信支付',
        unionpay: '银联支付',
        paypal: 'PayPal',
        stripe: 'Stripe',
      }
      return h(NTag, { type: 'info' }, () => typeMap[row.type] || row.type)
    },
  },
  {
    title: '状态',
    key: 'status',
    width: 80,
    render: (row) => {
      const status = row.status
      const type = status === 'active' ? 'success' : 'error'
      const icon = status === 'active' ? CheckmarkCircleOutline : CloseCircleOutline
      const text = status === 'active' ? '启用' : '禁用'

      return h(
        'div',
        {
          style: {
            display: 'flex',
            alignItems: 'center',
            gap: '4px',
          },
        },
        [
          h(NIcon, { color: status === 'active' ? '#52c41a' : '#f5222d' }, () => h(icon)),
          h(NTag, { type }, () => text),
        ],
      )
    },
  },
  {
    title: '描述',
    key: 'description',
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '排序',
    key: 'sortOrder',
    width: 80,
    sorter: (row1, row2) => row1.sortOrder - row2.sortOrder,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 160,
    sorter: (row1, row2) => new Date(row1.createdAt).getTime() - new Date(row2.createdAt).getTime(),
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    render: (row) => {
      return h('div', { class: 'action-buttons' }, [
        h(
          NButton,
          {
            size: 'small',
            type: 'info',
            quaternary: true,
            onClick: () => viewConfig(row),
          },
          {
            icon: () => h(NIcon, {}, () => h(EyeOutline)),
            default: () => '查看配置',
          },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            quaternary: true,
            style: 'margin-left: 8px',
            onClick: () => editChannel(row),
          },
          {
            icon: () => h(NIcon, {}, () => h(PencilOutline)),
            default: () => '编辑',
          },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'error',
            quaternary: true,
            style: 'margin-left: 8px',
            onClick: () => deleteChannel(row),
          },
          {
            icon: () => h(NIcon, {}, () => h(TrashOutline)),
            default: () => '删除',
          },
        ),
      ])
    },
  },
]

// 处理搜索
const handleSearch = () => {
  pagination.page = 1
}

// 处理筛选变化
const handleFilterChange = () => {
  pagination.page = 1
}

// 验证配置JSON格式
const validateConfig = () => {
  if (!channelForm.config.trim()) return

  try {
    JSON.parse(channelForm.config)
  } catch {
    message.warning('配置信息格式不正确，请检查JSON格式')
  }
}

// 打开添加模态框
const openAddModal = () => {
  isEditMode.value = false
  resetForm()
  showModal.value = true
}

// 关闭模态框
const closeModal = () => {
  showModal.value = false
  resetForm()
}

// 重置表单
const resetForm = () => {
  channelForm.id = ''
  channelForm.name = ''
  channelForm.code = ''
  channelForm.type = ''
  channelForm.config = ''
  channelForm.status = 'active'
  channelForm.description = ''
  channelForm.sortOrder = 0
}

// 编辑渠道
const editChannel = (channel: ChannelItem) => {
  isEditMode.value = true
  currentChannel.value = channel

  // 填充表单数据
  channelForm.id = channel.id
  channelForm.name = channel.name
  channelForm.code = channel.code
  channelForm.type = channel.type
  channelForm.config = channel.config
  channelForm.status = channel.status
  channelForm.description = channel.description
  channelForm.sortOrder = channel.sortOrder

  showModal.value = true
}

// 查看配置
const viewConfig = (channel: ChannelItem) => {
  currentChannel.value = channel
  showConfigModal.value = true
}

// 删除渠道
const deleteChannel = (channel: ChannelItem) => {
  // 这里可以添加确认对话框
  if (confirm(`确定要删除支付渠道 "${channel.name}" 吗？`)) {
    // 模拟删除操作
    const index = channelList.value.findIndex((item) => item.id === channel.id)
    if (index > -1) {
      channelList.value.splice(index, 1)
      message.success('删除成功')
    }
  }
}

// 提交表单
const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
  } catch {
    return // 验证失败，不提交
  }

  saving.value = true
  try {
    // 验证配置JSON格式
    if (channelForm.config.trim()) {
      JSON.parse(channelForm.config)
    }

    // 模拟API调用
    await new Promise((resolve) => setTimeout(resolve, 1000))

    if (isEditMode.value) {
      // 编辑模式
      const channel = channelList.value.find((item) => item.id === channelForm.id)
      if (channel) {
        Object.assign(channel, {
          name: channelForm.name,
          type: channelForm.type,
          typeLabel:
            typeOptions.find((opt) => opt.value === channelForm.type)?.label || channelForm.type,
          config: channelForm.config,
          status: channelForm.status,
          description: channelForm.description,
          sortOrder: channelForm.sortOrder,
          updatedAt: new Date().toLocaleString('zh-CN'),
        })
      }
      message.success('更新成功')
    } else {
      // 添加模式
      const newChannel: ChannelItem = {
        id: Date.now().toString(),
        name: channelForm.name,
        code: channelForm.code,
        type: channelForm.type,
        typeLabel:
          typeOptions.find((opt) => opt.value === channelForm.type)?.label || channelForm.type,
        config: channelForm.config,
        status: channelForm.status,
        description: channelForm.description,
        sortOrder: channelForm.sortOrder,
        createdAt: new Date().toLocaleString('zh-CN'),
        updatedAt: new Date().toLocaleString('zh-CN'),
      }
      channelList.value.unshift(newChannel)
      message.success('添加成功')
    }

    closeModal()
  } catch (error) {
    if (error instanceof Error && error.message.includes('JSON')) {
      message.error('配置信息格式不正确，请检查JSON格式')
    } else {
      message.error(isEditMode.value ? '更新失败' : '添加失败')
    }
  } finally {
    saving.value = false
  }
}

// 刷新数据
const handleRefresh = () => {
  loadChannelList()
}

// 加载渠道列表
const loadChannelList = async () => {
  loading.value = true
  try {
    // 模拟API调用
    await new Promise((resolve) => setTimeout(resolve, 800))

    // 模拟数据
    channelList.value = [
      {
        id: '1',
        name: '支付宝网页支付',
        code: 'alipay_web',
        type: 'alipay',
        typeLabel: '支付宝',
        config: JSON.stringify(
          {
            app_id: '2021000000000000',
            private_key: 'MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC...',
            public_key: 'MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...',
            sandbox: false,
          },
          null,
          2,
        ),
        status: 'active',
        description: '支付宝网页版支付接口',
        sortOrder: 1,
        createdAt: '2024-01-15 10:00:00',
        updatedAt: '2024-01-15 10:00:00',
      },
      {
        id: '2',
        name: '微信支付H5',
        code: 'wechat_h5',
        type: 'wechatpay',
        typeLabel: '微信支付',
        config: JSON.stringify(
          {
            app_id: 'wx1234567890abcdef',
            mch_id: '1234567890',
            api_key: 'abcdefghijklmnopqrstuvwxyz123456',
            sandbox: false,
          },
          null,
          2,
        ),
        status: 'active',
        description: '微信支付H5支付接口',
        sortOrder: 2,
        createdAt: '2024-01-15 10:30:00',
        updatedAt: '2024-01-15 10:30:00',
      },
      {
        id: '3',
        name: '银联支付',
        code: 'unionpay_web',
        type: 'unionpay',
        typeLabel: '银联支付',
        config: JSON.stringify(
          {
            mer_id: '123456789012345',
            cert_path: '/path/to/cert.pfx',
            cert_password: 'password123',
            sandbox: true,
          },
          null,
          2,
        ),
        status: 'inactive',
        description: '银联在线支付接口',
        sortOrder: 3,
        createdAt: '2024-01-15 11:00:00',
        updatedAt: '2024-01-15 11:00:00',
      },
    ]
  } catch {
    message.error('渠道列表加载失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadChannelList()
})
</script>

<style scoped>
.pay-channel-container {
  padding: 24px;
  min-height: 100vh;
  background-color: #f5f7fa;
}

.channel-card {
  max-width: 1400px;
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
}

.action-section {
  display: flex;
  align-items: center;
}

.channel-table {
  margin-top: 16px;
}

.modal-content {
  padding: 16px 0;
}

.config-modal-content {
  padding: 16px 0;
}

.action-buttons {
  display: flex;
  gap: 8px;
}

@media (max-width: 768px) {
  .pay-channel-container {
    padding: 16px;
  }

  .header-section {
    flex-direction: column;
    gap: 16px;
    align-items: flex-start;
  }

  .search-section {
    flex-wrap: wrap;
  }

  .action-section {
    width: 100%;
    justify-content: flex-end;
  }
}
</style>
