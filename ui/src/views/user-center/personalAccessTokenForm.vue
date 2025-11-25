<script setup lang="ts">
import dayjs from 'dayjs';
import type { FormInst, FormRules } from 'naive-ui';

const formRef = ref<FormInst>()
const formData = ref({
  name:"",
  expires:0,
  description:""
})

const rules = ref<FormRules>({
  name:[{required:true,message:"请输入令牌名称"}],
  expires:[{required:true,message:"过期时间"}]
})

onMounted(()=>{
  formData.value.expires= dayjs().add(1,'day').valueOf()
})

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          resolve(toRaw(formData.value))
        } else {
          reject(new Error("表单校验失败"))
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
    <n-form-item label="令牌名称" path="name">
      <n-input v-model:value="formData.name"/>
    </n-form-item>
    <n-form-item label="过期时间" path="expires">
      <n-date-picker v-model:value="formData.expires" type="datetime" clearable/>
    </n-form-item>
    <n-form-item label="介绍" path="description">
      <n-input v-model:value="formData.description" type="textarea"/>
    </n-form-item>
  </n-form>
</template>
