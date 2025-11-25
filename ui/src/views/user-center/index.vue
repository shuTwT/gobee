<script setup lang="ts">
import { NButton, NDescriptions,NDescriptionsItem, NSpace } from 'naive-ui'
import * as userApi from "@/api/system/user"
import personalAccessTokenForm from './personalAccessTokenForm.vue'
import { addDialog } from '@/components/dialog'
import type { TableColumn } from 'naive-ui/es/data-table/src/interface'
import dayjs from 'dayjs'
import { useClipboard } from '@vueuse/core'

const message = useMessage()
const colorOptions = ref([
  {
    value: 'light',
    label: '亮色',
  },
  {
    value: 'dark',
    label: '暗色',
  },
])
const personalInfomation = ref({
  username: '',
  email: '',
  nickname: '',
  bio: '',
})
const personalizedSetting = ref({
  theme: 'light',
  notifications: false,
})
const personalAccessTokenList = ref([])
const personalAccessTokenColumns = ref<TableColumn[]>([
  {
    title: '令牌名称',
    key: 'name',
  },
  {
    title: '过期时间',
    key: 'expires',
    render:(row:any)=>{
      return dayjs(row.expires).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    title:"操作",
    key:'actions',
    render:(row)=>{
      return h(NSpace,{},{
        default:()=>[
          h(NButton,{
            text:true,
            type:'primary',
            onClick:()=>{
              copyToken(row)
            }
          },{
            default:()=>'复制'
          })
        ]
      })
    }
  }
])

const activeTab = ref('profile')
const copyTokenOrigin = ref('')

const {copy,isSupported} = useClipboard({source:copyTokenOrigin})

const onSearchUserPersonalAccessToken = async()=>{
  const res= await userApi.getPersonalAccessTokenList()
  personalAccessTokenList.value =  res.data
}
const openPatEditDialog = (title='新增')=>{
  const formRef = ref()
  addDialog({
    title:`${title}个人令牌`,
    props:{},
    contentRenderer:()=>h(personalAccessTokenForm,{ref:formRef}),
    beforeSure:async (done)=>{
      const chores=()=>{
        message.success('创建成功')
        done()
        onSearchUserPersonalAccessToken()
      }
      try{
        const curData =await formRef.value.getData()
        await userApi.createPat(curData)
        chores()
      }catch{

      }

    }
  })
}
const copyToken = async (row:any)=>{
  try{
    const res = await userApi.getPersonalAccessToken(row.id)
    copyTokenOrigin.value = res.data.token
    if(isSupported.value){
      copy()
      message.success('已复制到剪切板')
    }else{
      message.error('剪切板不可用')
    }
  }catch{

  }

}
onMounted(()=>{
  onSearchUserPersonalAccessToken()
})
</script>
<template>
  <div class="p-4 sm:p-6 lg:p-8">
    <h2 class="text-2xl font-semibold text-gray-900 dark:text-white mb-6">用户中心</h2>

    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6 mb-6">
      <n-tabs v-model:value="activeTab" type="segment" animated>
        <n-tab-pane name="profile" tab="详情"></n-tab-pane>
        <n-tab-pane name="setting" tab="个性化"></n-tab-pane>
        <n-tab-pane name="personalAccessToken" tab="个人令牌"></n-tab-pane>
      </n-tabs>
    </div>
    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6 mb-6">
      <template v-if="activeTab === 'profile'">
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">个人信息</h3>

        <n-descriptions label-placement="left" bordered :columns="1">
          <n-descriptions-item>
            <template #label>
              用户名
            </template>
            {{ personalInfomation.username }}
          </n-descriptions-item>
          <n-descriptions-item>
            <template #label>
              邮箱
            </template>
            {{ personalInfomation.email }}
          </n-descriptions-item>
          <n-descriptions-item>
            <template #label>
              昵称
            </template>
            {{ personalInfomation.nickname }}
          </n-descriptions-item>
          <n-descriptions-item>
            <template #label>
              个人简介
            </template>
            {{ personalInfomation.bio }}
          </n-descriptions-item>
        </n-descriptions>
      </template>
      <template v-else-if="activeTab == 'setting'">
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">个性化设置</h3>
        <n-form label-placement="left">
          <div class="mb-4">
            <n-form-item label="主题">
              <n-select v-model:value="personalizedSetting.theme" :options="colorOptions" />
            </n-form-item>
          </div>
          <div class="mb-4">
            <n-form-item label="接收通知">
              <n-checkbox v-model:checked="personalizedSetting.notifications" />
            </n-form-item>
          </div>
          <div class="flex justify-end">
            <button
              type="submit"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-gray-800"
            >
              保存设置
            </button>
          </div>
        </n-form>
      </template>
      <template v-else-if="activeTab === 'personalAccessToken'">
        <h3 class="text-xl font-semibold text-gray-900 dar
        k:text-white mb-4">个人令牌</h3>
        <div class="flex justify-between">
          <p class="text-gray-700 dark:text-gray-300 mb-4">
            个人令牌用于访问 API。请妥善保管，不要分享给他人。
          </p>
          <div>
            <n-button type="primary" @click="openPatEditDialog('新增')">创建</n-button>
          </div>
        </div>

        <n-data-table :columns="personalAccessTokenColumns" :data="personalAccessTokenList" />
      </template>
    </div>
  </div>
</template>
