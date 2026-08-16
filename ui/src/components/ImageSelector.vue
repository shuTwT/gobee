<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { NUpload, NImage, NButton, NInput, NTag, NPagination, NIcon } from 'naive-ui';
import { Search, Refresh, CloudUploadOutline } from '@vicons/ionicons5';
import { apiClient, useApi } from '@/api';

import type { EntStorageStrategy } from '@hoshikuzu/api-client';

const props = withDefaults(defineProps<{
  visible: boolean;
  modelValue: string | string[];
  limit?: number;
}>(), {
  limit: 1
});

const emit = defineEmits<{
  (e: 'update:modelValue', value: string | string[]): void;
}>();

// 状态管理
const searchKeyword = ref('');
const storageFilter = ref('all');

// 分页
const currentPage = ref(1);
const pageSize = ref(60);
const total = ref(0);

// 已选择的图片
const selectedImages = ref<string[]>([]);

// 存储策略列表
const storageStrategies = ref<EntStorageStrategy[]>([]);

// 图片列表
const imageList = ref<any[]>([]);

// 加载状态
const loading = ref(false);

// 获取存储策略列表
const fetchStorageStrategies = async () => {
  try {
    const response = await useApi(apiClient.api.v1StorageStrategyListList);
    if (response.code === 200) {
      storageStrategies.value = (response.data || []) as EntStorageStrategy[];
    }
  } catch (error) {
    console.error('获取存储策略失败:', error);
  }
};

// 获取图片列表
const fetchImages = async () => {
  loading.value = true;
  try {
    const params: any = {
      page: currentPage.value,
      page_size: pageSize.value,
      type: 'image'
    };

    if (searchKeyword.value) {
      params.name = searchKeyword.value;
    }

    if (storageFilter.value !== 'all') {
      params.storage_strategy_id = storageFilter.value;
    }

    const response = await useApi(apiClient.api.v1FilePageList, params);
    if (response.code === 200) {
      imageList.value = response.data?.records || [];
      total.value = response.data?.total || 0;
    }
  } catch (error) {
    console.error('获取图片列表失败:', error);
    imageList.value = [];
    total.value = 0;
  } finally {
    loading.value = false;
  }
};

// 刷新图片列表
const refreshImages = () => {
  currentPage.value = 1;
  fetchImages();
};

// 监听分页和存储策略变化
watch([currentPage, pageSize], () => {
  fetchImages();
});

// 存储策略切换时回到第一页
watch(storageFilter, () => {
  currentPage.value = 1;
  fetchImages();
});

// 关键词搜索（防抖 300ms）
let searchTimer: number | undefined;
watch(searchKeyword, () => {
  window.clearTimeout(searchTimer);
  searchTimer = window.setTimeout(() => {
    currentPage.value = 1;
    fetchImages();
  }, 300);
});

// 外部传入的已选值同步到选中状态
watch(
  () => props.modelValue,
  (val) => {
    if (props.limit === 1) {
      selectedImages.value = typeof val === 'string' && val ? [val] : [];
    } else {
      selectedImages.value = Array.isArray(val) ? [...val] : [];
    }
  },
  { immediate: true },
);

// 组件挂载时获取数据
onMounted(() => {
  fetchStorageStrategies();
  fetchImages();
});

// 处理图片选择
const handleImageSelect = (imageUrl: string) => {
  if (props.limit === 1) {
    // 单选模式
    selectedImages.value = [imageUrl];
    emit('update:modelValue', imageUrl);
  } else {
    // 多选模式
    const index = selectedImages.value.indexOf(imageUrl);
    if (index > -1) {
      selectedImages.value.splice(index, 1);
    } else if (selectedImages.value.length < props.limit) {
      selectedImages.value.push(imageUrl);
    }
    emit('update:modelValue', [...selectedImages.value]);
  }
};

// 处理上传完成
const handleUploadFinish = () => {
  refreshImages();
};
</script>

