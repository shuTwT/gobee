<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import KeyValueEditor from './KeyValueEditor.vue'

const props = defineProps<{
  formInline: any
}>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)

const typeOptions = [
  { label: 'Cron表达式', value: 'cron' },
  { label: '时间间隔', value: 'interval' },
  { label: '一次性任务', value: 'once' },
]

const executionTypeOptions = [
  { label: 'HTTP请求', value: 'http' },
  { label: '内部服务', value: 'internal' },
  { label: '命令执行', value: 'command' },
  { label: '消息队列', value: 'mq' },
]

const httpMethodOptions = [
  { label: 'GET', value: 'GET' },
  { label: 'POST', value: 'POST' },
  { label: 'PUT', value: 'PUT' },
  { label: 'DELETE', value: 'DELETE' },
]

const isHttpExecution = computed(() => formData.value.execution_type === 'http')

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
  execution_type: [
    { required: true, message: '请选择执行类型', trigger: 'blur' }
  ],
  http_url: [
    {
      required: isHttpExecution.value,
      message: '请输入HTTP URL',
      trigger: 'blur',
      validator: (rule, value) => {
        if (isHttpExecution.value && !value) {
          return new Error('请输入HTTP URL')
        }
        return true
      }
    }
  ]
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
            execution_type: data.execution_type,
            http_method: data.http_method,
            http_url: data.http_url,
            http_headers: data.http_headers,
            http_body: data.http_body,
            http_timeout: data.http_timeout,
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
      :label="formData.type === 'cron' ? 'Cron表达式' : formData.type === 'interval' ? '时间间隔' : '执行时间'"
      path="expression"
    >
      <n-input
        v-model:value="formData.expression"
        :placeholder="formData.type === 'cron' ? '例如: 0 0 * * * (每天午夜执行)' : formData.type === 'interval' ? '例如: 5m (每5分钟执行)' : '例如: 2024-01-01 00:00:00'"
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

    <n-form-item label="执行类型" path="execution_type">
      <n-select v-model:value="formData.execution_type" :options="executionTypeOptions" placeholder="请选择执行类型" />
    </n-form-item>

    <template v-if="isHttpExecution">
      <n-form-item label="HTTP方法" path="http_method">
        <n-select v-model:value="formData.http_method" :options="httpMethodOptions" placeholder="请选择HTTP方法" />
      </n-form-item>

      <n-form-item label="HTTP URL" path="http_url">
        <n-input v-model:value="formData.http_url" placeholder="请输入HTTP URL" />
      </n-form-item>

      <n-form-item label="HTTP请求头" path="http_headers">
        <KeyValueEditor v-model="formData.http_headers" />
      </n-form-item>

      <n-form-item label="HTTP请求体" path="http_body">
        <n-input
          v-model:value="formData.http_body"
          type="textarea"
          placeholder="请输入HTTP请求体（JSON格式）"
          :autosize="{ minRows: 3, maxRows: 5 }"
        />
      </n-form-item>

      <n-form-item label="超时时间(秒)" path="http_timeout">
        <n-input-number v-model:value="formData.http_timeout" :min="1" :max="300" placeholder="请输入超时时间" />
      </n-form-item>
    </template>

    <n-form-item label="最大重试次数" path="max_retries">
      <n-input-number v-model:value="formData.max_retries" :min="0" :max="10" placeholder="请输入最大重试次数" />
    </n-form-item>

    <n-form-item label="失败通知" path="failure_notification">
      <n-switch v-model:value="formData.failure_notification" />
    </n-form-item>
  </n-form>
</template>
