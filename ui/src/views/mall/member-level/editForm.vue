<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FormProps } from './utils/types'

const props = defineProps<FormProps>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入等级名称',
    trigger: 'blur',
  },
  level: {
    required: true,
    type: 'number',
    message: '请输入等级级别',
    trigger: 'blur',
  },
  min_points: {
    required: true,
    type: 'number',
    message: '请输入最低积分',
    trigger: 'blur',
  },
  discount_rate: {
    required: true,
    type: 'number',
    message: '请输入折扣率',
    trigger: 'blur',
  },
}

const getData = () => {
  return new Promise((resolve, reject) => {
    console.log(formData.value)
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            name: data.name,
            level: data.level,
            min_points: data.min_points,
            discount_rate: data.discount_rate,
            active: data.active,
            sort_order: data.sort_order,
          }
          
          if (data.description) submitData.description = data.description
          if (data.icon) submitData.icon = data.icon
          if (data.privileges && data.privileges.length > 0) submitData.privileges = data.privileges
          
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
  <n-form :model="formData" :rules="rules" ref="formRef" label-placement="left" label-width="120px">
    <n-form-item label="等级名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入等级名称" />
    </n-form-item>
    <n-form-item label="等级级别" path="level">
      <n-input-number
        v-model:value="formData.level"
        :min="1"
        placeholder="请输入等级级别"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="最低积分" path="min_points">
      <n-input-number
        v-model:value="formData.min_points"
        :min="0"
        placeholder="请输入最低积分"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="折扣率(%)" path="discount_rate">
      <n-input-number
        v-model:value="formData.discount_rate"
        :min="0"
        :max="100"
        placeholder="请输入折扣率"
        style="width: 100%"
      />
      <template #feedback>
        <span style="font-size: 12px; color: #999;">100表示原价，小于100表示折扣</span>
      </template>
    </n-form-item>
    <n-form-item label="描述" path="description">
      <n-input
        v-model:value="formData.description"
        type="textarea"
        placeholder="请输入描述"
        :rows="4"
      />
    </n-form-item>
    <n-form-item label="会员特权" path="privileges">
      <n-dynamic-tags v-model:value="formData.privileges" />
    </n-form-item>
    <n-form-item label="图标" path="icon">
      <n-input v-model:value="formData.icon" placeholder="请输入图标URL" />
    </n-form-item>
    <n-form-item label="是否启用" path="active">
      <n-switch v-model:value="formData.active" />
    </n-form-item>
    <n-form-item label="排序" path="sort_order">
      <n-input-number
        v-model:value="formData.sort_order"
        :min="0"
        placeholder="请输入排序"
        style="width: 100%"
      />
    </n-form-item>
  </n-form>
</template>
