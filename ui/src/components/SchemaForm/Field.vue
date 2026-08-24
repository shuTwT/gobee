<script setup lang="ts">
import type { SettingFieldProps } from './types'

defineProps<{
  field: SettingFieldProps
}>()

const value = defineModel<any>('value')
</script>

<template>
  <div class="w-full">
    <n-input
      v-if="field.type === 'text'"
      v-model:value="value"
      :placeholder="field.placeholder || `请输入${field.label}`"
    />
    <n-input
      v-else-if="field.type === 'textarea'"
      v-model:value="value"
      type="textarea"
      :rows="3"
      :placeholder="field.placeholder || `请输入${field.label}`"
    />
    <n-input
      v-else-if="field.type === 'secret'"
      v-model:value="value"
      type="password"
      show-password-on="click"
      :placeholder="field.placeholder || `请输入${field.label}`"
    />
    <n-input-number
      v-else-if="field.type === 'number'"
      v-model:value="value"
      :min="field.min"
      :max="field.max"
      :placeholder="field.placeholder || `请输入${field.label}`"
      class="w-full"
    />
    <n-select
      v-else-if="field.type === 'select'"
      v-model:value="value"
      :options="field.options"
      :placeholder="field.placeholder || `请选择${field.label}`"
    />
    <n-radio-group v-else-if="field.type === 'radio'" v-model:value="value">
      <n-radio v-for="opt in field.options" :key="opt.value" :value="opt.value">{{
        opt.label
      }}</n-radio>
    </n-radio-group>
    <n-switch v-else-if="field.type === 'switch'" v-model:value="value" />
    <n-color-picker v-else-if="field.type === 'color'" v-model:value="value" />
    <n-date-picker
      v-else-if="field.type === 'date'"
      v-model:value="value"
      type="datetime"
      clearable
      class="w-full"
    />
    <div v-if="field.help" class="mt-1 text-xs text-gray-400">{{ field.help }}</div>
  </div>
</template>
