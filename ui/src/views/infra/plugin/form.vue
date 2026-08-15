<script setup lang="ts">
import { ref } from 'vue'
import { NUpload, NUploadDragger, NButton, NForm, NFormItem, NAlert, NSpin, type UploadSettledFileInfo, NIcon } from 'naive-ui'
import { CloudUploadOutline } from '@vicons/ionicons5'
import { createPlugin } from '@/api/infra/plugin'

const emit = defineEmits(['success'])

const uploadRef = ref()
const fileListRef = ref<File[]>([])
const uploading = ref(false)

const handleUploadChange = (options: {
    file: UploadSettledFileInfo;
    fileList: UploadSettledFileInfo[];
    event: ProgressEvent | Event | undefined;
}) => {
  const fileList = options.fileList
  if (fileList.length>0) {
    fileListRef.value = [fileList[0].file as File]
  }
}

const beforeUpload = () => {
  if (fileListRef.value.length > 0) {
    return false
  }
  return true
}

const handleUpload = async () => {
  if (fileListRef.value.length === 0) {
    window.$message?.error('请先选择文件')
    return
  }

  uploading.value = true

  try {
    await createPlugin(fileListRef.value[0])
    window.$message?.success('插件上传成功')
    emit('success')
    fileListRef.value = []
  } catch (error: any) {
    window.$message?.error('插件上传失败：' + error.message)
  } finally {
    uploading.value = false
  }
}

const handleRemove = () => {
  fileListRef.value = []
}
</script>

<template>
  <div class="plugin-upload">
    <n-alert type="warning">请勿安装来源不明的插件</n-alert>

    <n-form>
      <n-form-item label="选择插件压缩包">
        <n-upload
          ref="uploadRef"
          :max="1"
          accept=".zip"
          :show-file-list="false"
          :disabled="uploading"
          @change="handleUploadChange"
          @before-upload="beforeUpload"
        >
          <n-upload-dragger>
            <div class="upload-area">
              <n-spin :show="uploading">
                <div class="upload-inner">
                  <n-icon size="48" :component="CloudUploadOutline" />
                  <div class="upload-text">点击或拖拽上传插件压缩包</div>
                  <div class="upload-hint">支持 .zip 格式</div>
                </div>
              </n-spin>
            </div>
          </n-upload-dragger>
        </n-upload>
      </n-form-item>

      <n-form-item v-if="fileListRef.length > 0">
        <div class="file-info">
          <span>已选择文件：</span>
          <span class="file-name">{{ fileListRef[0].name }}</span>
          <span class="file-size">{{ (fileListRef[0].size / 1024).toFixed(2) }} KB</span>
          <n-button text type="error" @click="handleRemove">移除</n-button>
        </div>
      </n-form-item>

      <n-form-item>
        <n-button type="primary" :loading="uploading" :disabled="fileListRef.length === 0" @click="handleUpload">
          上传插件
        </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style scoped>
.plugin-upload {
  max-width: 800px;
  margin: 0 auto;
}

.upload-area {
  width: 100%;
  height: 300px;
  border: 2px dashed #d0d0d0;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #fafafa;
  transition: all 0.3s;
}

.upload-area:hover {
  border-color: #409eff;
  background-color: #f0f9ff;
}

.upload-inner {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 32px 48px;
}

.upload-text {
  margin-top: 16px;
  font-size: 16px;
  color: #666;
}

.upload-hint {
  margin-top: 8px;
  font-size: 12px;
  color: #999;
}

.file-info {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background-color: #f5f5f5;
  border-radius: 4px;
}

.file-name {
  font-weight: 500;
  color: #333;
}

.file-size {
  color: #999;
  font-size: 12px;
}
</style>
