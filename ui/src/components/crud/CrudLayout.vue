<script setup lang="ts">
import { ProSearchForm, type ProSearchFormColumns } from 'pro-naive-ui';
import type { CreateProSearchFormReturn } from './types';

const props = defineProps<{
  loading:false,
  form:CreateProSearchFormReturn<any>
}>()



const searchColumns:ProSearchFormColumns<any>=[] as any[]
</script>
<template>
  <div>
    <!-- 头部操作栏 -->
    <div class="header-section">
      <pro-search-form :form="form" :loading="false" :columns="searchColumns"/>
    </div>
    <!-- 内容 -->
     <n-data-table/>
  </div>
</template>
<style scoped>
</style>
