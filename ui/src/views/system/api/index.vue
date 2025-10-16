<template>
  <div class="api-management-container">
    <n-card title="API 接口管理" class="api-card">
      <!-- 搜索和操作栏 -->
      <div class="header-section">
        <div class="search-section">
          <n-input
            v-model:value="searchKeyword"
            placeholder="搜索API名称、路径或描述"
            clearable
            style="width: 300px"
            @input="handleSearch"
          >
            <template #prefix>
              <n-icon><search-outline /></n-icon>
            </template>
          </n-input>
          <n-select
            v-model:value="filterPermission"
            placeholder="权限筛选"
            clearable
            style="width: 150px; margin-left: 12px"
            :options="permissionOptions"
            @update:value="handleFilterChange"
          />
          <n-select
            v-model:value="filterMethod"
            placeholder="方法筛选"
            clearable
            style="width: 150px; margin-left: 12px"
            :options="methodOptions"
            @update:value="handleFilterChange"
          />
        </div>
        <div class="action-section">
          <n-button type="primary" @click="handleBatchOperation">
            <template #icon>
              <n-icon><settings-outline /></n-icon>
            </template>
            批量操作
          </n-button>
          <n-button @click="handleRefresh" style="margin-left: 12px">
            <template #icon>
              <n-icon><refresh-outline /></n-icon>
            </template>
            刷新
          </n-button>
        </div>
      </div>

      <!-- API列表表格 -->
      <n-data-table
        :columns="columns"
        :data="filteredApiList"
        :loading="loading"
        :pagination="pagination"
        :row-key="(row) => row.id"
        @update:checked-row-keys="handleCheck"
        class="api-table"
      />
    </n-card>

    <!-- 权限设置模态框 -->
    <n-modal
      v-model:show="showPermissionModal"
      preset="dialog"
      title="设置API权限"
      style="width: 600px"
    >
      <div class="permission-modal-content">
        <n-form
          ref="permissionFormRef"
          :model="permissionForm"
          label-placement="left"
          label-width="100px"
        >
          <n-form-item label="API名称">
            <n-input :value="currentApi?.name" readonly />
          </n-form-item>
          <n-form-item label="API路径">
            <n-input :value="currentApi?.path" readonly />
          </n-form-item>
          <n-form-item label="权限类型" path="permissionType">
            <n-radio-group v-model:value="permissionForm.permissionType">
              <n-radio-button value="public">公开</n-radio-button>
              <n-radio-button value="private">私有</n-radio-button>
            </n-radio-group>
          </n-form-item>

          <div v-if="permissionForm.permissionType === 'private'" class="role-permission-section">
            <n-divider />
            <n-form-item label="角色权限">
              <div class="role-tree-container">
                <n-tree
                  :data="roleTreeData"
                  :checked-keys="permissionForm.allowedRoles"
                  :on-update:checked-keys="handleRoleCheck"
                  checkable
                  cascade
                  block-line
                />
              </div>
            </n-form-item>
            <n-form-item label="权限说明">
              <n-alert type="info" :show-icon="false">
                选中的角色将拥有访问此API的权限，未选中的角色将被拒绝访问
              </n-alert>
            </n-form-item>
          </div>
        </n-form>
      </div>
      <template #action>
        <n-space>
          <n-button @click="showPermissionModal = false">取消</n-button>
          <n-button type="primary" @click="savePermissionSettings" :loading="savingPermission">
            保存
          </n-button>
        </n-space>
      </template>
    </n-modal>

    <!-- API测试模态框 -->
    <n-modal
      v-model:show="showTestModal"
      preset="dialog"
      title="API测试"
      style="width: 800px"
    >
      <div class="test-modal-content">
        <n-form
          ref="testFormRef"
          :model="testForm"
          label-placement="left"
          label-width="100px"
        >
          <n-form-item label="API名称">
            <n-input :value="currentApi?.name" readonly />
          </n-form-item>
          <n-form-item label="请求方法">
            <n-tag :type="getMethodType(currentApi?.method)">
              {{ currentApi?.method }}
            </n-tag>
          </n-form-item>
          <n-form-item label="请求路径">
            <n-input :value="currentApi?.path" readonly />
          </n-form-item>
          <n-form-item label="测试参数">
            <n-input
              v-model:value="testForm.params"
              type="textarea"
              placeholder='请输入测试参数，例如：{"id": 1, "name": "test"}'
              :rows="4"
            />
          </n-form-item>
          <n-form-item label="测试结果" v-if="testResult">
            <n-code :code="testResult" language="json" :word-wrap="true" />
          </n-form-item>
        </n-form>
      </div>
      <template #action>
        <n-space>
          <n-button @click="showTestModal = false">关闭</n-button>
          <n-button type="primary" @click="executeTest" :loading="testingApi">
            执行测试
          </n-button>
        </n-space>
      </template>
    </n-modal>

    <!-- 批量操作模态框 -->
    <n-modal
      v-model:show="showBatchModal"
      preset="dialog"
      title="批量设置权限"
      style="width: 600px"
    >
      <div class="batch-modal-content">
        <n-alert type="warning" style="margin-bottom: 16px">
          已选择 {{ checkedRowKeys.length }} 个API接口进行批量操作
        </n-alert>
        <n-form
          ref="batchFormRef"
          :model="batchForm"
          label-placement="left"
          label-width="100px"
        >
          <n-form-item label="权限类型" path="permissionType">
            <n-radio-group v-model:value="batchForm.permissionType">
              <n-radio-button value="public">全部设为公开</n-radio-button>
              <n-radio-button value="private">全部设为私有</n-radio-button>
            </n-radio-group>
          </n-form-item>

          <div v-if="batchForm.permissionType === 'private'" class="batch-role-section">
            <n-divider />
            <n-form-item label="角色权限">
              <div class="role-tree-container">
                <n-tree
                  :data="roleTreeData"
                  :checked-keys="batchForm.allowedRoles"
                  :on-update:checked-keys="handleBatchRoleCheck"
                  checkable
                  cascade
                  block-line
                />
              </div>
            </n-form-item>
          </div>
        </n-form>
      </div>
      <template #action>
        <n-space>
          <n-button @click="showBatchModal = false">取消</n-button>
          <n-button type="primary" @click="executeBatchOperation" :loading="batchOperating">
            批量应用
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, h } from 'vue'
import { NButton, NIcon, NTag, useMessage } from 'naive-ui'
import type { DataTableColumns, TreeOption, FormInst } from 'naive-ui'
import {
  SearchOutline,
  SettingsOutline,
  RefreshOutline,
  ShieldOutline,
  GlobeOutline,
  LockClosedOutline,
  PlayOutline,
  PencilOutline
} from '@vicons/ionicons5'

