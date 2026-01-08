<template>
  <div class="pay-order-container p-6">
    <n-card title="支付订单管理" class="order-card">
      <!-- 头部操作栏 -->
      <div class="header-section">
        <div class="search-section">
          <n-input
            v-model:value="searchKeyword"
            placeholder="搜索订单号、商户订单号或标题"
            clearable
            style="width: 250px"
            @input="handleSearch"
          >
            <template #prefix>
              <n-icon><search-outline /></n-icon>
            </template>
          </n-input>
          <n-select
            v-model:value="filterStatus"
            placeholder="筛选状态"
            clearable
            style="width: 150px; margin-left: 12px"
            :options="statusOptions"
            @update:value="handleFilterChange"
          />
          <n-select
            v-model:value="filterChannel"
            placeholder="筛选支付渠道"
            clearable
            style="width: 150px; margin-left: 12px"
            :options="channelOptions"
            @update:value="handleFilterChange"
          />
          <n-date-picker
            v-model:value="dateRange"
            type="daterange"
            clearable
            style="width: 240px; margin-left: 12px"
            placeholder="选择日期范围"
            @update:value="handleFilterChange"
          />
        </div>
        <div class="action-section">
          <n-button type="primary" @click="exportOrders" style="margin-right: 12px">
            <template #icon>
              <n-icon><download-outline /></n-icon>
            </template>
            导出订单
          </n-button>
          <n-button @click="handleRefresh">
            <template #icon>
              <n-icon><refresh-outline /></n-icon>
            </template>
            刷新
          </n-button>
        </div>
      </div>

      <!-- 统计信息卡片 -->
      <div class="stats-section">
        <n-grid :cols="4" :x-gap="16" :y-gap="16">
          <n-gi>
            <n-statistic label="今日订单" :value="todayStats.total">
              <template #prefix>
                <n-icon color="#52c41a"><document-text-outline /></n-icon>
              </template>
            </n-statistic>
          </n-gi>
          <n-gi>
            <n-statistic label="今日收入" :value="todayStats.amount" precision="2">
              <template #prefix>
                <n-icon color="#1890ff"><cash-outline /></n-icon>
              </template>
              <template #suffix>元</template>
            </n-statistic>
          </n-gi>
          <n-gi>
            <n-statistic label="成功率" :value="todayStats.successRate" precision="2">
              <template #prefix>
                <n-icon color="#722ed1"><trending-up-outline /></n-icon>
              </template>
              <template #suffix>%</template>
            </n-statistic>
          </n-gi>
          <n-gi>
            <n-statistic label="待处理" :value="todayStats.pending">
              <template #prefix>
                <n-icon color="#fa8c16"><time-outline /></n-icon>
              </template>
            </n-statistic>
          </n-gi>
        </n-grid>
      </div>

      <!-- 订单列表表格 -->
      <n-data-table
        :columns="columns"
        :data="filteredOrderList"
        :loading="loading"
        :pagination="pagination"
        :row-key="(row) => row.id"
        :scroll-x="1500"
        :remote="true"
        class="order-table"
      />
    </n-card>

    <!-- 订单详情模态框 -->
    <n-modal v-model:show="showDetailModal" preset="dialog" title="订单详情" style="width: 800px">
      <div class="order-detail-content" v-if="currentOrder">
        <n-tabs type="line" animated>
          <n-tab-pane name="basic" tab="基本信息">
            <n-descriptions :columns="2" bordered>
              <n-descriptions-item label="订单ID">{{ currentOrder.id }}</n-descriptions-item>
              <n-descriptions-item label="商户订单号">{{
                currentOrder.merchantOrderNo
              }}</n-descriptions-item>
              <n-descriptions-item label="订单标题">{{ currentOrder.title }}</n-descriptions-item>
              <n-descriptions-item label="订单金额">
                <n-text type="primary" strong>¥{{ currentOrder.amount }}</n-text>
              </n-descriptions-item>
              <n-descriptions-item label="实际支付">
                <n-text type="success" strong>¥{{ currentOrder.paidAmount }}</n-text>
              </n-descriptions-item>
              <n-descriptions-item label="支付渠道">{{
                currentOrder.channelName
              }}</n-descriptions-item>
              <n-descriptions-item label="订单状态">
                <n-tag :type="getStatusType(currentOrder.status)">{{
                  getStatusText(currentOrder.status)
                }}</n-tag>
              </n-descriptions-item>
              <n-descriptions-item label="创建时间">{{
                currentOrder.createdAt
              }}</n-descriptions-item>
              <n-descriptions-item label="支付时间" v-if="currentOrder.paidAt">{{
                currentOrder.paidAt
              }}</n-descriptions-item>
              <n-descriptions-item label="过期时间">{{
                currentOrder.expireAt
              }}</n-descriptions-item>
              <n-descriptions-item label="用户ID">{{ currentOrder.userId }}</n-descriptions-item>
            </n-descriptions>
          </n-tab-pane>

          <n-tab-pane name="payment" tab="支付信息">
            <n-descriptions :columns="2" bordered>
              <n-descriptions-item label="支付流水号" v-if="currentOrder.tradeNo">{{
                currentOrder.tradeNo
              }}</n-descriptions-item>
              <n-descriptions-item label="第三方订单号" v-if="currentOrder.thirdPartyOrderNo">{{
                currentOrder.thirdPartyOrderNo
              }}</n-descriptions-item>
              <n-descriptions-item label="支付状态">{{
                getStatusText(currentOrder.status)
              }}</n-descriptions-item>
              <n-descriptions-item label="支付时间" v-if="currentOrder.paidAt">{{
                currentOrder.paidAt
              }}</n-descriptions-item>
              <n-descriptions-item label="支付金额" v-if="currentOrder.paidAmount"
                >¥{{ currentOrder.paidAmount }}</n-descriptions-item
              >
              <n-descriptions-item label="手续费" v-if="currentOrder.fee"
                >¥{{ currentOrder.fee }}</n-descriptions-item
              >
              <n-descriptions-item label="失败原因" v-if="currentOrder.failReason" :span="2">
                <n-text type="error">{{ currentOrder.failReason }}</n-text>
              </n-descriptions-item>
            </n-descriptions>
          </n-tab-pane>

          <n-tab-pane name="extra" tab="扩展信息">
            <n-descriptions :columns="1" bordered>
              <n-descriptions-item label="客户端IP">{{
                currentOrder.clientIp
              }}</n-descriptions-item>
              <n-descriptions-item label="用户代理">{{
                currentOrder.userAgent
              }}</n-descriptions-item>
              <n-descriptions-item label="回调地址">{{
                currentOrder.notifyUrl
              }}</n-descriptions-item>
              <n-descriptions-item label="跳转地址">{{
                currentOrder.returnUrl
              }}</n-descriptions-item>
              <n-descriptions-item label="扩展参数" v-if="currentOrder.extra">
                <n-code :code="formatExtra(currentOrder.extra)" language="json" :word-wrap="true" />
              </n-descriptions-item>
            </n-descriptions>
          </n-tab-pane>
        </n-tabs>
      </div>
      <template #action>
        <n-space justify="end">
          <n-button @click="showDetailModal = false">关闭</n-button>
          <n-button
            v-if="currentOrder?.status === 'pending'"
            type="warning"
            @click="handleCancelOrder"
          >
            取消订单
          </n-button>
          <n-button v-if="currentOrder?.status === 'paid'" type="error" @click="handleRefundOrder">
            退款
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, h } from 'vue'
import { NButton, NIcon, NTag, NText, useMessage } from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import {
  SearchOutline,
  RefreshOutline,
  DownloadOutline,
  DocumentTextOutline,
  CashOutline,
  TrendingUpOutline,
  TimeOutline,
  EyeOutline,
} from '@vicons/ionicons5'
import * as payOrderApi from '@/api/pay/payOrder'

