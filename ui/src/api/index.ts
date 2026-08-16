import { Api, HttpClient, type RequestParams } from '@hoshikuzu/api-client'
import type { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import { formatToken, getToken } from '@/utils/auth'
import { refreshAccessToken, useUserStoreHook } from '@/stores/modules/user'

const httpClient = new HttpClient({
  baseURL: '',
})

/** 共享的刷新 Promise，并发请求只触发一次刷新 */
let refreshPromise: Promise<string> | null = null

/** 防止重复弹出登录过期提示 */
let isLoggingOut = false

/** 执行刷新并返回新的 accessToken，并发调用共享同一 Promise */
function getRefreshedToken(): Promise<string> {
  if (refreshPromise) {
    return refreshPromise
  }
  refreshPromise = refreshAccessToken().finally(() => {
    refreshPromise = null
  })
  return refreshPromise
}

/** 刷新失败或刷新无效时统一登出 */
function handleAuthFailure() {
  if (isLoggingOut) return
  isLoggingOut = true
  const doLogout = () => {
    isLoggingOut = false
    useUserStoreHook().logOut()
  }
  if (window.$dialog) {
    window.$dialog.warning({
      title: '登录过期',
      content: '登录状态已失效，请重新登录',
      positiveText: '确定',
      closable: false,
      maskClosable: false,
      onPositiveClick: doLogout,
      onClose: doLogout,
    })
  } else {
    doLogout()
  }
}

/** 判断是否为免 token 的鉴权接口 */
function isWhitelist(url: string | undefined): boolean {
  const whiteList = ['/refresh-token', '/login', '/logout']
  return !!url && whiteList.some((item) => url.endsWith(item))
}

// 请求拦截：附带 access token，过期时先刷新
httpClient.instance.interceptors.request.use(async (config) => {
  // 如果数据是 FormData，删除 Content-Type 让浏览器自动设置
  if (config.data instanceof FormData && config.headers) {
    delete config.headers['Content-Type']
  }

  if (isWhitelist(config.url)) {
    return config
  }

  const data = getToken()
  if (data) {
    const now = Date.now()
    if (Number(data.expires) - now <= 0) {
      // token 已过期，先刷新再挂载新 token
      try {
        const newToken = await getRefreshedToken()
        config.headers.set('Authorization', formatToken(newToken))
      } catch {
        handleAuthFailure()
        return Promise.reject(new Error('登录已过期，请重新登录'))
      }
    } else {
      config.headers.set('Authorization', formatToken(data.accessToken))
    }
  }
  return config
})

/** 尝试用刷新后的 token 重放原始请求（仅一次） */
async function replayWithFreshToken(config: AxiosRequestConfig): Promise<AxiosResponse> {
  const newToken = await getRefreshedToken()
  config.headers = config.headers ?? {}
  config.headers['Authorization'] = formatToken(newToken)
  return httpClient.instance(config)
}

// 响应拦截：业务码非 200 弹错；401 时静默刷新并重放一次
httpClient.instance.interceptors.response.use(
  async (response) => {
    const code = response.data?.code || 200
    let msg = response.data?.msg || response.statusText || '未知错误'

    if (code === 401) {
      // 刷新接口本身返回 401，直接登出，不再重试
      if (isWhitelist(response.config.url)) {
        handleAuthFailure()
        return Promise.reject(new Error('登录已过期，请重新登录'))
      }
      const config = response.config as AxiosRequestConfig & { __isRetry?: boolean }
      if (!config.__isRetry) {
        config.__isRetry = true
        try {
          return await replayWithFreshToken(config)
        } catch {
          // 刷新失败，落到下方登出
        }
      }
      msg = '登录过期，请重新登录'
      handleAuthFailure()
      return Promise.reject(new Error(msg))
    } else if (code !== 200) {
      if (window.$message) window.$message.error(msg)
      return Promise.reject(new Error(msg))
    }
    // 生成客户端的 request() 期望拿到完整 AxiosResponse，由 useApi 解包 .data
    return response
  },
  async (error: AxiosError) => {
    const status = error.response?.status
    const config = error.config as (AxiosRequestConfig & { __isRetry?: boolean }) | undefined

    // HTTP 401：尝试静默刷新并重放一次
    if (status === 401 && config && !isWhitelist(config.url) && !config.__isRetry) {
      config.__isRetry = true
      try {
        return await replayWithFreshToken(config)
      } catch {
        handleAuthFailure()
        return Promise.reject(error)
      }
    }

    if (status === 401) {
      handleAuthFailure()
      return Promise.reject(error)
    }

    if (window.$message) {
      const respData = error.response?.data as { msg?: string } | undefined
      window.$message.error(respData?.msg || error.message || '请求失败')
    }
    return Promise.reject(error)
  }
)

const apiClient = new Api(httpClient)

type RequestFn<P extends any[], T = any> = (
  ...args: [...P, RequestParams | undefined]
) => Promise<AxiosResponse<T>>

/**
 * 调用 API 函数并解包响应体（{code, msg, data}）
 * @param requestFn API 函数
 * @param req 请求参数
 * @returns 后端响应体
 */
async function useApi<P extends any[], T = any>(requestFn: RequestFn<P, T>, ...args: P): Promise<T> {
  const res = await requestFn(...args, undefined)
  return res.data
}

export { apiClient, useApi }
