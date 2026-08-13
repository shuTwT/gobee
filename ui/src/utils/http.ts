import axios, { AxiosError, type AxiosRequestConfig, type Method } from 'axios'
import { formatToken, getToken } from './auth'
import { useUserStoreHook, refreshAccessToken } from '@/stores/modules/user'


export type RequestMethods = Extract<
  Method,
  'get' | 'post' | 'put' | 'delete' | 'patch' | 'option' | 'head'
>

class HttpService {

  constructor(){
    this.httpInterceptorsRequest()
    this.httpInterceptorsResponse()
  }

  /** 共享的刷新 Promise，并发请求只触发一次刷新 */
  private static refreshPromise: Promise<string> | null = null

  /** 防止重复弹出登录过期提示 */
  private static isLoggingOut = false

  /** 保存当前`Axios`实例对象 */
  private static axiosInstance = axios.create({
    timeout: 10000,
    headers: {
      Accept: 'application/json, text/plain, */*',
      'Content-Type': 'application/json',
      'X-Requested-With': 'XMLHttpRequest',
    },
  })

  /** 执行刷新并返回新的 accessToken，并发调用共享同一 Promise */
  private static getRefreshedToken(): Promise<string> {
    if (HttpService.refreshPromise) {
      return HttpService.refreshPromise
    }
    HttpService.refreshPromise = refreshAccessToken().finally(() => {
      HttpService.refreshPromise = null
    })
    return HttpService.refreshPromise
  }

  /** 刷新失败或刷新无效时统一登出 */
  private static handleAuthFailure() {
    if (HttpService.isLoggingOut) return
    HttpService.isLoggingOut = true
    const doLogout = () => {
      HttpService.isLoggingOut = false
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
  private static isWhitelist(url: string | undefined): boolean {
    const whiteList = ['/refresh-token', '/login', '/logout']
    return !!url && whiteList.some((item) => url.endsWith(item))
  }

  /** 请求拦截 */
  private httpInterceptorsRequest(): void {
    HttpService.axiosInstance.interceptors.request.use(async (config):Promise<any> => {
      // 如果数据是 FormData，删除 Content-Type 让浏览器自动设置
      if (config.data instanceof FormData && config.headers) {
        delete config.headers['Content-Type']
      }

      if (HttpService.isWhitelist(config.url)) {
        return config
      }

      const data = getToken()
      if (data) {
        const now = Date.now()
        if (!config.headers) {
          config.headers = {}
        }
        if (Number(data.expires) - now <= 0) {
          // token 已过期，先刷新再挂载新 token
          try {
            const newToken = await HttpService.getRefreshedToken()
            config.headers['Authorization'] = formatToken(newToken)
          } catch {
            HttpService.handleAuthFailure()
            return Promise.reject(new Error('登录已过期，请重新登录'))
          }
        } else {
          config.headers['Authorization'] = formatToken(data.accessToken)
        }
      }
      return config
    },
    error=>{
      return Promise.reject(error)
    }
  )
  }

  /** 尝试用刷新后的 token 重放原始请求（仅一次） */
  private static async replayWithFreshToken(config: AxiosRequestConfig): Promise<any> {
    const newToken = await HttpService.getRefreshedToken()
    if (!config.headers) {
      config.headers = {}
    }
    config.headers['Authorization'] = formatToken(newToken)
    return HttpService.axiosInstance(config)
  }

  private  httpInterceptorsResponse():void{
    const instance = HttpService.axiosInstance;
    instance.interceptors.response.use(
      async (response) =>{
        const code = response.data.code || 200
        let msg = response.data.msg || response.statusText || '未知错误'

        if (code === 401) {
          // 刷新接口本身返回 401，直接登出，不再重试
          if (HttpService.isWhitelist(response.config.url)) {
            HttpService.handleAuthFailure()
            return Promise.reject(new Error('登录已过期，请重新登录'))
          }
          const config = response.config as AxiosRequestConfig & { __isRetry?: boolean }
          if (!config.__isRetry) {
            config.__isRetry = true
            try {
              return await HttpService.replayWithFreshToken(config)
            } catch {
              // 刷新失败，落到下方登出
            }
          }
          msg = '登录过期，请重新登录'
          HttpService.handleAuthFailure()
          return Promise.reject(new Error(msg))
        } else if (code !== 200) {
          if (window.$message) window.$message.error(msg)
          return Promise.reject(new Error(msg))
        }
        return response.data
      },
      async (error: AxiosError) => {
        const status = error.response?.status
        const config = error.config as (AxiosRequestConfig & { __isRetry?: boolean }) | undefined

        // HTTP 401：尝试静默刷新并重放一次
        if (status === 401 && config && !HttpService.isWhitelist(config.url) && !config.__isRetry) {
          config.__isRetry = true
          try {
            return await HttpService.replayWithFreshToken(config)
          } catch {
            HttpService.handleAuthFailure()
            return Promise.reject(error)
          }
        }

        if (status === 401) {
          HttpService.handleAuthFailure()
          return Promise.reject(error)
        }

        if (window.$message) {
          const respData = error.response?.data as any
          window.$message.error(respData?.msg || respData || error.message || '请求失败')
        }
        return Promise.reject(error)
      }
    )
  }

  public request<T=any>(method: RequestMethods, url: string, params?: AxiosRequestConfig): Promise<T> {
    const config = {
      method,
      url,
      ...params,
    }
    return new Promise((resolve, reject) => {
      HttpService.axiosInstance
        .request<T>(config)
        .then((response:any) => {
          resolve(response)
        })
        .catch((error) => {
          reject(error)
        })
    })
  }
  public post<T=any, P=any>(url: string, params?: AxiosRequestConfig<P>): Promise<T> {
    return this.request<T>('post', url, params)
  }
  public get<T=any, P=any>(url: string, params?: AxiosRequestConfig<P>): Promise<T> {
    return this.request<T>('get', url, params)
  }
}

export const http = new HttpService()
