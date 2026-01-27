<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FormProps } from './utils/types'

const props = defineProps<FormProps>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入标签名称',
    trigger: 'blur',
  },
  sort_order: {
    required: true,
    type: 'number',
    message: '请输入排序',
    trigger: 'blur',
  },
}

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            name: data.name,
            sort_order: data.sort_order,
            active: data.active,
          }
          
          if (data.slug) submitData.slug = data.slug
          if (data.description) submitData.description = data.description
          if (data.color) submitData.color = data.color
          
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
    <n-form-item label="标签名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入标签名称" />
    </n-form-item>
    <n-form-item label="标签别名" path="slug">
      <n-input v-model:value="formData.slug" placeholder="请输入标签别名" />
    </n-form-item>
    <n-form-item label="描述" path="description">
      <n-input
        v-model:value="formData.description"
        type="textarea"
        placeholder="请输入描述"
        :rows="4"
      />
    </n-form-item>
    <n-form-item label="标签颜色" path="color">
      <n-color-picker v-model:value="formData.color" />
    </n-form-item>
    <n-form-item label="排序" path="sort_order">
      <n-input-number
        v-model:value="formData.sort_order"
        :min="0"
        placeholder="请输入排序"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="是否启用" path="active">
      <n-switch v-model:value="formData.active" />
    </n-form-item>
  </n-form>
</template>
