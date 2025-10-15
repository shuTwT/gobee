<script setup lang="ts">
const albumFormData = ref({
  name: '',
})
</script>
<template>
  <div class="p-4 sm:p-6 lg:p-8">
    <h2 class="text-2xl font-semibold text-gray-900 dark:text-white mb-6">相册管理</h2>

    <div class="flex flex-col md:flex-row gap-6">
      <!-- 左侧相册列表 -->
      <div class="w-full md:w-1/3 bg-white dark:bg-gray-800 shadow-md rounded-lg p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">相册列表</h3>
          <n-button id="addAlbumBtn" type="primary">
            <i class="fas fa-plus mr-2"></i>新增相册
          </n-button>
        </div>
        <ul id="albumList" class="space-y-2">
          <!-- 示例相册 -->
          <li
            class="flex justify-between items-center p-2 bg-gray-50 dark:bg-gray-700 rounded-md cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-600 album-item"
            data-album-id="1"
          >
            <span class="text-gray-800 dark:text-gray-200">我的照片</span>
            <span class="text-sm text-gray-500 dark:text-gray-400">50 张</span>
          </li>
          <li
            class="flex justify-between items-center p-2 bg-gray-50 dark:bg-gray-700 rounded-md cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-600 album-item"
            data-album-id="2"
          >
            <span class="text-gray-800 dark:text-gray-200">旅行回忆</span>
            <span class="text-sm text-gray-500 dark:text-gray-400">120 张</span>
          </li>
        </ul>
      </div>

      <!-- 右侧照片列表 -->
      <div class="w-full md:w-2/3 bg-white dark:bg-gray-800 shadow-md rounded-lg p-4">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">照片列表</h3>
          <button
            id="uploadPhotoBtn"
            class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 dark:focus:ring-offset-gray-800"
          >
            <i class="fas fa-upload mr-2"></i>上传照片
          </button>
        </div>
        <div id="photoList" class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-4">
          <!-- 示例照片卡片 -->
          <div class="bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm overflow-hidden">
            <img
              src="https://via.placeholder.com/150"
              alt="Photo 1"
              class="w-full h-32 object-cover"
            />
            <div class="p-3">
              <p class="text-sm font-medium text-gray-800 dark:text-gray-200 truncate">
                照片 1.jpg
              </p>
              <p class="text-xs text-gray-500 dark:text-gray-400">2023-01-01</p>
            </div>
          </div>
          <div class="bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm overflow-hidden">
            <img
              src="https://via.placeholder.com/150"
              alt="Photo 2"
              class="w-full h-32 object-cover"
            />
            <div class="p-3">
              <p class="text-sm font-medium text-gray-800 dark:text-gray-200 truncate">
                照片 2.png
              </p>
              <p class="text-xs text-gray-500 dark:text-gray-400">2023-01-02</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 新增相册模态框 -->
    <n-modal>
      <div
        class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800"
      >
        <h3 class="text-lg font-medium leading-6 text-gray-900 dark:text-white mb-4">新增相册</h3>
        <n-form>
          <n-form-item label="相册名称">
            <n-input v-model:value="albumFormData.name" />
          </n-form-item>
        </n-form>

        <div class="items-center px-4 py-3 mt-4">
          <n-button id="saveAlbumBtn" type="primary"> 保存 </n-button>
          <n-button id="cancelAlbumBtn" type="default"> 取消 </n-button>
        </div>
      </div>
    </n-modal>

    <!-- 上传照片模态框 -->
    <n-modal>
      <div
        class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800"
      >
        <h3 class="text-lg font-medium leading-6 text-gray-900 dark:text-white mb-4">上传照片</h3>
        <div class="mt-2">
          <input
            type="file"
            id="photoFile"
            accept="image/*"
            class="mt-1 block w-full text-gray-900 dark:text-white"
          />
        </div>
        <div class="items-center px-4 py-3 mt-4">
          <n-button id="confirmUploadBtn" type="primary"> 上传 </n-button>
          <n-button id="cancelUploadBtn" type="default"> 取消 </n-button>
        </div>
      </div>
    </n-modal>
  </div>
</template>
