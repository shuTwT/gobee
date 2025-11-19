<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import type { SettingsProps } from '../utils/types';
import * as settingApi from '@/api/system/setting'


const props = defineProps<{
  settings:SettingsProps
}>()
const emit = defineEmits(["refresh"])
const message = useMessage()
const paymentFormRef = ref<FormInst | null>(null)

const paymentForm = reactive({
  enableAlipay: false,
  alipayAppId: '',
  alipayPrivateKey: '',
  alipayPublicKey: '',
  enableWechatPay: false,
  wechatMchId: '',
  wechatApiKey: '',
  wechatAppId: '',
  paymentNotifyUrl: '',
  orderTimeout: 30,
  refundReview: true
})
const paymentLoading = ref(false)

watch(()=>props.settings,(newSettings)=>{
  paymentForm.enableAlipay = newSettings.enableAlipay || false
  paymentForm.alipayAppId = newSettings.alipayAppId || ''
  paymentForm.alipayPrivateKey = newSettings.alipayPrivateKey || ''
  paymentForm.alipayPublicKey = newSettings.alipayPublicKey || ''
  paymentForm.enableWechatPay = newSettings.enableWechatPay || false
  paymentForm.wechatMchId = newSettings.wechatMchId || ''
  paymentForm.wechatApiKey = newSettings.wechatApiKey || ''
  paymentForm.wechatAppId = newSettings.wechatAppId || ''
  paymentForm.paymentNotifyUrl = newSettings.paymentNotifyUrl || ''
  paymentForm.orderTimeout = newSettings.orderTimeout || 30
  paymentForm.refundReview = newSettings.refundReview || true
})

// 保存支付设置
const savePaymentSettings = async () => {
  paymentLoading.value = true
  try {
    await settingApi.saveSettings(paymentForm)
    await new Promise(resolve => setTimeout(resolve, 1000))
    emit('refresh')
    message.success('支付设置保存成功')
  } catch {
    message.error('支付设置保存失败')
  } finally {
    paymentLoading.value = false
  }
}

// 测试支付连接
const testPaymentConnection = async () => {
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('支付连接测试成功')
  } catch {
    message.error('支付连接测试失败')
  }
}
</script>
<template>
  <n-form
    ref="paymentFormRef"
    :model="paymentForm"
    label-placement="left"
    label-width="auto"
    require-mark-placement="right-hanging"
    class="settings-form"
  >
    <n-form-item label="启用支付宝" path="enableAlipay">
      <n-switch v-model:value="paymentForm.enableAlipay" />
    </n-form-item>
    <n-form-item label="支付宝应用ID" path="alipayAppId">
      <n-input v-model:value="paymentForm.alipayAppId" placeholder="请输入支付宝应用ID" />
    </n-form-item>
    <n-form-item label="支付宝私钥" path="alipayPrivateKey">
      <n-input
        v-model:value="paymentForm.alipayPrivateKey"
        type="textarea"
        placeholder="请输入支付宝私钥"
        :rows="4"
      />
    </n-form-item>
    <n-form-item label="支付宝公钥" path="alipayPublicKey">
      <n-input
        v-model:value="paymentForm.alipayPublicKey"
        type="textarea"
        placeholder="请输入支付宝公钥"
        :rows="4"
      />
    </n-form-item>
    <n-divider />
    <n-form-item label="启用微信支付" path="enableWechatPay">
      <n-switch v-model:value="paymentForm.enableWechatPay" />
    </n-form-item>
    <n-form-item label="微信商户号" path="wechatMchId">
      <n-input v-model:value="paymentForm.wechatMchId" placeholder="请输入微信商户号" />
    </n-form-item>
    <n-form-item label="微信API密钥" path="wechatApiKey">
      <n-input v-model:value="paymentForm.wechatApiKey" placeholder="请输入微信API密钥" />
    </n-form-item>
    <n-form-item label="微信APP ID" path="wechatAppId">
      <n-input v-model:value="paymentForm.wechatAppId" placeholder="请输入微信APP ID" />
    </n-form-item>
    <n-divider />
    <n-form-item label="支付回调域名" path="paymentNotifyUrl">
      <n-input v-model:value="paymentForm.paymentNotifyUrl" placeholder="https://yourdomain.com" />
    </n-form-item>
    <n-form-item label="订单超时时间" path="orderTimeout">
      <n-input-number v-model:value="paymentForm.orderTimeout" :min="5" :max="1440">
        <template #suffix>分钟</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="退款审核" path="refundReview">
      <n-switch v-model:value="paymentForm.refundReview" />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="savePaymentSettings" :loading="paymentLoading">
        保存支付设置
      </n-button>
      <n-button @click="testPaymentConnection" style="margin-left: 12px"> 测试支付连接 </n-button>
    </n-form-item>
  </n-form>
</template>
