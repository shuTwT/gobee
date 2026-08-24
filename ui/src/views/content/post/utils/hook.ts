import { NButton, NPopconfirm, type MessageReactive } from 'naive-ui'
import { apiClient, useApi } from '@/api'
import { addDialog } from '@/components/dialog'
import type { FormProps } from './types'
import SettingForm from '../settingForm.vue'
import ImportForm from '../importForm.vue'
import ShareDialog from '../shareDialog.vue'

export function usePostHook() {
  const message = useMessage()
  const dialog = useDialog()

  // 保存文章
  const savePost = (row: any) => {
    return new Promise((resolve, reject) => {
      useApi(apiClient.api.v1PostUpdateContentUpdate, row.id, {
          content: row.content,
          md_content: row.md_content,
          html_content: row.html_content,
        } as Parameters<typeof apiClient.api.v1PostUpdateContentUpdate>[1])
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
      useApi(apiClient.api.v1PostQueryDetail, row.id)
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
                categoryIds: res.data.categoryIds ?? [],
                tagIds: res.data.tagIds ?? [],
                is_autogen_summary: res.data.is_autogen_summary ?? false,
                author: res.data.author ?? '',
                is_allow_comment: res.data.is_allow_comment ?? true,
                is_pin_to_top: res.data.is_pin_to_top ?? false,
                is_visible: res.data.is_visible ?? true,
                is_visible_after_comment: res.data.is_visible_after_comment ?? false,
                is_visible_after_pay: res.data.is_visible_after_pay ?? false,
                price: res.data.price ?? 0,
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
                const chores = () => {
                  message.success('更新成功喵~')
                  done()
                  resolve(true)
                }
                useApi(apiClient.api.v1PostUpdateSettingUpdate, row.id, curData).then(() => {
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
          useApi(apiClient.api.v1PostPublishUpdate, row.id)
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
          useApi(apiClient.api.v1PostUnpublishUpdate, row.id)
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
    return new Promise<void>((resolve, reject) => {
      Promise.all([
        useApi(apiClient.api.v1PostQueryDetail, row.id),
        useApi(apiClient.api.v1SettingsJsonDetail, 'basic'),
      ])
        .then(([postRes, settingRes]) => {
          const post = postRes.data
          const siteUrl = (settingRes.data?.siteUrl || settingRes.data?.site_url || '').replace(
            /\/$/,
            '',
          )
          if (!siteUrl) {
            message.warning('请先在系统设置中配置站点地址')
            reject(new Error('site_url not configured'))
            return
          }
          const slug = post.slug || post.id
          const shareUrl = `${siteUrl}/post/${slug}`

          const dialogRef = ref()
          addDialog({
            title: '分享文章',
            contentRenderer: ({ options }) =>
              h(ShareDialog, {
                ref: dialogRef,
                formInline: {
                  title: post.title,
                  shareUrl,
                },
              }),
            beforeSure: (done) => {
              done()
              resolve()
            },
            positiveText: '关闭',
            negativeText: undefined,
          })
        })
        .catch((err) => {
          message.error('获取文章信息失败')
          reject(err)
        })
    })
  }

  // 导出文章
  const exportPost = (row: any) => {
    return new Promise<void>((resolve, reject) => {
      useApi(apiClient.api.v1PostQueryDetail, row.id)
        .then((res) => {
          const post = res.data
          const content = post.md_content || post.content || ''
          if (!content.trim()) {
            message.warning('文章内容为空，无法导出')
            reject(new Error('empty content'))
            return
          }

          const safeTitle = post.title.replace(/[\\/:*?"<>|\r\n\t]/g, '_')
          const fileName = `${safeTitle}.md`

          const blob = new Blob([content], { type: 'text/markdown;charset=utf-8' })
          const url = URL.createObjectURL(blob)
          const a = document.createElement('a')
          a.href = url
          a.download = fileName
          document.body.appendChild(a)
          a.click()
          document.body.removeChild(a)
          URL.revokeObjectURL(url)

          message.success('导出成功')
          resolve()
        })
        .catch((err) => {
          message.error('导出失败')
          reject(err)
        })
    })
  }

  // 复制文章内容
  const copyPostContent = (row: any) => {
    return new Promise<void>((resolve, reject) => {
      useApi(apiClient.api.v1PostQueryDetail, row.id)
        .then((res) => {
          const post = res.data
          const content = post.md_content || post.content || ''
          if (!content.trim()) {
            message.warning('文章内容为空，无法复制')
            reject(new Error('empty content'))
            return
          }

          const doCopy = async () => {
            try {
              await navigator.clipboard.writeText(content)
            } catch {
              const textarea = document.createElement('textarea')
              textarea.value = content
              textarea.style.position = 'fixed'
              textarea.style.opacity = '0'
              document.body.appendChild(textarea)
              textarea.select()
              document.execCommand('copy')
              document.body.removeChild(textarea)
            }
            message.success('内容已复制到剪贴板')
            resolve()
          }
          doCopy()
        })
        .catch((err) => {
          message.error('复制失败')
          reject(err)
        })
    })
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
