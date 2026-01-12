<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import * as couponApi from '@/api/mall/coupon'
import type { FormItemProps,FormProps } from '@/views/mall/coupon/utils/types'

const props = defineProps<FormProps>()
const emit = defineEmits(['confirm'])

const formRef = ref<FormInst | null>(null)
const formData = ref<FormItemProps>(props.formInline)

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入优惠券名称',
    trigger: 'blur',
  },
  code: {
    required: true,
    message: '请输入优惠券代码',
    trigger: 'blur',
  },
  type: {
    required: true,
    type: 'number',
    message: '请选择优惠券类型',
    trigger: 'change',
  },
  value: {
    required: true,
    type: 'number',
    message: '请输入优惠券值',
    trigger: 'blur',
  },
  total_count: {
    required: true,
    type: 'number',
    message: '请输入发放总数',
    trigger: 'blur',
  },
  start_time: {
    required: true,
    type: 'string',
    message: '请选择开始时间',
    trigger: 'change',
  },
  end_time: {
    required: true,
    type: 'string',
    message: '请选择结束时间',
    trigger: 'change',
  },
}

const couponTypeOptions = [
  { label: '满减', value: 0 },
  { label: '折扣', value: 1 },
  { label: '无门槛', value: 2 },
]

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            name: data.name,
            code: data.code,
            type: data.type,
            value: data.value,
            total_count: data.total_count,
            per_user_limit: data.per_user_limit,
            start_time: data.start_time,
            end_time: data.end_time,
            active: data.active,
          }
          
          if (data.description) submitData.description = data.description
          if (data.min_amount) submitData.min_amount = data.min_amount
          if (data.max_discount) submitData.max_discount = data.max_discount
          if (data.image) submitData.image = data.image
          if (data.product_ids && data.product_ids.length > 0) submitData.product_ids = data.product_ids
          if (data.category_ids && data.category_ids.length > 0) submitData.category_ids = data.category_ids
          
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
    <n-form-item label="优惠券名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入优惠券名称" />
    </n-form-item>
    <n-form-item label="优惠券代码" path="code">
      <n-input v-model:value="formData.code" placeholder="请输入优惠券代码" />
    </n-form-item>
    <n-form-item label="优惠券类型" path="type">
      <n-select v-model:value="formData.type" :options="couponTypeOptions" placeholder="请选择优惠券类型" />
    </n-form-item>
    <n-form-item label="优惠券值" path="value">
      <n-input-number
        v-model:value="formData.value"
        :min="0"
        placeholder="请输入优惠券值(分)"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="最低消费金额(分)" path="min_amount">
      <n-input-number
        v-model:value="formData.min_amount"
        :min="0"
        placeholder="请输入最低消费金额"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="最大优惠金额(分)" path="max_discount">
      <n-input-number
        v-model:value="formData.max_discount"
        :min="0"
        placeholder="请输入最大优惠金额"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="发放总数" path="total_count">
      <n-input-number
        v-model:value="formData.total_count"
        :min="0"
        placeholder="请输入发放总数"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="每用户限领数量" path="per_user_limit">
      <n-input-number
        v-model:value="formData.per_user_limit"
        :min="1"
        placeholder="请输入每用户限领数量"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="开始时间" path="start_time">
      <n-date-picker
        v-model:value="formData.start_time"
        type="datetime"
        placeholder="请选择开始时间"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="结束时间" path="end_time">
      <n-date-picker
        v-model:value="formData.end_time"
        type="datetime"
        placeholder="请选择结束时间"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="优惠券描述" path="description">
      <n-input
        v-model:value="formData.description"
        type="textarea"
        placeholder="请输入优惠券描述"
        :rows="3"
      />
    </n-form-item>
    <n-form-item label="优惠券图片" path="image">
      <n-input v-model:value="formData.image" placeholder="请输入优惠券图片URL" />
    </n-form-item>
    <n-form-item label="是否启用" path="active">
      <n-switch v-model:value="formData.active" />
    </n-form-item>
  </n-form>
</template>
