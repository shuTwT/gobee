<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FormProps } from './utils/types'

const props = defineProps<FormProps>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入主题名称',
    trigger: 'blur',
  },
  display_name: {
    required: true,
    message: '请输入显示名称',
    trigger: 'blur',
  },
  version: {
    required: true,
    message: '请输入版本号',
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
            display_name: data.display_name,
            version: data.version,
          }
          
          if (data.description) submitData.description = data.description
          if (data.author_name) submitData.author_name = data.author_name
          if (data.author_email) submitData.author_email = data.author_email
          if (data.logo) submitData.logo = data.logo
          if (data.homepage) submitData.homepage = data.homepage
          if (data.repo) submitData.repo = data.repo
          if (data.issue) submitData.issue = data.issue
          if (data.setting_name) submitData.setting_name = data.setting_name
          if (data.config_map_name) submitData.config_map_name = data.config_map_name
          if (data.require) submitData.require = data.require
          if (data.license) submitData.license = data.license
          
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
    <n-form-item label="主题名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入主题名称" disabled />
    </n-form-item>
    <n-form-item label="显示名称" path="display_name">
      <n-input v-model:value="formData.display_name" placeholder="请输入显示名称" disabled />
    </n-form-item>
    <n-form-item label="版本号" path="version">
      <n-input v-model:value="formData.version" placeholder="请输入版本号" disabled />
    </n-form-item>
    <n-form-item label="描述" path="description">
      <n-input
        v-model:value="formData.description"
        type="textarea"
        placeholder="请输入描述"
        :rows="4"
        disabled
      />
    </n-form-item>
    <n-form-item label="作者姓名" path="author_name">
      <n-input v-model:value="formData.author_name" placeholder="请输入作者姓名" disabled />
    </n-form-item>
    <n-form-item label="作者邮箱" path="author_email">
      <n-input v-model:value="formData.author_email" placeholder="请输入作者邮箱" disabled />
    </n-form-item>
    <n-form-item label="Logo" path="logo">
      <n-input v-model:value="formData.logo" placeholder="请输入Logo地址" disabled />
    </n-form-item>
    <n-form-item label="主页" path="homepage">
      <n-input v-model:value="formData.homepage" placeholder="请输入主页地址" disabled />
    </n-form-item>
    <n-form-item label="仓库地址" path="repo">
      <n-input v-model:value="formData.repo" placeholder="请输入仓库地址" disabled />
    </n-form-item>
    <n-form-item label="问题反馈" path="issue">
      <n-input v-model:value="formData.issue" placeholder="请输入问题反馈地址" disabled />
    </n-form-item>
    <n-form-item label="设置名称" path="setting_name">
      <n-input v-model:value="formData.setting_name" placeholder="请输入设置名称" disabled />
    </n-form-item>
    <n-form-item label="配置映射" path="config_map_name">
      <n-input v-model:value="formData.config_map_name" placeholder="请输入配置映射名称" disabled />
    </n-form-item>
    <n-form-item label="要求版本" path="require">
      <n-input v-model:value="formData.require" placeholder="请输入要求的gobee版本" disabled />
    </n-form-item>
    <n-form-item label="许可证" path="license">
      <n-input v-model:value="formData.license" placeholder="请输入许可证" disabled />
    </n-form-item>
    <n-form-item label="文件路径" path="path">
      <n-input v-model:value="formData.path" placeholder="文件路径" disabled />
    </n-form-item>
    <n-form-item label="是否启用" path="enabled">
      <n-switch v-model:value="formData.enabled" disabled />
    </n-form-item>
  </n-form>
</template>
