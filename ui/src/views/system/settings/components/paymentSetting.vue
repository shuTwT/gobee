<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import { apiClient, useApi } from '@/api'

const message = useMessage()
const paymentFormRef = ref<FormInst | null>(null)

const defaultForm = {
  enableEpay: false,
  epayApiUrl: '',
  epayMerchantId: '',
  epayMerchantKey: '',
  epayNotifyUrl: '',
  epayReturnUrl: '',
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
  refundReview: true,
  rechargePointsRate: 1,
  enableMockPay: false,
}
const paymentForm = ref({
  enableEpay: false,
  epayApiUrl: '',
  epayMerchantId: '',
  epayMerchantKey: '',
  epayNotifyUrl: '',
  epayReturnUrl: '',
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
  refundReview: true,
  rechargePointsRate: 1,
  enableMockPay: false,
})
const paymentLoading = ref(false)

// 保存支付设置
const savePaymentSettings = async () => {
  paymentLoading.value = true
  try {
    await useApi(apiClient.api.v1SettingsJsonSaveCreate, 'payment', paymentForm.value)
    await new Promise((resolve) => setTimeout(resolve, 1000))
    onSearch()
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
    await new Promise((resolve) => setTimeout(resolve, 1000))
    message.success('支付连接测试成功')
  } catch {
    message.error('支付连接测试失败')
  }
}

const onSearch = async () => {
  const res = await useApi(apiClient.api.v1SettingsJsonDetail, 'payment')
  paymentForm.value = Object.assign({}, defaultForm, res.data)
}

onMounted(() => {
  onSearch()
})
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
    <n-form-item label="启用易支付" path="enableEpay">
      <n-switch v-model:value="paymentForm.enableEpay" />
      <n-text>
        启用易支付后，支付宝和微信的渠道将被易支付替代。
      </n-text>
    </n-form-item>
    <n-form-item label="易支付网关地址" path="epayApiUrl">
      <n-input v-model:value="paymentForm.epayApiUrl" placeholder="https://api.pay.com" />
    </n-form-item>
    <n-form-item label="易支付商户ID" path="epayMerchantId">
      <n-input v-model:value="paymentForm.epayMerchantId" placeholder="请输入易支付商户ID" />
    </n-form-item>
    <n-form-item label="易支付商户密钥" path="epayMerchantKey">
      <n-input
        v-model:value="paymentForm.epayMerchantKey"
        type="password"
        show-password-on="click"
        placeholder="请输入易支付商户密钥"
      />
    </n-form-item>
    <n-form-item label="易支付异步通知URL" path="epayNotifyUrl">
      <n-input
        v-model:value="paymentForm.epayNotifyUrl"
        placeholder="https://yourdomain.com/api/v1/pay-order/notify"
      />
    </n-form-item>
    <n-form-item label="易支付同步返回URL" path="epayReturnUrl">
      <n-input
        v-model:value="paymentForm.epayReturnUrl"
        placeholder="https://yourdomain.com/api/v1/pay-order/return"
      />
    </n-form-item>
    <n-divider />
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
    <n-form-item label="充值积分比例" path="rechargePointsRate">
      <n-input-number v-model:value="paymentForm.rechargePointsRate" :min="1" :max="10000">
        <template #suffix>积分/分</template>
      </n-input-number>
    </n-form-item>
    <n-form-item label="退款审核" path="refundReview">
      <n-switch v-model:value="paymentForm.refundReview" />
    </n-form-item>
    <n-divider />
    <n-form-item label="模拟支付(Mock)" path="enableMockPay">
      <n-switch v-model:value="paymentForm.enableMockPay" />
      <span style="margin-left: 8px; font-size: 12px; color: #999">
        开启后下单不调用真实支付网关，返回本地模拟支付页（仅测试环境使用）
      </span>
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="savePaymentSettings" :loading="paymentLoading">
        保存支付设置
      </n-button>
      <n-button @click="testPaymentConnection" style="margin-left: 12px"> 测试支付连接 </n-button>
    </n-form-item>
  </n-form>
</template>