const message = useMessage()

// 搜索和筛选
const searchKeyword = ref('')
const filterPermission = ref<string | null>(null)
const filterMethod = ref<string | null>(null)

// 表格数据
const loading = ref(false)
const apiList = ref<ApiItem[]>([])
const checkedRowKeys = ref<string[]>([])

// 模态框状态
const showPermissionModal = ref(false)
const showTestModal = ref(false)
const showBatchModal = ref(false)
const savingPermission = ref(false)
const testingApi = ref(false)
const batchOperating = ref(false)

// 当前操作的API
const currentApi = ref<ApiItem | null>(null)

// 表单引用
const permissionFormRef = ref<FormInst | null>(null)
const testFormRef = ref<FormInst | null>(null)
const batchFormRef = ref<FormInst | null>(null)

// 权限表单
const permissionForm = reactive({
  permissionType: 'public' as 'public' | 'private',
  allowedRoles: [] as string[]
})

// 测试表单
const testForm = reactive({
  params: '',
  result: ''
})

// 批量操作表单
const batchForm = reactive({
  permissionType: 'public' as 'public' | 'private',
  allowedRoles: [] as string[]
})

// 测试结果
const testResult = ref('')

// 权限选项
const permissionOptions = [
  { label: '全部', value: null as any },
  { label: '公开', value: 'public' },
  { label: '私有', value: 'private' }
]

// HTTP方法选项
const methodOptions = [
  { label: '全部', value: null as any },
  { label: 'GET', value: 'GET' },
  { label: 'POST', value: 'POST' },
  { label: 'PUT', value: 'PUT' },
  { label: 'DELETE', value: 'DELETE' },
  { label: 'PATCH', value: 'PATCH' }
]

