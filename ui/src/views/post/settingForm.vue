<script setup lang="ts">

const formData = reactive({
  title:"",
  alias:"",
  categorys:[],
  tags:[],
  autogenSummary:false,
  author:"",
  allowComments:true,
  pinToTop:false,
  visible:true,
  cover:""
})

const categoryOptions = ref([])

const tagOptions = ref([])

</script>
<template>
  <n-form>
    <n-form-item label="标题">
      <n-input v-model:value="formData.title"/>
    </n-form-item>
    <n-form-item label="别名">
      <n-input v-model:value="formData.alias"/>
    </n-form-item>
    <n-form-item label="分类">
      <n-select v-model:value="formData.categorys" :multiple="true" :filterable="true" :options="categoryOptions"/>
    </n-form-item>
    <n-form-item label="标签">
      <n-select v-model:value="formData.tags" :multiple="true" :filterable="true" :options="tagOptions"/>
    </n-form-item>
    <n-form-item label="自动生成摘要">
      <n-checkbox v-model:checked="formData.autogenSummary"/>
    </n-form-item>
    <n-form-item label="作者">
      <n-input v-model:value="formData.author"/>
    </n-form-item>
    <n-form-item label="允许评论">
      <n-checkbox v-model:checked="formData.allowComments"/>
    </n-form-item>
    <n-form-item label="是否置顶">
      <n-checkbox v-model:checked="formData.pinToTop"/>
    </n-form-item>
    <n-form-item label="可见性">
      <n-checkbox v-model:checked="formData.visible"/>
    </n-form-item>
    <n-form-item label="封面">
      <n-input v-model:value="formData.cover"/>
    </n-form-item>
  </n-form>
</template>
