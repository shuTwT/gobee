import { NButton, NPopconfirm, type MessageReactive } from 'naive-ui'
import * as postApi from '@/api/content/post'
import { addDialog } from '@/components/dialog'
import type { FormProps } from './types'
import SettingForm from '../settingForm.vue'
import ImportForm from '../importForm.vue'

export function usePostHook() {
  const message = useMessage()
  const dialog = useDialog()

  // 保存文章
  const savePost = (row: any) => {
    return new Promise((resolve, reject) => {
      postApi
        .updatePostContent(row.id, {
          content: row.content,
        })
        .then((res) => {
          message.success('保存成功')
          resolve(true)
        })
        .catch((err) => {
          reject(err)
        })
    })
  }

  // 文章设置
  const settingPost = (row: any) => {
    let messsageReactive: MessageReactive | null = null
    messsageReactive = message.info('加载中...', {
      duration: 0,
    })
    return new Promise((resolve, reject) => {
      postApi
        .queryPost(row.id)
        .then((res) => {
          const formRef = ref()
          addDialog<FormProps>({
            title: `文章设置`,
            scroll: true,
            scrollbarHeight: '600px',
            props: {
              formInline: {
                id: res.data.id ?? undefined,
                title: res.data.title ?? '',
                alias: res.data.alias ?? '',
                content: res.data.content ?? '',
                md_content: res.data.md_content ?? '',
                html_content: res.data.html_content ?? '',
                content_type: res.data.content_type ?? 'markdown',
                status: res.data.status ?? 'draft',
                categories: res.data.categories ?? [],
                tags: res.data.tags ?? [],
                is_autogen_summary: res.data.is_autogen_summary ?? false,
                author: res.data.author ?? '',
                is_allow_comment: res.data.is_allow_comment ?? true,
                is_pin_to_top: res.data.is_pin_to_top ?? false,
                is_visible: res.data.is_visible ?? true,
                is_visible_after_comment: res.data.is_visible_after_comment ?? false,
                is_visible_after_pay: res.data.is_visible_after_pay ?? false,
                price: (res.data.price ?? 0).toString(),
                cover: res.data.cover ?? '',
                keywords: res.data.keywords ?? '',
                copyright: res.data.copyright ?? '',
                summary: res.data.summary ?? '',
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
                  resolve(true)
                }
                postApi.updatePostSetting(row.id, curData).then(() => {
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
    })
  }

  // 发布文章
  const publishPost = (row: any) => {
    return new Promise((resolve, reject) => {
      dialog.info({
        title: '确认',
        content: '确定要发布该文章吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          postApi
            .publishPost(row.id)
            .then(() => {
              message.success('发布成功')
              resolve(true)
            })
            .catch(() => {
              message.error('发布失败')
              reject(false)
            })
        },
      })
    })
  }

  // 取消发布文章
  const unpublishPost = (row: any) => {
    return new Promise((resolve, reject) => {
      dialog.info({
        title: '确认',
        content: '确定要取消发布该文章吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          postApi
            .unpublishPost(row.id)
            .then(() => {
              message.success('取消发布成功')
              resolve(true)
            })
            .catch(() => {
              message.error('取消发布失败')
              reject(false)
            })
        },
      })
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

  /**
   * 导入文章
   * @param row
   */
  const importPost = (row: any) => {
    const formRef = ref()
    return new Promise((resolve, reject) => {
      addDialog({
        title: '导入文章',
        contentRenderer: () => h(ImportForm,{ref:formRef}),
        beforeSure: async (done) => {
          try {
            const curData = await formRef.value?.getData()
            resolve(curData)
            done()
          } catch {}
        },
      })
    })
  }

  return {
    savePost,
    settingPost,
    publishPost,
    unpublishPost,
    sharePost,
    exportPost,
    copyPostContent,
    importPost,
  }
}
