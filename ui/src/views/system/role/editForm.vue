<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FormProps } from './utils/types';

const props = defineProps<FormProps>()

const formRef = ref<FormInst|null>(null)
const formData = ref(props.formInline)

const rules: FormRules = {
  name: [
    { required: true, message: '请输入角色名称', trigger: 'blur' }
  ],
  code: [
    { required: true, message: '请输入角色代码', trigger: 'blur' },
    { pattern: /^[a-zA-Z0-9_-]+$/, message: '角色代码只能包含字母、数字、下划线和横线', trigger: ['blur', 'input'] }
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
            code: data.code,
            description: data.description
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
    <n-form-item label="角色名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入角色名称" />
    </n-form-item>
    <n-form-item label="角色代码" path="code">
      <n-input v-model:value="formData.code" placeholder="请输入角色代码（如：admin、user）" :disabled="!!formData.id" />
    </n-form-item>
    <n-form-item label="角色描述" path="description">
      <n-input
        v-model:value="formData.description"
        type="textarea"
        placeholder="请输入角色描述"
        :rows="3"
      />
    </n-form-item>
  </n-form>
</template>
