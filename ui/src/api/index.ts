import { Api, HttpClient, type RequestParams } from '@gobee/api-client'
import type { AxiosResponse } from 'axios'

const httpClient = new HttpClient({
  baseURL: '',
})


const apiClient = new Api(httpClient)

type RequestFn<R, T> =
  | ((req: R, params?: RequestParams) => Promise<AxiosResponse<T>>)
  | ((req?: R, params?: RequestParams) => Promise<AxiosResponse<T>>)

/**
 * 调用 API 函数
 * @param requestFn API 函数
 * @param req 请求参数
 * @param params 请求参数
 * @returns API 响应数据
 */
async function useApi<R = any, T = any>(
  requestFn: RequestFn<R, T>,
  req?: R,
  params?: RequestParams,
): Promise<T> {
  const res = await requestFn(req as R, params)
  return res.data
}

export { apiClient, useApi }
