import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getCategoryList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/category/list`)
}

export const getCategoryPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/category/page`, { params })
}

export const queryCategory = (id: number) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/category/query/${id}`)
}

export const createCategory = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/category/create`, { data })
}

export const updateCategory = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/category/update/${id}`, { data })
}

export const deleteCategory = (id: number) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/category/delete/${id}`)
}
