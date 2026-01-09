import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getMemberLevelPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/member-level/page`, { params })
}

export const getMemberLevelList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/member-level/list`)
}

export const queryMemberLevel = (id: number) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/member-level/query/${id}`)
}

export const createMemberLevel = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/member-level/create`, { data })
}

export const updateMemberLevel = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/member-level/update/${id}`, { data })
}

export const deleteMemberLevel = (id: number) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/member-level/delete/${id}`)
}
