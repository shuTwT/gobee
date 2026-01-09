<script setup lang="ts">
import * as friendCircleRecordApi from '@/api/content/friendCircle'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'

dayjs.extend(utc)

const dataList = ref<any[]>([])

const onSearch = async () => {
  try {
    const res = await friendCircleRecordApi.getFriendCircleRecordPage()
    dataList.value = res.data.records
  } catch {}
}
onMounted(() => {
  onSearch()
})
</script>
<template>
  <div class="container-fluid p-6">
    <n-card class="max-w-[1600px] mx-auto">
      <div>
        <div class="w-[800px] mx-auto">
          <div class="flex w-full justify-between">
            <div class="w-[80px]">活跃友链数</div>
            <div>
              <progress></progress>
            </div>
            <div>0/0</div>
          </div>
          <div class="flex w-full justify-between">
            <div class="w-[80px]">失联友链数</div>
            <div>
              <progress></progress>
            </div>
            <div>0/0</div>
          </div>
          <div class="flex w-full justify-between">
            <div class="w-[80px]">当前库存</div>
            <div>
              <progress></progress>
            </div>
            <div>0/0</div>
          </div>
        </div>
        <div class="article-list flex flex-row flex-wrap justify-start">
          <div v-for="(item, index) in dataList" :key="index" class="w-[25%]">
            <div class="article relative flex flex-col p-4 m-2 justify-between">
              <a class="title" :href="item.link_url" target="_blank">{{ item.title }}</a>
              <div class="flex absolute bottom-0 right-0 opacity-10 rounded-4xl pointer-events-none">
                <n-avatar :src="item.avatar_url" :size="90"/>
              </div>

              <div class="flex w-full justify-between">
                <a :href="item.site_url" :title="item.author">
                  <span class="author">{{ item.author }}</span>
                </a>
                <span class="time">{{
                  item.published_at
                    ? dayjs.utc(item.published_at).format('YYYY-MM-DD HH:mm:ss')
                    : ''
                }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </n-card>
  </div>
</template>
<style scoped>
.article {
  background: #fff;
  border: 1px solid #e3e8e9;
  height: 160px;
}
</style>
