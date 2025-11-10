<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import type { SettingsProps } from '../utils/types';

const props = defineProps<{
  settings:SettingsProps
}>()
const emit = defineEmits(["refresh"])
const message = useMessage()

const backupFormRef = ref<FormInst | null>(null)

const backupLoading = ref(false)
const backupForm = reactive({
  enableAutoBackup: true,
  backupFrequency: 'daily',
  backupRetentionDays: 30,
  backupStorageLocation: 'local',
  backupDatabase: true,
  backupUploads: true,
  backupConfig: true,
  enableRemoteBackup: false,
  remoteStorageType: 's3',
  remoteStorageConfig: '',
})
watch(()=>props.settings,(newSettings)=>{
  backupForm.enableAutoBackup = newSettings.enableAutoBackup ?? true
  backupForm.backupFrequency = newSettings.backupFrequency ?? 'daily'
  backupForm.backupRetentionDays= newSettings.backupRetentionDays ?? 30
  backupForm.backupStorageLocation= newSettings.backupStorageLocation='local'
  backupForm.backupDatabase = newSettings.backupDatabase=true
  backupForm.backupUploads = newSettings.backupUploads=true
  backupForm.backupConfig = newSettings.backupConfig=true
  backupForm.enableRemoteBackup = newSettings.enableRemoteBackup=false
  backupForm.remoteStorageType = newSettings.remoteStorageType='s3'
  backupForm.remoteStorageConfig = newSettings.remoteStorageConfig=''

})
const backupFrequencyOptions = [
  { label: '每小时', value: 'hourly' },
  { label: '每天', value: 'daily' },
  { label: '每周', value: 'weekly' },
  { label: '每月', value: 'monthly' },
]

const backupStorageOptions = [
  { label: '本地存储', value: 'local' },
  { label: '云存储', value: 'cloud' },
  { label: '混合存储', value: 'hybrid' },
]

const remoteStorageOptions = [
  { label: 'Amazon S3', value: 's3' },
  { label: '阿里云OSS', value: 'aliyun' },
  { label: '腾讯云COS', value: 'tencent' },
  { label: '七牛云', value: 'qiniu' },
]
// 保存备份设置
const saveBackupSettings = async () => {
  backupLoading.value = true
  try {
    await new Promise((resolve) => setTimeout(resolve, 1000))
    emit('refresh')
    message.success('备份设置保存成功')
  } catch {
    message.error('备份设置保存失败')
  } finally {
    backupLoading.value = false
  }
}

// 手动备份
const manualBackup = async () => {
  try {
    await new Promise((resolve) => setTimeout(resolve, 2000))
    message.success('手动备份成功')
  } catch {
    message.error('手动备份失败')
  }
}

// 查看备份历史
const viewBackupHistory = async () => {
  message.info('备份历史功能开发中...')
}
</script>
<template>
  <n-form
    ref="backupFormRef"
    :model="backupForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="启用自动备份" path="enableAutoBackup">
      <n-switch v-model:value="backupForm.enableAutoBackup" />
    </n-form-item>
    <n-form-item label="备份频率" path="backupFrequency">
      <n-select v-model:value="backupForm.backupFrequency" :options="backupFrequencyOptions" />
    </n-form-item>
    <n-form-item label="备份保留天数" path="backupRetentionDays">
      <n-input-number v-model:value="backupForm.backupRetentionDays" :min="1" :max="365">
        <template #suffix>天</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="备份存储位置" path="backupStorageLocation">
      <n-select v-model:value="backupForm.backupStorageLocation" :options="backupStorageOptions" />
    </n-form-item>
    <n-divider>备份内容</n-divider>
    <n-form-item label="备份数据库" path="backupDatabase">
      <n-switch v-model:value="backupForm.backupDatabase" />
    </n-form-item>
    <n-form-item label="备份上传文件" path="backupUploads">
      <n-switch v-model:value="backupForm.backupUploads" />
    </n-form-item>
    <n-form-item label="备份配置文件" path="backupConfig">
      <n-switch v-model:value="backupForm.backupConfig" />
    </n-form-item>
    <n-divider>远程存储</n-divider>
    <n-form-item label="启用远程备份" path="enableRemoteBackup">
      <n-switch v-model:value="backupForm.enableRemoteBackup" />
    </n-form-item>
    <n-form-item label="远程存储类型" path="remoteStorageType">
      <n-select v-model:value="backupForm.remoteStorageType" :options="remoteStorageOptions" />
    </n-form-item>
    <n-form-item label="远程存储配置" path="remoteStorageConfig">
      <n-input
        v-model:value="backupForm.remoteStorageConfig"
        type="textarea"
        placeholder="请输入远程存储配置(JSON格式)"
        :rows="4"
      />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveBackupSettings" :loading="backupLoading">
        保存备份设置
      </n-button>
      <n-button @click="manualBackup" style="margin-left: 12px"> 立即备份 </n-button>
      <n-button @click="viewBackupHistory" style="margin-left: 12px"> 查看备份历史 </n-button>
    </n-form-item>
  </n-form>
</template>
