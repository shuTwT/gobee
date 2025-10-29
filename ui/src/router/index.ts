import { createRouter, createWebHistory, type Router, type RouteRecordRaw } from 'vue-router'
import BaseLayout from '@/layout/index.vue'
import remainingRouter from './modules/remaining'
import { useSettingsStoreHook } from '@/stores/modules/settings'
import { multipleTabsKey, removeToken, userKey, type DataInfo } from '@/utils/auth'
import { isUrl, useStorageLocal } from '@/utils/utils'
import Cookies from 'js-cookie'
import { ascending, formatFlatteningRoutes, formatTwoStageRoutes, initRouter, isOneOfArray } from './utils'
import { usePermissionStoreHook } from '@/stores/modules/permission'
import { buildHierarchyTree } from '@/utils/tree'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const modules: Record<string, any> = import.meta.glob(
  ['./modules/**/*.ts', '!./modules/**/remaining.ts'],
  {
    eager: true,
  },
)

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const routes: any[] = []

Object.keys(modules).forEach((key) => {
  routes.push(...modules[key].default)
})

/** 导出处理后的静态路由（三级及以上的路由全部拍成二级） */
export const constantRoutes: Array<RouteRecordRaw> = formatTwoStageRoutes(
  formatFlatteningRoutes(buildHierarchyTree(ascending(routes.flat(Infinity))))
);

export const constantMenus:Array<RouteRecordRaw> = ascending(routes.flat(Infinity)).concat(...remainingRouter)

/** 不参与菜单的路由 */
export const remainingPaths = Object.keys(remainingRouter).map((v:any) => {
  return remainingRouter[v].path;
});

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes.concat(...(remainingRouter as any)),
})

const whiteList = ['/login', '/initialize']

router.beforeEach((to, _from, next) => {
  const userInfo = useStorageLocal().getItem<DataInfo<number>>(userKey)
  const externalLink = isUrl(to.path)

  /** 如果已经登录并存在登录信息后不能跳转到路由白名单，而是继续保持在当前页面 */
  function toCorrectRoute() {
    if (whiteList.includes(to.fullPath)) {
      next(_from.fullPath)
    } else {
      next()
    }
  }

  if (Cookies.get(multipleTabsKey) && userInfo) {
    // 无权限跳转 403 页面
    if (to.meta?.roles && !isOneOfArray(to.meta?.roles as string[] | undefined, userInfo.roles)) {
      next({ path: '/error/403' })
    }

    if (_from?.name) {
      if (externalLink) {
        window.open(to?.name as string)
      } else {
        toCorrectRoute()
      }
    } else {

      // 刷新
      if(usePermissionStoreHook().wholeMenus.length===0 &&to.path!=="/login"){
        initRouter().then((router:Router)=>{
          if(typeof to.name==="string"&&to.name!==""){
            router.push(to.fullPath)
          }
        })
      }
      toCorrectRoute()
    }
  } else {
    if (to.path === '/initialize') {
      const settingsStore = useSettingsStoreHook()
      if (settingsStore.initialized) {
        next({ path: '/' })
      } else {
        next()
      }
    } else if (to.path !== '/login') {
      if (whiteList.indexOf(to.path) !== -1) {
        next()
      } else {
        removeToken()
        next({ path: '/login' })
      }
    } else {
      next()
    }
  }
})

export function resetRouter(){
  router.getRoutes().forEach((route)=>{
    const {name,meta} = route;
    if(name && router.hasRoute(name)&&meta?.backstage){
      router.removeRoute(name)

    }
  })
}

export default router
