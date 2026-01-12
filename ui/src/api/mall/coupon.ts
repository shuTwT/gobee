import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getCouponList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/coupon/list`)
}

export const getCouponPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/coupon/page`, { params })
}

export const createCoupon = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/coupon/create`, { data })
}

export const updateCoupon = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/coupon/update/${id}`, { data })
}

export const queryCoupon = (id: number) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/coupon/query/${id}`)
}

export const deleteCoupon = (id: number) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/coupon/delete/${id}`)
}

export const batchUpdateCoupons = (ids: number[], data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/coupon/batch`, { data: { ids, ...data } })
}

export const batchDeleteCoupons = (ids: number[]) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/coupon/batch/delete`, { data: { ids } })
}

export const searchCoupons = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/coupon/search`, { params })
}
