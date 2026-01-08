<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui';
import type { FormItemProps } from './utils/types'

const props = defineProps<{
  formInline: FormItemProps
}>()

const formRef = ref<FormInst|null>()
const formData = ref(props.formInline)
const rules:FormRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
}

const categoryOptions = ref([])

const tagOptions = ref([])

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
  <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100px">
    <n-form-item label="标题" path="title">
      <n-input v-model:value="formData.title"/>
    </n-form-item>
    <n-form-item label="别名" path="alias">
      <n-input v-model:value="formData.alias"/>
    </n-form-item>
    <n-form-item label="分类" path="categorys">
      <n-select v-model:value="formData.categorys" :multiple="true" :filterable="true" :options="categoryOptions"/>
    </n-form-item>
    <n-form-item label="标签" path="tags">
      <n-select v-model:value="formData.tags" :multiple="true" :filterable="true" :options="tagOptions"/>
    </n-form-item>
    <n-form-item label="自动生成摘要" path="is_autogen_summary">
      <n-checkbox v-model:checked="formData.is_autogen_summary"/>
    </n-form-item>
    <n-form-item label="作者" path="author">
      <n-input v-model:value="formData.author"/>
    </n-form-item>
    <n-form-item label="允许评论" path="is_allow_comments">
      <n-checkbox v-model:checked="formData.is_allow_comments"/>
    </n-form-item>
    <n-form-item label="是否置顶" path="is_pin_to_top">
      <n-checkbox v-model:checked="formData.is_pin_to_top"/>
    </n-form-item>
    <n-form-item label="可见性" path="is_visible">
      <n-checkbox v-model:checked="formData.is_visible"/>
    </n-form-item>
    <n-form-item label="封面" path="cover">
      <n-input v-model:value="formData.cover"/>
    </n-form-item>
  </n-form>
</template>
