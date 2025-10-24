import './assets/main.css'
import {getSettings} from './api/setting'

import { createApp } from 'vue'
import { setupStore } from './stores'

import App from './App.vue'
import router from './router'
import { useSettingsStoreHook } from './stores/modules/settings'

getSettings().then(res=>{
    const app = createApp(App)

    setupStore(app)
    app.use(router)

    app.mount('#app')
    if(!res.data.initialized){
      router.push('/initialize')
    }else{
      useSettingsStoreHook().initialize()
    }
})

