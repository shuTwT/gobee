<script lang="ts" setup>
import type { FormInst, FormRules } from 'naive-ui'
import type{ FormItemProps, FormProps } from './utils/types'

const props = defineProps<FormProps>()

const formRef = ref<FormInst | null>(null)
const formData = ref<FormItemProps>(props.formInline)

const rules: FormRules = {
  name: [{ required: true, message: '请输入策略名称' }],
  type: [{ required: true, message: '请选择存储类型' }],
  base_path: [{ required: true, message: '请输入基础路径' }],
  node_id: [{ required: true, message: '请输入节点 ID' }],
  access_key: [{ required: true, message: '请输入Access Key' }],
  secret_key: [{ required: true, message: '请输入Secret Key' }],
  bucket: [{ required: true, message: '请输入Bucket' }],
  region: [{ required: true, message: '请输入Region' }],
  domain: [{ required: true, message: '请输入域名' }],
}

// 存储类型选项
const storageTypes = [
  { label: '本地存储', value: 'local' as const },
  { label: 'S3', value: 's3' as const },
]

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          resolve(toRaw(formData.value))
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
  <n-form
    ref="formRef"
    :model="formData"
    :rules="rules"
    label-placement="left"
    label-width="100"
    require-mark-placement="right-hanging"
  >
    <n-form-item label="策略名称" path="name" >
      <n-input v-model:value="formData.name" placeholder="请输入策略名称" />
    </n-form-item>
    <n-form-item label="存储类型" path="type">
      <n-select
        v-model:value="formData.type"
        :options="storageTypes"
        placeholder="请选择存储类型"
      />
    </n-form-item>
    <n-form-item
      v-if="formData.type === 'local'"
      label="基础路径"
      path="base_path"
    >
      <n-input v-model:value="formData.base_path" placeholder="请输入基础路径" />
    </n-form-item>
    <n-form-item
      v-if="formData.type === 's3'"
      label="节点 ID"
      path="node_id"
    >
      <n-input v-model:value="formData.node_id" placeholder="请输入节点 ID" />
    </n-form-item>

    <template v-if="formData.type === 's3'">
      <n-form-item label="Access Key" path="access_key">
        <n-input v-model:value="formData.access_key" placeholder="请输入Access Key" />
      </n-form-item>
      <n-form-item label="Secret Key" path="secret_key">
        <n-input
          v-model:value="formData.secret_key"
          type="password"
          show-password-on="click"
          placeholder="请输入Secret Key"
        />
      </n-form-item>
    </template>
    <template v-if="formData.type === 's3'">
      <n-form-item label="Bucket" path="bucket">
        <n-input v-model:value="formData.bucket" placeholder="请输入Bucket" />
      </n-form-item>
    </template>
    <n-form-item label="访问域名" path="domain">
      <n-input v-model:value="formData.domain" placeholder="请输入访问域名" />
    </n-form-item>
    <n-form-item label="设为默认">
      <n-switch v-model:value="formData.master" />
    </n-form-item>
  </n-form>
</template>
