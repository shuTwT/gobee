import {http} from '@/utils/http'
import type { ApiResponse } from './utils'

export interface StorageStrategy {
  id: number
  name: string
  type: 'local' | 'aliyun' | 'tencent' | 'qiniu'
  domain: string
  access_key?: string
  secret_key?: string
  bucket?: string
  region?: string
  is_default: boolean
  created_at: string
  updated_at: string
}

export interface CreateStorageStrategyParams {
  name: string
  type: StorageStrategy['type']
  domain: string
  access_key?: string
  secret_key?: string
  bucket?: string
  region?: string
  is_default: boolean
}

export interface UpdateStorageStrategyParams extends Partial<CreateStorageStrategyParams> {
  id: number
}

// 获取存储策略列表
export const getStorageStrategyList = () => {
  return http.request<ApiResponse<StorageStrategy[]>>('get','/api/storage/strategies')
}

// 创建存储策略
export const createStorageStrategy = (data: CreateStorageStrategyParams) => {
  return http.request('post','/api/storage/strategies', {data})
}

// 更新存储策略
export const updateStorageStrategy = (data: UpdateStorageStrategyParams) => {
  return http.request('put',`/api/storage/strategies/${data.id}`, {data})
}

// 删除存储策略
export const deleteStorageStrategy = (id: number) => {
  return http.request('delete',`/api/storage/strategies/${id}`)
}

// 设置默认存储策略
export const setDefaultStorageStrategy = (id: number) => {
  return http.request('put',`/api/storage/strategies/${id}/default`)
}
