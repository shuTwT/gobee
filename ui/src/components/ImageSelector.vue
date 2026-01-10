<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue';
import { NUpload, NImage, NButton, NInput, NSelect, NTag, NPagination } from 'naive-ui';
import { Search, Refresh, Grid, List } from '@vicons/ionicons5';
import { getStorageStrategyListAll } from '@/api/infra/storage';
import { getFilePage } from '@/api/infra/file';
import type { StorageStrategy } from '@/api/infra/storage';

const props = withDefaults(defineProps<{
  visible: boolean;
  modelValue: string;
  limit?: number;
}>(), {
  limit: 1
});

const emit = defineEmits<{
  (e: 'update:visible', visible: boolean): void;
  (e: 'update:modelValue', value: string): void;
  (e: 'confirm', value: string): void;
  (e: 'cancel'): void;
}>();

// 状态管理
const activeTab = ref('attachment'); // attachment 或 imageStream
const searchKeyword = ref('');
const sortOption = ref('default');
const viewMode = ref('grid'); // grid 或 list
const storageFilter = ref('all');

// 分页
const currentPage = ref(1);
const pageSize = ref(60);
const total = ref(0);

// 已选择的图片
const selectedImages = ref<string[]>([]);

// 存储策略列表
const storageStrategies = ref<StorageStrategy[]>([]);

// 图片列表
const imageList = ref<any[]>([]);

// 加载状态
const loading = ref(false);

// 获取存储策略列表
const fetchStorageStrategies = async () => {
  try {
    const response = await getStorageStrategyListAll();
    if (response.code === 200) {
      storageStrategies.value = response.data || [];
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
      params.keyword = searchKeyword.value;
    }
    
    if (storageFilter.value !== 'all') {
      params.storage_strategy_id = storageFilter.value;
    }
    
    const response = await getFilePage(params);
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

// 监听分页和筛选条件变化
watch([currentPage, pageSize, searchKeyword, storageFilter], () => {
  fetchImages();
});

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
  } else {
    // 多选模式
    const index = selectedImages.value.indexOf(imageUrl);
    if (index > -1) {
      selectedImages.value.splice(index, 1);
    } else if (selectedImages.value.length < props.limit) {
      selectedImages.value.push(imageUrl);
    }
  }
};

// 处理确定按钮点击
const handleConfirm = () => {
  if (selectedImages.value.length > 0) {
    emit('update:modelValue', selectedImages.value[0]);
    emit('confirm', selectedImages.value[0]);
    emit('update:visible', false);
  }
};

// 处理取消按钮点击
const handleCancel = () => {
  emit('update:visible', false);
  emit('cancel');
};

// 处理上传完成
const handleUploadFinish = () => {
  refreshImages();
};
</script>

<template>
  <div class="image-selector">
    <!-- 标签页 -->
    <div class="tabs">
      <div 
        class="tab-item" 
        :class="{ active: activeTab === 'attachment' }"
        @click="activeTab = 'attachment'"
      >
        附件库
      </div>
      <div 
        class="tab-item" 
        :class="{ active: activeTab === 'imageStream' }"
        @click="activeTab = 'imageStream'"
      >
        Image Stream
      </div>
    </div>

    <!-- 搜索和排序 -->
    <div class="header">
      <div class="search-box">
        <n-input
          v-model:value="searchKeyword"
          placeholder="输入关键词搜索"
          clearable
          prefix-icon="Search"
          size="small"
        />
      </div>
      <div class="header-actions">
        <div class="sort-box">
          <span>排序:</span>
          <n-select
            v-model:value="sortOption"
            :options="[{ label: '默认', value: 'default' }, { label: '最新', value: 'latest' }, { label: '最热', value: 'hottest' }]"
            size="small"
            style="width: 80px;"
          />
        </div>
        <n-button 
          quaternary 
          circle 
          size="small" 
          @click="viewMode = viewMode === 'grid' ? 'list' : 'grid'"
        >
          <component :is="viewMode === 'grid' ? List : Grid" />
        </n-button>
        <n-button 
          quaternary 
          circle 
          size="small" 
          @click="refreshImages"
        >
          <refresh />
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

    <!-- 上传按钮 -->
    <div class="upload-section">
      <n-button type="primary" size="small" prefix-icon="Upload">
        上传
      </n-button>
      <!-- 这里可以添加上传组件 -->
      <n-upload
        v-show="false"
        accept="image/*"
        action="/api/v1/file/upload"
        list-type="image"
        @finish="handleUploadFinish"
      />
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

    <!-- 底部按钮 -->
    <div class="footer">
      <div class="selection-info">
        确定 (已选择 {{ selectedImages.length }}/{{ limit }} 项)
      </div>
      <div class="footer-buttons">
        <n-button size="small" @click="handleCancel">取消</n-button>
        <n-button type="primary" size="small" @click="handleConfirm" :disabled="selectedImages.length === 0">确定</n-button>
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

/* 标签页 */
.tabs {
  display: flex;
  border-bottom: 1px solid #e5e7eb;
  margin-bottom: 16px;
}

.tab-item {
  padding: 8px 16px;
  cursor: pointer;
  color: #6b7280;
  font-size: 14px;
  border-bottom: 2px solid transparent;
  transition: all 0.2s;
}

.tab-item:hover {
  color: #3b82f6;
}

.tab-item.active {
  color: #3b82f6;
  border-bottom-color: #3b82f6;
  font-weight: 500;
}

/* 头部搜索和排序 */
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

.sort-box {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #6b7280;
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

/* 底部按钮 */
.footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 16px;
  border-top: 1px solid #e5e7eb;
}

.selection-info {
  font-size: 14px;
  color: #6b7280;
}

.footer-buttons {
  display: flex;
  gap: 8px;
}
</style>