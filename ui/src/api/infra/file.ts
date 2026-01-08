import {http} from '@/utils/http'
import { BASE_URL, type ApiResponse,type TableResponse } from './utils'

type FileListResponse = {
  id: number;
  name: string;
  url: string;
  type: string;
}

export const getFileList = () =>{
  return http.request<ApiResponse<FileListResponse[]>>('get',`${BASE_URL}/v1/file/list`)
}

export const getFilePage = (params:any) =>{
  return http.request<TableResponse<FileListResponse[]>>('get',`${BASE_URL}/v1/file/page`,{params})
}

export const queryFile = (id: number) =>{
  return http.request<ApiResponse<FileListResponse>>('get',`${BASE_URL}/v1/file/query/${id}`)
}

export const deleteFile = (id: number) =>{
  return http.request<ApiResponse<FileListResponse>>('delete',`${BASE_URL}/v1/file/delete/${id}`)
}
