<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import * as couponUsageApi from '@/api/mall/couponUsage'

interface Props {
  couponUsageId?: number
  couponUsageData: any
}

const props = defineProps<Props>()
const emit = defineEmits(['confirm'])

const formRef = ref<FormInst | null>(null)
const formData = ref({
  coupon_code: props.couponUsageData.coupon_code || '',
  user_id: props.couponUsageData.user_id || 0,
  order_id: props.couponUsageData.order_id || null,
  status: props.couponUsageData.status !== undefined ? props.couponUsageData.status : 0,
  used_at: props.couponUsageData.used_at || '',
  discount_amount: props.couponUsageData.discount_amount || 0,
  expire_at: props.couponUsageData.expire_at || '',
  remark: props.couponUsageData.remark || '',
})

const rules: FormRules = {
  coupon_code: {
    required: true,
    message: '请输入优惠券代码',
    trigger: 'blur',
  },
  user_id: {
    required: true,
    type: 'number',
    message: '请输入用户ID',
    trigger: 'blur',
  },
  discount_amount: {
    required: true,
    type: 'number',
    message: '请输入优惠金额',
    trigger: 'blur',
  },
  expire_at: {
    required: true,
    type: 'string',
    message: '请选择过期时间',
    trigger: 'change',
  },
}

const statusOptions = [
  { label: '未使用', value: 0 },
  { label: '已使用', value: 1 },
  { label: '已过期', value: 2 },
]

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            coupon_code: data.coupon_code,
            user_id: data.user_id,
            discount_amount: data.discount_amount,
            expire_at: data.expire_at,
          }
          
          if (data.order_id) submitData.order_id = data.order_id
          if (data.status !== undefined) submitData.status = data.status
          if (data.used_at) submitData.used_at = data.used_at
          if (data.remark) submitData.remark = data.remark
          
          resolve(submitData)
        } else {
          reject(errors)
        }
      })
    } else {
      reject(new Error('表单实例不存在'))
    }
  })
}

defineExpose({ getData })
</script>

<template>
  <n-form :rules="rules" ref="formRef" label-placement="left" label-width="120px">
    <n-form-item label="优惠券代码" path="coupon_code">
      <n-input v-model:value="formData.coupon_code" placeholder="请输入优惠券代码" />
    </n-form-item>
    <n-form-item label="用户ID" path="user_id">
      <n-input-number
        v-model:value="formData.user_id"
        :min="1"
        placeholder="请输入用户ID"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="订单ID" path="order_id">
      <n-input-number
        v-model:value="formData.order_id"
        :min="1"
        placeholder="请输入订单ID"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="状态" path="status">
      <n-select v-model:value="formData.status" :options="statusOptions" placeholder="请选择状态" />
    </n-form-item>
    <n-form-item label="使用时间" path="used_at">
      <n-date-picker
        v-model:value="formData.used_at"
        type="datetime"
        placeholder="请选择使用时间"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="优惠金额(分)" path="discount_amount">
      <n-input-number
        v-model:value="formData.discount_amount"
        :min="0"
        placeholder="请输入优惠金额"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="过期时间" path="expire_at">
      <n-date-picker
        v-model:value="formData.expire_at"
        type="datetime"
        placeholder="请选择过期时间"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="备注" path="remark">
      <n-input
        v-model:value="formData.remark"
        type="textarea"
        placeholder="请输入备注"
        :rows="3"
      />
    </n-form-item>
  </n-form>
</template>
