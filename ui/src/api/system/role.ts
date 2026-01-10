import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getRoleList = () => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/role/list`)
}

export const getRolePage = (params?:any) => {
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/role/page`,{params})
}

export const createRole = (data:any) => {
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/role/create`,{data})
}

export const updateRole = (id:number,data:any) => {
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/role/update/${id}`,{data})
}

export const queryRole = (id:number) => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/role/query/${id}`)
}

export const deleteRole = (id:number) => {
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/role/delete/${id}`)
}
