<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import * as settingApi from '@/api/system/setting'

const message = useMessage()

const notificationFormRef = ref<FormInst | null>(null)
const defaultForm = {
  enableEmailNotification: true,
  enableSmsNotification: false,
  enableInAppNotification: true,
  notifyNewUserRegistration: true,
  notifyLoginAbnormal: true,
  notifyPasswordChange: true,
  notifyNewOrder: true,
  notifyOrderPaid: true,
  notifyOrderShipped: true,
  notifyOrderCompleted: true,
  notifySystemError: true,
  notifyMaintenance: true,
  notificationEmail: ''
}
const notificationForm = ref({
  enableEmailNotification: true,
  enableSmsNotification: false,
  enableInAppNotification: true,
  notifyNewUserRegistration: true,
  notifyLoginAbnormal: true,
  notifyPasswordChange: true,
  notifyNewOrder: true,
  notifyOrderPaid: true,
  notifyOrderShipped: true,
  notifyOrderCompleted: true,
  notifySystemError: true,
  notifyMaintenance: true,
  notificationEmail: ''
})
const notificationLoading = ref(false)


// 保存通知设置
const saveNotificationSettings = async () => {
  notificationLoading.value = true
  try {
    await settingApi.saveSettings('notify',notificationForm.value)
    await new Promise(resolve => setTimeout(resolve, 1000))
    onSearch()
    message.success('通知设置保存成功')
  } catch {
    message.error('通知设置保存失败')
  } finally {
    notificationLoading.value = false
  }
}

// 测试通知
const testNotification = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('通知测试成功')
  } catch {
    message.error('通知测试失败')
  }
}

const onSearch = async()=>{
  const res = await settingApi.getSettingsMap('notify')
  notificationForm.value = Object.assign({},defaultForm,res.data)
}

onMounted(()=>{
  onSearch()
})
</script>
<template>
  <n-form
    ref="notificationFormRef"
    :model="notificationForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="启用邮件通知" path="enableEmailNotification">
      <n-switch v-model:value="notificationForm.enableEmailNotification" />
    </n-form-item>
    <n-form-item label="启用短信通知" path="enableSmsNotification">
      <n-switch v-model:value="notificationForm.enableSmsNotification" />
    </n-form-item>
    <n-form-item label="启用站内通知" path="enableInAppNotification">
      <n-switch v-model:value="notificationForm.enableInAppNotification" />
    </n-form-item>
    <n-divider>用户相关通知</n-divider>
    <n-form-item label="新用户注册通知" path="notifyNewUserRegistration">
      <n-switch v-model:value="notificationForm.notifyNewUserRegistration" />
    </n-form-item>
    <n-form-item label="用户登录异常通知" path="notifyLoginAbnormal">
      <n-switch v-model:value="notificationForm.notifyLoginAbnormal" />
    </n-form-item>
    <n-form-item label="密码修改通知" path="notifyPasswordChange">
      <n-switch v-model:value="notificationForm.notifyPasswordChange" />
    </n-form-item>
    <n-divider>订单相关通知</n-divider>
    <n-form-item label="新订单通知" path="notifyNewOrder">
      <n-switch v-model:value="notificationForm.notifyNewOrder" />
    </n-form-item>
    <n-form-item label="订单支付成功通知" path="notifyOrderPaid">
      <n-switch v-model:value="notificationForm.notifyOrderPaid" />
    </n-form-item>
    <n-form-item label="订单发货通知" path="notifyOrderShipped">
      <n-switch v-model:value="notificationForm.notifyOrderShipped" />
    </n-form-item>
    <n-form-item label="订单完成通知" path="notifyOrderCompleted">
      <n-switch v-model:value="notificationForm.notifyOrderCompleted" />
    </n-form-item>
    <n-divider>系统相关通知</n-divider>
    <n-form-item label="系统异常通知" path="notifySystemError">
      <n-switch v-model:value="notificationForm.notifySystemError" />
    </n-form-item>
    <n-form-item label="系统维护通知" path="notifyMaintenance">
      <n-switch v-model:value="notificationForm.notifyMaintenance" />
    </n-form-item>
    <n-form-item label="通知接收邮箱" path="notificationEmail">
      <n-input
        v-model:value="notificationForm.notificationEmail"
        placeholder="管理员接收通知的邮箱"
      />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveNotificationSettings" :loading="notificationLoading">
        保存通知设置
      </n-button>
      <n-button @click="testNotification" style="margin-left: 12px"> 测试通知 </n-button>
    </n-form-item>
  </n-form>
</template>
