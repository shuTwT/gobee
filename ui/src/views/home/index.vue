<script setup lang="ts">
import { apiClient, useApi } from "@/api"
import { getNotificationPage } from "@/api/system/notification"
import { RouterLink } from "vue-router"
import dayjs from "dayjs"
import relativeTime from "dayjs/plugin/relativeTime"
import "dayjs/locale/zh-cn"
import {
  DocumentTextOutline,
  PeopleOutline,
  ChatbubblesOutline,
  EyeOutline,
  TimeOutline,
  NotificationsOffOutline,
  InformationCircleOutline,
  CheckmarkCircleOutline,
  WarningOutline,
  AlertCircleOutline,
  NotificationsOutline,
  FolderOutline,
  SettingsOutline,
  CartOutline,
} from "@vicons/ionicons5"

dayjs.extend(relativeTime)
dayjs.locale("zh-cn")

const homeStatistic = ref({
  commentCount: 0,
  postCount: 0,
  userCount: 0,
  visitCount: 0,
})

const notificationList = ref<any[]>([])
const notificationLoading = ref(false)
const unreadCount = ref(0)

type quickLink = {
  name: string
  icon: string
  color: "blue" | "green" | "yellow" | "purple" | "red" | "orange"
  path: string
}

const quickLinks: quickLink[] = [
  {
    name: "文章管理",
    icon: "document-text-outline",
    color: "blue",
    path: "/content/post",
  },
  {
    name: "用户管理",
    icon: "people-outline",
    color: "green",
    path: "/system/user",
  },
  {
    name: "文件管理",
    icon: "folder-outline",
    color: "yellow",
    path: "/infra/file",
  },
  {
    name: "AI助手",
    icon: "chatbubbles-outline",
    color: "purple",
    path: "/ai/chat",
  },
  {
    name: "商品列表",
    icon: "cart-outline",
    color: "red",
    path: "/mall/product",
  },
  {
    name: "通知管理",
    icon: "notifications-outline",
    color: "orange",
    path: "/system/notification",
  },
]

const loadNotifications = async () => {
  notificationLoading.value = true
  try {
    const res = await getNotificationPage({
      page: 1,
      page_size: 10,
      is_read: false,
    })
    if (res.code === 200) {
      notificationList.value = res.data.records || []
      unreadCount.value = res.data.total || 0
    }
  } catch (error) {
    console.error("加载通知失败:", error)
  } finally {
    notificationLoading.value = false
  }
}

const formatTime = (time: string) => {
  if (!time) return "-"
  return dayjs(time).fromNow()
}

const getNotificationIcon = (type: string) => {
  switch (type) {
    case "info":
      return InformationCircleOutline
    case "success":
      return CheckmarkCircleOutline
    case "warning":
      return WarningOutline
    case "error":
      return AlertCircleOutline
    default:
      return NotificationsOutline
  }
}

const getIcon = (iconName: string) => {
  const iconMap: Record<string, any> = {
    "document-text-outline": DocumentTextOutline,
    "people-outline": PeopleOutline,
    "folder-outline": FolderOutline,
    "chatbubbles-outline": ChatbubblesOutline,
    "cart-outline": CartOutline,
    "notifications-outline": NotificationsOutline,
  }
  return iconMap[iconName] || NotificationsOutline
}

const getNotificationColor = (type: string) => {
  switch (type) {
    case "info":
      return "blue"
    case "success":
      return "green"
    case "warning":
      return "yellow"
    case "error":
      return "red"
    default:
      return "gray"
  }
}

onMounted(() => {
  useApi(apiClient.api.v1CommonStatisticList).then((res) => {
    if (res.code === 200) {
      homeStatistic.value.commentCount = res.data.commentCount
      homeStatistic.value.postCount = res.data.postCount
      homeStatistic.value.userCount = res.data.userCount
      homeStatistic.value.visitCount = res.data.visitCount
    }
  })
  loadNotifications()
})
</script>

