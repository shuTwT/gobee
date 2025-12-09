<script setup lang="ts">
import { RouterView } from 'vue-router'
import { zhCN, dateZhCN, enUS,dateEnUS } from 'naive-ui/es/locales'
import hljs from 'highlight.js/lib/core'
// import javascriptLang from 'highlight.js/lib/languages/javscript'
import jsonLang from 'highlight.js/lib/languages/json'
import DialogManager from '@/components/dialog/DialogManager.vue'
import { useAppStore } from './stores/modules/app'
import { Locales } from './plugins/i18n'

const appStore = useAppStore()
// hljs.registerLanguage('javascript',javascriptLang)
hljs.registerLanguage('json',jsonLang)

const locale = computed(()=>{
  if(appStore.language === Locales.zh){
    return zhCN
  }
  return enUS
})
const dateLocale = computed(()=>{
  if(appStore.language === Locales.zh){
    return dateZhCN
  }
  return dateEnUS
})
</script>

<template>
  <n-config-provider class="w-full h-full" :locale="locale" :date-locale="dateLocale" :hljs="hljs">
    <n-notification-provider>
      <n-message-provider>
        <n-modal-provider>
          <n-dialog-provider>
            <RouterView />
            <DialogManager />
          </n-dialog-provider>
        </n-modal-provider>
      </n-message-provider>
    </n-notification-provider>
  </n-config-provider>
</template>