const message = useMessage()

// 搜索和筛选
const searchKeyword = ref('')
const filterStatus = ref<string | null>(null)
const filterChannel = ref<string | null>(null)
const dateRange = ref<[number, number] | null>(null)

// 表格数据
const loading = ref(false)
const orderList = ref<PayOrderItem[]>([])

// 模态框状态
const showDetailModal = ref(false)
const currentOrder = ref<PayOrderItem | null>(null)

// 分页配置
const pagination = reactive({
  page: 1,
  pageSize: 10,
  showSizePicker: true,
  pageSizes: [10, 20, 50, 100],
  onChange: (page: number) => {
    pagination.page = page
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.pageSize = pageSize
    pagination.page = 1
  },
})

// 订单状态选项
const statusOptions = [
  { label: '全部状态', value: null as any },
  { label: '待支付', value: 'pending' },
  { label: '支付成功', value: 'paid' },
  { label: '支付失败', value: 'failed' },
  { label: '已取消', value: 'cancelled' },
  { label: '已退款', value: 'refunded' },
]

// 支付渠道选项
const channelOptions = [
  { label: '全部渠道', value: null as any },
  { label: '支付宝', value: 'alipay' },
  { label: '微信支付', value: 'wechatpay' },
  { label: '银联支付', value: 'unionpay' },
  { label: 'PayPal', value: 'paypal' },
  { label: 'Stripe', value: 'stripe' },
]

