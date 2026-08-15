<script setup lang="ts">
import { NButton, NTag, useThemeVars, type DataTableColumns } from 'naive-ui'
import * as aiApi from '@/api/ai'
import type { AIProvider, AIModelItem, AIProviderInput, AIModelInput } from '@/api/ai'

const message = useMessage()
const dialog = useDialog()
const themeVars = useThemeVars()

const cssVars = computed(() => ({
  '--selected-color': themeVars.value.primaryColor,
}))

// ---------- 提供商列表 ----------

const providers = ref<AIProvider[]>([])
const loading = ref(false)
const selectedProviderId = ref<number | null>(null)

const selectedProvider = computed(
  () => providers.value.find(p => p.id === selectedProviderId.value) || null,
)

const loadProviders = async () => {
  loading.value = true
  try {
    const res = await aiApi.listProviders()
    providers.value = res.data
    if (!providers.value.some(p => p.id === selectedProviderId.value)) {
      selectedProviderId.value = providers.value[0]?.id ?? null
    }
  } catch {
    message.error('加载 AI 提供商列表失败')
  } finally {
    loading.value = false
  }
}

// 常见提供商预设：选择类型时自动填充 API 地址
const providerPresets: { label: string; value: string; baseUrl: string }[] = [
  { label: 'OpenAI', value: 'openai', baseUrl: 'https://api.openai.com/v1' },
  { label: 'DeepSeek', value: 'deepseek', baseUrl: 'https://api.deepseek.com/v1' },
  { label: 'Moonshot (Kimi)', value: 'moonshot', baseUrl: 'https://api.moonshot.cn/v1' },
  { label: '智谱 GLM', value: 'zhipu', baseUrl: 'https://open.bigmodel.cn/api/paas/v4' },
  { label: '火山方舟 (豆包)', value: 'volcengine', baseUrl: 'https://ark.cn-beijing.volces.com/api/v3' },
  { label: '硅基流动', value: 'siliconflow', baseUrl: 'https://api.siliconflow.cn/v1' },
  { label: 'Ollama (本地)', value: 'ollama', baseUrl: 'http://localhost:11434/v1' },
  { label: '自定义', value: 'custom', baseUrl: '' },
]
const providerTypeOptions = providerPresets.map(p => ({ label: p.label, value: p.value }))

const providerRowProps = (row: AIProvider) => ({
  style: { cursor: 'pointer' },
  onClick: () => {
    selectedProviderId.value = row.id
  },
})
const providerRowClassName = (row: AIProvider) =>
  row.id === selectedProviderId.value ? 'selected-row' : ''

const providerColumns: DataTableColumns<AIProvider> = [
  {
    title: '名称',
    key: 'name',
    width: 170,
    render: row =>
      h('div', { class: 'flex items-center gap-2' }, [
        h('span', { class: 'truncate' }, row.name),
        row.is_default
          ? h(NTag, { size: 'small', type: 'info', bordered: false }, { default: () => '默认' })
          : null,
      ]),
  },
  {
    title: '类型',
    key: 'provider_type',
    width: 110,
    render: row => h(NTag, { size: 'small', bordered: false }, { default: () => row.provider_type }),
  },
  { title: 'API 地址', key: 'base_url', ellipsis: { tooltip: true } },
  {
    title: 'Key',
    key: 'api_key_configured',
    width: 90,
    render: row =>
      row.api_key_configured
        ? h(NTag, { size: 'small', type: 'success', bordered: false }, { default: () => '已配置' })
        : h(NTag, { size: 'small', type: 'warning', bordered: false }, { default: () => '未配置' }),
  },
  {
    title: '状态',
    key: 'is_enabled',
    width: 80,
    render: row =>
      row.is_enabled
        ? h(NTag, { size: 'small', type: 'success', bordered: false }, { default: () => '启用' })
        : h(NTag, { size: 'small', type: 'default', bordered: false }, { default: () => '停用' }),
  },
  {
    title: '操作',
    key: 'actions',
    width: 230,
    render: row =>
      h('div', { class: 'flex gap-2' }, [
        h(NButton, { size: 'small', onClick: () => openEditProvider(row) }, { default: () => '编辑' }),
        h(
          NButton,
          {
            size: 'small',
            disabled: row.is_default,
            onClick: () => handleSetDefault(row),
          },
          { default: () => '设为默认' },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'error',
            quaternary: true,
            onClick: () => handleDeleteProvider(row),
          },
          { default: () => '删除' },
        ),
      ]),
  },
]

