<script setup lang="ts">
import type { UploadFileInfo } from 'naive-ui'

const props = withDefaults(
  defineProps<{
    fileList?: string[]
    limit?: number
  }>(),
  {
    fileList: () => [],
    limit: 1,
  },
)

const emit = defineEmits<{
  (e: 'update:fileList', fileList: string[]): void
}>()

// 受控文件列表，供 n-upload 展示缩略图
const uploadFileList = ref<UploadFileInfo[]>([])

watch(
  () => props.fileList,
  (list) => {
    const urls = list || []
    // 保留仍在上传中（尚无 url）的条目，避免其从受控列表消失
    const pending = uploadFileList.value.filter((f) => !f.url || !urls.includes(f.url))
    const finished = urls.map((url) => ({
      id: url,
      name: url.split('/').pop() || url,
      url,
      status: 'finished' as const,
    }))
    uploadFileList.value = [...pending, ...finished]
  },
  { immediate: true },
)

// 上传完成：naive-ui 默认不会解析响应，需从 xhr 响应中取出 url
// 后端返回 { code, msg, data: [{ url, ... }] }
const handleUploadFinish = ({ file, event }: { file: UploadFileInfo; event?: ProgressEvent }) => {
  const response = (event?.target as XMLHttpRequest)?.response
  let data: Array<{ url?: string }> | undefined
  try {
    data = JSON.parse(typeof response === 'string' ? response : '')?.data
  } catch {
    data = undefined
  }
  const url = data?.[0]?.url
  if (url) {
    file.url = url
  }
}

// 列表变化（新增/删除/状态更新）时同步到外部
const handleFileListChange = (list: UploadFileInfo[]) => {
  uploadFileList.value = list
  emit('update:fileList', list.map((f) => f.url || '').filter(Boolean))
}
</script>
<template>
  <n-upload
    :file-list="uploadFileList"
    :multiple="limit > 1"
    accept="image/*"
    action="/api/v1/file/upload"
    list-type="image"
    :max="limit"
    @finish="handleUploadFinish"
    @update:file-list="handleFileListChange"
  >
    <n-upload-dragger v-if="uploadFileList.length === 0">
      <div style="margin-bottom: 12px">
        <n-icon size="48" :depth="3">
          <svg
            class="mx-auto h-12 w-12 text-gray-400"
            stroke="currentColor"
            fill="none"
            viewBox="0 0 48 48"
            aria-hidden="true"
          >
            <path
              d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </n-icon>
      </div>
      <n-text style="font-size: 16px"> 点击或者拖动图片到该区域来上传 </n-text>
      <n-p depth="3" style="margin: 8px 0 0 0">
        PNG, JPG, GIF 最大 10MB
      </n-p>
    </n-upload-dragger>
  </n-upload>
</template>
