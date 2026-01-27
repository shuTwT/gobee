<script setup lang="ts">
import { ref } from 'vue';
import { MdEditor, type ToolbarNames } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import MyToolbar from './MyToolbar.vue';
import CustomPreviewComponent from './CustomPreviewComponent.vue';
import "./config"

const props = withDefaults(defineProps<{
    preview?: boolean
}>(),{
    preview:false
})

const emit = defineEmits<{

}>()

const toolbars = [
    'bold',
  'underline',
  'italic',
  '-',
  'title',
  'strikeThrough',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
  'task',
  '-',
  'codeRow',
  'code',
  'link',
  'image',
  'table',
  'mermaid',
  'katex',
  0,
  '-',
  'revoke',
  'next',
  'save',
  '=',
  'pageFullscreen',
  'fullscreen',
  'preview',
  'previewOnly',
  'htmlPreview',
  'catalog',
  'github',
] satisfies ToolbarNames[]

/**
 * 不展示的工具栏
 */
const toolbarsExclude= [
    'save',
    'github',
    'catalog',
    'htmlPreview'
] satisfies ToolbarNames[]

const floatingToolbars = [
] satisfies ToolbarNames[]

const text = ref('# Hello Editor');
</script>
<template>
    <MdEditor v-model="text" :preview-component="CustomPreviewComponent" :preview="$props.preview" :show-toolbar-name="true" :toolbars="toolbars" :toolbars-exclude="toolbarsExclude" :floating-toolbars="floatingToolbars" style="height: 100%;" >
        <template #defToolbars>
            <MyToolbar/>
        </template>
    </MdEditor>
</template>