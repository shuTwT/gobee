<script setup lang="ts">

const formData = reactive({
  title:""
})
</script>
<template>
  <n-form>
    <n-form-item label="标题">
      <n-input v-model:value="formData.title"/>
    </n-form-item>
    <n-form-item label="别名"></n-form-item>
    <n-form-item label="分类"></n-form-item>
    <n-form-item label="标签"></n-form-item>
    <n-form-item label="自动生成摘要"></n-form-item>
    <n-form-item label="作者"></n-form-item>
    <n-form-item label="允许评论"></n-form-item>
    <n-form-item label="是否置顶"></n-form-item>
    <n-form-item label="可见性"></n-form-item>
    <n-form-item label="封面"></n-form-item>
  </n-form>
</template>
