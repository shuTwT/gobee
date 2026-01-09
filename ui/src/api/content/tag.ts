import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getTagList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/tag/list`)
}

export const getTagPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/tag/page`, { params })
}

export const queryTag = (id: number) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/tag/query/${id}`)
}

export const createTag = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/tag/create`, { data })
}

export const updateTag = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/tag/update/${id}`, { data })
}

export const deleteTag = (id: number) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/tag/delete/${id}`)
}
