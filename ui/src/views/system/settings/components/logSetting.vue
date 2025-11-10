<script setup lang="ts">
import type { FormInst } from 'naive-ui';
import type { SettingsProps } from '../utils/types';

const props = defineProps<{
  settings:SettingsProps
}>()
const emit = defineEmits(["refresh"])
const message = useMessage()

const logsFormRef = ref<FormInst | null>(null)

const logsLoading = ref(false)
const logsForm = reactive({
  enableSystemLogs: true,
  logLevel: 'info',
  logRetentionDays: 30,
  maxLogFileSize: 100,
  logUserActions: true,
  logSystemErrors: true,
  logDatabaseQueries: false,
  logApiCalls: true,
  logPaymentActions: true,
  logStorageType: 'file',
  externalLogService: 'none',
  externalLogConfig: ''
})

watch(()=>props.settings,(newSettings)=>{
  logsForm.enableSystemLogs = newSettings.enableSystemLogs || true
  logsForm.logLevel = newSettings.logLevel || 'info'
  logsForm.logRetentionDays = newSettings.logRetentionDays || 30
  logsForm.maxLogFileSize = newSettings.maxLogFileSize || 100
  logsForm.logUserActions = newSettings.logUserActions || true
  logsForm.logSystemErrors = newSettings.logSystemErrors || true
  logsForm.logDatabaseQueries = newSettings.logDatabaseQueries || false
  logsForm.logApiCalls = newSettings.logApiCalls || true
  logsForm.logPaymentActions = newSettings.logPaymentActions || true
  logsForm.logStorageType = newSettings.logStorageType || 'file'
  logsForm.externalLogService = newSettings.externalLogService || 'none'
  logsForm.externalLogConfig = newSettings.externalLogConfig || ''
})

const logLevelOptions = [
  { label: '调试', value: 'debug' },
  { label: '信息', value: 'info' },
  { label: '警告', value: 'warning' },
  { label: '错误', value: 'error' },
  { label: '致命', value: 'fatal' }
]

const logStorageOptions = [
  { label: '文件存储', value: 'file' },
  { label: '数据库存储', value: 'database' },
  { label: '混合存储', value: 'hybrid' }
]

const externalLogServiceOptions = [
  { label: '不使用', value: 'none' },
  { label: 'ELK Stack', value: 'elk' },
  { label: 'Graylog', value: 'graylog' },
  { label: 'Splunk', value: 'splunk' },
  { label: 'Datadog', value: 'datadog' }
]

// 保存日志设置
const saveLogsSettings = async () => {
  logsLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    emit('refresh')
    message.success('日志设置保存成功')
  } catch {
    message.error('日志设置保存失败')
  } finally {
    logsLoading.value = false
  }
}

// 查看当前日志
const viewCurrentLogs = async () => {
  message.info('当前日志功能开发中...')
}

// 导出日志
const exportLogs = async () => {
  try {
    // 模拟导出日志
    const logData = '系统日志数据...'
    const blob = new Blob([logData], { type: 'text/plain' })
    const url = URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = `system-logs-${new Date().toISOString().split('T')[0]}.txt`
    link.click()
    URL.revokeObjectURL(url)
    message.success('日志导出成功')
  } catch {
    message.error('日志导出失败')
  }
}
</script>
<template>
  <n-form
    ref="logsFormRef"
    :model="logsForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="启用系统日志" path="enableSystemLogs">
      <n-switch v-model:value="logsForm.enableSystemLogs" />
    </n-form-item>
    <n-form-item label="日志级别" path="logLevel">
      <n-select v-model:value="logsForm.logLevel" :options="logLevelOptions" />
    </n-form-item>
    <n-form-item label="日志保留天数" path="logRetentionDays">
      <n-input-number v-model:value="logsForm.logRetentionDays" :min="1" :max="365">
        <template #suffix>天</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="日志文件最大大小" path="maxLogFileSize">
      <n-input-number v-model:value="logsForm.maxLogFileSize" :min="1" :max="1000">
        <template #suffix>MB</template>
      </n-input-number>
    </n-form-item>
    <n-divider>日志类型</n-divider>
    <n-form-item label="记录用户操作日志" path="logUserActions">
      <n-switch v-model:value="logsForm.logUserActions" />
    </n-form-item>
    <n-form-item label="记录系统错误日志" path="logSystemErrors">
      <n-switch v-model:value="logsForm.logSystemErrors" />
    </n-form-item>
    <n-form-item label="记录数据库查询日志" path="logDatabaseQueries">
      <n-switch v-model:value="logsForm.logDatabaseQueries" />
    </n-form-item>
    <n-form-item label="记录API调用日志" path="logApiCalls">
      <n-switch v-model:value="logsForm.logApiCalls" />
    </n-form-item>
    <n-form-item label="记录支付相关日志" path="logPaymentActions">
      <n-switch v-model:value="logsForm.logPaymentActions" />
    </n-form-item>
    <n-divider>日志存储</n-divider>
    <n-form-item label="日志存储方式" path="logStorageType">
      <n-select v-model:value="logsForm.logStorageType" :options="logStorageOptions" />
    </n-form-item>
    <n-form-item label="外部日志服务" path="externalLogService">
      <n-select v-model:value="logsForm.externalLogService" :options="externalLogServiceOptions" />
    </n-form-item>
    <n-form-item label="外部日志配置" path="externalLogConfig">
      <n-input
        v-model:value="logsForm.externalLogConfig"
        type="textarea"
        placeholder="请输入外部日志服务配置(JSON格式)"
        :rows="4"
      />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveLogsSettings" :loading="logsLoading">
        保存日志设置
      </n-button>
      <n-button @click="viewCurrentLogs" style="margin-left: 12px"> 查看当前日志 </n-button>
      <n-button @click="exportLogs" style="margin-left: 12px"> 导出日志 </n-button>
    </n-form-item>
  </n-form>
</template>
