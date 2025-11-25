<script setup lang="ts">
import '@wangeditor-next/editor/dist/css/style.css' // 引入 css
import { Editor, Toolbar } from '@wangeditor-next/editor-for-vue'
import { useRoute } from 'vue-router'
import * as postApi from '@/api/post'
import { usePostHook } from '../utils/hook'
import {SlateNode} from "@wangeditor-next/editor"
import type { IDomEditor, IEditorConfig,IToolbarConfig } from '@wangeditor-next/editor'
import { BASE_URL } from '@/api/utils'

const { settingPost,savePost,publishPost,unpublishPost } = usePostHook()

const route = useRoute()
const mode = 'default'
const editorRef = shallowRef()
const showPreview = ref(false)

const publishStatus = ref(false)

const valueHtml = ref('<p>hello</p>')
const previewHtml = ref('')

const catalogHtml = ref('')

const toolbarConfig:Partial<IToolbarConfig> = {
}
const editorConfig:Partial<IEditorConfig> = {
  MENU_CONF:{
    'uploadImage':{
      server:BASE_URL+"/v1/file/upload",
      metaWithUrl:true,
      base64LimitSize:1024*1024*2,
      onSuccess:(file,res)=>{

      },
      onError:(res)=>{
        console.log(res)
      },
      onFailed:(res)=>{
        console.log(res)
      },
      customInsert:(res,insertFn)=>{
        if(res.code==200&&res.data.length>0){
          insertFn(res.data[0].url)
        }
      },
    }
  },
  placeholder: '请输入内容...',
 }

const handleEditorChange = (editor: IDomEditor) => {
  const headers = editor.getElemsByTypePrefix('header')
    catalogHtml.value = headers.map(header => {
            const text = SlateNode.string(header)
            const { id, type } = header
            return `<li id="${id}" type="${type}">${text}</li>`
          }).join('')

}

const handleCreated = (editor: any) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}

const openSettingDialog = () => {
  settingPost({id:route.query.id})
}

const handlePreview = () => {
  previewHtml.value = editorRef.value.getHtml()
  showPreview.value = true
}

const handleSave=()=>{
  savePost({
    id:route.query.id,
    content:valueHtml.value
  }).then(()=>{
    getPostData()
  })
}

const handlePublish = () => {
  publishPost({id:route.query.id}).then(()=>{
    publishStatus.value = true
  })
}

const handleUnpublish = () => {
  unpublishPost({id:route.query.id}).then(()=>{
    publishStatus.value = false
  })
}

const getPostData = ()=>{
  const id = route.query.id
  if(id){
    postApi.queryPost(id+"").then((res)=>{
      valueHtml.value = res.data.content
      if(res.data.status=='draft'){
        publishStatus.value= false
      }else if(res.data.status=='published'){
        publishStatus.value = true
      }
    })
  }
}

onMounted(()=>{
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
      <Toolbar
        style="border-bottom: 1px solid #ccc"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        :mode="mode"
      />
      <Editor
        style="height: var(--editor-height); overflow-y: hidden"
        v-model="valueHtml"
        :defaultConfig="editorConfig"
        :mode="mode"
        @onCreated="handleCreated"
        @onChange="handleEditorChange"
      />
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
          <n-button v-if="!publishStatus" type="primary" size="large" @click="handlePublish"> 发布 </n-button>
          <n-button v-else type="primary" size="large" @click="handleUnpublish"> 取消发布 </n-button>
        </div>
        <n-divider />
        <n-tabs type="segment" animated>
          <n-tab-pane name="chap1" tab="大纲">
            <ul v-html="catalogHtml" class="p-2">

            </ul>
          </n-tab-pane>
          <n-tab-pane name="chap2" tab="详情"></n-tab-pane>
        </n-tabs>
      </div>
    </div>
    <n-modal v-model:show="showPreview" preset="card" style="height: 100vh;">
      <div class="w-full h-full" v-html="previewHtml" tabindex="1">

      </div>
    </n-modal>
  </div>
</template>
<style scoped>
.editor-wrapper {
  --editor-wrapper-height: calc(
    100vh - var(--pro-layout-footer-height) - var(--pro-layout-nav-height) - 48px
  );
  --editor-height: calc(var(--editor-wrapper-height) - 42px);
  height: var(--editor-wrapper-height);
}
</style>
