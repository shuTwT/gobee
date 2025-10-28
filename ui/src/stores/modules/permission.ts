import { defineStore } from "pinia";
import {constantMenus} from "@/router"
import { store } from '../index'
import {ascending, filterTree, filterNoPermissionTree, formatFlatteningRoutes} from "@/router/utils"
import type { RouteRecordRaw } from "vue-router";

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
  function handleWholeMenus(routes: RouteRecordRaw[]){
    //!TODO 用户角色过滤无权限菜单
    wholeMenus.value = filterNoPermissionTree(
      filterTree(ascending(constantMenusRef.value.concat(routes)))
    )
    flatteningRoutes.value = formatFlatteningRoutes(
      constantMenusRef.value.concat(routes)
    )
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
