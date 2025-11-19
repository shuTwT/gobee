<script setup lang="ts">
import { NDescriptions,NDescriptionsItem } from 'naive-ui'
const colorOptions = ref([
  {
    value: 'light',
    label: '亮色',
  },
  {
    value: 'dark',
    label: '暗色',
  },
])
const personalInfomation = ref({
  username: '',
  email: '',
  nickname: '',
  bio: '',
})
const personalizedSetting = ref({
  theme: 'light',
  notifications: false,
})
const personalAccessTokenList = ref([])
const personalAccessTokenColumns = ref([
  {
    title: '令牌名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '令牌值',
    dataIndex: 'token',
    key: 'token',
  },
])

const activeTab = ref('profile')
</script>
<template>
  <div class="p-4 sm:p-6 lg:p-8">
    <h2 class="text-2xl font-semibold text-gray-900 dark:text-white mb-6">用户中心</h2>

    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6 mb-6">
      <n-tabs v-model:value="activeTab" type="segment" animated>
        <n-tab-pane name="profile" tab="详情"></n-tab-pane>
        <n-tab-pane name="setting" tab="个性化"></n-tab-pane>
        <n-tab-pane name="personalAccessToken" tab="个人令牌"></n-tab-pane>
      </n-tabs>
    </div>
    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6 mb-6">
      <template v-if="activeTab === 'profile'">
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">个人信息</h3>

        <n-descriptions label-placement="left" bordered :columns="1">
          <n-descriptions-item>
            <template #label>
              用户名
            </template>
            {{ personalInfomation.username }}
          </n-descriptions-item>
          <n-descriptions-item>
            <template #label>
              邮箱
            </template>
            {{ personalInfomation.email }}
          </n-descriptions-item>
          <n-descriptions-item>
            <template #label>
              昵称
            </template>
            {{ personalInfomation.nickname }}
          </n-descriptions-item>
          <n-descriptions-item>
            <template #label>
              个人简介
            </template>
            {{ personalInfomation.bio }}
          </n-descriptions-item>
        </n-descriptions>
      </template>
      <template v-else-if="activeTab == 'setting'">
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">个性化设置</h3>
        <n-form label-placement="left">
          <div class="mb-4">
            <n-form-item label="主题">
              <n-select v-model:value="personalizedSetting.theme" :options="colorOptions" />
            </n-form-item>
          </div>
          <div class="mb-4">
            <n-form-item label="接收通知">
              <n-checkbox v-model:checked="personalizedSetting.notifications" />
            </n-form-item>
          </div>
          <div class="flex justify-end">
            <button
              type="submit"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-gray-800"
            >
              保存设置
            </button>
          </div>
        </n-form>
      </template>
      <template v-else-if="activeTab === 'personalAccessToken'">
        <h3 class="text-xl font-semibold text-gray-900 dar
        k:text-white mb-4">个人令牌</h3>
        <div class="flex justify-between">
          <p class="text-gray-700 dark:text-gray-300 mb-4">
            个人令牌用于访问 API。请妥善保管，不要分享给他人。
          </p>
          <div>
            <n-button type="primary">创建</n-button>
          </div>
        </div>

        <n-data-table :columns="personalAccessTokenColumns" :data="personalAccessTokenList" />
      </template>
    </div>
  </div>
</template>
