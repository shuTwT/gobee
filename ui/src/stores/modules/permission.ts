import { defineStore } from "pinia";
import {constantMenus} from "@/router"
import { store } from '../index'
import {ascending, filterTree} from "@/router/utils"

export const usePermissionStore = defineStore('permission',()=>{
  /**
   * 静态路由生成的菜单
   */
  const constantMenusRef = ref(constantMenus)
  /**
   * 整体路由生成的菜单(静态、动态)
   */
  const wholeMenus = ref<unknown[]>([])
  /**
   * 整体路由（一维数组格式）
   */
  const flatteningRoutes= ref([])
  /**
   * 缓存页面keepAlive
   */
  const cachePageList= ref([])
  function handleWholeMenus(menus: unknown[]){
    //!TODO 用户角色过滤无权限菜单
    wholeMenus.value = filterTree(ascending(menus))
  }
  /**
   * 清空缓存页面
   */
  function clearAllCachePage(){
    wholeMenus.value = []
    cachePageList.value = []
  }
  return {
    constantMenusRef,
    wholeMenus,
    flatteningRoutes,
    cachePageList,
    handleWholeMenus,
    clearAllCachePage
  }
})

export function usePermissionStoreHook(){
  return usePermissionStore(store)
}
