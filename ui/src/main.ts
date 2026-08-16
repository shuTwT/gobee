import './assets/main.css'
import { apiClient, useApi } from '@/api'

import { createApp } from 'vue'
import { setupStore } from './stores'

import App from './App.vue'
import router from './router'
import { useSettingsStoreHook } from './stores/modules/settings'

const app = createApp(App)

useApi(apiClient.api.v1SettingsList).then(async(res) => {
  setupStore(app)
  app.use(router)

  await router.isReady()

  app.mount('#app')
  if (!res.data.initialized) {
    router.push('/initialize')
  } else {
    useSettingsStoreHook().initialize()
  }
})
