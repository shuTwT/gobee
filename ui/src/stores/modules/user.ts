import { defineStore } from 'pinia'
import { store } from '..'
import { useStorageLocal } from '@/utils/utils'
import { removeToken, setToken, userKey, type DataInfo } from '@/utils/auth'
import { passwordLogin } from '@/api/auth'
import { resetRouter } from '@/router'

export const useUserStore = defineStore('user', () => {
  /**
   * 头像
   */
  const avatar = ref(useStorageLocal().getItem<DataInfo<number>>(userKey)?.avatar ?? '')
  /**
   * 用户名
   */
  const username = ref(useStorageLocal().getItem<DataInfo<number>>(userKey)?.username ?? '')
  /**
   * 昵称
   */
  const nickname = ref(useStorageLocal().getItem<DataInfo<number>>(userKey)?.nickname ?? '')
  /**
   * 角色
   */
  const roles = ref(useStorageLocal().getItem<DataInfo<number>>(userKey)?.roles ?? [])
  /**
   * 权限
   */
  const permissions = ref(useStorageLocal().getItem<DataInfo<number>>(userKey)?.permissions ?? [])
  // 前端生成的验证码（按实际需求替换）
  const verifyCode = ref('')
  // 判断登录页面显示哪个组件（0：登录（默认）、1：手机登录、2：二维码登录、3：注册、4：忘记密码）
  const currentPage = ref(0)
  /**
   * 记住我
   */
  const isRemembered = ref(false)
  /**
   * 免登录存储时间
   */
  const loginDay = ref(7)

  /** 存储头像 */
  function SET_AVATAR(_avatar: string) {
    avatar.value = _avatar
  }
  /** 存储用户名 */
  function SET_USERNAME(_username: string) {
    username.value = _username
  }
  /** 存储昵称 */
  function SET_NICKNAME(_nickname: string) {
    nickname.value = _nickname
  }
  /** 存储角色 */
  function SET_ROLES(_roles: Array<string>) {
    roles.value = _roles
  }
  /** 存储按钮级别权限 */
  function SET_PERMS(_permissions: Array<string>) {
    permissions.value = _permissions
  }
  /** 存储前端生成的验证码 */
  function SET_VERIFYCODE(_verifyCode: string) {
    verifyCode.value = _verifyCode
  }
  /** 存储登录页面显示哪个组件 */
  function SET_CURRENTPAGE(_value: number) {
    currentPage.value = _value
  }
  /** 存储是否勾选了登录页的免登录 */
  function SET_ISREMEMBERED(_bool: boolean) {
    isRemembered.value = _bool
  }
  /** 设置登录页的免登录存储几天 */
  function SET_LOGINDAY(_value: number) {
    loginDay.value = Number(_value)
  }

  async function loginByUsername(data:any){
    return new Promise<any>((resolve, reject) => {
      passwordLogin(data).then(({data}) => {
        setToken(data)
        resolve(data)
      }).catch(err => {
        reject(err)
      })
    })
  }

  async function logOut(){
    username.value = ""
    roles.value = []
    permissions.value = []
    removeToken()
    resetRouter()
  }

  async function handleRefreshToken(data:any){

  }

  return {
    avatar,
    username,
    nickname,
    roles,
    permissions,
    isRemembered,
    loginDay,
    verifyCode,
    currentPage,
    SET_AVATAR,
    SET_USERNAME,
    SET_NICKNAME,
    SET_ROLES,
    SET_PERMS,
    SET_VERIFYCODE,
    SET_CURRENTPAGE,
    SET_ISREMEMBERED,
    SET_LOGINDAY,
    loginByUsername,
    logOut,
    handleRefreshToken,
  }
})

export function useUserStoreHook() {
  return useUserStore(store)
}
