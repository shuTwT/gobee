<script setup lang="ts">
import { FileTraySharp } from "@vicons/ionicons5"
import type { AlbumPhotoFormProps } from "./utils/types";
import type { FormInst, FormRules } from "naive-ui";
import { addDialog } from "@/components/dialog";
import ImageSelector from "@/components/ImageSelector.vue";

const props = defineProps<AlbumPhotoFormProps>()

const formRef = ref<FormInst|null>()
const formData = ref(props.formInline)
const rules:FormRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
}

const openUploadImage = ()=>{
  const selectedImage = ref('')
  addDialog({
    title: '选择附件',
    width: '80%',
    height: '80%',
    contentRenderer:()=>h(ImageSelector, {
      visible: true,
      modelValue: selectedImage.value,
      'onUpdate:modelValue': (value) => {
        selectedImage.value = value
      }
    }),
    beforeSure:async(done)=>{
      if (selectedImage.value) {
        formData.value.image_url = selectedImage.value
      }
      done()
    }
  })
}

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          resolve(toRaw(formData.value))
        } else {
          reject(errors)
        }
      })
    } else {
      reject(new Error('表单实例不存在'))
    }
  })
}

defineExpose({getData})
</script>
<template>
  <n-form ref="formRef" :rules="rules" :model="formData">
    <n-form-item label="相片名称" path="name">
      <n-input v-model:value="formData.name" />
    </n-form-item>
    <n-form-item label="图片地址" path="image_url">
      <n-input-group >
        <n-input v-model:value="formData.image_url"/>
        <n-button @click="openUploadImage">
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
