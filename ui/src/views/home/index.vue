<script setup lang="ts">
import * as commonApi from '@/api/common'
const homeStatistic = ref({
  commentCount: 0,
  postCount: 0,
  userCount: 0,
  visitCount: 0,
})
const noticeList = ref([
  {
    title: '新文章 "如何使用 Tailwind CSS" 已发布。',
    time: '11天前',
  },
  {
    title: '新用户注册',
    time: '10天前',
  },
])

onMounted(() => {
  commonApi.getHomeStatistic().then((res) => {
    if (res.code === 200) {
      homeStatistic.value.commentCount = res.data.commentCount
      homeStatistic.value.postCount = res.data.postCount
      homeStatistic.value.userCount = res.data.userCount
      homeStatistic.value.visitCount = res.data.visitCount
    }
  })
})
</script>
<template>
  <div class="p-6">
    <div class="p-6 max-w-[1600px] mx-auto">
      <!-- 统计概览 -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div
          class="bg-blue-100 dark:bg-blue-900 p-5 rounded-lg shadow-sm flex items-center justify-between"
        >
          <div>
            <p class="text-sm font-medium text-blue-600 dark:text-blue-300">文章数</p>
            <p class="text-3xl font-bold text-blue-800 dark:text-blue-100" id="articleCount">
              {{ homeStatistic.postCount }}
            </p>
          </div>
          <i class="fas fa-newspaper text-blue-500 dark:text-blue-400 text-4xl"></i>
        </div>
        <div
          class="bg-green-100 dark:bg-green-900 p-5 rounded-lg shadow-sm flex items-center justify-between"
        >
          <div>
            <p class="text-sm font-medium text-green-600 dark:text-green-300">用户数</p>
            <p class="text-3xl font-bold text-green-800 dark:text-green-100" id="userCount">
              {{ homeStatistic.userCount }}
            </p>
          </div>
          <i class="fas fa-users text-green-500 dark:text-green-400 text-4xl"></i>
        </div>
        <div
          class="bg-yellow-100 dark:bg-yellow-900 p-5 rounded-lg shadow-sm flex items-center justify-between"
        >
          <div>
            <p class="text-sm font-medium text-yellow-600 dark:text-yellow-300">评论数</p>
            <p class="text-3xl font-bold text-yellow-800 dark:text-yellow-100" id="commentCount">
              {{ homeStatistic.commentCount }}
            </p>
          </div>
          <i class="fas fa-comments text-yellow-500 dark:text-yellow-400 text-4xl"></i>
        </div>
        <div
          class="bg-purple-100 dark:bg-purple-900 p-5 rounded-lg shadow-sm flex items-center justify-between"
        >
          <div>
            <p class="text-sm font-medium text-purple-600 dark:text-purple-300">浏览量</p>
            <p class="text-3xl font-bold text-purple-800 dark:text-purple-100" id="pageViews">
              {{ homeStatistic.visitCount }}
            </p>
          </div>
          <i class="fas fa-eye text-purple-500 dark:text-purple-400 text-4xl"></i>
        </div>
      </div>

      <!-- 快捷访问与通知容器 -->
      <div class="flex flex-col md:flex-row gap-6">
        <!-- 快捷访问 -->
        <div class="w-full md:w-1/2">
          <n-card title="快捷访问">
            <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
              <a
                href="/console/model-content"
                class="flex flex-col items-center justify-center p-4 bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
              >
                <i class="fas fa-edit text-blue-500 text-2xl mb-2"></i>
                <span class="text-sm font-medium text-gray-700 dark:text-gray-200">内容管理</span>
              </a>
              <a
                href="/console/user"
                class="flex flex-col items-center justify-center p-4 bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
              >
                <i class="fas fa-users-cog text-green-500 text-2xl mb-2"></i>
                <span class="text-sm font-medium text-gray-700 dark:text-gray-200">用户管理</span>
              </a>
              <a
                href="/console/file"
                class="flex flex-col items-center justify-center p-4 bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
              >
                <i class="fas fa-folder-open text-yellow-500 text-2xl mb-2"></i>
                <span class="text-sm font-medium text-gray-700 dark:text-gray-200">文件管理</span>
              </a>
              <a
                href="/console/webhook"
                class="flex flex-col items-center justify-center p-4 bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
              >
                <i class="fas fa-webhook text-purple-500 text-2xl mb-2"></i>
                <span class="text-sm font-medium text-gray-700 dark:text-gray-200">Webhook</span>
              </a>
              <a
                href="/console/model-schema"
                class="flex flex-col items-center justify-center p-4 bg-gray-50 dark:bg-gray-700 rounded-lg shadow-sm hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
              >
                <i class="fas fa-cogs text-red-500 text-2xl mb-2"></i>
                <span class="text-sm font-medium text-gray-700 dark:text-gray-200">模型管理</span>
              </a>
            </div>
          </n-card>
        </div>

        <!-- 最新通知 -->
        <div class="w-full md:w-1/2">
          <n-card title="最新通知">
            <n-list>
              <!-- 最多显示10条通知 -->
              <n-list-item v-for="(item, index) in noticeList" :key="index">
                <p>{{ item.title }}</p>
                <span>{{ item.time }}</span>
              </n-list-item>
            </n-list>
          </n-card>
        </div>
      </div>
    </div>
  </div>
</template>
