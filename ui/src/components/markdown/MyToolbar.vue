<script setup lang="ts">
import type { PropType } from 'vue';
import { NormalToolbar } from 'md-editor-v3';
import type { Insert } from 'md-editor-v3';
import GoodsIcon from '@/assets/imgs/svg/goods.svg?component'

const props = defineProps({
  /**
   * `insert`方法会由编辑器自动向组件的组件注入。
   */
  insert: {
    type: Function as PropType<Insert>,
    default: () => {
      //
    },
  },
});

const handler = () => {
  props.insert((selectedText) => {
    /**
     * targetValue    待插入内容
     * select         插入后是否自动选中内容，默认：true
     * deviationStart 插入后选中内容鼠标开始位置，默认：0
     * deviationEnd   插入后选中内容鼠标结束位置，默认：0
     */
    return {
      targetValue: `\`\`\`product-card
name ${selectedText}
price 100
description 测试
\`\`\``,
      select: false,
      deviationStart: 0,
      deviationEnd: 0,
    };
  });
};
</script>
<template>
    <NormalToolbar title="商品" @onClick="handler">
      <GoodsIcon class="md-editor-icon" aria-hidden="true" />
      <div class="md-editor-toolbar-item-name">商品</div>
  </NormalToolbar>
</template>