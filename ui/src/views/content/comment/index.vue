<script setup lang="ts">
import { ref, onMounted } from 'vue'
import {
  NDataTable,
  NButton,
  NSpace,
  NModal,
  NDescriptions,
  NDescriptionsItem,
  NTag,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import {
  getCommentPage,
  getCommentDetail,
  approveComment,
  rejectComment,
  deleteComment,
  type Comment,
  CommentStatus,
} from '@/api/content/comment'
import {apiClient,useApi} from "@/api"

const message = useMessage()

// 表格数据
const loading = ref(false)
const data = ref<Comment[]>([])
const pagination = ref({
  page: 1,
  pageSize: 10,
  itemCount: 0,
  showSizePicker: true,
  pageSizes: [10, 20, 30, 40],
})

// 详情弹窗
const showModal = ref(false)
const currentComment = ref<Comment | null>(null)

// 获取评论列表
const fetchComments = async () => {
  loading.value = true
  try {
    const res = await useApi(apiClient.api.v1CommentPageList,{
      page: pagination.value.page,
      pageSize: pagination.value.pageSize,
    })
    data.value = res.data.records
    pagination.value.itemCount = res.data.total
  } catch {
    message.error('获取评论列表失败')
  } finally {
    loading.value = false
  }
}

// 查看评论详情
const handleView = async (row: Comment) => {
  try {
    const res = await useApi(apiClient.api.v1CommentDetail,row.id)
    currentComment.value = res.data
    showModal.value = true
  } catch {
    message.error('获取评论详情失败')
  }
}

// 审核评论
const handleApprove = async (row: Comment) => {
  try {
    await approveComment(row.id)
    message.success('审核通过成功')
    fetchComments()
  } catch  {
    message.error('审核失败')
  }
}

// 拒绝评论
const handleReject = async (row: Comment) => {
  try {
    await rejectComment(row.id)
    message.success('拒绝评论成功')
    fetchComments()
  } catch {
    message.error('操作失败')
  }
}

// 删除评论
const handleDelete = async (row: Comment) => {
  try {
    await deleteComment(row.id)
    message.success('删除成功')
    fetchComments()
  } catch {
    message.error('删除失败')
  }
}

// 表格列定义
const columns: DataTableColumns<Comment> = [
  {
    title: 'ID',
    key: 'id',
    width: 80,
  },
  {
    title: '评论内容',
    key: 'content',
    ellipsis: true,
  },
  {
    title: '评论者',
    key: 'user',
    render: (row) => row.user?.name || '未知用户',
  },
  {
    title: '评论文章',
    key: 'article',
    render: (row) => row.article?.title || '未知文章',
  },
  {
    title: '评论时间',
    key: 'created_at',
    render: (row) => new Date(row.created_at).toLocaleString(),
  },
  {
    title: '状态',
    key: 'status',
    render: (row) => {
      if (row.status === CommentStatus.Pending) {
        return h(NTag, { type: 'warning' }, { default: () => '待审核' })
      } else if (row.status === CommentStatus.Approved) {
        return h(NTag, { type: 'success' }, { default: () => '已通过' })
      }
      return '-'
    },
  },
  {
    title: '操作',
    key: 'actions',
    render: (row) => {
      return h(
        NSpace,
        {},
        {
          default: () => [
            h(
              NButton,
              {
                size: 'small',
                onClick: () => handleView(row),
              },
              { default: () => '详情' },
            ),
            row.status === CommentStatus.Pending
              ? [
                  h(
                    NButton,
                    {
                      size: 'small',
                      type: 'primary',
                      onClick: () => handleApprove(row),
                    },
                    { default: () => '通过' },
                  ),
                  h(
                    NButton,
                    {
                      size: 'small',
                      type: 'warning',
                      onClick: () => handleReject(row),
                    },
                    { default: () => '拒绝' },
                  ),
                ]
              : null,
            h(
              NButton,
              {
                size: 'small',
                type: 'error',
                onClick: () => handleDelete(row),
              },
              { default: () => '删除' },
            ),
          ],
        },
      )
    },
  },
]

// 监听分页变化
const handlePageChange = (page: number) => {
  pagination.value.page = page
  fetchComments()
}

const handlePageSizeChange = (pageSize: number) => {
  pagination.value.pageSize = pageSize
  pagination.value.page = 1
  fetchComments()
}

// 初始化
onMounted(() => {
  fetchComments()
})
</script>

<template>
  <div class="container-fluid p-6">
    <n-card title="评论管理" class="comment-card">
      <!-- 头部操作栏 -->
       <div class="header-section">
        <div class="search-section"></div>
        <div class="action-section"></div>
       </div>
      <!-- 文章列表 -->
      <n-data-table
        :loading="loading"
        :columns="columns"
        :data="data"
        :pagination="pagination"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
        :remote="true"
      />
    </n-card>

    <!-- 评论详情弹窗 -->
    <n-modal v-model:show="showModal" preset="card" title="评论详情" style="width: 600px">
      <n-descriptions v-if="currentComment" bordered tabindex="1">
        <n-descriptions-item label="评论者">
          {{ currentComment.user?.name || '未知用户' }}
        </n-descriptions-item>
        <n-descriptions-item label="评论文章">
          {{ currentComment.article?.title || '未知文章' }}
        </n-descriptions-item>
        <n-descriptions-item label="评论内容">
          {{ currentComment.content }}
        </n-descriptions-item>
        <n-descriptions-item label="IP地址">
          {{ currentComment.ip_address }}
        </n-descriptions-item>
        <n-descriptions-item label="IP位置">
          {{ currentComment.ip_location || '-' }}
        </n-descriptions-item>
        <n-descriptions-item label="浏览器信息">
          {{ currentComment.user_agent || '-' }}
        </n-descriptions-item>
        <n-descriptions-item label="评论时间">
          {{ new Date(currentComment.created_at).toLocaleString() }}
        </n-descriptions-item>
        <n-descriptions-item label="状态">
          <n-tag :type="currentComment.status === CommentStatus.Approved ? 'success' : 'warning'">
            {{ currentComment.status === CommentStatus.Approved ? '已通过' : '待审核' }}
          </n-tag>
        </n-descriptions-item>
      </n-descriptions>
    </n-modal>
  </div>
</template>

<style scoped>
.comment-manager {
  padding: 16px;
}
.comment-card {
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
</style>
