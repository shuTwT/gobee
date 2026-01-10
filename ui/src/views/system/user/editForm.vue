<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FormProps } from './utils/types';
import * as roleApi from '@/api/system/role'

const props = defineProps<FormProps>()

const formRef = ref<FormInst|null>(null)
const formData = ref(props.formInline)
const roleOptions = ref<any[]>([])

const rules: FormRules = {
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入有效的邮箱地址', trigger: ['blur', 'input'] }
  ],
  name: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { 
      required: !formData.value.id, 
      message: '请输入密码', 
      trigger: 'blur',
      validator: (rule, value) => {
        if (!formData.value.id && !value) {
          return new Error('请输入密码')
        }
        if (value && value.length < 8) {
          return new Error('密码长度不能少于8位')
        }
        return true
      }
    }
  ]
}

const loadRoles = () => {
  roleApi.getRoleList().then(res => {
    if (res.code === 200) {
      roleOptions.value = res.data.map((role: any) => ({
        label: role.name,
        value: role.id
      }))
    }
  })
}

loadRoles()

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            email: data.email,
            name: data.name,
            phone_number: data.phone_number,
            role_id: data.role_id
          }
          if (data.password) {
            submitData.password = data.password
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
    <n-form-item label="邮箱" path="email">
      <n-input v-model:value="formData.email" placeholder="请输入邮箱" />
    </n-form-item>
    <n-form-item label="用户名" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入用户名" />
    </n-form-item>
    <n-form-item label="手机号" path="phone_number">
      <n-input v-model:value="formData.phone_number" placeholder="请输入手机号" />
    </n-form-item>
    <n-form-item label="角色" path="role_id">
      <n-select
        v-model:value="formData.role_id"
        :options="roleOptions"
        placeholder="请选择角色"
        clearable
      />
    </n-form-item>
    <n-form-item label="密码" path="password">
      <n-input
        v-model:value="formData.password"
        type="password"
        show-password-on="click"
        :placeholder="formData.id ? '留空则不修改密码' : '请输入密码（至少8位）'"
      />
    </n-form-item>
  </n-form>
</template>
