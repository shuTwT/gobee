<template>
  <div class="p-4">
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">支付订单管理</h1>
      <div class="flex space-x-4">
        <select id="stateFilter" onchange="loadOrders()" class="border rounded py-2 px-4">
          <option value="">全部状态</option>
          <option value="1">待支付</option>
          <option value="2">支付成功</option>
          <option value="3">支付失败</option>
        </select>
        <n-button onclick="loadOrders()" type="info"> 刷新列表 </n-button>
      </div>
    </div>

    <div class="overflow-x-auto">
      <n-table>
        <thead>
          <tr>
            <th>订单ID</th>
            <th>商户订单号</th>
            <th>订单金额</th>
            <th>订单标题</th>
            <th>支付渠道</th>
            <th>支付状态</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <!-- 订单列表将通过JavaScript动态加载 -->
        </tbody>
      </n-table>
    </div>

    <!-- 订单详情模态框 -->
    <n-modal >
      <template #header>
          订单详情
      </template>
      <div class="flex items-center justify-center min-h-screen">
        <div class="bg-white p-6 rounded-lg w-full max-w-4xl">
          <div id="orderDetails" class="grid grid-cols-2 gap-4">
            <!-- 订单详情将通过JavaScript动态加载 -->
          </div>
          <div class="mt-6 flex justify-end">
            <n-button
              onclick="closeOrderModal()"
              type="default"
            >
              关闭
            </n-button>
          </div>
        </div>
      </div>
    </n-modal>
  </div>
</template>
