import axios, { type AxiosRequestConfig, type Method } from 'axios'

export type RequestMethods = Extract<
  Method,
  'get' | 'post' | 'put' | 'delete' | 'patch' | 'option' | 'head'
>

class HttpService {
  private static axiosInstance = axios.create({
    timeout: 10000,
    headers: {
      Accept: 'application/json, text/plain, */*',
      'Content-Type': 'application/json',
      'X-Requested-With': 'XMLHttpRequest',
    },
  })
  public request<T>(method: RequestMethods, url: string, params?: AxiosRequestConfig):Promise<T> {
    const config = {
      method,
      url,
      ...params,
    }
    return new Promise((resolve, reject) => {
      HttpService.axiosInstance
        .request<T>(config)
        .then((response) => {
          resolve(response.data)
        })
        .catch((error) => {
          reject(error)
        })
    })
  }
  public post<T, P>(url: string, params?: AxiosRequestConfig<P>):Promise<T> {
    return this.request<T>('post', url, params)
  }
  public get<T,P>(url:string,params?:AxiosRequestConfig<P>):Promise<T> {
    return this.request<T>('get', url, params)
  }
}

export const http = new HttpService()
