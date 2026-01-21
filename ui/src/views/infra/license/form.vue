<script setup lang="ts">
import { ref, computed } from 'vue'
import { NForm, NFormItem, NInput, NDatePicker, NSelect, NButton, NSpace } from 'naive-ui'
import { createLicense, updateLicense, type LicenseCreateReq, type LicenseUpdateReq } from '@/api/infra/license'

const props = defineProps<{
  license?: any
}>()

const emit = defineEmits(['success'])

const formRef = ref()
const loading = ref(false)

const formData = ref<LicenseCreateReq & LicenseUpdateReq>({
  domain: '',
  license_key: '',
  customer_name: '',
  expire_date: '',
  status: 1,
})

const isEdit = computed(() => !!props.license?.id)

const statusOptions = [
  { label: '有效', value: 1 },
  { label: '过期', value: 2 },
  { label: '禁用', value: 3 },
]

const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    loading.value = true

    if (isEdit.value) {
      await updateLicense(props.license.id, formData.value)
      window.$message?.success('授权更新成功')
    } else {
      await createLicense(formData.value)
      window.$message?.success('授权创建成功')
    }
    emit('success')
  } catch (error: any) {
    if (error?.errors) {
      window.$message?.error('请填写完整信息')
    } else {
      window.$message?.error((isEdit.value ? '授权更新' : '授权创建') + '失败：' + error.message)
    }
  } finally {
    loading.value = false
  }
}

const initForm = () => {
  if (props.license) {
    formData.value = {
      domain: props.license.domain || '',
      license_key: props.license.license_key || '',
      customer_name: props.license.customer_name || '',
      expire_date: props.license.expire_date || '',
      status: props.license.status || 1,
    }
  }
}

initForm()
</script>

<template>
  <div class="license-form">
    <NForm ref="formRef" :model="formData" label-placement="left" label-width="100px">
      <NFormItem label="域名" path="domain" :rule="{ required: true, message: '请输入域名' }">
        <NInput v-model:value="formData.domain" placeholder="请输入域名" />
      </NFormItem>

      <NFormItem label="授权密钥" path="license_key" :rule="{ required: true, message: '请输入授权密钥' }">
        <NInput v-model:value="formData.license_key" placeholder="请输入授权密钥" />
      </NFormItem>

      <NFormItem label="客户名称" path="customer_name">
        <NInput v-model:value="formData.customer_name" placeholder="请输入客户名称" />
      </NFormItem>

      <NFormItem label="过期日期" path="expire_date" :rule="{ required: true, message: '请选择过期日期' }">
        <NDatePicker
          v-model:value="formData.expire_date"
          type="datetime"
          placeholder="请选择过期日期"
          format="yyyy-MM-dd HH:mm:ss"
          value-format="yyyy-MM-dd HH:mm:ss"
          style="width: 100%"
        />
      </NFormItem>

      <NFormItem v-if="isEdit" label="状态" path="status">
        <NSelect
          v-model:value="formData.status"
          :options="statusOptions"
          placeholder="请选择状态"
        />
      </NFormItem>

      <NFormItem>
        <NSpace>
          <NButton type="primary" :loading="loading" @click="handleSubmit">
            {{ isEdit ? '更新' : '创建' }}
          </NButton>
        </NSpace>
      </NFormItem>
    </NForm>
  </div>
</template>

<style scoped>
.license-form {
  max-width: 600px;
  margin: 0 auto;
}
</style>
