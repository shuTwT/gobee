
<script setup lang="ts">
const webhookFormData = ref({
  webhookUrl:'',
  webhookEvents:[]
})
const eventOptions = ref([
  {
    value:'content.create',
    label:'内容创建'
  },
  {
    value:'content.update',
    label:'内容更新'
  }
])
</script>
<template>
  <div class="p-6 bg-white dark:bg-gray-800 shadow-md rounded-lg">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-semibold text-gray-900 dark:text-white">Webhook 管理</h2>
      <div class="flex space-x-4">
        <n-button> 添加 Webhook </n-button>
        <div class="relative">
          <input
            type="text"
            id="webhookSearchInput"
            placeholder="搜索 Webhook..."
            class="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white"
          />
          <svg
            class="absolute right-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            ></path>
          </svg>
        </div>
      </div>
    </div>

    <div class="overflow-x-auto">
      <n-table :bordered="false" :single-line="false">
        <thead>
          <tr>
            <th>URL</th>
            <th>事件</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <!-- Webhook 列表将在这里动态加载 -->
          <tr>
            <td>https://example.com/webhook1</td>
            <td>内容创建, 内容更新</td>
            <td>
              <n-button type="primary"> 编辑 </n-button>
              <n-button type="error"> 删除 </n-button>
            </td>
          </tr>
          <!-- 更多 Webhook 行 -->
        </tbody>
      </n-table>
    </div>

    <!-- 分页组件 -->
    <div class="mt-6 flex justify-end">
      <n-pagination :total="10" :page-size="5" :page-size-options="[5, 10, 20]" />
    </div>

    <!-- Add/Edit Webhook Modal -->
    <n-modal>
      <template #header>
        添加 Webhook
      </template>
      <n-form>
        <n-form-item label="url">
          <n-input v-model:value="webhookFormData.webhookUrl" placeholder="请输入 Webhook URL" />
        </n-form-item>
        <n-form-item label="事件">
          <n-checkbox-group v-model:value="webhookFormData.webhookEvents" :options="eventOptions" />
        </n-form-item>
      </n-form>
      <template #action>
        <div class="items-center px-4 py-3">
          <n-button
            type="primary"
          >
            保存
          </n-button>
          <n-button
            type="default"
          >
            取消
          </n-button>
        </div>
      </template>
    </n-modal>
  </div>
</template>
