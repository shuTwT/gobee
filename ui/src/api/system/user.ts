import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getUserList = () => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/user/list`)
}

export const getUserPage = (params?:any) => {
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/user/page`,{params})
}

export const createUser = (data:any) => {
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/user/create`,{data})
}

export const updateUser = (id:number,data:any) => {
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/user/update/${id}`,{data})
}

export const queryUser = (id:number) => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/user/query/${id}`)
}

export const deleteUser = (id:number) => {
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/user/delete/${id}`)
}

export const getPersonalAccessTokenList = ()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/user/personal-access-token/list`)
}

export const getPersonalAccessToken = (id:number)=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/user/personal-access-token/query/${id}`)
}

export const createPat = (data:any)=>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/user/personal-access-token/create`,{data})
}

export const getUserProfile = ()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/user/profile`)
}
