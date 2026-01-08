<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import * as productApi from '@/api/mall/product'

interface Props {
  productId?: number
  productData: any
}

const props = defineProps<Props>()
const emit = defineEmits(['confirm'])

const formRef = ref<FormInst | null>(null)
const formData = ref({
  name: props.productData.name || '',
  sku: props.productData.sku || '',
  price: props.productData.price || 0,
  original_price: props.productData.original_price || 0,
  cost_price: props.productData.cost_price || 0,
  stock: props.productData.stock || 0,
  min_stock: props.productData.min_stock || 0,
  category_id: props.productData.category_id || null,
  brand: props.productData.brand || '',
  unit: props.productData.unit || '',
  weight: props.productData.weight || 0,
  volume: props.productData.volume || 0,
  description: props.productData.description || '',
  short_description: props.productData.short_description || '',
  images: props.productData.images || [],
  attributes: props.productData.attributes || {},
  tags: props.productData.tags || [],
  active: props.productData.active !== undefined ? props.productData.active : true,
  featured: props.productData.featured || false,
  digital: props.productData.digital || false,
  meta_title: props.productData.meta_title || '',
  meta_description: props.productData.meta_description || '',
  meta_keywords: props.productData.meta_keywords || '',
  sort_order: props.productData.sort_order || 0,
})

const rules: FormRules = {
  name: {
    required: true,
    message: '请输入商品名称',
    trigger: 'blur',
  },
  sku: {
    required: true,
    message: '请输入商品SKU',
    trigger: 'blur',
  },
  price: {
    required: true,
    type: 'number',
    message: '请输入商品价格',
    trigger: 'blur',
  },
  stock: {
    required: true,
    type: 'number',
    message: '请输入库存数量',
    trigger: 'blur',
  },
}

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          const data = toRaw(formData.value)
          const submitData: any = {
            name: data.name,
            sku: data.sku,
            price: data.price,
            stock: data.stock,
            active: data.active,
            featured: data.featured,
            digital: data.digital,
          }
          
          if (data.original_price) submitData.original_price = data.original_price
          if (data.cost_price) submitData.cost_price = data.cost_price
          if (data.min_stock) submitData.min_stock = data.min_stock
          if (data.category_id) submitData.category_id = data.category_id
          if (data.brand) submitData.brand = data.brand
          if (data.unit) submitData.unit = data.unit
          if (data.weight) submitData.weight = data.weight
          if (data.volume) submitData.volume = data.volume
          if (data.description) submitData.description = data.description
          if (data.short_description) submitData.short_description = data.short_description
          if (data.images && data.images.length > 0) submitData.images = data.images
          if (data.attributes && Object.keys(data.attributes).length > 0) submitData.attributes = data.attributes
          if (data.tags && data.tags.length > 0) submitData.tags = data.tags
          if (data.meta_title) submitData.meta_title = data.meta_title
          if (data.meta_description) submitData.meta_description = data.meta_description
          if (data.meta_keywords) submitData.meta_keywords = data.meta_keywords
          if (data.sort_order) submitData.sort_order = data.sort_order
          
          resolve(submitData)
        } else {
          reject(errors)
        }
      })
    } else {
      reject(new Error('表单实例不存在'))
    }
  })
}

const handleSubmit = async () => {
  try {
    const data = await getData()
    
    if (props.productId) {
      await productApi.updateProduct(props.productId, data)
      window.$message?.success('更新成功')
    } else {
      await productApi.createProduct(data)
      window.$message?.success('创建成功')
    }
    
    emit('confirm')
  } catch (error) {
    console.error('提交失败:', error)
  }
}

defineExpose({ getData, handleSubmit })
</script>

<template>
  <n-form :rules="rules" ref="formRef" label-placement="left" label-width="120px">
    <n-form-item label="商品名称" path="name">
      <n-input v-model:value="formData.name" placeholder="请输入商品名称" />
    </n-form-item>
    <n-form-item label="商品SKU" path="sku">
      <n-input v-model:value="formData.sku" placeholder="请输入商品SKU" />
    </n-form-item>
    <n-form-item label="商品价格(元)" path="price">
      <n-input-number
        v-model:value="formData.price"
        :min="0"
        :precision="2"
        placeholder="请输入商品价格"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="原价(元)" path="original_price">
      <n-input-number
        v-model:value="formData.original_price"
        :min="0"
        :precision="2"
        placeholder="请输入原价"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="成本价(元)" path="cost_price">
      <n-input-number
        v-model:value="formData.cost_price"
        :min="0"
        :precision="2"
        placeholder="请输入成本价"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="库存数量" path="stock">
      <n-input-number
        v-model:value="formData.stock"
        :min="0"
        placeholder="请输入库存数量"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="最低库存预警" path="min_stock">
      <n-input-number
        v-model:value="formData.min_stock"
        :min="0"
        placeholder="请输入最低库存预警"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="品牌" path="brand">
      <n-input v-model:value="formData.brand" placeholder="请输入品牌" />
    </n-form-item>
    <n-form-item label="单位" path="unit">
      <n-input v-model:value="formData.unit" placeholder="请输入单位" />
    </n-form-item>
    <n-form-item label="重量(kg)" path="weight">
      <n-input-number
        v-model:value="formData.weight"
        :min="0"
        :precision="2"
        placeholder="请输入重量"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="体积(立方米)" path="volume">
      <n-input-number
        v-model:value="formData.volume"
        :min="0"
        :precision="3"
        placeholder="请输入体积"
        style="width: 100%"
      />
    </n-form-item>
    <n-form-item label="简短描述" path="short_description">
      <n-input
        v-model:value="formData.short_description"
        type="textarea"
        placeholder="请输入简短描述"
        :rows="2"
      />
    </n-form-item>
    <n-form-item label="商品描述" path="description">
      <n-input
        v-model:value="formData.description"
        type="textarea"
        placeholder="请输入商品描述"
        :rows="4"
      />
    </n-form-item>
    <n-form-item label="商品标签" path="tags">
      <n-dynamic-tags v-model:value="formData.tags" />
    </n-form-item>
    <n-form-item label="是否上架" path="active">
      <n-switch v-model:value="formData.active" />
    </n-form-item>
    <n-form-item label="是否推荐" path="featured">
      <n-switch v-model:value="formData.featured" />
    </n-form-item>
    <n-form-item label="是否数字商品" path="digital">
      <n-switch v-model:value="formData.digital" />
    </n-form-item>
    <n-form-item label="SEO标题" path="meta_title">
      <n-input v-model:value="formData.meta_title" placeholder="请输入SEO标题" />
    </n-form-item>
    <n-form-item label="SEO描述" path="meta_description">
      <n-input
        v-model:value="formData.meta_description"
        type="textarea"
        placeholder="请输入SEO描述"
        :rows="2"
      />
    </n-form-item>
    <n-form-item label="SEO关键词" path="meta_keywords">
      <n-input v-model:value="formData.meta_keywords" placeholder="请输入SEO关键词" />
    </n-form-item>
    <n-form-item label="排序" path="sort_order">
      <n-input-number
        v-model:value="formData.sort_order"
        :min="0"
        placeholder="请输入排序"
        style="width: 100%"
      />
    </n-form-item>
  </n-form>
</template>
