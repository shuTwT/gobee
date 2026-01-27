<template>
  <div style="padding: 20px">
    <n-upload
      accept=".zip"
      :max="1"
      :custom-request="customRequest"
    >
      <n-button>上传文件</n-button>
    </n-upload>
  </div>
</template>

<script setup lang="ts">
import { NUpload, NButton, type UploadCustomRequestOptions } from 'naive-ui'
import * as themeApi from '@/api/infra/theme'

const emit = defineEmits(['success', 'error'])

const customRequest = ({ file, onFinish, onError }: UploadCustomRequestOptions) => {
  const formData = new FormData()
  formData.append('file', file.file as File)
  
  themeApi.uploadTheme(formData).then(res => {
    window.$message?.success('上传成功')
    onFinish()
    emit('success')
  }).catch(err => {
    window.$message?.error('上传失败')
    onError()
    emit('error')
  })
}
</script>
