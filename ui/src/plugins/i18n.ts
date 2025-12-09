import type { App } from 'vue'

import type { NLocale } from 'naive-ui/es/locales/common/enUS'
import enLocales from 'naive-ui/es/locales/common/enUS'
import zhLocales from 'naive-ui/es/locales/common/zhCN'

const siphonI18n = (function () {
  const cache = Object.fromEntries(
    Object.entries(import.meta.glob('../../locales/*.y(a)?ml', { eager: true })).map(
      ([key, value]: any) => {
        const matched = key.match(/([A-Za-z0-9-_]+)\./i)[1]
        return [matched, value.default]
      },
    ),
  )
  return (prefix="zh-CN")=>{
    return cache[prefix]
  }
})()

export const localesConfigs:Record<string,NLocale> = {
  zh:{
    ...siphonI18n("zh-CN"),
    ...zhLocales
  },
  en:{
    ...siphonI18n("en-US"),
    ...enLocales
  }
}

export function useI18n(app: App) {}

export const enum Locales {
  zh = "zh-CN",
  en = "en-US",
}
