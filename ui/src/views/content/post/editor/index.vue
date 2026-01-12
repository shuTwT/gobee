<script setup lang="ts">
import Editor from '@tinymce/tinymce-vue'
await import ('tinymce/tinymce')
import '@/utils/tinymce/langs/zh_CN'
// DOM model
 import 'tinymce/models/dom/model'
// Theme
 import 'tinymce/themes/silver'
// Toolbar icons
 import 'tinymce/icons/default'
// Editor styles
 import 'tinymce/skins/ui/oxide/skin.min.css'
// Import plugins
import 'tinymce/plugins/advlist'
import 'tinymce/plugins/autolink'
import 'tinymce/plugins/link'
import 'tinymce/plugins/image'
import 'tinymce/plugins/lists'
import 'tinymce/plugins/table'
import 'tinymce/plugins/code'
import 'tinymce/plugins/codesample'
import 'tinymce/plugins/wordcount'
await import ('tinymce/skins/content/default/content.js')
await import ('tinymce/skins/ui/oxide/content.js')
import { useRoute } from 'vue-router'
import * as postApi from '@/api/content/post'
import { usePostHook } from '../utils/hook'

const { settingPost, savePost, publishPost, unpublishPost, importPost } = usePostHook()

const route = useRoute()
const editorRef = shallowRef()
const showPreview = ref(false)

const publishStatus = ref(false)

const valueHtml = ref('<p>hello</p>')
const previewHtml = ref('')

const catalogHtml = ref('')

const editorInit = {
  selector:'textarea',
  content: '<p>This is the initial content of the editor.</p>',
  language: 'zh_CN',
  height: '100%',
  resize:false,
  promotion: false, // 官方推荐的关闭方式（见下文）
  branding:false,
  plugins: 'advlist autolink lists link image table code wordcount codesample',
  toolbar:
    'undo redo | blocks | bold italic | alignleft aligncenter alignright | bullist numlist | image | table | code |wordcount | codesample',

  // menubar: false,
  image_list:[],
  file_picker_callback:(callback:any,value:any,meta:any)=>{
     if (meta.filetype == 'file') {
      callback('mypage.html', { text: 'My text' });
    }

    // Provide image and alt text for the image dialog
    if (meta.filetype == 'image') {
      callback('myimage.jpg', { alt: 'My alt text' });
    }

    // Provide alternative source and posted for the media dialog
    if (meta.filetype == 'media') {
      callback('movie.mp4', { source2: 'alt.ogg', poster: 'image.jpg' });
    }
  }
}

const openSettingDialog = () => {
  settingPost({ id: route.query.id })
}

const handlePreview = () => {
  previewHtml.value = valueHtml.value
  showPreview.value = true
}

const handleSave = () => {
  savePost({
    id: route.query.id,
    content: valueHtml.value,
  }).then(() => {
    getPostData()
  })
}

const handlePublish = () => {
  publishPost({ id: route.query.id }).then(() => {
    publishStatus.value = true
  })
}

const handleUnpublish = () => {
  unpublishPost({ id: route.query.id }).then(() => {
    publishStatus.value = false
  })
}

const handleImport = () =>{
  importPost({
    id:route.query.id
  }).then((res:any)=>{
    valueHtml.value = res.importContent
  })
}

const getPostData = () => {
  const id = route.query.id
  if (id) {
    postApi.queryPost(id + '').then((res) => {
      valueHtml.value = res.data.content
      if (res.data.status == 'draft') {
        publishStatus.value = false
      } else if (res.data.status == 'published') {
        publishStatus.value = true
      }
    })
  }
}

onMounted(() => {
  getPostData()
})

onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})
</script>
<template>
  <div class="p-6 flex gap-1">
    <div class="editor-wrapper w-[calc(100%_-_400px)] border border-[#ccc]">
      <Editor v-model:modelValue="valueHtml" id="uuid" licenseKey="gpl" :init="editorInit" />
    </div>
    <div class="w-[400px]">
      <div class="border border-[#ccc] bg-white flex flex-col">
        <div class="pt-6 flex justify-end pr-6 items-center">
          <n-button style="margin-right: 10px"> 历史版本 </n-button>
          <n-button style="margin-right: 10px" @click="handlePreview"> 预览 </n-button>
          <n-button style="margin-right: 10px" @click="openSettingDialog"> 设置 </n-button>
          <n-button @click="handleSave"> 保存 </n-button>
        </div>
        <div class="pt-6 flex justify-end pr-6 items-center">
          <n-button style="margin-right: 10px;" @click="handleImport">导入</n-button>
          <n-button v-if="!publishStatus" type="primary" @click="handlePublish">
            发布
          </n-button>
          <n-button v-else type="primary" @click="handleUnpublish">
            取消发布
          </n-button>
        </div>
        <n-divider />
        <n-tabs type="segment" animated>
          <n-tab-pane name="chap1" tab="大纲">
            <ul v-html="catalogHtml" class="p-2"></ul>
          </n-tab-pane>
          <n-tab-pane name="chap2" tab="详情"></n-tab-pane>
        </n-tabs>
      </div>
    </div>
    <n-modal v-model:show="showPreview" preset="card" style="height: 100vh">
      <div class="w-full h-full"  tabindex="1">
        <n-scrollbar style="height: calc(100vh - 80px);">
          <div  v-html="previewHtml"></div>
      </n-scrollbar>
      </div>
    </n-modal>
  </div>
</template>
<style scoped>
.editor-wrapper {
  --editor-wrapper-height: calc(
    100vh - var(--pro-layout-footer-height) - var(--pro-layout-nav-height) - 48px
  );
  --editor-height: calc(var(--editor-wrapper-height) - 83px);
  height: var(--editor-wrapper-height);
}
</style>
