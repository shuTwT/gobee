<script setup lang="ts">
import { apiClient, useApi } from '@/api'
import SchemaForm from '@/components/SchemaForm/index.vue'
import type { SettingFormGroupProps } from '@/components/SchemaForm/types'

const props = defineProps<{
  theme: {
    id?: number
    name?: string
    display_name?: string
    type?: string
  }
}>()

const loading = ref(false)
const forms = ref<SettingFormGroupProps[]>([])
const values = ref<Record<string, any>>({})
const schemaFormRef = ref<InstanceType<typeof SchemaForm> | null>(null)

onMounted(async () => {
  if (!props.theme.id) return
  loading.value = true
  try {
    const res = await useApi(apiClient.api.v1ThemeSettingList, props.theme.id)
    forms.value = res.data?.forms || []
    values.value = res.data?.values || {}
  } catch (error) {
    console.error('获取主题设置失败:', error)
    window.$message?.error('获取主题设置失败')
  } finally {
    loading.value = false
  }
})

const hasSetting = computed(() => forms.value.some((form) => (form.formSchema || []).length > 0))

const getData = () => {
  if (!hasSetting.value) return Promise.resolve(null)
  return schemaFormRef.value?.getData()
}

defineExpose({ getData })
</script>

<template>
  <div style="padding: 20px 20px 0">
    <n-spin :show="loading">
      <n-empty
        v-if="!loading && !hasSetting"
        description="该主题没有可配置项"
        style="padding: 40px 0"
      />
      <SchemaForm v-else-if="!loading" ref="schemaFormRef" :forms="forms" :model-value="values" />
    </n-spin>
  </div>
</template>
