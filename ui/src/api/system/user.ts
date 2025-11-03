import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "../utils";

export const getUserList = () => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/user/list`)
}

export const createUser = () => {
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/user/create`)
}

export const updateUser = () => {
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/user/update`)
}

export const queryUser = () => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/user/query`)
}

export const deleteUser = () => {
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/user/delete`)
}
