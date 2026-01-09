<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import * as memberlevelService from '@/api/member/memberLevel'
import type {FormProps,FormItemProps} from '@/views/member/member/utils/types'



const props = defineProps<FormProps>()
const emit = defineEmits(['confirm'])

const formRef = ref<FormInst | null>(null)
const formData = ref<FormItemProps>(props.formInline)

const memberLevels = ref<any[]>([])

const rules: FormRules = {
  member_level: {
    required: true,
    type: 'number',
    message: '请选择会员等级',
    trigger: 'change',
  },
  balance: {
    required: true,
    type: 'number',
    message: '请输入余额',
    trigger: 'blur',
  },
  discount_rate: {
    required: true,
    type: 'number',
    message: '请输入倍率',
    trigger: 'blur',
  },
}

const getMemberLevels = async () => {
  try {
    const data = await memberlevelService.getMemberLevelList()
    if (data.code === 200) {
      memberLevels.value = data.data || []
    }
  } catch (error) {
    console.error('获取会员等级失败:', error)
  }
}

const getData = (): Promise<typeof formData.value> => {
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

const handleSubmit = async () => {
  try {
    const data = await getData()
    const updateData = {
      member_level: data.member_level,
    }
    
    await fetch(`/api/v1/member/update/${props.formInline.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(updateData),
    })
    
    emit('confirm')
  } catch (error) {
    console.error('提交失败:', error)
  }
}

defineExpose({ getData, handleSubmit })

onMounted(() => {
  getMemberLevels()
})
</script>

<template>
  <n-form :model="formData" :rules="rules" ref="formRef" label-placement="left" label-width="100px">
    <n-form-item label="会员等级" path="member_level">
      <n-select
        v-model:value="formData.member_level"
        :options="memberLevels.map(level => ({
          label: level.name,
          value: level.id,
        }))"
        placeholder="请选择会员等级"
      />
    </n-form-item>
    <n-form-item label="余额(元)" path="balance">
      <n-input-number
        v-model:value="formData.balance"
        :min="0"
        :precision="2"
        placeholder="请输入余额"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="倍率(%)" path="discount_rate">
      <n-input-number
        v-model:value="formData.discount_rate"
        :min="0"
        :max="200"
        :precision="0"
        placeholder="请输入倍率"
        style="width: 100%"
      />
      <template #feedback>
        <span style="font-size: 12px; color: #999;">100表示原价，大于100表示加价，小于100表示折扣</span>
      </template>
    </n-form-item>
  </n-form>
</template>