// 角色树数据
const roleTreeData: TreeOption[] = [
  {
    label: '管理员角色',
    key: 'admin',
    children: [
      { label: '超级管理员', key: 'super_admin' },
      { label: '系统管理员', key: 'system_admin' },
      { label: '内容管理员', key: 'content_admin' }
    ]
  },
  {
    label: '用户角色',
    key: 'user',
    children: [
      { label: '普通用户', key: 'normal_user' },
      { label: 'VIP用户', key: 'vip_user' },
      { label: '高级用户', key: 'premium_user' }
    ]
  },
  {
    label: '访客角色',
    key: 'guest',
    children: [
      { label: '游客', key: 'guest_user' },
      { label: '未验证用户', key: 'unverified_user' }
    ]
  }
]

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
  }
})

// API接口类型定义
interface ApiItem {
  id: string
  name: string
  path: string
  method: string
  description: string
  permission: 'public' | 'private'
  allowedRoles: string[]
  category: string
  status: 'active' | 'inactive'
  createdAt: string
  updatedAt: string
}

// 表格列定义
const columns: DataTableColumns<ApiItem> = [
  {
    type: 'selection',
    width: 50
  },
  {
    title: 'API名称',
    key: 'name',
    width: 200,
    ellipsis: {
      tooltip: true
    }
  },
  {
    title: '请求方法',
    key: 'method',
    width: 100,
    render: (row) => {
      return h(
        'div',
        {
          style: {
            display: 'flex',
            alignItems: 'center',
            gap: '4px'
          }
        },
        [
          h('div', {
            style: {
              width: '8px',
              height: '8px',
              borderRadius: '50%',
              backgroundColor: getMethodColor(row.method)
            }
          }),
          h('span', {}, row.method)
        ]
      )
    }
  },
  {
    title: 'API路径',
    key: 'path',
    width: 300,
    ellipsis: {
      tooltip: true
    }
  },
  {
    title: '描述',
    key: 'description',
    ellipsis: {
      tooltip: true
    }
  },
  {
    title: '权限',
    key: 'permission',
    width: 100,
    render: (row) => {
      const icon = row.permission === 'public' ? GlobeOutline : LockClosedOutline
      const type = row.permission === 'public' ? 'success' : 'warning'
      return h(
        'div',
        {
          style: {
            display: 'flex',
            alignItems: 'center',
            gap: '4px'
          }
        },
        [
          h(NIcon, {},()=> h(icon)),
          h(NTag, { type }, ()=>row.permission === 'public' ? '公开' : '私有')
        ]
      )
    }
  },
  {
    title: '分类',
    key: 'category',
    width: 120
  },
  {
    title: '状态',
    key: 'status',
    width: 80,
    render: (row) => {
      const type = row.status === 'active' ? 'success' : 'error'
      return h(NTag, { type }, ()=>row.status === 'active' ? '启用' : '禁用')
    }
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
            type: 'primary',
            quaternary: true,
            onClick: () => openPermissionModal(row)
          },
          {
            icon: () => h(NIcon, {}, ()=>h(ShieldOutline)),
            default: () => '权限'
          }
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'info',
            quaternary: true,
            style: 'margin-left: 8px',
            onClick: () => openTestModal(row)
          },
          {
            icon: () => h(NIcon, {}, ()=>h(PlayOutline)),
            default: () => '测试'
          }
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'warning',
            quaternary: true,
            style: 'margin-left: 8px',
            onClick: () => editApi(row)
          },
          {
            icon: () => h(NIcon, {},()=> h(PencilOutline)),
            default: () => '编辑'
          }
        )
      ])
    }
  }
]

// 筛选后的API列表
const filteredApiList = computed(() => {
  let filtered = apiList.value

  // 关键词搜索
  if (searchKeyword.value) {
    const keyword = searchKeyword.value.toLowerCase()
    filtered = filtered.filter(api =>
      api.name.toLowerCase().includes(keyword) ||
      api.path.toLowerCase().includes(keyword) ||
      api.description.toLowerCase().includes(keyword)
    )
  }

  // 权限筛选
  if (filterPermission.value) {
    filtered = filtered.filter(api => api.permission === filterPermission.value)
  }

  // 方法筛选
  if (filterMethod.value) {
    filtered = filtered.filter(api => api.method === filterMethod.value)
  }

  return filtered
})

