import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getProductList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/products/list`)
}

export const getProductPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/products/page`, { params })
}

export const createProduct = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/products/create`, { data })
}

export const updateProduct = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/products/update/${id}`, { data })
}

export const queryProduct = (id: number) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/products/query/${id}`)
}

export const deleteProduct = (id: number) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/products/delete/${id}`)
}

export const batchUpdateProducts = (ids: number[], data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/products/batch`, { data: { ids, ...data } })
}

export const batchDeleteProducts = (ids: number[]) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/products/batch/delete`, { data: { ids } })
}
