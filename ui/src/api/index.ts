import { Api, HttpClient, type RequestParams } from '@hoshikuzu/api-client'
import type { AxiosError, AxiosResponse, InternalAxiosRequestConfig } from 'axios'
import { getToken } from "@/utils/auth"
import { refreshAccessToken } from "@/stores/modules/user"

const httpClient = new HttpClient({
  baseURL: '',
})

// 请求拦截：附带 access token，过期时先刷新
httpClient.instance.interceptors.request.use(async (config) => {
  const data = getToken()
  if (data) {
    const now = Date.now()
    if (Number(data.expires) - now <= 0) {
      try {
        const newToken = await refreshAccessToken()
        config.headers['Authorization'] = `Bearer ${newToken}`
      } catch {
        // 刷新失败交由响应拦截/路由守卫处理，这里仍尝试用旧 token 发起
        config.headers['Authorization'] = `Bearer ${data.accessToken}`
      }
    } else {
      config.headers['Authorization'] = `Bearer ${data.accessToken}`
    }
  }
  return config
})

// 响应拦截：401 时静默刷新并重放一次
httpClient.instance.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const status = error.response?.status
    const config = error.config as (InternalAxiosRequestConfig & { __isRetry?: boolean }) | undefined
    if (status === 401 && config && !config.__isRetry) {
      config.__isRetry = true
      try {
        const newToken = await refreshAccessToken()
        config.headers['Authorization'] = `Bearer ${newToken}`
        return httpClient.instance(config)
      } catch {
        return Promise.reject(error)
      }
    }
    return Promise.reject(error)
  }
)


const apiClient = new Api(httpClient)

type RequestFn<P extends any[], T=any> = (...args:[...P,RequestParams|undefined]) => Promise<AxiosResponse<T>>

/**
 * 调用 API 函数
 * @param requestFn API 函数
 * @param req 请求参数
 * @param params 请求参数
 * @returns API 响应数据
 */
async function useApi<P extends any[], T = any>(
  requestFn: RequestFn<P, T>,
  ...args:P
): Promise<T> {
  const res = await requestFn(...args,undefined)
  return res.data
}

export { apiClient, useApi }
