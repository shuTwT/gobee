<template>
  <div class="p-6 bg-white dark:bg-gray-800 shadow-md rounded-lg">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-semibold text-gray-900 dark:text-white">文件管理</h2>
      <div class="flex space-x-4">
        <n-button id="uploadFileBtn" type="primary"> 上传文件 </n-button>
        <div class="relative">
          <input
            type="text"
            id="fileSearchInput"
            placeholder="搜索文件..."
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
      <n-table>
        <thead>
          <tr>
            <th>文件名</th>
            <th>类型</th>
            <th>大小</th>
            <th>上传日期</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <!-- 文件列表将在这里动态加载 -->
          <tr>
            <td>example.jpg</td>
            <td>图片</td>
            <td>1.2 MB</td>
            <td>2023-10-26</td>
            <td>
              <n-button type="primary">查看</n-button>
              <n-button type="danger">删除</n-button>
            </td>
          </tr>
          <!-- 更多文件行 -->
        </tbody>
      </n-table>
    </div>

    <!-- 分页组件 -->
    <div class="mt-6 flex justify-end">
      <n-pagination :total="100" :page-size="10" :page-size-options="[10, 20, 30, 40]" />
    </div>

    <!-- 文件上传 Modal -->
    <n-modal>
      <div
        class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800"
      >
        <div class="mt-3 text-center">
          <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white">上传文件</h3>
          <div class="mt-2 px-7 py-3">
            <input
              type="file"
              id="fileInput"
              class="block w-full text-sm text-gray-500 dark:text-gray-400 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100 dark:file:bg-blue-900/30 dark:file:text-blue-400 dark:hover:file:bg-blue-800"
            />
          </div>
          <div class="items-center px-4 py-3">
            <button
              id="doUploadBtn"
              class="px-4 py-2 bg-blue-600 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
            >
              上传
            </button>
            <button
              id="closeUploadModalBtn"
              class="mt-3 px-4 py-2 bg-gray-200 text-gray-800 text-base font-medium rounded-md w-full shadow-sm hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-300 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white dark:hover:bg-gray-600"
            >
              取消
            </button>
          </div>
        </div>
      </div>
    </n-modal>
  </div>
</template>