// 今日统计
const todayStats = reactive({
  total: 0,
  amount: 0,
  successRate: 0,
  pending: 0,
})

// 订单数据类型
interface PayOrderItem {
  id: string
  merchantOrderNo: string
  title: string
  amount: number
  paidAmount: number
  fee: number
  channel: string
  channelName: string
  status: 'pending' | 'paid' | 'failed' | 'cancelled' | 'refunded'
  tradeNo?: string
  thirdPartyOrderNo?: string
  userId: string
  clientIp: string
  userAgent: string
  notifyUrl: string
  returnUrl: string
  extra?: string
  failReason?: string
  createdAt: string
  paidAt?: string
  expireAt: string
  updatedAt: string
}

// 筛选后的订单列表
const filteredOrderList = computed(() => {
  let filtered = orderList.value

  // 关键词搜索
  if (searchKeyword.value) {
    const keyword = searchKeyword.value.toLowerCase()
    filtered = filtered.filter(
      (order) =>
        order.id.toLowerCase().includes(keyword) ||
        order.merchantOrderNo.toLowerCase().includes(keyword) ||
        order.title.toLowerCase().includes(keyword),
    )
  }

  // 状态筛选
  if (filterStatus.value) {
    filtered = filtered.filter((order) => order.status === filterStatus.value)
  }

  // 渠道筛选
  if (filterChannel.value) {
    filtered = filtered.filter((order) => order.channel === filterChannel.value)
  }

  // 日期范围筛选
  if (dateRange.value && dateRange.value.length === 2) {
    const [startDate, endDate] = dateRange.value
    filtered = filtered.filter((order) => {
      const orderDate = new Date(order.createdAt).getTime()
      return orderDate >= startDate && orderDate <= endDate
    })
  }

  return filtered
})

// 获取状态标签类型
const getStatusType = (status: string) => {
  const typeMap: Record<string, any> = {
    pending: 'warning',
    paid: 'success',
    failed: 'error',
    cancelled: 'default',
    refunded: 'info',
  }
  return typeMap[status] || 'default'
}

// 获取状态文本
const getStatusText = (status: string) => {
  const textMap: Record<string, string> = {
    pending: '待支付',
    paid: '支付成功',
    failed: '支付失败',
    cancelled: '已取消',
    refunded: '已退款',
  }
  return textMap[status] || status
}

// 格式化扩展信息
const formatExtra = (extra: string) => {
  try {
    const parsed = JSON.parse(extra)
    return JSON.stringify(parsed, null, 2)
  } catch {
    return extra
  }
}

