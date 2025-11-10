<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FormProps } from './utils/types';

const props = defineProps<FormProps>()

const formRef = ref<FormInst|null>(null)
const formData = ref(props.formInline)

const rules: FormRules = {}

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
  <n-form :rules="rules" ref="formRef">
    <n-form-item label="邮箱" path="email">
      <n-input v-model:value="formData.email" />
    </n-form-item>
    <n-form-item label="用户名" path="username">
      <n-input v-model:value="formData.username" />
    </n-form-item>
    <n-form-item label="手机号" path="phone_number">
      <n-input v-model:value="formData.phone_number" />
    </n-form-item>
    <n-form-item label="密码" path="password">
      <n-input v-model:value="formData.password" />
    </n-form-item>
  </n-form>
</template>
