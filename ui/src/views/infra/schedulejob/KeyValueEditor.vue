<script setup lang="ts">
import { ref, watch } from 'vue'
import { NButton, NInput, NSpace, NIcon } from 'naive-ui'
import { CloseOutline, AddOutline } from '@vicons/ionicons5'

const props = defineProps<{
  modelValue: Record<string, string>
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: Record<string, string>): void
}>()

interface HeaderItem {
  key: string
  value: string
}

const items = ref<HeaderItem[]>([])

watch(() => props.modelValue, (newVal) => {
  items.value = Object.entries(newVal || {}).map(([key, value]) => ({ key, value }))
}, { immediate: true, deep: true })

const addItem = () => {
  items.value.push({ key: '', value: '' })
}

const removeItem = (index: number) => {
  items.value.splice(index, 1)
}

const updateKey = (index: number, value: string) => {
  items.value[index].key = value
  emitUpdate()
}

const updateValue = (index: number, value: string) => {
  items.value[index].value = value
  emitUpdate()
}

const emitUpdate = () => {
  const result: Record<string, string> = {}
  items.value.forEach(item => {
    if (item.key && item.value) {
      result[item.key] = item.value
    }
  })
  emit('update:modelValue', result)
}

defineExpose({
  addItem,
})
</script>

<template>
  <div class="key-value-editor">
    <div v-for="(item, index) in items" :key="index" class="item-row">
      <n-input
        :value="item.key"
        placeholder="键"
        @update:value="(val) => updateKey(index, val)"
        class="key-input"
      />
      <n-input
        :value="item.value"
        placeholder="值"
        @update:value="(val) => updateValue(index, val)"
        class="value-input"
      />
      <n-button
        quaternary
        type="error"
        size="small"
        @click="removeItem(index)"
        class="remove-btn"
      >
        <template #icon>
          <n-icon>
            <close-outline />
          </n-icon>
        </template>
      </n-button>
    </div>
    <n-button dashed @click="addItem" class="add-btn">
      <template #icon>
        <n-icon>
          <add-outline />
        </n-icon>
      </template>
      添加请求头
    </n-button>
  </div>
</template>

<style scoped>
.key-value-editor {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.item-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.key-input {
  flex: 1;
}

.value-input {
  flex: 2;
}

.remove-btn {
  flex-shrink: 0;
}

.add-btn {
  width: 100%;
}
</style>
