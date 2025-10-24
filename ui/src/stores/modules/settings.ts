import { defineStore } from "pinia";
import { store } from "..";

export const useSettingsStore = defineStore("settings",()=> {
  /**
   * 是否初始化
   */
  const initialized = ref(false)

  /**
   * 初始化设置
   */
  function initialize() {
    initialized.value = true
  }

  return {
    initialized,
    initialize,
  }
})

export function useSettingsStoreHook() {
  return useSettingsStore(store)
}
