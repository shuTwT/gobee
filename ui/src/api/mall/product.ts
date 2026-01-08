import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getProductList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/product/list`)
}

export const getProductPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/product/page`, { params })
}

export const createProduct = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/product/create`, { data })
}

export const updateProduct = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/product/update/${id}`, { data })
}

export const queryProduct = (id: number) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/product/query/${id}`)
}

export const deleteProduct = (id: number) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/product/delete/${id}`)
}

export const batchUpdateProducts = (ids: number[], data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/product/batch`, { data: { ids, ...data } })
}

export const batchDeleteProducts = (ids: number[]) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/product/batch/delete`, { data: { ids } })
}
