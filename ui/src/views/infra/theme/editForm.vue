<template>
  <div style="padding: 20px">
    <n-form :model="formData" label-placement="left" label-width="100px" ref="formRef">
      <n-form-item label="主题类型" path="type">
        <n-radio-group v-model:value="formData.type">
          <n-radio value="internal">内部主题</n-radio>
          <n-radio value="external">外部主题</n-radio>
        </n-radio-group>
      </n-form-item>
      
      <template v-if="formData.type === 'internal'">
        <n-form-item label="上传文件" path="file">
          <n-upload
            accept=".zip"
            :max="1"
            :custom-request="handleUploadFile"
            :show-file-list="true"
            :file-list="fileList"
          >
            <n-button>选择文件</n-button>
          </n-upload>
        </n-form-item>
      </template>
      
      <template v-if="formData.type === 'external'">
        <n-form-item label="主题名称" path="name">
          <n-input v-model:value="formData.name" placeholder="请输入主题名称" />
        </n-form-item>
        <n-form-item label="显示名称" path="display_name">
          <n-input v-model:value="formData.display_name" placeholder="请输入显示名称" />
        </n-form-item>
        <n-form-item label="版本号" path="version">
          <n-input v-model:value="formData.version" placeholder="请输入版本号" />
        </n-form-item>
        <n-form-item label="外部URL" path="external_url">
          <n-input v-model:value="formData.external_url" placeholder="请输入外部主题URL地址" />
        </n-form-item>
        <n-form-item label="描述" path="description">
          <n-input
            v-model:value="formData.description"
            type="textarea"
            placeholder="请输入描述"
            :rows="4"
          />
        </n-form-item>
      </template>
    </n-form>
  </div>
</template>

<script setup lang="ts">
import type { FormInst } from 'naive-ui'
import { NUpload, NButton, NForm, NFormItem, NInput, NRadioGroup, NRadio, type UploadCustomRequestOptions } from 'naive-ui'
import * as themeApi from '@/api/infra/theme'

const props = defineProps({
  formInline: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['success', 'error'])

const formRef = ref<FormInst | null>(null)
const formData = ref(props.formInline)
const uploadedFilePath = ref('')
const fileList = ref<any[]>([])

const handleUploadFile = ({ file, onFinish, onError }: UploadCustomRequestOptions) => {
  const formDataUpload = new FormData()
  formDataUpload.append('file', file.file as File)  
  themeApi.uploadThemeFile(formDataUpload).then(res => {
    window.$message?.success('文件上传成功')
    uploadedFilePath.value = res.data.file_path
    fileList.value = [{
      id: Date.now(),
      name: file.name,
      status: 'finished'
    }]
    onFinish()
  }).catch(err => {
    window.$message?.error('文件上传失败')
    onError()
    emit('error')
  })
}

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          if (formData.value.type === 'internal') {
            if (!uploadedFilePath.value) {
              reject(new Error('请先上传主题文件'))
              return
            }
            
            resolve({
              type: 'internal',
              file_path: uploadedFilePath.value
            })
          } else if (formData.value.type === 'external') {
            if (!formData.value.name) {
              reject(new Error('请输入主题名称'))
              return
            }
            if (!formData.value.display_name) {
              reject(new Error('请输入显示名称'))
              return
            }
            if (!formData.value.version) {
              reject(new Error('请输入版本号'))
              return
            }
            if (!formData.value.external_url) {
              reject(new Error('请输入外部主题URL地址'))
              return
            }
            
            resolve({
              type: 'external',
              name: formData.value.name,
              display_name: formData.value.display_name,
              version: formData.value.version,
              external_url: formData.value.external_url,
              description: formData.value.description
            })
          }
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
