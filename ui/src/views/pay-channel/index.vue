<script setup lang="ts">
const channelFormData = ref({
  name: '',
  code: '',
  type: '',
  config: '',
})
const typeOptions = [
  {
    label: '支付宝',
    value: 'alipay',
  },
  {
    label: '微信支付',
    value: 'wechatpay',
  },
  {
    label: '银联支付',
    value: 'unionpay',
  },
]
</script>
<template>
  <div class="p-4">
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">支付渠道管理</h1>
      <n-button onclick="openAddModal()" type="primary"> 添加渠道 </n-button>
    </div>

    <div class="overflow-x-auto">
      <n-table>
        <thead>
          <tr>
            <th>渠道名称</th>
            <th>渠道代码</th>
            <th>渠道类型</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <!-- 渠道列表将通过JavaScript动态加载 -->
        </tbody>
      </n-table>
    </div>

    <!-- 添加/编辑渠道模态框 -->
    <n-modal>
      <template #header> 添加支付渠道 </template>
      <n-form>
        <n-form-item label="渠道名称">
          <n-input v-model:value="channelFormData.name" />
        </n-form-item>
        <n-form-item label="渠道代码">
          <n-input v-model:value="channelFormData.code" />
        </n-form-item>
        <n-form-item label="渠道类型">
          <n-select v-model:value="channelFormData.type" :options="typeOptions" />
        </n-form-item>
      </n-form>
      <template #action>
        <div class="flex items-center justify-end">
          <n-button
            type="default"
            onclick="closeModal()"

          >
            取消
          </n-button>
          <n-button
            type="primary"
          >
            保存
          </n-button>
        </div>
      </template>
    </n-modal>
  </div>
</template>