// 获取方法颜色
const getMethodColor = (method?: string) => {
  switch (method?.toUpperCase()) {
    case 'GET': return '#52c41a'
    case 'POST': return '#1890ff'
    case 'PUT': return '#faad14'
    case 'DELETE': return '#f5222d'
    case 'PATCH': return '#722ed1'
    default: return '#d9d9d9'
  }
}

// 获取方法标签类型
const getMethodType = (method?: string) => {
  switch (method?.toUpperCase()) {
    case 'GET': return 'success'
    case 'POST': return 'info'
    case 'PUT': return 'warning'
    case 'DELETE': return 'error'
    case 'PATCH': return 'default'
    default: return 'default'
  }
}

// 处理搜索
const handleSearch = () => {
  pagination.page = 1
}

// 处理筛选变化
const handleFilterChange = () => {
  pagination.page = 1
}

// 处理表格选择
const handleCheck = (keys: string[]) => {
  checkedRowKeys.value = keys
}

// 打开权限设置模态框
const openPermissionModal = (api: ApiItem) => {
  currentApi.value = api
  permissionForm.permissionType = api.permission
  permissionForm.allowedRoles = [...api.allowedRoles]
  showPermissionModal.value = true
}

// 打开测试模态框
const openTestModal = (api: ApiItem) => {
  currentApi.value = api
  testForm.params = ''
  testResult.value = ''
  showTestModal.value = true
}

// 处理角色选择
const handleRoleCheck = (keys: string[]) => {
  permissionForm.allowedRoles = keys
}

// 处理批量角色选择
const handleBatchRoleCheck = (keys: string[]) => {
  batchForm.allowedRoles = keys
}

// 保存权限设置
const savePermissionSettings = async () => {
  if (!currentApi.value) return

  savingPermission.value = true
  try {
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 1000))

    // 更新本地数据
    const api = apiList.value.find(item => item.id === currentApi.value!.id)
    if (api) {
      api.permission = permissionForm.permissionType
      api.allowedRoles = [...permissionForm.allowedRoles]
    }

    message.success('权限设置保存成功')
    showPermissionModal.value = false
  } catch {
    message.error('权限设置保存失败')
  } finally {
    savingPermission.value = false
  }
}

// 执行API测试
const executeTest = async () => {
  if (!currentApi.value) return

  testingApi.value = true
  try {
    // 模拟API测试
    await new Promise(resolve => setTimeout(resolve, 1500))

    const result = {
      success: true,
      code: 200,
      message: 'API测试成功',
      data: {
        method: currentApi.value.method,
        path: currentApi.value.path,
        params: testForm.params || '无参数',
        responseTime: '150ms',
        timestamp: new Date().toISOString()
      }
    }

    testResult.value = JSON.stringify(result, null, 2)
    message.success('API测试完成')
  } catch (error) {
    const errorResult = {
      success: false,
      code: 500,
      message: 'API测试失败',
      error: error instanceof Error ? error.message : '未知错误'
    }
    testResult.value = JSON.stringify(errorResult, null, 2)
    message.error('API测试失败')
  } finally {
    testingApi.value = false
  }
}

// 编辑API
const editApi = (api: ApiItem) => {
  message.info(`编辑API功能开发中: ${api.name}`)
}

// 批量操作
const handleBatchOperation = () => {
  if (checkedRowKeys.value.length === 0) {
    message.warning('请先选择要操作的API')
    return
  }

  batchForm.permissionType = 'public'
  batchForm.allowedRoles = []
  showBatchModal.value = true
}

// 执行批量操作
const executeBatchOperation = async () => {
  batchOperating.value = true
  try {
    // 模拟批量API调用
    await new Promise(resolve => setTimeout(resolve, 1500))

    // 更新本地数据
    checkedRowKeys.value.forEach(key => {
      const api = apiList.value.find(item => item.id === key)
      if (api) {
        api.permission = batchForm.permissionType
        if (batchForm.permissionType === 'private') {
          api.allowedRoles = [...batchForm.allowedRoles]
        } else {
          api.allowedRoles = []
        }
      }
    })

    message.success(`批量设置成功，已更新 ${checkedRowKeys.value.length} 个API`)
    showBatchModal.value = false
    checkedRowKeys.value = []
  } catch {
    message.error('批量设置失败')
  } finally {
    batchOperating.value = false
  }
}