<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900">
    <div class="p-6 max-w-[1600px] mx-auto">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
          欢迎回来
        </h1>
        <p class="text-gray-600 dark:text-gray-400">
          查看系统概览和最新动态
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div
          class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 p-6 transition-all duration-300 hover:shadow-md hover:scale-[1.02]">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">
                文章数
              </p>
              <p class="text-3xl font-bold text-gray-900 dark:text-white">
                {{ homeStatistic.postCount }}
              </p>
            </div>
            <div class="w-14 h-14 bg-blue-100 dark:bg-blue-900/30 rounded-xl flex items-center justify-center">
              <n-icon :component="DocumentTextOutline" :size="28" class="text-blue-600 dark:text-blue-400" />
            </div>
          </div>
        </div>

        <div
          class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 p-6 transition-all duration-300 hover:shadow-md hover:scale-[1.02]">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">
                用户数
              </p>
              <p class="text-3xl font-bold text-gray-900 dark:text-white">
                {{ homeStatistic.userCount }}
              </p>
            </div>
            <div class="w-14 h-14 bg-green-100 dark:bg-green-900/30 rounded-xl flex items-center justify-center">
              <n-icon :component="PeopleOutline" :size="28" class="text-green-600 dark:text-green-400" />
            </div>
          </div>
        </div>

        <div
          class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 p-6 transition-all duration-300 hover:shadow-md hover:scale-[1.02]">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">
                评论数
              </p>
              <p class="text-3xl font-bold text-gray-900 dark:text-white">
                {{ homeStatistic.commentCount }}
              </p>
            </div>
            <div class="w-14 h-14 bg-yellow-100 dark:bg-yellow-900/30 rounded-xl flex items-center justify-center">
              <n-icon :component="ChatbubblesOutline" :size="28" class="text-yellow-600 dark:text-yellow-400" />
            </div>
          </div>
        </div>

        <div
          class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 p-6 transition-all duration-300 hover:shadow-md hover:scale-[1.02]">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">
                浏览量
              </p>
              <p class="text-3xl font-bold text-gray-900 dark:text-white">
                {{ homeStatistic.visitCount }}
              </p>
            </div>
            <div class="w-14 h-14 bg-purple-100 dark:bg-purple-900/30 rounded-xl flex items-center justify-center">
              <n-icon :component="EyeOutline" :size="28" class="text-purple-600 dark:text-purple-400" />
            </div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <n-card title="快捷访问" class="shadow-sm border-gray-200 dark:border-gray-700">
          <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
            <RouterLink v-for="link in quickLinks" :key="link.name" :to="link.path"
              class="group flex flex-col items-center justify-center p-6 bg-gray-50 dark:bg-gray-700/50 rounded-xl border border-gray-200 dark:border-gray-600 hover:border-gray-300 dark:hover:border-gray-500 transition-all duration-200 hover:shadow-md hover:-translate-y-1 cursor-pointer">
              <div
                class="w-12 h-12 rounded-xl flex items-center justify-center mb-3 group-hover:scale-110 transition-transform duration-200"
                :class="{
                  'bg-blue-100 dark:bg-blue-900/30': link.color === 'blue',
                  'bg-green-100 dark:bg-green-900/30': link.color === 'green',
                  'bg-yellow-100 dark:bg-yellow-900/30': link.color === 'yellow',
                  'bg-purple-100 dark:bg-purple-900/30': link.color === 'purple',
                  'bg-red-100 dark:bg-red-900/30': link.color === 'red',
                  'bg-orange-100 dark:bg-orange-900/30': link.color === 'orange',
                }">
                <n-icon :component="getIcon(link.icon)" :size="24" :class="{
                  'text-blue-600 dark:text-blue-400': link.color === 'blue',
                  'text-green-600 dark:text-green-400': link.color === 'green',
                  'text-yellow-600 dark:text-yellow-400': link.color === 'yellow',
                  'text-purple-600 dark:text-purple-400': link.color === 'purple',
                  'text-red-600 dark:text-red-400': link.color === 'red',
                  'text-orange-600 dark:text-orange-400': link.color === 'orange',
                }" />
              </div>
              <span
                class="text-sm font-medium text-gray-700 dark:text-gray-200 group-hover:text-gray-900 dark:group-hover:text-white transition-colors">
                {{ link.name }}
              </span>
            </RouterLink>
          </div>
        </n-card>

        <n-card title="最新通知" class="shadow-sm border-gray-200 dark:border-gray-700">
          <template #header-extra>
            <RouterLink to="/system/notification"
              class="text-sm text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300 transition-colors">
              查看全部
            </RouterLink>
          </template>

          <n-spin :show="notificationLoading">
            <div v-if="notificationList.length > 0" class="space-y-3">
              <div v-for="notification in notificationList" :key="notification.id"
                class="flex items-start gap-3 p-4 bg-gray-50 dark:bg-gray-700/30 rounded-lg border border-gray-200 dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700/50 transition-colors cursor-pointer">
                <div
                  :class="`w-10 h-10 bg-${getNotificationColor(notification.type)}-100 dark:bg-${getNotificationColor(notification.type)}-900/30 rounded-lg flex items-center justify-center flex-shrink-0`">
                  <n-icon :component="getNotificationIcon(notification.type)" :size="20"
                    :class="`text-${getNotificationColor(notification.type)}-600 dark:text-${getNotificationColor(notification.type)}-400`" />
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-medium text-gray-900 dark:text-white mb-1 line-clamp-2">
                    {{ notification.title }}
                  </p>
                  <p v-if="notification.content" class="text-xs text-gray-600 dark:text-gray-400 mb-2 line-clamp-2">
                    {{ notification.content }}
                  </p>
                  <span class="text-xs text-gray-500 dark:text-gray-500 flex items-center gap-1">
                    <n-icon :component="TimeOutline" :size="12" />
                    {{ formatTime(notification.created_at) }}
                  </span>
                </div>
              </div>
            </div>

            <n-empty v-else description="暂无新通知" class="py-8">
              <template #icon>
                <n-icon :component="NotificationsOffOutline" :size="48" />
              </template>
            </n-empty>
          </n-spin>
        </n-card>
      </div>
    </div>
  </div>
</template>