// 表格列定义
const columns: DataTableColumns<PayOrderItem> = [
  {
    title: '订单ID',
    key: 'id',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '商户订单号',
    key: 'merchantOrderNo',
    width: 180,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '订单标题',
    key: 'title',
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: '订单金额',
    key: 'amount',
    width: 100,
    render: (row) => {
      return h(NText, { type: 'primary', strong: true }, () => `¥${row.amount}`)
    },
    sorter: (row1, row2) => row1.amount - row2.amount,
  },
  {
    title: '支付金额',
    key: 'paidAmount',
    width: 100,
    render: (row) => {
      if (row.paidAmount > 0) {
        return h(NText, { type: 'success', strong: true }, () => `¥${row.paidAmount}`)
      }
      return h(NText, { type: 'default' }, () => '-')
    },
    sorter: (row1, row2) => row1.paidAmount - row2.paidAmount,
  },
  {
    title: '支付渠道',
    key: 'channelName',
    width: 120,
    render: (row) => {
      return h(NTag, { type: 'info' }, () => row.channelName)
    },
  },
  {
    title: '订单状态',
    key: 'status',
    width: 100,
    render: (row) => {
      return h(NTag, { type: getStatusType(row.status) }, () => getStatusText(row.status))
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 160,
    sorter: (row1, row2) => new Date(row1.createdAt).getTime() - new Date(row2.createdAt).getTime(),
  },
  {
    title: '支付时间',
    key: 'paidAt',
    width: 160,
    render: (row) => {
      return row.paidAt || '-'
    },
    sorter: (row1, row2) => {
      const time1 = row1.paidAt ? new Date(row1.paidAt).getTime() : 0
      const time2 = row2.paidAt ? new Date(row2.paidAt).getTime() : 0
      return time1 - time2
    },
  },
  {
    title: '操作',
    fixed:"right",
    key: 'actions',
    width: 100,
    render: (row) => {
      return h(
        NButton,
        {
          size: 'small',
          type: 'primary',
          quaternary: true,
          onClick: () => viewOrderDetail(row),
        },
        {
          icon: () => h(NIcon, {}, () => h(EyeOutline)),
          default: () => '详情',
        },
      )
    },
  },
]

// 处理搜索
const handleSearch = () => {
  pagination.page = 1
}

// 处理筛选变化
const handleFilterChange = () => {
  pagination.page = 1
}

// 查看订单详情
const viewOrderDetail = (order: PayOrderItem) => {
  currentOrder.value = order
  showDetailModal.value = true
}

// 取消订单
const handleCancelOrder = () => {
  if (currentOrder.value && currentOrder.value.status === 'pending') {
    // 模拟API调用
    setTimeout(() => {
      const order = orderList.value.find((item) => item.id === currentOrder.value?.id)
      if (order) {
        order.status = 'cancelled'
        order.updatedAt = new Date().toLocaleString('zh-CN')
      }
      message.success('订单已取消')
      showDetailModal.value = false
    }, 500)
  }
}

// 退款订单
const handleRefundOrder = () => {
  if (currentOrder.value && currentOrder.value.status === 'paid') {
    // 模拟API调用
    setTimeout(() => {
      const order = orderList.value.find((item) => item.id === currentOrder.value?.id)
      if (order) {
        order.status = 'refunded'
        order.updatedAt = new Date().toLocaleString('zh-CN')
      }
      message.success('退款申请已提交')
      showDetailModal.value = false
    }, 500)
  }
}

// 导出订单
const exportOrders = () => {
  message.info('订单导出功能开发中...')
}

// 刷新数据
const handleRefresh = () => {
  loadOrderList()
}

// 计算今日统计
const calculateTodayStats = () => {
  const today = new Date().toDateString()
  const todayOrders = orderList.value.filter(
    (order) => new Date(order.createdAt).toDateString() === today,
  )

  todayStats.total = todayOrders.length
  todayStats.amount = todayOrders.reduce((sum, order) => sum + order.paidAmount, 0)
  todayStats.pending = todayOrders.filter((order) => order.status === 'pending').length

  const successOrders = todayOrders.filter((order) => order.status === 'paid')
  todayStats.successRate =
    todayOrders.length > 0 ? (successOrders.length / todayOrders.length) * 100 : 0
}

const onSearch = async ()=>{
  const res = await payOrderApi.getPayOrderList()
}

// 加载订单列表
const loadOrderList = async () => {
  loading.value = true
  try {
    // 模拟API调用
    await new Promise((resolve) => setTimeout(resolve, 800))

    // 模拟数据
    orderList.value = [
      {
        id: 'ORD202401150001',
        merchantOrderNo: 'MER202401150001',
        title: 'VIP会员年费',
        amount: 299.0,
        paidAmount: 299.0,
        fee: 2.99,
        channel: 'alipay',
        channelName: '支付宝',
        status: 'paid',
        tradeNo: '2024011522001234567890',
        thirdPartyOrderNo: '2024011531f2e4r5t6y7u8i9o0p',
        userId: 'USER123456',
        clientIp: '192.168.1.100',
        userAgent: 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36',
        notifyUrl: 'https://example.com/api/pay/notify/alipay',
        returnUrl: 'https://example.com/pay/success',
        extra: JSON.stringify({ product_id: 'VIP001', quantity: 1 }, null, 2),
        createdAt: '2024-01-15 14:30:00',
        paidAt: '2024-01-15 14:31:25',
        expireAt: '2024-01-15 15:00:00',
        updatedAt: '2024-01-15 14:31:25',
      },
      {
        id: 'ORD202401150002',
        merchantOrderNo: 'MER202401150002',
        title: '月度会员',
        amount: 29.9,
        paidAmount: 0,
        fee: 0,
        channel: 'wechatpay',
        channelName: '微信支付',
        status: 'pending',
        userId: 'USER123457',
        clientIp: '192.168.1.101',
        userAgent: 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X)',
        notifyUrl: 'https://example.com/api/pay/notify/wechat',
        returnUrl: 'https://example.com/pay/success',
        extra: JSON.stringify({ product_id: 'MONTH001', quantity: 1 }, null, 2),
        createdAt: '2024-01-15 14:35:00',
        expireAt: '2024-01-15 15:05:00',
        updatedAt: '2024-01-15 14:35:00',
      },
      {
        id: 'ORD202401150003',
        merchantOrderNo: 'MER202401150003',
        title: '高级功能包',
        amount: 99.0,
        paidAmount: 0,
        fee: 0,
        channel: 'alipay',
        channelName: '支付宝',
        status: 'failed',
        failReason: '用户支付超时',
        userId: 'USER123458',
        clientIp: '192.168.1.102',
        userAgent: 'Mozilla/5.0 (Android 11; Mobile; rv:68.0)',
        notifyUrl: 'https://example.com/api/pay/notify/alipay',
        returnUrl: 'https://example.com/pay/success',
        extra: JSON.stringify({ product_id: 'ADV001', quantity: 1 }, null, 2),
        createdAt: '2024-01-15 14:40:00',
        expireAt: '2024-01-15 15:10:00',
        updatedAt: '2024-01-15 14:45:30',
      },
      {
        id: 'ORD202401150004',
        merchantOrderNo: 'MER202401150004',
        title: '企业版套餐',
        amount: 999.0,
        paidAmount: 999.0,
        fee: 9.99,
        channel: 'stripe',
        channelName: 'Stripe',
        status: 'refunded',
        tradeNo: 'pi_3O4H5I6J7K8L9M0N',
        thirdPartyOrderNo: 'ch_1O4H5I6J7K8L9M0N',
        userId: 'USER123459',
        clientIp: '192.168.1.103',
        userAgent: 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)',
        notifyUrl: 'https://example.com/api/pay/notify/stripe',
        returnUrl: 'https://example.com/pay/success',
        extra: JSON.stringify({ product_id: 'ENT001', quantity: 1 }, null, 2),
        createdAt: '2024-01-15 14:50:00',
        paidAt: '2024-01-15 14:51:15',
        expireAt: '2024-01-15 15:20:00',
        updatedAt: '2024-01-15 15:30:45',
      },
      {
        id: 'ORD202401150005',
        merchantOrderNo: 'MER202401150005',
        title: '季度会员',
        amount: 79.9,
        paidAmount: 79.9,
        fee: 0.8,
        channel: 'unionpay',
        channelName: '银联支付',
        status: 'paid',
        tradeNo: '2024011531040000000000',
        thirdPartyOrderNo: 'UP2024011531f2e4r5t6y7u8i9o1',
        userId: 'USER123460',
        clientIp: '192.168.1.104',
        userAgent: 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36',
        notifyUrl: 'https://example.com/api/pay/notify/unionpay',
        returnUrl: 'https://example.com/pay/success',
        extra: JSON.stringify({ product_id: 'QUARTER001', quantity: 1 }, null, 2),
        createdAt: '2024-01-15 15:00:00',
        paidAt: '2024-01-15 15:02:30',
        expireAt: '2024-01-15 15:30:00',
        updatedAt: '2024-01-15 15:02:30',
      },
    ]

    // 计算今日统计
    calculateTodayStats()
  } catch {
    message.error('订单列表加载失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadOrderList()
  onSearch()
})
</script>

<style scoped>
.pay-order-container {
  background-color: #f5f7fa;
}

.order-card {
  max-width: 1600px;
  margin: 0 auto;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding: 16px 0;
  border-bottom: 1px solid #f0f0f0;
}

.search-section {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: wrap;
}

.action-section {
  display: flex;
  align-items: center;
}

.stats-section {
  margin-bottom: 24px;
  padding: 20px;
  background-color: #fafafa;
  border-radius: 8px;
}

.order-table {
  margin-top: 16px;
}

.order-detail-content {
  padding: 16px 0;
}

@media (max-width: 768px) {
  .pay-order-container {
    padding: 16px;
  }

  .header-section {
    flex-direction: column;
    gap: 16px;
    align-items: flex-start;
  }

  .search-section {
    width: 100%;
  }

  .action-section {
    width: 100%;
    justify-content: flex-end;
  }
}
</style>
