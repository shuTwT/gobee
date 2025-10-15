<script setup lang="ts">
type ModelSchema = {
  id: string
  name: string
  model: string
  fields: Array<{
    name: string
    key: string
    type: string
    required: boolean
  }>
}
const activeModelSchema = ref<ModelSchema>()
const modelSchemaList = ref<ModelSchema[]>([
  {
    id: '1',
    name: '文章',
    model: 'article',
    fields: [
      {
        name: '标题',
        key: 'title',
        type: 'string',
        required: true,
      },
      {
        name: '内容',
        key: 'content',
        type: 'string',
        required: true,
      },
    ],
  },
  {
    id: '2',
    name: '产品',
    model: 'product',
    fields: [
      {
        name: '用户名',
        key: 'username',
        type: 'string',
        required: true,
      },
      {
        name: '邮箱',
        key: 'email',
        type: 'string',
        required: true,
      },
    ],
  },
])
const languageOptions = [
  {
    value: 'zh-CN',
    label: '中文',
  },
  {
    value: 'en-US',
    label: '英文',
  },
  {
    value: 'ja-JP',
    label: '日文',
  },
]
const contentStateOptions = [
  {
    value: 'draft',
    label: '草稿',
  },
  {
    value: 'published',
    label: '已发布',
  },
  {
    value: 'archived',
    label: '已归档',
  },
]
const modelContentFormData = ref<Record<string, any>>({})

const showEditModal = ref(false)

