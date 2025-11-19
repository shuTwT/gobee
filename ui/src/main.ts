import './assets/main.css'
import { getSettings } from './api/system/setting'

import { createApp } from 'vue'
import { setupStore } from './stores'

import App from './App.vue'
import router from './router'
import { useSettingsStoreHook } from './stores/modules/settings'

const app = createApp(App)

getSettings().then(async(res) => {
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