// 刷新数据
const handleRefresh = () => {
  loadApiList()
}

// 加载API列表
const loadApiList = async () => {
  loading.value = true
  try {
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 1000))

    // 模拟数据
    apiList.value = [
      {
        id: '1',
        name: '用户登录',
        path: '/api/auth/login',
        method: 'POST',
        description: '用户登录接口，验证用户身份',
        permission: 'public',
        allowedRoles: [],
        category: '认证',
        status: 'active',
        createdAt: '2024-01-15 10:00:00',
        updatedAt: '2024-01-15 10:00:00'
      },
      {
        id: '2',
        name: '获取用户信息',
        path: '/api/user/profile',
        method: 'GET',
        description: '获取当前登录用户的详细信息',
        permission: 'private',
        allowedRoles: ['normal_user', 'vip_user', 'super_admin'],
        category: '用户',
        status: 'active',
        createdAt: '2024-01-15 10:30:00',
        updatedAt: '2024-01-15 10:30:00'
      },
      {
        id: '3',
        name: '创建文章',
        path: '/api/article/create',
        method: 'POST',
        description: '创建新的文章内容',
        permission: 'private',
        allowedRoles: ['content_admin', 'super_admin'],
        category: '内容',
        status: 'active',
        createdAt: '2024-01-15 11:00:00',
        updatedAt: '2024-01-15 11:00:00'
      },
      {
        id: '4',
        name: '获取文章列表',
        path: '/api/article/list',
        method: 'GET',
        description: '获取文章列表，支持分页和筛选',
        permission: 'public',
        allowedRoles: [],
        category: '内容',
        status: 'active',
        createdAt: '2024-01-15 11:30:00',
        updatedAt: '2024-01-15 11:30:00'
      },
      {
        id: '5',
        name: '删除文章',
        path: '/api/article/delete',
        method: 'DELETE',
        description: '删除指定的文章',
        permission: 'private',
        allowedRoles: ['content_admin', 'super_admin'],
        category: '内容',
        status: 'active',
        createdAt: '2024-01-15 12:00:00',
        updatedAt: '2024-01-15 12:00:00'
      },
      {
        id: '6',
        name: '文件上传',
        path: '/api/file/upload',
        method: 'POST',
        description: '上传文件到服务器',
        permission: 'private',
        allowedRoles: ['normal_user', 'vip_user', 'super_admin'],
        category: '文件',
        status: 'active',
        createdAt: '2024-01-15 12:30:00',
        updatedAt: '2024-01-15 12:30:00'
      },
      {
        id: '7',
        name: '获取系统设置',
        path: '/api/system/settings',
        method: 'GET',
        description: '获取系统配置信息',
        permission: 'private',
        allowedRoles: ['system_admin', 'super_admin'],
        category: '系统',
        status: 'active',
        createdAt: '2024-01-15 13:00:00',
        updatedAt: '2024-01-15 13:00:00'
      },
      {
        id: '8',
        name: '更新系统设置',
        path: '/api/system/settings',
        method: 'PUT',
        description: '更新系统配置信息',
        permission: 'private',
        allowedRoles: ['super_admin'],
        category: '系统',
        status: 'active',
        createdAt: '2024-01-15 13:30:00',
        updatedAt: '2024-01-15 13:30:00'
      }
    ]
  } catch {
    message.error('API列表加载失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadApiList()
})
</script>

<style scoped>
.api-management-container {
  padding: 24px;
  min-height: 100vh;
  background-color: #f5f7fa;
}

.api-card {
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

.api-table {
  margin-top: 16px;
}

.action-buttons {
  display: flex;
  gap: 8px;
}

.permission-modal-content,
.test-modal-content,
.batch-modal-content {
  padding: 16px 0;
}

.role-tree-container {
  border: 1px solid #d9d9d9;
  border-radius: 6px;
  padding: 12px;
  background-color: #fafafa;
  max-height: 300px;
  overflow-y: auto;
}

.role-permission-section,
.batch-role-section {
  margin-top: 16px;
}

.slider-value {
  margin-left: 12px;
  min-width: 40px;
  text-align: center;
  font-weight: 500;
}

@media (max-width: 768px) {
  .api-management-container {
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
