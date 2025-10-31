<script setup lang="ts">
import { ref } from 'vue'
import {
  NLayout,
  NLayoutSider,
  NLayoutContent,
  NMenu,
  NCard,
  NDataTable,
  NButton,
  NModal,
  NForm,
  NFormItem,
  NInput,
  useMessage,
  NImage,
  NIcon,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { Pencil } from '@vicons/ionicons5'

const message = useMessage()

// 分组相关
const selectedGroup = ref('all')
const groups = ref([
  {
    label: '全部',
    key: 'all',
  },
])

// 友链表格数据
const loading = ref(false)
const data = ref([])
const columns: DataTableColumns = [
  {
    title: '网站名称',
    key: 'name',
  },
  {
    title: 'Logo',
    key: 'logo',
    render: (row: any) => {
      return row.logo ? h(NImage, { src: row.logo }) : '暂无'
    },
  },
  {
    title: '网站地址',
    key: 'url',
  },
  {
    title: '描述',
    key: 'description',
  },
  {
    title: '操作',
    key: 'actions',
    render: () => {
      return h(
        NButton,
        {
          size: 'small',
          type: 'primary',
          quaternary: true,
        },
        {
          icon: () => h(NIcon, {}, () => h(Pencil)),
          default: () => '编辑',
        },
      )
    },
  },
]

// 编辑弹窗相关
const showModal = ref(false)
const modalTitle = ref('')
const formRef = ref(null)
const formValue = ref({
  name: '',
  url: '',
  logo: '',
  description: '',
  group: 'all',
})

const handleAdd = () => {
  modalTitle.value = '添加友链'
  formValue.value = {
    name: '',
    url: '',
    logo: '',
    description: '',
    group: selectedGroup.value,
  }
  showModal.value = true
}

const handleEdit = (row: any) => {
  modalTitle.value = '编辑友链'
  formValue.value = { ...row }
  showModal.value = true
}

const handleDelete = async (row: any) => {
  // TODO: 实现删除逻辑
  message.success('删除成功喵~')
}

const handleSubmit = async () => {
  // TODO: 实现提交逻辑
  showModal.value = false
  message.success('保存成功喵~')
}
</script>

<template>
  <div class="container-fluid p-6">
    <n-card title="友链管理">
      <n-layout has-sider>
        <n-layout-sider
          bordered
          collapse-mode="width"
          :collapsed-width="64"
          :width="240"
          :native-scrollbar="false"
          style="height: 100%"
        >
          <n-menu
            :value="selectedGroup"
            :options="groups"
            @update:value="(key) => (selectedGroup = key)"
          />
        </n-layout-sider>
        <n-layout-content content-style="padding: 24px;">
          <n-card>
            <template #header>
              <div class="header-section">
                <div class="title">友链管理</div>
                <n-button type="primary" @click="handleAdd">添加友链</n-button>
              </div>
            </template>
            <n-data-table
              :loading="loading"
              :columns="columns"
              :data="data"
              :pagination="{ pageSize: 10 }"
            />
          </n-card>
        </n-layout-content>
      </n-layout>
    </n-card>
    <!-- 编辑弹窗 -->
    <n-modal
      v-model:show="showModal"
      :title="modalTitle"
      preset="dialog"
      @positive-click="handleSubmit"
    >
      <n-form ref="formRef" :model="formValue" label-placement="left" label-width="80">
        <n-form-item label="网站名称" path="name">
          <n-input v-model:value="formValue.name" placeholder="请输入网站名称" />
        </n-form-item>
        <n-form-item label="网站地址" path="url">
          <n-input v-model:value="formValue.url" placeholder="请输入网站地址" />
        </n-form-item>
        <n-form-item label="Logo" path="logo">
          <n-input v-model:value="formValue.logo" placeholder="请输入Logo地址" />
        </n-form-item>
        <n-form-item label="描述" path="description">
          <n-input
            v-model:value="formValue.description"
            type="textarea"
            placeholder="请输入网站描述"
          />
        </n-form-item>
      </n-form>
    </n-modal>
  </div>
</template>

<style scoped>
.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title {
  font-size: 16px;
  font-weight: 500;
}
</style>
