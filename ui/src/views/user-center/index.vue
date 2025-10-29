<script setup lang="ts">
const colorOptions = ref([
  {
    value:'light',
    label:'亮色'
  },
  {
    value:'dark',
    label:'暗色'
  }
])
const personalInfomation = ref({
  username:'',
  email:'',
  nickname:'',
  bio:''
})
const personalizedSetting = ref({
  theme:'light',
  notifications:false
})
</script>
<template>
  <div class="p-4 sm:p-6 lg:p-8">
    <h2 class="text-2xl font-semibold text-gray-900 dark:text-white mb-6">用户中心</h2>

    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6 mb-6">
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">个人信息</h3>
        <n-form>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                  <n-form-item label="用户名">
                    <n-input v-model:value="personalInfomation.username"/>
                  </n-form-item>
                  </div>
                <div>
                  <n-form-item label="邮箱">
                    <n-input v-model:value="personalInfomation.email"/>
                  </n-form-item>
                </div>
            </div>
            <div class="mb-4">
                <n-form-item label="昵称">
                  <n-input v-model:value="personalInfomation.nickname"/>
                </n-form-item>
            </div>
            <div class="mb-4">
                <n-form-item label="个人简介">
                  <n-textarea v-model:value="personalInfomation.bio"/>
                </n-form-item>
            </div>
            <div class="flex justify-end">
                <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-gray-800">
                    保存个人信息
                </button>
            </div>
        </n-form>
    </div>

    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6">
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">个性化设置</h3>
        <n-form id="settingsForm">
            <div class="mb-4">
                <n-form-item label="主题">
                  <n-select v-model:value="personalizedSetting.theme" :options="colorOptions"/>
                </n-form-item>
            </div>
            <div class="mb-4">
                <n-form-item label="接收通知">
                  <n-checkbox v-model:checked="personalizedSetting.notifications"/>
                </n-form-item>
            </div>
            <div class="flex justify-end">
                <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-gray-800">
                    保存设置
                </button>
            </div>
        </n-form>
    </div>
</div>
  </template>
