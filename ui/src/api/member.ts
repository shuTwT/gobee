import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getMemberPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/member/page`, { params })
}

export const updateMember = (id: number, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/member/update/${id}`, { data })
}
