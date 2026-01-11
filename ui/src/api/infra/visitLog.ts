import {http} from '@/utils/http'
import { BASE_URL, type ApiResponse,type TableResponse } from '@/api/utils'

export type VisitLogResponse = {
  id: number;
  ip: string;
  user_agent: string;
  path: string;
  os: string;
  browser: string;
  device: string;
  created_at: string;
  updated_at: string;
}

export const getVisitLogPage = (params:any) =>{
  return http.request<TableResponse<VisitLogResponse>>('get',`${BASE_URL}/v1/visit-log/page`,{params})
}

export const queryVisitLog = (id: number) =>{
  return http.request<ApiResponse<VisitLogResponse>>('get',`${BASE_URL}/v1/visit-log/query/${id}`)
}

export const deleteVisitLog = (id: number) =>{
  return http.request<ApiResponse<null>>('delete',`${BASE_URL}/v1/visit-log/delete/${id}`)
}

export const batchDeleteVisitLog = (ids: number[]) =>{
  return http.request<ApiResponse<null>>('post',`${BASE_URL}/v1/visit-log/batch/delete`,{data:{ids}})
}