<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import type { SettingsProps } from '../utils/types';
import * as settingApi from '@/api/system/setting'

const props = defineProps<{
  settings:SettingsProps
}>()
const emit = defineEmits(["refresh"])
const message = useMessage()

const securityFormRef = ref<FormInst | null>(null)
const securityLoading = ref(false)

const securityForm = reactive({
  enableTwoFactor: false,
  maxLoginAttempts: 5,
  lockoutDuration: 15,
  sessionTimeout: 120,
  minPasswordLength: 8,
  passwordComplexity: ['uppercase', 'lowercase', 'number'],
  apiRateLimit: 100,
  ipWhitelist: [] as string[],
  ipBlacklist: [] as string[]
})

watch(()=>props.settings,(newSettings)=>{
  securityForm.enableTwoFactor = newSettings.enableTwoFactor || false
  securityForm.maxLoginAttempts = newSettings.maxLoginAttempts || 5
  securityForm.lockoutDuration = newSettings.lockoutDuration || 15
  securityForm.sessionTimeout = newSettings.sessionTimeout || 120
  securityForm.minPasswordLength = newSettings.minPasswordLength || 8
  securityForm.passwordComplexity = newSettings.passwordComplexity || ['uppercase', 'lowercase', 'number']
  securityForm.apiRateLimit = newSettings.apiRateLimit || 100
  securityForm.ipWhitelist = newSettings.ipWhitelist || []
  securityForm.ipBlacklist = newSettings.ipBlacklist || []
})

// 保存安全设置
const saveSecuritySettings = async () => {
  securityLoading.value = true
  try {
    await settingApi.saveSettings(securityForm)
    await new Promise(resolve => setTimeout(resolve, 1000))
    emit('refresh')
    message.success('安全设置保存成功')
  } catch {
    message.error('安全设置保存失败')
  } finally {
    securityLoading.value = false
  }
}
</script>
<template>
  <n-form
    ref="securityFormRef"
    :model="securityForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="启用两步验证" path="enableTwoFactor">
      <n-switch v-model:value="securityForm.enableTwoFactor" />
    </n-form-item>
    <n-form-item label="登录失败重试次数" path="maxLoginAttempts">
      <n-input-number v-model:value="securityForm.maxLoginAttempts" :min="3" :max="10" />
    </n-form-item>
    <n-form-item label="账户锁定时间" path="lockoutDuration">
      <n-input-number v-model:value="securityForm.lockoutDuration" :min="5" :max="60">
        <template #suffix>分钟</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="会话超时时间" path="sessionTimeout">
      <n-input-number v-model:value="securityForm.sessionTimeout" :min="30" :max="1440">
        <template #suffix>分钟</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="密码最小长度" path="minPasswordLength">
      <n-input-number v-model:value="securityForm.minPasswordLength" :min="6" :max="32" />
    </n-form-item>
    <n-form-item label="密码复杂度要求" path="passwordComplexity">
      <n-checkbox-group v-model:value="securityForm.passwordComplexity">
        <n-space vertical>
          <n-checkbox value="uppercase">包含大写字母</n-checkbox>
          <n-checkbox value="lowercase">包含小写字母</n-checkbox>
          <n-checkbox value="number">包含数字</n-checkbox>
          <n-checkbox value="special">包含特殊字符</n-checkbox>
        </n-space>
      </n-checkbox-group>
    </n-form-item>
    <n-form-item label="API访问限制" path="apiRateLimit">
      <n-input-number v-model:value="securityForm.apiRateLimit" :min="10" :max="1000">
        <template #suffix>次/分钟</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="IP白名单" path="ipWhitelist">
      <n-dynamic-tags v-model:value="securityForm.ipWhitelist" />
    </n-form-item>
    <n-form-item label="IP黑名单" path="ipBlacklist">
      <n-dynamic-tags v-model:value="securityForm.ipBlacklist" />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveSecuritySettings" :loading="securityLoading">
        保存安全设置
      </n-button>
    </n-form-item>
  </n-form>
</template>
