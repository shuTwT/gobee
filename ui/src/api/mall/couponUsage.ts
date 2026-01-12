import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getCouponUsageList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/coupon-usage/list`)
}

export const getCouponUsagePage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/coupon-usage/page`, { params })
}

export const createCouponUsage = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/coupon-usage/create`, { data })
}

export const updateCouponUsage = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/coupon-usage/update/${id}`, { data })
}

export const queryCouponUsage = (id: number) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/coupon-usage/query/${id}`)
}

export const deleteCouponUsage = (id: number) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/coupon-usage/delete/${id}`)
}

export const batchUpdateCouponUsages = (ids: number[], data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/coupon-usage/batch`, { data: { ids, ...data } })
}

export const batchDeleteCouponUsages = (ids: number[]) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/coupon-usage/batch/delete`, { data: { ids } })
}

export const searchCouponUsages = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/coupon-usage/search`, { params })
}
