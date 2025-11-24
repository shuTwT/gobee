<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui';
import type { AlbumFormProps } from './utils/types';

const props = defineProps<AlbumFormProps>()
const formRef = ref<FormInst|null>()
const formData = ref(props.formInline)
const rules:FormRules = {
  name: [{ required: true, message: '请输入相册名称', trigger: 'blur' }],
}

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

defineExpose({getData})
</script>
<template>
  <n-form ref="formRef" :rules="rules" :model="formData">
    <n-form-item label="相册名称" path="name">
      <n-input v-model:value="formData.name" />
    </n-form-item>
    <n-form-item label="相册描述" path="description">
      <n-input v-model:value="formData.description" />
    </n-form-item>
  </n-form>
</template>
