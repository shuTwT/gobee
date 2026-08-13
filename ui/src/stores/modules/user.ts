import { defineStore } from 'pinia'
import { store } from '..'
import { useStorageLocal } from '@/utils/utils'
import { removeToken, setToken, getToken, userKey, type DataInfo } from '@/utils/auth'
import { passwordLogin, refreshToken as refreshTokenApi, logout as logoutApi } from '@/api/system/auth'
import router, { resetRouter } from '@/router'

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

  async function logOut() {
    // 尽力通知后端吊销 refresh token，失败不阻塞登出
    const refreshTokenValue = getToken()?.refreshToken
    if (refreshTokenValue) {
      try {
        await logoutApi(refreshTokenValue)
      } catch {
        // 忽略网络错误，继续本地登出
      }
    }
    username.value = ''
    roles.value = []
    permissions.value = []
    removeToken()
    resetRouter()
    router.push({path: '/login'})
  }

  // 刷新 access token，返回新的令牌数据供调用方使用
  async function handleRefreshToken() {
    await refreshAccessToken()
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

/**
 * 共享的 access token 刷新 Promise。
 * 多个 HTTP 客户端在 token 过期时共享同一次刷新，避免并发刷新导致 refresh token 轮换冲突。
 */
let accessTokenRefreshPromise: Promise<string> | null = null

export function refreshAccessToken(): Promise<string> {
  if (accessTokenRefreshPromise) {
    return accessTokenRefreshPromise
  }
  const data = getToken()
  if (!data?.refreshToken) {
    return Promise.reject(new Error('缺少 refreshToken'))
  }
  accessTokenRefreshPromise = (async () => {
    try {
      const { data: tokenData } = await refreshTokenApi({ refreshToken: data.refreshToken })
      setToken(tokenData)
      return tokenData.accessToken as string
    } finally {
      accessTokenRefreshPromise = null
    }
  })()
  return accessTokenRefreshPromise
}

export function useUserStoreHook() {
  return useUserStore(store)
}
