import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { store } from '../index'
import type { ProLayoutMode } from 'pro-naive-ui'

export const useAppStore = defineStore('app', () => {
  const layoutMode = ref<ProLayoutMode>('vertical')
  /**
   * 是否显示移动端布局
   */
  const isMobile = ref(false)
  /**
   * 侧边栏是否折叠
   */
  const collapsed = ref(false)
  /**
   * 是否显示 logo
   */
  const showLogo = ref(true)
  /**
   * 是否显示侧边栏
   */
  const showSidebar = ref(true)
  /**
   * 是否显示侧边栏额外内容
   */
  const showSidebarExtra = ref(true)
  /**
   * 是否显示顶栏
   */
  const showNav = ref(true)
  /**
   * 是否显示标签栏
   */
  const showTabbar = ref(true)
  /**
   * 是否显示底部
   */
  const showFooter = ref(true)
  /**
   * 顶栏是否固定
   */
  const navFixed = ref(true)
  /**
   * 底部是否固定
   */
  const footerFixed = ref(true)

  return {
    isMobile,
    collapsed,
    showLogo,
    showSidebar,
    showSidebarExtra,
    showNav,
    showTabbar,
    showFooter,
    navFixed,
    footerFixed,
    layoutMode,
  }
})

export function useAppStoreHook() {
  return useAppStore(store)
}
