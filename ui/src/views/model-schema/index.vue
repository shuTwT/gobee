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

const fieldTypeOptions = [
  {
    label: '字符串',
    value: 'string',
  },
  {
    label: '数字',
    value: 'number',
  },
  {
    label: '布尔',
    value: 'boolean',
  },
  {
    label:'文本域',
    value:'textarea',
  },
  {
    label:'富文本',
    value:'richtext',
  },
  {
    label: '日期时间',
    value: 'datetime',
  }
]

const showModalSchemaModal = ref(false)
const showModalSchemaFieldModal = ref(false)
const modelFieldFormData = ref({
  name:'',
  key:'',
  type:''
})

const updateActiveModelSchema = (modelSchema: ModelSchema) => {
  activeModelSchema.value = modelSchema
}

const openModelSchemaModal = () => {
  showModalSchemaModal.value = true
}
const openModelSchemaFieldModal = () => {
  showModalSchemaFieldModal.value = true
}
</script>
<template>
  <div class="flex h-full">
    <!-- 左侧列表 -->
    <div class="w-64 bg-white border-r border-gray-200 flex flex-col">
      <div class="p-4 border-b border-gray-200 flex justify-between items-center">
        <h2 class="text-lg font-medium text-gray-900">内容模型</h2>
        <button id="add-model-btn" class="p-1 rounded-md text-blue-600 hover:bg-blue-50" @click="openModelSchemaModal">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
              clip-rule="evenodd"
            />
          </svg>
        </button>
      </div>
      <div class="overflow-y-auto flex-1">
        <ul id="model-list" class="divide-y divide-gray-200">
          <!-- 示例模型项 -->
          <li
            v-for="(item, index) in modelSchemaList"
            :key="index"
            class="model-item px-4 py-3 hover:bg-gray-50 cursor-pointer"
          >
            <div class="flex justify-between items-center" @click="updateActiveModelSchema(item)">
              <div>
                <h3 class="text-sm font-medium text-gray-900">{{ item.name }}</h3>
                <p class="text-xs text-gray-500">{{ item.model }}</p>
              </div>
              <div class="flex space-x-1">
                <button class="text-gray-400 hover:text-gray-500">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"
                    />
                  </svg>
                </button>
                <button class="text-gray-400 hover:text-red-500">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </button>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>

    <!-- 右侧编辑区域 -->
    <div class="flex-1 flex flex-col">
      <!-- 选择模型后的编辑区域 -->
      <div v-if="activeModelSchema" class="flex-1 flex flex-col">
        <div class="p-4 border-b border-gray-200 flex justify-between items-center">
          <div>
            <h2 id="model-title" class="text-lg font-medium text-gray-900">文章</h2>
            <p id="model-identifier" class="text-sm text-gray-500">article</p>
          </div>
          <n-button type="primary" @click="openModelSchemaFieldModal"> 添加字段 </n-button>
        </div>
        <div class="p-4 flex-1 overflow-auto">
          <n-table>
            <thead class="bg-gray-50">
              <tr>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  字段名称
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  字段标识
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  字段类型
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  必填
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  操作
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <!-- 示例字段 -->
              <tr v-for="(item, index) in activeModelSchema.fields" :key="index">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ item.name }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ item.key }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ item.type }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <span
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800"
                    >是</span
                  >
                </td>
                <td>
                  <n-button type="primary">编辑</n-button>
                  <n-button type="error">删除</n-button>
                </td>
              </tr>
            </tbody>
          </n-table>
        </div>
      </div>

      <!-- 未选择模型时的提示 -->
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
            d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
          />
        </svg>
        <h3 class="text-lg font-medium text-gray-900 mb-2">未选择内容模型</h3>
        <p class="text-sm text-gray-500 max-w-md">
          请从左侧列表选择一个内容模型进行编辑，或者点击左上角的"+"按钮创建一个新的内容模型。
        </p>
      </div>
    </div>
  </div>

  <!-- 添加模型弹窗 -->
  <n-modal v-model:show="showModalSchemaModal">
    <div class="bg-white rounded-lg shadow-xl max-w-md w-full">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">添加内容模型</h3>
      </div>
      <form id="add-model-form" class="px-6 py-4">
        <div class="mb-4">
          <label for="model-name" class="block text-sm font-medium text-gray-700 mb-1"
            >模型名称</label
          >
          <input
            type="text"
            id="model-name"
            name="model-name"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="例如：文章、产品"
          />
        </div>
        <div class="mb-4">
          <label for="model-identifier" class="block text-sm font-medium text-gray-700 mb-1"
            >模型标识</label
          >
          <input
            type="text"
            id="model-identifier"
            name="model-identifier"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="例如：article、product"
          />
          <p class="mt-1 text-xs text-gray-500">只能包含小写字母、数字和下划线，不能以数字开头</p>
        </div>
        <div class="mb-4">
          <label for="model-description" class="block text-sm font-medium text-gray-700 mb-1"
            >模型描述（可选）</label
          >
          <textarea
            id="model-description"
            name="model-description"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="描述这个内容模型的用途..."
          ></textarea>
        </div>
      </form>
      <div class="px-6 py-4 border-t border-gray-200 flex justify-end">
        <n-button
          type="primary"
        >
          取消
        </n-button>
        <n-button
          type="primary"
        >
          确定
        </n-button>
      </div>
    </div>
  </n-modal>

  <!-- 添加字段弹窗 -->
  <n-modal v-model:show="showModalSchemaFieldModal">
    <div class="bg-white rounded-lg shadow-xl max-w-md w-full">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">添加字段</h3>
      </div>
      <form id="add-field-form" class="px-6 py-4">
        <div class="mb-4">
          <label for="field-name" class="block text-sm font-medium text-gray-700 mb-1"
            >字段名称</label
          >
          <n-input
            v-model:value="modelFieldFormData.name"
            placeholder="例如：标题、内容"
          />
        </div>
        <div class="mb-4">
          <label for="field-identifier" class="block text-sm font-medium text-gray-700 mb-1"
            >字段标识</label
          >
          <n-input
            v-model:value="modelFieldFormData.key"
            placeholder="例如：title、content"
          />
          <p class="mt-1 text-xs text-gray-500">只能包含小写字母、数字和下划线，不能以数字开头</p>
        </div>
        <div class="mb-4">
          <label for="field-type" class="block text-sm font-medium text-gray-700 mb-1"
            >字段类型</label
          >
          <n-select :options="fieldTypeOptions" v-model:value="modelFieldFormData.type" />
        </div>
        <div class="mb-4">
          <div class="flex items-center">
            <input
              type="checkbox"
              id="field-required"
              name="field-required"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="field-required" class="ml-2 block text-sm text-gray-900">必填字段</label>
          </div>
        </div>
      </form>
      <div class="px-6 py-4 border-t border-gray-200 flex justify-end">
        <n-button
          type="default"
        >
          取消
        </n-button>
        <n-button
          type="primary"
        >
          确定
        </n-button>
      </div>
    </div>
  </n-modal>
</template>
