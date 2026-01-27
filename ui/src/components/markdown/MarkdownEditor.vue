<script setup lang="ts">
import { ref } from 'vue';
import { MdEditor, type ToolbarNames } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

const props = withDefaults(defineProps<{
    preview?: boolean
}>(),{
    preview:false
})

const emit = defineEmits<{

}>()

/**
 * 不展示的工具栏
 */
const toolbarsExclude= [
    'save'
] satisfies ToolbarNames[]

const text = ref('# Hello Editor');
</script>
<template>
    <MdEditor v-model="text" :preview="$props.preview" :toolbars-exclude="toolbarsExclude" style="height: 100%;"/>
</template>