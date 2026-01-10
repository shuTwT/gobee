<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui';
import type { FormItemProps } from './utils/types'
import { getCategoryList } from '@/api/content/category'
import { getTagList } from '@/api/content/tag'

const props = defineProps<{
  formInline: FormItemProps
}>()

const formRef = ref<FormInst|null>()
const formData = ref(props.formInline)
const rules:FormRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  content_type: [{ required: true, message: '请选择内容类型', trigger: 'change' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }],
  author: [{ required: true, message: '请输入作者', trigger: 'blur' }],
}

const categoryOptions = ref([])

const tagOptions = ref([])

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

const loadCategoryOptions = async () => {
  try {
    const res = await getCategoryList()
    if (res.code === 200 && res.data) {
      categoryOptions.value = res.data.map((item: any) => ({
        value: item.id,
        label: item.name
      }))
    }
  } catch (error) {
    console.error('加载分类列表失败:', error)
  }
}

const loadTagOptions = async () => {
  try {
    const res = await getTagList()
    if (res.code === 200 && res.data) {
      tagOptions.value = res.data.map((item: any) => ({
        value: item.id,
        label: item.name
      }))
    }
  } catch (error) {
    console.error('加载标签列表失败:', error)
  }
}

onMounted(() => {
  loadCategoryOptions()
  loadTagOptions()
})

defineExpose({getData})

</script>
<template>
  <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="160px">
    <n-form-item label="标题" path="title">
      <n-input v-model:value="formData.title" placeholder="请输入文章标题"/>
    </n-form-item>
    <n-form-item label="别名" path="alias">
      <n-input v-model:value="formData.alias" placeholder="请输入文章别名"/>
    </n-form-item>
    
    <n-grid :cols="2" :x-gap="16">
      <n-gi>
        <n-form-item label="内容类型" path="content_type">
          <n-select 
            v-model:value="formData.content_type" 
            placeholder="请选择内容类型"
            :options="[
              { value: 'markdown', label: 'Markdown' },
              { value: 'html', label: 'HTML' }
            ]"
          />
        </n-form-item>
      </n-gi>
      <n-gi>
        <n-form-item label="状态" path="status">
          <n-select 
            v-model:value="formData.status" 
            placeholder="请选择状态"
            :options="[
              { value: 'draft', label: '草稿' },
              { value: 'published', label: '已发布' },
              { value: 'archived', label: '已归档' }
            ]"
          />
        </n-form-item>
      </n-gi>
    </n-grid>
    
    <n-form-item label="分类" path="categories">
      <n-select v-model:value="formData.categories" :multiple="true" :filterable="true" :options="categoryOptions" placeholder="请选择分类"/>
    </n-form-item>
    <n-form-item label="标签" path="tags">
      <n-select v-model:value="formData.tags" :multiple="true" :filterable="true" :options="tagOptions" placeholder="请选择标签"/>
    </n-form-item>
    
    <n-grid :cols="2" :x-gap="16">
      <n-gi>
        <n-form-item label="作者" path="author">
          <n-input v-model:value="formData.author" placeholder="请输入作者"/>
        </n-form-item>
      </n-gi>
      <n-gi>
        <n-form-item label="封面" path="cover">
          <n-input v-model:value="formData.cover" placeholder="请输入封面URL"/>
        </n-form-item>
      </n-gi>
    </n-grid>
    
    <n-form-item label="关键词" path="keywords">
      <n-input v-model:value="formData.keywords" placeholder="请输入关键词，用逗号分隔"/>
    </n-form-item>
    <n-form-item label="版权" path="copyright">
      <n-input v-model:value="formData.copyright" placeholder="请输入版权信息"/>
    </n-form-item>
    
    <n-form-item label="自动生成摘要" path="is_autogen_summary">
      <n-checkbox v-model:checked="formData.is_autogen_summary"/>
    </n-form-item>
    <n-form-item v-if="!formData.is_autogen_summary" label="摘要" path="summary">
      <n-input v-model:value="formData.summary" type="textarea" :autosize="{ minRows: 3, maxRows: 6 }" placeholder="请输入文章摘要"/>
    </n-form-item>
    
    <n-form-item label="是否置顶" path="is_pin_to_top">
      <n-checkbox v-model:checked="formData.is_pin_to_top"/>
    </n-form-item>
    <n-form-item label="是否可见" path="is_visible">
      <n-checkbox v-model:checked="formData.is_visible"/>
    </n-form-item>
    <n-form-item label="允许评论" path="is_allow_comment">
      <n-checkbox v-model:checked="formData.is_allow_comment"/>
    </n-form-item>
    <n-form-item label="评论后可见" path="is_visible_after_comment">
      <n-checkbox v-model:checked="formData.is_visible_after_comment"/>
    </n-form-item>
    <n-form-item label="支付后可见" path="is_visible_after_pay">
      <n-checkbox v-model:checked="formData.is_visible_after_pay"/>
    </n-form-item>
    <n-form-item v-if="formData.is_visible_after_pay" label="支付金额" path="price">
      <n-input-number v-model:value="formData.price" :min="0" :precision="2" placeholder="请输入支付金额"/>
    </n-form-item>
    
    <n-grid :cols="2" :x-gap="16">
      <n-gi>
        <n-form-item label="浏览次数" path="view_count">
          <n-input-number v-model:value="formData.view_count" :disabled="true" placeholder="浏览次数"/>
        </n-form-item>
      </n-gi>
      <n-gi>
        <n-form-item label="评论次数" path="comment_count">
          <n-input-number v-model:value="formData.comment_count" :disabled="true" placeholder="评论次数"/>
        </n-form-item>
      </n-gi>
    </n-grid>
  </n-form>
</template>
