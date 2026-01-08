<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import type { FlinkFormProps } from './utils/types'
import * as FriendCircleRuleApi from "@/api/friend-circle/friendCircleRule"
import { apiClient, useApi } from '@/api';

const props = defineProps<FlinkFormProps>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)
const rules: FormRules = {
  name:[{required:true,message:"请输入名称"}],
  url:[{required:true,message:"请输入网站地址"}],
  logo:[{required:true,message:"请输入Logo"}],
  description:[{required:true,message:"请输入网站描述"}],
  group_id:[{required:true,message:"请选择分组"}]
}

const flinkGroups = ref<any[]>([])

const friendCircleRuleOptions = ref([])

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

onMounted(async ()=>{
  try{
    const res1 = await useApi(apiClient.api.v1FlinkGroupListList)
    flinkGroups.value = res1.data.map((item)=>({
      label: item.name,
      value: item.id
    }))
    const res2= await FriendCircleRuleApi.getFriendCircleRuleList()
    friendCircleRuleOptions.value = res2.data.map((item:any)=>({
      value: item.id,
      label: item.name
    }))
  }catch{

  }

})

defineExpose({ getData })
</script>
<template>
  <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="120">
    <n-form-item label="网站名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入网站名称" />
    </n-form-item>
    <n-form-item label="网站地址" path="url">
      <n-input v-model:value="formData.url" placeholder="请输入网站地址" />
    </n-form-item>
    <n-form-item label="头像地址" path="avatar_url">
      <n-input v-model:value="formData.avatar_url" placeholder="请输入头像地址" />
    </n-form-item>
    <n-form-item label="描述" path="description">
      <n-input v-model:value="formData.description" type="textarea" placeholder="请输入网站描述" />
    </n-form-item>
    <n-form-item label="网站封面地址" path="cover_url">
      <n-input v-model:value="formData.cover_url" placeholder="请输入封面地址" />
    </n-form-item>
    <n-form-item label="网站截图地址" path="snapshot_url">
      <n-input v-model:value="formData.snapshot_url" placeholder="请输入网站截图地址" />
    </n-form-item>
    <n-form-item label="开启友链朋友圈" path="enable_friend_circle">
      <n-checkbox v-model:checked="formData.enable_friend_circle"/>
    </n-form-item>
    <n-form-item label="朋友圈解析规则" path="friend_circle_rule_id">
      <n-select v-model:value="formData.friend_circle_rule_id" :options="friendCircleRuleOptions"/>
    </n-form-item>
    <n-form-item label="分组" path="group_id">
      <n-select v-model:value="formData.group_id" :options="flinkGroups"/>
    </n-form-item>
  </n-form>
</template>
