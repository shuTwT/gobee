<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FlinkFormProps } from './utils/types'

const props = defineProps<FlinkFormProps>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)
const rules: FormRules = {
  name:[{required:true,message:"请输入名称"}],
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

defineExpose({ getData })
</script>
<template>
  <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="80">
    <n-form-item label="名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入名称" />
    </n-form-item>
  </n-form>
</template>