const handleSetDefault = async (row: AIProvider) => {
  try {
    await aiApi.updateProvider(row.id, {
      name: row.name,
      provider_type: row.provider_type,
      base_url: row.base_url,
      api_key: '',
      temperature: row.temperature,
      max_tokens: row.max_tokens,
      top_p: row.top_p,
      frequency_penalty: row.frequency_penalty,
      presence_penalty: row.presence_penalty,
      is_default: true,
      is_enabled: row.is_enabled,
      sort: row.sort,
      remark: row.remark,
    })
    message.success('已设为默认提供商')
    await loadProviders()
  } catch {
    message.error('设置默认提供商失败')
  }
}

const handleDeleteProvider = (row: AIProvider) => {
  dialog.warning({
    title: '删除提供商',
    content: `确定删除提供商「${row.name}」吗？其下的 ${row.models.length} 个模型将一并删除。`,
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      try {
        await aiApi.deleteProvider(row.id)
        message.success('提供商已删除')
        await loadProviders()
      } catch {
        message.error('删除提供商失败')
      }
    },
  })
}

// ---------- 提供商表单 ----------

const providerFormVisible = ref(false)
const providerFormRef = ref<FormInst | null>(null)
const editingProviderId = ref<number | null>(null)
const providerSaving = ref(false)
const providerTesting = ref(false)

const defaultProviderForm = (): AIProviderInput => ({
  name: '',
  provider_type: 'openai',
  base_url: 'https://api.openai.com/v1',
  api_key: '',
  temperature: 0.7,
  max_tokens: 2048,
  top_p: 1,
  frequency_penalty: 0,
  presence_penalty: 0,
  is_default: false,
  is_enabled: true,
  sort: 0,
  remark: '',
})

const providerForm = ref<AIProviderInput>(defaultProviderForm())

const providerRules = {
  name: { required: true, message: '请输入提供商名称', trigger: ['blur', 'input'] },
  base_url: { required: true, message: '请输入 API 地址', trigger: ['blur', 'input'] },
}

const openCreateProvider = () => {
  editingProviderId.value = null
  providerForm.value = defaultProviderForm()
  providerFormVisible.value = true
}

const openEditProvider = (row: AIProvider) => {
  editingProviderId.value = row.id
  providerForm.value = {
    name: row.name,
    provider_type: row.provider_type,
    base_url: row.base_url,
    api_key: '',
    temperature: row.temperature,
    max_tokens: row.max_tokens,
    top_p: row.top_p,
    frequency_penalty: row.frequency_penalty,
    presence_penalty: row.presence_penalty,
    is_default: row.is_default,
    is_enabled: row.is_enabled,
    sort: row.sort,
    remark: row.remark,
  }
  providerFormVisible.value = true
}

const onProviderTypeChange = (value: string) => {
  const preset = providerPresets.find(p => p.value === value)
  if (preset) providerForm.value.base_url = preset.baseUrl
}

const saveProvider = async () => {
  if (!providerFormRef.value) return
  try {
    await providerFormRef.value.validate()
  } catch {
    return
  }
  providerSaving.value = true
  try {
    if (editingProviderId.value === null) {
      await aiApi.createProvider({ ...providerForm.value })
      message.success('提供商创建成功')
    } else {
      await aiApi.updateProvider(editingProviderId.value, { ...providerForm.value })
      message.success('提供商保存成功')
    }
    providerFormVisible.value = false
    await loadProviders()
  } catch {
    message.error('保存提供商失败')
  } finally {
    providerSaving.value = false
  }
}

