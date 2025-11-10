import type { MessageReactive } from 'naive-ui'
import * as postApi from '@/api/post'
import { addDialog } from '@/components/dialog'
import type { FormProps } from './types'
import SettingForm from '../settingForm.vue'

export function usePostHook() {
  const message = useMessage()
  const dialog = useDialog()

  // 保存文章
  const savePost = (row:any)=>{
    postApi.updatePostContent(row.id,{
      content:row.content
    }).then((res)=>{
      message.success('保存成功')
    })
  }

  // 文章设置
  const settingPost = (row: any) => {
    let messsageReactive: MessageReactive | null = null
    messsageReactive = message.info('加载中...', {
      duration: 0,
    })
    postApi
      .queryPost(row.id)
      .then((res) => {
        const formRef = ref()
        addDialog<FormProps>({
          title: `文章设置`,
          props: {
            formInline: {
              id: row.id ?? undefined,
              title: res.data.title ?? '',
              alias: res.data.alias ?? '',
              categorys: res.data.categorys ?? [],
              tags: res.data.tags ?? [],
              autogenSummary: res.data.autogenSummary ?? false,
              author: res.data.author ?? '',
              allowComments: res.data.allowComments ?? false,
              pinToTop: res.data.pinToTop ?? false,
              visible: res.data.visible ?? false,
              cover: res.data.cover ?? '',
            },
          },
          contentRenderer: ({ options }) =>
            h(SettingForm, { ref: formRef, formInline: options.props!.formInline }),
          beforeSure: async (done) => {
            try {
              const curData = await formRef.value?.getData()
              console.log(curData)
              const chores = () => {
                message.success('更新成功喵~')
                done()

              }
              postApi.updatePostSetting(row.id,curData).then(() => {
                chores()
              })

            } catch {}
          },
        })
      })
      .catch((err) => {})
      .finally(() => {
        messsageReactive.destroy()
      })
  }

  // 发布文章
  const publishPost = (row: any) => {
    dialog.info({
      title: '确认',
      content: '确定要发布该文章吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        row.status = 'published'
        message.success('发布成功')
      },
    })
  }

  // 分享文章
  const sharePost = (row: any) => {
    // TODO: 实现分享功能
    message.info('分享功能开发中')
  }

  // 导出文章
  const exportPost = (row: any) => {
    // TODO: 实现导出功能
    message.info('导出功能开发中')
  }

  // 复制文章内容
  const copyPostContent = (row: any) => {
    // TODO: 实现复制内容功能
    message.info('复制内容功能开发中')
  }

  return {
    savePost,
    settingPost,
    publishPost,
    sharePost,
    exportPost,
    copyPostContent,
  }
}
