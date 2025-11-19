import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getRoleList = () => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/role/list`)
}

export const getRolePage = (params?:any) => {
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/role/page`,{params})
}

export const createRole = () => {
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/role/create`)
}

export const updateRole = () => {
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/role/update`)
}

export const queryRole = () => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/role/query`)
}

export const deleteRole = () => {
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/role/delete`)
}
