<script setup lang="ts">
import { NButton, NInput } from 'naive-ui'
import { CopyOutline } from '@vicons/ionicons5'
import { h, ref } from 'vue'

const props = defineProps<{
  formInline: {
    title: string
    shareUrl: string
  }
}>()

const emit = defineEmits<{
  (e: 'confirm'): void
}>()

const message = useMessage()

const copyUrl = async () => {
  try {
    await navigator.clipboard.writeText(props.formInline.shareUrl)
    message.success('链接已复制到剪贴板')
  } catch {
    // 降级方案
    const textarea = document.createElement('textarea')
    textarea.value = props.formInline.shareUrl
    textarea.style.position = 'fixed'
    textarea.style.opacity = '0'
    document.body.appendChild(textarea)
    textarea.select()
    try {
      document.execCommand('copy')
      message.success('链接已复制到剪贴板')
    } catch {
      message.error('复制失败，请手动复制')
    }
    document.body.removeChild(textarea)
  }
}

const getData = () => {
  emit('confirm')
  return Promise.resolve()
}

defineExpose({
  getData,
})
</script>

<template>
  <div class="share-dialog">
    <div class="share-title">{{ formInline.title }}</div>
    <div class="share-label">分享链接</div>
    <div class="share-url-row">
      <n-input :value="formInline.shareUrl" readonly />
      <n-button type="primary" @click="copyUrl">
        <template #icon>
          <n-icon :component="CopyOutline" />
        </template>
        复制链接
      </n-button>
    </div>
    <div class="share-tip">将链接分享给好友，对方即可访问该文章</div>
  </div>
</template>

<style scoped lang="scss">
.share-dialog {
  .share-title {
    font-size: 16px;
    font-weight: 600;
    margin-bottom: 16px;
    color: var(--n-text-color);
  }

  .share-label {
    font-size: 13px;
    color: var(--n-text-color-2);
    margin-bottom: 8px;
  }

  .share-url-row {
    display: flex;
    gap: 8px;
    margin-bottom: 12px;
  }

  .share-tip {
    font-size: 12px;
    color: var(--n-text-color-3);
  }
}
</style>