<template>
  <div v-if="visible" class="image-selector">
    <!-- 搜索 -->
    <div class="header">
      <div class="search-box">
        <n-input v-model:value="searchKeyword" placeholder="输入名称搜索" clearable size="small">
          <template #prefix>
            <n-icon><Search /></n-icon>
          </template>
        </n-input>
      </div>
      <div class="header-actions">
        <n-button quaternary circle size="small" @click="refreshImages">
          <template #icon>
            <n-icon><Refresh /></n-icon>
          </template>
        </n-button>
      </div>
    </div>

    <!-- 筛选条件 -->
    <div class="filters">
      <!-- 存储策略 -->
      <div class="filter-section">
        <div class="filter-label">存储策略:</div>
        <div class="filter-tags">
          <n-tag type="info" round checkable :checked="storageFilter === 'all'" @update:checked="storageFilter = 'all'">全部</n-tag>
          <n-tag 
            v-for="strategy in storageStrategies" 
            :key="strategy.id" 
            type="info" 
            round 
            checkable 
            :checked="storageFilter === String(strategy.id)" 
            @update:checked="storageFilter = String(strategy.id)"
          >
            {{ strategy.name }}
          </n-tag>
        </div>
      </div>
    </div>

    <!-- 上传 -->
    <div class="upload-section">
      <n-upload
        accept="image/*"
        action="/api/v1/file/upload"
        :show-file-list="false"
        @finish="handleUploadFinish"
      >
        <n-button type="primary" size="small">
          <template #icon>
            <n-icon><CloudUploadOutline /></n-icon>
          </template>
          上传
        </n-button>
      </n-upload>
    </div>

    <!-- 已选择提示 -->
    <div class="selection-info">
      已选择 {{ selectedImages.length }}/{{ limit }} 项
    </div>

    <!-- 图片网格 -->
    <div class="image-grid">
      <div 
        v-for="image in imageList" 
        :key="image.id"
        class="image-item"
        :class="{ selected: selectedImages.includes(image.url) }"
        @click="handleImageSelect(image.url)"
      >
        <n-image
          :src="image.url"
          fit="cover"
          width="100%"
          height="100%"
          :preview-disabled="true"
        />
        <div class="image-mask">
          <div class="image-name">{{ image.name }}</div>
        </div>
        <div class="image-check" v-if="selectedImages.includes(image.url)">
          ✓
        </div>
      </div>
      <div v-if="loading" class="loading-mask">
        <span>加载中...</span>
      </div>
      <div v-if="!loading && imageList.length === 0" class="empty-mask">
        <span>暂无图片</span>
      </div>
    </div>

    <!-- 分页 -->
    <div class="pagination">
      <div class="total">
        共 {{ total }} 项数据
      </div>
      <div class="pagination-controls">
        <n-pagination
          v-model:page="currentPage"
          v-model:page-size="pageSize"
          :page-count="Math.ceil(total / pageSize)"
          :page-sizes="[20, 40, 60, 100]"
          show-size-picker
          show-quick-jumper
          show-total
          :total="total"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.image-selector {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* 头部搜索 */
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.search-box {
  flex: 1;
  max-width: 400px;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

/* 筛选条件 */
.filters {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 16px;
}

.filter-section {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filter-label {
  font-size: 14px;
  color: #6b7280;
  min-width: 60px;
}

.filter-tags {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  align-items: center;
}

/* 上传按钮 */
.upload-section {
  margin-bottom: 16px;
}

/* 已选择提示 */
.selection-info {
  font-size: 13px;
  color: #6b7280;
  margin-bottom: 8px;
}

/* 图片网格 */
.image-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 12px;
  overflow-y: auto;
  flex: 1;
  margin-bottom: 16px;
  padding: 8px;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
}

.image-item {
  position: relative;
  aspect-ratio: 1;
  border: 1px solid #e5e7eb;
  border-radius: 4px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s;
}

.image-item:hover {
  border-color: #3b82f6;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.2);
}

.image-item.selected {
  border-color: #3b82f6;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}

.image-mask {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  padding: 4px 8px;
  font-size: 12px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  opacity: 0;
  transition: opacity 0.2s;
}

.image-item:hover .image-mask {
  opacity: 1;
}

.image-name {
  font-size: 12px;
}

.image-check {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 20px;
  height: 20px;
  background: #3b82f6;
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: bold;
}

.loading-mask,
.empty-mask {
  grid-column: 1 / -1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px;
  color: #9ca3af;
  font-size: 14px;
}

/* 分页 */
.pagination {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 0 8px;
}

.total {
  font-size: 14px;
  color: #6b7280;
}
</style>
