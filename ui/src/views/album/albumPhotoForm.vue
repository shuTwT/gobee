<script setup lang="ts">
import { FileTraySharp } from "@vicons/ionicons5"

const formData = ref({
  name:"",
  image_url:"",
  description:""
})
</script>
<template>
  <n-form>
    <n-form-item label="相片名称" path="name">
      <n-input v-model:value="formData.name" />
    </n-form-item>
    <n-form-item label="图片地址" path="image_url">
      <n-input-group v-model:value="formData.image_url" >
        <n-input/>
        <n-button>
          <n-icon>
            <file-tray-sharp/>
          </n-icon>
        </n-button>
      </n-input-group>
    </n-form-item>
    <n-form-item label="相片描述" path="description">
      <n-input v-model:value="formData.description" />
    </n-form-item>
  </n-form>
</template>
