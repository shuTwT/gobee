import { defineStore } from "pinia";

export const usePermissionStore = defineStore('permission',()=>{
  /**
   * 静态路由生成的菜单
   */
  const constantMenus = ref([])
  /**
   * 整体路由生成的菜单(静态、动态)
   */
  const wholeMenus = ref([])
  /**
   * 整体路由（一维数组格式）
   */
  const flatteningRoutes= ref([])
  /**
   * 缓存页面keepAlive
   */
  const cachePageList= ref([])
  /**
   * 清空缓存页面
   */
  function clearAllCachePage(){
    wholeMenus.value = []
    cachePageList.value = []
  }
  return {
    constantMenus,
    wholeMenus,
    flatteningRoutes,
    cachePageList,
    clearAllCachePage
  }
})
