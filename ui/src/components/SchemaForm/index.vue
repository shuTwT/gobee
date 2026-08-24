<script setup lang="ts">
import type { FormInst, FormRules } from 'naive-ui'
import Field from './Field.vue'
import type { SettingFormGroupProps } from './types'

const props = defineProps<{
  forms: SettingFormGroupProps[]
  modelValue: Record<string, any>
}>()

const formRef = ref<FormInst | null>(null)
const formData = ref(props.modelValue)

// 值对象中缺失的字段按类型补齐初始键，保证控件绑定生效
watch(
  () => props.forms,
  (forms) => {
    for (const form of forms) {
      for (const field of form.formSchema || []) {
        if (!(field.name in formData.value)) {
          formData.value[field.name] =
            field.default ?? (field.type === 'switch' ? false : field.type === 'number' ? null : '')
        }
      }
    }
  },
  { immediate: true },
)

const rules = computed<FormRules>(() => {
  const result: FormRules = {}
  for (const form of props.forms) {
    for (const field of form.formSchema || []) {
      if (field.required) {
        result[field.name] = [
          {
            required: true,
            message: `请填写${field.label}`,
            trigger: ['blur', 'change'],
          },
        ]
      }
    }
  }
  return result
})

const isSingleForm = computed(() => props.forms.length === 1)

const getData = () => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((errors) => {
        if (!errors) {
          resolve({ ...toRaw(formData.value) })
        } else {
          reject(errors)
        }
      })
    } else {
      reject(new Error('表单实例不存在'))
    }
  })
}

defineExpose({ getData })
</script>

<template>
  <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="120px">
    <template v-if="isSingleForm">
      <n-form-item
        v-for="field in forms[0].formSchema || []"
        :key="field.name"
        :label="field.label"
        :path="field.name"
      >
        <Field v-model:value="formData[field.name]" :field="field" />
      </n-form-item>
    </template>
    <n-tabs v-else type="line" animated>
      <n-tab-pane
        v-for="form in forms"
        :key="form.group"
        :name="form.group"
        :tab="form.label"
        display-directive="show"
      >
        <n-form-item
          v-for="field in form.formSchema || []"
          :key="field.name"
          :label="field.label"
          :path="field.name"
        >
          <Field v-model:value="formData[field.name]" :field="field" />
        </n-form-item>
      </n-tab-pane>
    </n-tabs>
  </n-form>
</template>