const updateActiveModelSchema = (modelSchema: ModelSchema) => {
  activeModelSchema.value = modelSchema
}
</script>
<template>
  <div class="flex h-full">
    <!-- 左侧内容模型列表 -->
    <div class="w-64 bg-white border-r border-gray-200 flex flex-col">
      <div class="p-4 border-b border-gray-200">
        <h2 class="text-lg font-medium text-gray-900">内容模型</h2>
      </div>
      <div class="overflow-y-auto flex-1">
        <ul id="model-list" class="divide-y divide-gray-200">
          <!-- 示例内容模型项 -->
          <li
            v-for="(item, index) in modelSchemaList"
            :key="index"
            class="model-item px-4 py-3 hover:bg-gray-50 cursor-pointer"
            data-model-id="article"
            data-model-name="文章"
          >
            <div class="flex justify-between items-center" @click="updateActiveModelSchema(item)">
              <div>
                <h3 class="text-sm font-medium text-gray-900">{{ item.name }}</h3>
                <p class="text-xs text-gray-500">{{ item.model }}</p>
              </div>
              <span
                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
              >
                {{ item.fields.length }}
              </span>
            </div>
          </li>
        </ul>
      </div>
    </div>

    <!-- 右侧内容区域 -->
    <div class="flex-1 flex flex-col">
      <!-- 内容列表区域 -->
      <div v-if="activeModelSchema" id="content-list-container" class="flex-1 flex flex-col">
        <div class="p-4 border-b border-gray-200 flex justify-between items-center">
          <h2 id="content-list-title" class="text-lg font-medium text-gray-900">
            {{ activeModelSchema.name }}列表
          </h2>
          <n-button type="primary" @click="showEditModal = true"> 新建内容 </n-button>
        </div>
        <div class="p-4">
          <div class="mb-4">
            <div class="relative">
              <input
                type="text"
                id="search-content"
                placeholder="搜索内容..."
                class="w-full px-3 py-2 pl-10 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 text-gray-400"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
            </div>
          </div>
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th
                    scope="col"
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    标题
                  </th>
                  <th
                    scope="col"
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    作者
                  </th>
                  <th
                    scope="col"
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    状态
                  </th>
                  <th
                    scope="col"
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    发布日期
                  </th>
                  <th
                    scope="col"
                    class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    操作
                  </th>
                </tr>
              </thead>
              <tbody id="content-list" class="bg-white divide-y divide-gray-200">
                <!-- 示例内容项 -->
                <tr class="hover:bg-gray-50">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm font-medium text-gray-900">如何使用Headless CMS</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-500">张三</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span
                      class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800"
                      >已发布</span
                    >
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">2023-05-15</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button
                      class="edit-content-btn text-blue-600 hover:text-blue-900 mr-2"
                      data-content-id="1"
                    >
                      编辑
                    </button>
                    <button class="text-red-600 hover:text-red-900">删除</button>
                  </td>
                </tr>
                <tr class="hover:bg-gray-50">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm font-medium text-gray-900">内容模型设计指南</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-500">李四</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span
                      class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800"
                      >草稿</span
                    >
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button
                      class="edit-content-btn text-blue-600 hover:text-blue-900 mr-2"
                      data-content-id="2"
                    >
                      编辑
                    </button>
                    <button class="text-red-600 hover:text-red-900">删除</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="mt-4 flex justify-between items-center">
            <div class="text-sm text-gray-500">显示 1-2 共 2 条</div>
            <div class="flex space-x-2">
              <button
                class="px-3 py-1 border border-gray-300 rounded-md text-sm text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
                disabled
              >
                上一页
              </button>
              <button
                class="px-3 py-1 border border-gray-300 rounded-md text-sm text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
                disabled
              >
                下一页
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 未选择内容时的提示 -->
      <div v-else class="flex-1 flex flex-col items-center justify-center p-8 text-center">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-16 w-16 text-gray-300 mb-4"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
          />
        </svg>
        <h3 class="text-lg font-medium text-gray-900 mb-2">未选择内容模型</h3>
        <p class="text-sm text-gray-500 max-w-md">请从左侧列表选择一个内容模型查看其内容列表。</p>
      </div>

      <!-- 内容编辑表单 -->
      <n-modal v-model:show="showEditModal" preset="dialog">
        <template #header>
          <!-- 表单头部 -->
          <div class="p-4 border-b border-gray-200 flex justify-between items-center">
            <div class="flex items-center">
              <h2 id="content-title" class="text-lg font-medium text-gray-900 mr-4">编辑内容</h2>
              <div class="relative">
                <n-select
                  v-model:value="modelContentFormData!.language"
                  placeholder="选择语言"
                  :options="languageOptions"
                  style="width: 200px"
                />
                <div
                  class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
                >
                  <svg
                    class="h-4 w-4"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </template>
        <!-- 表单内容 -->
        <div class="p-6 flex-1 overflow-auto">
          <n-form>
            <n-form-item label="标题">
              <n-input v-model:value="modelContentFormData!.title" />
            </n-form-item>

            <n-form-item label="摘要">
              <n-input v-model:value="modelContentFormData!.summary" type="textarea" />
            </n-form-item>
            <n-form-item label="内容">
              <n-input v-model:value="modelContentFormData!.content" type="textarea" />
            </n-form-item>
            <n-form-item label="发布日期">
              <n-date-picker v-model:value="modelContentFormData!.publishDate" type="datetime" />
            </n-form-item>

            <n-form-item label="作者">
              <n-input v-model:value="modelContentFormData!.author" />
            </n-form-item>
            <n-form-item label="标签">
              <n-input v-model:value="modelContentFormData!.tags" />
            </n-form-item>

            <!-- 封面图片字段 -->
            <n-form-item label="封面图片">
              <image-upload v-model:file-list="modelContentFormData!.coverImage" :limit="1"/>
            </n-form-item>

            <n-form-item label="状态">
              <n-select
                :options="contentStateOptions"
                v-model:value="modelContentFormData!.status"
              />
            </n-form-item>
          </n-form>
        </div>
        <template #action>
          <div class="flex space-x-2">
            <n-button type="default"> 返回列表 </n-button>
            <n-button type="default"> 保存草稿 </n-button>
            <n-button type="primary"> 发布 </n-button>
          </div>
        </template>
      </n-modal>
    </div>
  </div>
</template>