const testProvider = async () => {
  if (!providerForm.value.base_url.trim()) {
    message.warning('请先填写 API 地址')
    return
  }
  if (!providerForm.value.api_key.trim() && !providerForm.value.base_url.includes('localhost')) {
    message.warning('测试连接前请输入 API Key（本地服务可留空）')
    return
  }
  providerTesting.value = true
  try {
    await aiApi.testProvider({
      base_url: providerForm.value.base_url,
      api_key: providerForm.value.api_key,
    })
    message.success('连接测试成功')
  } catch {
    message.error('连接测试失败，请检查地址与密钥')
  } finally {
    providerTesting.value = false
  }
}

// ---------- 模型子表 ----------

const modelColumns: DataTableColumns<AIModelItem> = [
  { title: '模型名称', key: 'model_name' },
  {
    title: '显示名称',
    key: 'display_name',
    render: row => row.display_name || '-',
  },
  {
    title: '状态',
    key: 'is_enabled',
    width: 90,
    render: row =>
      row.is_enabled
        ? h(NTag, { size: 'small', type: 'success', bordered: false }, { default: () => '启用' })
        : h(NTag, { size: 'small', type: 'default', bordered: false }, { default: () => '停用' }),
  },
  { title: '排序', key: 'sort', width: 70 },
  {
    title: '操作',
    key: 'actions',
    width: 130,
    render: row =>
      h('div', { class: 'flex gap-2' }, [
        h(NButton, { size: 'small', onClick: () => openEditModel(row) }, { default: () => '编辑' }),
        h(
          NButton,
          {
            size: 'small',
            type: 'error',
            quaternary: true,
            onClick: () => handleDeleteModel(row),
          },
          { default: () => '删除' },
        ),
      ]),
  },
]

const modelFormVisible = ref(false)
const modelFormRef = ref<FormInst | null>(null)
const editingModelId = ref<number | null>(null)
const modelSaving = ref(false)
const syncingModels = ref(false)

const defaultModelForm = (): AIModelInput => ({
  model_name: '',
  display_name: '',
  is_enabled: true,
  sort: 0,
})

const modelForm = ref<AIModelInput>(defaultModelForm())

const modelRules = {
  model_name: { required: true, message: '请输入模型名称', trigger: ['blur', 'input'] },
}

const openCreateModel = () => {
  if (!selectedProvider.value) {
    message.warning('请先在左侧选择提供商')
    return
  }
  editingModelId.value = null
  modelForm.value = defaultModelForm()
  modelFormVisible.value = true
}

const openEditModel = (row: AIModelItem) => {
  editingModelId.value = row.id
  modelForm.value = {
    model_name: row.model_name,
    display_name: row.display_name || '',
    is_enabled: row.is_enabled,
    sort: row.sort,
  }
  modelFormVisible.value = true
}

const saveModel = async () => {
  if (!selectedProvider.value || !modelFormRef.value) return
  try {
    await modelFormRef.value.validate()
  } catch {
    return
  }
  modelSaving.value = true
  try {
    if (editingModelId.value === null) {
      await aiApi.createProviderModel(selectedProvider.value.id, { ...modelForm.value })
      message.success('模型添加成功')
    } else {
      await aiApi.updateProviderModel(selectedProvider.value.id, editingModelId.value, {
        ...modelForm.value,
      })
      message.success('模型保存成功')
    }
    modelFormVisible.value = false
    await loadProviders()
  } catch {
    message.error('保存模型失败')
  } finally {
    modelSaving.value = false
  }
}

