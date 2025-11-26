<script setup lang="ts">
import { NCard, NDataTable, NButton, NSpace, useMessage, NIcon } from 'naive-ui'
import { RefreshOutline,Pencil } from '@vicons/ionicons5'
import type { DataTableColumns } from 'naive-ui'
import * as friendCircleRuleApi from "@/api/friendCircle/friendCircleRule"
import FormComponent from "./form.vue"
import { addDialog } from '@/components/dialog'
const message = useMessage()
const dialog = useDialog()

// 表格数据
const loading = ref(false)
const dataList = ref<any[]>([])
const columns: DataTableColumns<any> = [
  {
    title: '规则名称',
    key: 'name',
  },
  {
    title: '操作',
    key: 'actions',
    render(row) {
      return h(
        NSpace,
        {},
        {
          default: () => [
            h(
              NButton,
              {
                size: 'small',
                quaternary: true,
                type: 'primary',
                onClick: () => openEditDialog('编辑',row),
              },
              {
                icon: () => h(NIcon, null, { default: () => h(Pencil) }),
                default: () => '编辑',
              },
            ),
            h(
              NButton,
              {
                size: 'small',
                quaternary: true,
                type: 'error',
                disabled: row.master,
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


// 获取存储策略列表
const onSearch = async () => {
  loading.value = true
  try {
    const res = await friendCircleRuleApi.getFriendCircleRulePage()
    dataList.value = res.data.records
  } catch (error) {
    message.error('获取规则列表失败：' + (error as Error).message)
  } finally {
    loading.value = false
  }
}

const openEditDialog = (title='新增',row?:any) => {
  const formRef = ref<any>()
  addDialog({
    title:`${title}规则`,
    props: {
      formInline: {
        id: row?.id ?? undefined,
        name: row?.name ?? '',
        title_selector:row?.title_selector ?? '',
        link_selector:row?.link_selector ??'',
        created_selector:row?.created_selector??'',
        updated_selector:row?.updated_selector??''
      },
    },
    contentRenderer: ({ options }) =>
      h(FormComponent, { ref: formRef, formInline: options.props!.formInline }),
    beforeSure: async (done) => {
      try {
        const curData = await formRef.value?.getData()
        console.log(curData)
        const chores = () => {
          message.success('创建成功喵~')
          done()
          onSearch()
        }
        if(title=='新增'){
          friendCircleRuleApi.createFriendCircleRule(curData).then(() => {
            chores()
          })
        }else{
          friendCircleRuleApi.updateFriendCircleRule(row.id,curData).then(() => {
            chores()
          })
        }

      } catch {}
    },
  })
}

const handleDelete = async (row: any) => {
  dialog.warning({
    title: '确认删除',
    content: '确定要删除该规则吗？删除后无法恢复！',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: async () => {
      try {
        await friendCircleRuleApi.deleteFriendCircleRule(row.id)
        message.success('删除成功喵~')
        onSearch()
      } catch (error) {
        message.error('删除失败：' + (error as Error).message)
      }
    },
  })
}

onMounted(async ()=>{
  try{
    const res =await friendCircleRuleApi.getFriendCircleRulePage()
    dataList.value = res.data.records
  }catch{

  }
})
</script>
<template>
  <div class="container-fluid p-6">
    <n-card title="朋友圈规则管理" class="rule-card">
      <!-- 头部操作栏 -->
       <div class="header-section">
        <div class="search-section"></div>
        <div class="action-section">
          <n-button type="primary" @click="openEditDialog('新增')" style="margin-right: 12px">添加规则</n-button>
          <n-button @click="onSearch()">
            <template #icon>
              <n-icon><refresh-outline /></n-icon>
            </template>
            刷新
          </n-button>
        </div>
       </div>
      <n-data-table
        :loading="loading"
        :columns="columns"
        :data="dataList"
        :pagination="{ pageSize: 10 }"
      />
    </n-card>
  </div>
</template>

<style scoped>
.rule-card {
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
