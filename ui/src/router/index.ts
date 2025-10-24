import { createRouter, createWebHistory } from 'vue-router'
import BaseLayout from '@/layout/index.vue'
import remainingRouter from "./modules/remaining";
import { useSettingsStoreHook } from '@/stores/modules/settings';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const modules: Record<string, any> = import.meta.glob(
  [
    "./modules/**/*.ts",
    "!./modules/**/remaining.ts",
  ],
  {
    eager: true
  }
);

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const routes:any[] = []

Object.keys(modules).forEach(key => {
  routes.push(modules[key].default);
});

export const constantMenus = routes.concat(...remainingRouter);

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes:constantMenus
})

const whiteList = ['/login','/initialize']

router.beforeEach((to,_from,next)=>{

  if(to.path === '/initialize'){
    const settingsStore = useSettingsStoreHook()
    if(settingsStore.initialized){
      next({path:'/'})
    }else{
      next()
    }
  }else if(to.path !== '/login'){
    if(whiteList.indexOf(to.path) !== -1){
      next()
    }else{
      next({path:'/login'})
    }
  }else{
    next()
  }
})

export default router
