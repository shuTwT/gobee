<script setup lang="ts">
const props = defineProps<{
  html: string;
  id?: string;
  className?: string;
}>();
onMounted(()=>{
})
</script>
<template>
    <div :id="id" :class="className" v-html="html" />
</template>