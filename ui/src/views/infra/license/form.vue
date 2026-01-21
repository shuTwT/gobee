<script setup lang="ts">
import { toRaw, computed } from 'vue'
import dayjs from 'dayjs'
import type { FormInst, FormRules } from 'naive-ui'
import type { FormProps } from './utils/types'
import { NForm, NFormItem, NInput, NDatePicker, NSelect, NButton, NSpace } from 'naive-ui'

const props = defineProps<FormProps>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)

const isEdit = computed(() => !!props.formInline.id)

const statusOptions = [
  { label: '有效', value: 1 },
  { label: '过期', value: 2 },
  { label: '禁用', value: 3 },
]

const rules: FormRules = {
  domain: [
    { required: true, message: '请输入域名', trigger: 'blur' },
    { type: 'url', message: '请输入有效的域名地址', trigger: ['blur', 'input'] }
  ],
  customer_name: [
    { required: true, message: '请输入客户名称', trigger: 'blur' }
  ],
  expire_date: [
    { required: true, message: '请选择过期日期', trigger: 'blur', type: 'number' }
  ],
  status: [
    { required: true, message: '请选择状态', trigger: 'blur', type: 'number' }
  ]
}

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            domain: data.domain,
            customer_name: data.customer_name,
            expire_date: data.expire_date,
          }
          if (isEdit.value) {
            submitData.status = data.status
          }
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
  <n-form :model="formData" :rules="rules" ref="formRef">
    <n-form-item label="域名" path="domain">
      <n-input v-model:value="formData.domain" placeholder="请输入域名" />
    </n-form-item>

    <n-form-item label="客户名称" path="customer_name">
      <n-input v-model:value="formData.customer_name" placeholder="请输入客户名称" />
    </n-form-item>

    <n-form-item label="过期日期" path="expire_date">
      <n-date-picker
        v-model:value="formData.expire_date"
        type="datetime"
        placeholder="请选择过期日期"
        format="yyyy-MM-dd HH:mm:ss"
        value-format="timestamp"
        style="width: 100%"
      />
    </n-form-item>

    <n-form-item v-if="isEdit" label="状态" path="status">
      <n-select
        v-model:value="formData.status"
        :options="statusOptions"
        placeholder="请选择状态"
      />
    </n-form-item>
  </n-form>
</template>
