<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'

const props = defineProps<{
  formInline: any
}>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)

const typeOptions = [
  { label: 'Cron表达式', value: 'cron' },
  { label: '时间间隔', value: 'interval' },
]

const rules: FormRules = {
  name: [
    { required: true, message: '请输入任务名称', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请选择任务类型', trigger: 'blur' }
  ],
  expression: [
    { required: true, message: '请输入调度表达式', trigger: 'blur' }
  ],
  job_name: [
    { required: true, message: '请选择内部任务', trigger: 'blur' }
  ],
}

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            name: data.name,
            type: data.type,
            expression: data.expression,
            description: data.description,
            enabled: data.enabled,
            job_name: data.job_name,
            max_retries: data.max_retries,
            failure_notification: data.failure_notification,
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

defineExpose({
  getData,
})
</script>

<template>
  <n-form :model="formData" :rules="rules" ref="formRef" label-placement="left" label-width="120px">
    <n-form-item label="任务名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入任务名称" />
    </n-form-item>

    <n-form-item label="任务类型" path="type">
      <n-select v-model:value="formData.type" :options="typeOptions" placeholder="请选择任务类型" />
    </n-form-item>

    <n-form-item
      :label="formData.type === 'cron' ? 'Cron表达式' : '时间间隔'"
      path="expression"
    >
      <n-input
        v-model:value="formData.expression"
        :placeholder="formData.type === 'cron' ? '例如: 0 0 * * * (每天午夜执行)' : '例如: 24h (每24小时执行)'"
      />
    </n-form-item>

    <n-form-item label="任务描述" path="description">
      <n-input
        v-model:value="formData.description"
        type="textarea"
        placeholder="请输入任务描述"
        :autosize="{ minRows: 3, maxRows: 5 }"
      />
    </n-form-item>

    <n-form-item label="是否启用" path="enabled">
      <n-switch v-model:value="formData.enabled" />
    </n-form-item>

    <n-form-item label="内部任务" path="job_name">
      <n-input v-model:value="formData.job_name" placeholder="请输入内部任务名称，例如: friendCircle" />
    </n-form-item>

    <n-form-item label="最大重试次数" path="max_retries">
      <n-input-number v-model:value="formData.max_retries" :min="0" :max="10" placeholder="请输入最大重试次数" />
    </n-form-item>

    <n-form-item label="失败通知" path="failure_notification">
      <n-switch v-model:value="formData.failure_notification" />
    </n-form-item>
  </n-form>
</template>
