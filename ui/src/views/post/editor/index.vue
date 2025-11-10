<script setup lang="ts">
import '@wangeditor-next/editor/dist/css/style.css' // 引入 css
import { Editor, Toolbar } from '@wangeditor-next/editor-for-vue'
import { useRoute } from 'vue-router'
import * as postApi from '@/api/post'
import { usePostHook } from '../utils/hook'

const { settingPost,savePost,publishPost } = usePostHook()

const route = useRoute()
const mode = 'default'
const editorRef = shallowRef()

const valueHtml = ref('<p>hello</p>')

const toolbarConfig = {}
const editorConfig = { placeholder: '请输入内容...' }

const handleCreated = (editor: any) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}

const openSettingDialog = () => {
  settingPost({id:route.query.id})
}

const handleSave=()=>{
  savePost({
    id:route.query.id,
    content:valueHtml.value
  })
}

const handlePublish = () => {
  publishPost({id:route.query.id})
}

onMounted(()=>{
  const id = route.query.id
  if(id){
    postApi.queryPost(id+"").then((res)=>{
      valueHtml.value = res.data.content
    })
  }
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
      />
    </div>
    <div class="w-[400px]">
      <div class="border border-[#ccc] bg-white flex flex-col">
        <div class="pt-6 flex justify-end pr-6 items-center">
          <n-button style="margin-right: 10px"> 历史版本 </n-button>
          <n-button style="margin-right: 10px"> 预览 </n-button>
          <n-button style="margin-right: 10px" @click="openSettingDialog"> 设置 </n-button>
          <n-button style="margin-right: 10px" @click="handleSave"> 保存 </n-button>
          <n-button type="primary" size="large" @click="handlePublish"> 发布 </n-button>
        </div>
        <n-divider />
        <n-tabs type="segment" animated>
          <n-tab-pane name="chap1" tab="大纲"></n-tab-pane>
          <n-tab-pane name="chap2" tab="详情"></n-tab-pane>
        </n-tabs>
      </div>
    </div>
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
