<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'

const props = defineProps<{ nickname?: string; bio?: string }>()

const formRef = ref<FormInst>()
const formData = ref({
  nickname: props.nickname ?? '',
  bio: props.bio ?? '',
})

const rules = ref<FormRules>({
  nickname: [{ max: 50, message: '昵称最多 50 个字符', trigger: 'input' }],
  bio: [{ max: 255, message: '个人简介最多 255 个字符', trigger: 'input' }],
})

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          resolve(toRaw(formData.value))
        } else {
          reject(new Error('表单校验失败'))
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
    <n-form-item label="昵称" path="nickname">
      <n-input v-model:value="formData.nickname" placeholder="请输入昵称" clearable />
    </n-form-item>
    <n-form-item label="个人简介" path="bio">
      <n-input
        v-model:value="formData.bio"
        type="textarea"
        placeholder="介绍一下自己"
        :autosize="{ minRows: 3, maxRows: 6 }"
        clearable
      />
    </n-form-item>
  </n-form>
</template>