const handleDeleteModel = (row: AIModelItem) => {
  dialog.warning({
    title: '删除模型',
    content: `确定删除模型「${row.model_name}」吗？`,
    positiveText: '删除',
    negativeText: '取消',
    onPositiveClick: async () => {
      if (!selectedProvider.value) return
      try {
        await aiApi.deleteProviderModel(selectedProvider.value.id, row.id)
        message.success('模型已删除')
        await loadProviders()
      } catch {
        message.error('删除模型失败')
      }
    },
  })
}

const handleSyncModels = async () => {
  if (!selectedProvider.value) return
  syncingModels.value = true
  try {
    await aiApi.syncProviderModels(selectedProvider.value.id)
    message.success('模型同步成功')
    await loadProviders()
  } catch {
    message.error('模型同步失败，请检查提供商配置')
  } finally {
    syncingModels.value = false
  }
}

onMounted(loadProviders)
</script>

<template>
  <div class="container-fluid p-6">
    <div :style="cssVars" class="flex items-start gap-4">
      <!-- 左：AI 提供商 -->
      <n-card title="AI 提供商" :style="{ width: '560px' }" class="shrink-0">
        <template #header-extra>
          <n-button type="primary" size="small" @click="openCreateProvider">添加提供商</n-button>
        </template>
        <n-data-table
          :columns="providerColumns"
          :data="providers"
          :loading="loading"
          :row-props="providerRowProps"
          :row-class-name="providerRowClassName"
          size="small"
          :max-height="480"
        />
        <p class="mt-3 text-xs leading-5 text-gray-400">
          点击左侧某一行可在右侧管理其模型。默认提供商用于 AI 聊天与文章摘要。
        </p>
      </n-card>

      <!-- 右：模型管理子表 -->
      <n-card class="min-w-0 flex-1">
        <template #header>
          <div class="flex items-center justify-between gap-4">
            <span class="truncate">模型管理{{ selectedProvider ? ` - ${selectedProvider.name}` : '' }}</span>
            <div class="flex shrink-0 gap-2">
              <n-button
                size="small"
                :loading="syncingModels"
                :disabled="!selectedProvider"
                @click="handleSyncModels"
              >
                从供应商同步模型
              </n-button>
              <n-button
                size="small"
                type="primary"
                :disabled="!selectedProvider"
                @click="openCreateModel"
              >
                添加模型
              </n-button>
            </div>
          </div>
        </template>
        <n-empty v-if="!selectedProvider" description="请先在左侧选择提供商" class="py-20" />
        <n-data-table
          v-else
          :columns="modelColumns"
          :data="selectedProvider.models"
          size="small"
          :max-height="480"
        />
      </n-card>

      <!-- 提供商弹窗 -->
      <n-modal
        v-model:show="providerFormVisible"
        preset="card"
        :title="editingProviderId === null ? '添加提供商' : '编辑提供商'"
        style="width: 640px"
      >
        <n-form
          ref="providerFormRef"
          :model="providerForm"
          :rules="providerRules"
          label-placement="left"
          label-width="110px"
        >
          <n-form-item label="提供商名称" path="name">
            <n-input v-model:value="providerForm.name" placeholder="如 DeepSeek、OpenAI" />
          </n-form-item>
          <n-form-item label="提供商类型" path="provider_type">
            <n-select
              v-model:value="providerForm.provider_type"
              :options="providerTypeOptions"
              @update:value="onProviderTypeChange"
            />
          </n-form-item>
          <n-form-item label="API 地址" path="base_url">
            <n-input v-model:value="providerForm.base_url" placeholder="https://api.openai.com/v1" />
          </n-form-item>
          <n-form-item label="API Key" path="api_key">
            <n-input
              v-model:value="providerForm.api_key"
              type="password"
              show-password-on="mousedown"
              :placeholder="editingProviderId === null ? '本地服务（如 Ollama）可留空' : '留空表示不修改'"
            />
          </n-form-item>
          <n-form-item label="状态" path="is_enabled">
            <n-switch v-model:value="providerForm.is_enabled" />
            <span class="ml-2 text-xs text-gray-400">停用后 AI 聊天将不会使用该提供商</span>
          </n-form-item>
          <n-form-item label="设为默认" path="is_default">
            <n-switch v-model:value="providerForm.is_default" />
            <span class="ml-2 text-xs text-gray-400">默认提供商用于 AI 聊天与文章摘要</span>
          </n-form-item>
          <n-form-item label="排序" path="sort">
            <n-input-number v-model:value="providerForm.sort" :min="0" />
          </n-form-item>
          <n-form-item label="备注" path="remark">
            <n-input v-model:value="providerForm.remark" type="textarea" :rows="2" placeholder="可选" />
          </n-form-item>
          <n-collapse>
            <n-collapse-item title="高级参数（采样）" name="advanced">
              <n-form-item label="温度" path="temperature">
                <n-slider v-model:value="providerForm.temperature" :min="0" :max="2" :step="0.1" />
                <span class="slider-value">{{ providerForm.temperature }}</span>
              </n-form-item>
              <n-form-item label="最大令牌数" path="max_tokens">
                <n-input-number v-model:value="providerForm.max_tokens" :min="1" :max="8192" />
              </n-form-item>
              <n-form-item label="Top P" path="top_p">
                <n-slider v-model:value="providerForm.top_p" :min="0" :max="1" :step="0.1" />
                <span class="slider-value">{{ providerForm.top_p }}</span>
              </n-form-item>
              <n-form-item label="频率惩罚" path="frequency_penalty">
                <n-slider
                  v-model:value="providerForm.frequency_penalty"
                  :min="-2"
                  :max="2"
                  :step="0.1"
                />
                <span class="slider-value">{{ providerForm.frequency_penalty }}</span>
              </n-form-item>
              <n-form-item label="存在惩罚" path="presence_penalty">
                <n-slider
                  v-model:value="providerForm.presence_penalty"
                  :min="-2"
                  :max="2"
                  :step="0.1"
                />
                <span class="slider-value">{{ providerForm.presence_penalty }}</span>
              </n-form-item>
            </n-collapse-item>
          </n-collapse>
        </n-form>
        <template #footer>
          <div class="flex justify-end gap-3">
            <n-button :loading="providerTesting" @click="testProvider">测试连接</n-button>
            <n-button type="primary" :loading="providerSaving" @click="saveProvider">保存</n-button>
          </div>
        </template>
      </n-modal>

      <!-- 模型弹窗 -->
      <n-modal
        v-model:show="modelFormVisible"
        preset="card"
        :title="editingModelId === null ? '添加模型' : '编辑模型'"
        style="width: 480px"
      >
        <n-form
          ref="modelFormRef"
          :model="modelForm"
          :rules="modelRules"
          label-placement="left"
          label-width="90px"
        >
          <n-form-item label="模型名称" path="model_name">
            <n-input v-model:value="modelForm.model_name" placeholder="如 gpt-4o、deepseek-chat" />
          </n-form-item>
          <n-form-item label="显示名称" path="display_name">
            <n-input v-model:value="modelForm.display_name" placeholder="可选" />
          </n-form-item>
          <n-form-item label="启用" path="is_enabled">
            <n-switch v-model:value="modelForm.is_enabled" />
          </n-form-item>
          <n-form-item label="排序" path="sort">
            <n-input-number v-model:value="modelForm.sort" :min="0" />
          </n-form-item>
        </n-form>
        <template #footer>
          <div class="flex justify-end gap-3">
            <n-button @click="modelFormVisible = false">取消</n-button>
            <n-button type="primary" :loading="modelSaving" @click="saveModel">保存</n-button>
          </div>
        </template>
      </n-modal>
    </div>
  </div>
</template>

<style scoped>
:deep(.selected-row td) {
  background-color: color-mix(in srgb, var(--selected-color) 10%, transparent);
}
</style>
