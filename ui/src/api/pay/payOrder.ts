import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "../utils";

export const getPayOrderList = ()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/pay-order/page`);
}

export const createPayOrder = () =>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/pay-order/create`)
}

export const updatePayOrder = () =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/pay-order/update`)
}

export const queryPayOrder = () =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/pay-order/query`)
}

export const deletePayOrder = () =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/pay-order/delete`)
}
