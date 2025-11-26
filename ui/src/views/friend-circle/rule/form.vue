<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui';
import type { FormItemProps, FormProps } from './utils/types';


const props = defineProps<FormProps>()

const formRef = ref<FormInst|null>(null)
const formData = ref<FormItemProps>(props.formInline)
const rules: FormRules = {
  name: [{ required: true, message: '请输入名称' }],
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
    <n-form
    ref="formRef"
    :model="formData"
    :rules="rules"
    label-placement="top"
    label-width="130"
    require-mark-placement="right-hanging"
  >
    <n-form-item label="名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入名称" />
    </n-form-item>
    <n-form-item label="标题选择器" path="title_selector">
      <n-input v-model:value="formData.title_selector" type="textarea" placeholder="请输入标题选择器" />
    </n-form-item>
    <n-form-item label="链接选择器" path="link_selector">
      <n-input v-model:value="formData.link_selector" type="textarea" placeholder="请输入链接选择器" />
    </n-form-item>
    <n-form-item label="创建时间选择器" path="created_selector">
      <n-input v-model:value="formData.created_selector" type="textarea" placeholder="请输入创建时间选择器" />
    </n-form-item>
    <n-form-item label="更新时间选择器" path="updated_selector">
      <n-input v-model:value="formData.updated_selector" type="textarea" placeholder="请输入更新时间选择器" />
    </n-form-item>
  </n-form>
</template>
