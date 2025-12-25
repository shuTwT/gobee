<script lang="ts" setup>
import { ArchiveOutline as ArchiveIcon } from '@vicons/ionicons5'
import { parse } from 'marked'
import DOMPurify from 'dompurify'
const formData = ref({
  importType: 'markdown',
  importName: '',
  importContent: '',
})
const handleUpload = ({ event, file }: any) => {
  if (event) {
    formData.value.importName = file.name
    const reader = new FileReader()
    reader.onload = async (e) => {
      try {
        const resultText = e.target!.result as string
        if (formData.value.importType === 'markdown') {
          const rawHtml = await parse(resultText)
          const cleanHtml = DOMPurify.sanitize(rawHtml)
          formData.value.importContent = cleanHtml
        }else if(formData.value.importType==='html'){
          const cleanHtml = DOMPurify.sanitize(resultText)
          formData.value.importContent = cleanHtml
        }else{
          const cleanHtml = DOMPurify.sanitize(resultText)
          formData.value.importContent = `<p>${cleanHtml}</p>`
        }
      } catch (err) {
        console.error(err)
      }
    }
    reader.readAsText(file.file)
  }
}

const getData = () => {
  return formData.value
}

defineExpose({ getData })
</script>
<template>
  <n-form>
    <n-form-item label="文件类型">
      <n-radio-group v-model:value="formData.importType" name="importTypeGroup">
        <n-radio-button label="Markdown" value="markdown" />

        <n-radio-button label="Html" value="html" />
        <!--<n-radio-button label="Text" value="text" />-->
      </n-radio-group>
    </n-form-item>
    <n-form-item label="文件">
      <n-upload multiple directory-dnd :default-upload="false" :max="1" @change="handleUpload">
        <n-upload-dragger>
          <div style="margin-bottom: 12px">
            <n-icon size="48" :depth="3">
              <ArchiveIcon />
            </n-icon>
          </div>
          <n-text style="font-size: 16px"> 点击或者拖动文件到该区域来上传 </n-text>
          <n-p depth="3" style="margin: 8px 0 0 0"> 请上传指定的格式 </n-p>
        </n-upload-dragger>
      </n-upload>
    </n-form-item>
  </n-form>
</template>
