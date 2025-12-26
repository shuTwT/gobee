<script setup lang="ts">
import type { UploadFileInfo } from 'naive-ui';

// import type { UploadFileInfo } from 'naive-ui';

const props  = withDefaults(defineProps<{
  fileList?: string[],
  limit?:number
}>(),{
  fileList: ()=>[],
  limit:1
})

const fileListRef = ref(props.fileList)

const emit = defineEmits<{
  (e: 'update:fileList', fileList: string[]): void
}>()

// const defaultUploadFileList = ref<UploadFileInfo[]>([])
// const uploadFileList = ref<UploadFileInfo[]>([])

const handleUploadFinish=({file}:{file:UploadFileInfo,event?:ProgressEvent})=>{
  fileListRef.value.push(file.url||'')
  emit('update:fileList',fileListRef.value)
}

const getFileList = ()=>{
  return fileListRef.value
}
defineExpose({getFileList})
</script>
<template>
  <n-upload
    :multiple="props.limit > 1"
    accept="image/*"
    action="/api/v1/file/upload"
    list-type="image"
    :max="props.limit"
    @finish="handleUploadFinish"
  >
  <template v-if="fileListRef&&fileListRef.length>0">
    <div v-for="(item,index) in fileListRef" :key="index">
      <n-image :src="item"/>
      <div class="upload-toolbar" @click.stop></div>
    </div>
  </template>
    <n-upload-dragger v-else>
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
