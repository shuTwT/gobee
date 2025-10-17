import { defineStore } from "pinia";
import { store } from "..";

export const useUserStore = defineStore('user',()=>{
  /**
   * 头像
   */
  const avatar = ref('')
  /**
   * 用户名
   */
  const username = ref('')
  /**
   * 昵称
   */
  const nickname = ref('')
  /**
   * 角色
   */
  const roles = ref([])
  /**
   * 权限
   */
  const permissions = ref([])
  /**
   * 记住我
   */
  const isRemember = ref(false)
  /**
   * 免登录存储时间
   */
  const loginDay = ref(7)

  return {
    avatar,
    username,
    nickname,
    roles,
    permissions,
    isRemember,
    loginDay,
  }
})

export function useUserStoreHook(){
  return useUserStore(store)
}
