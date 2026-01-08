import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getMemberLevelPage = (params?: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/member-level/page`, { params })
}

export const getMemberLevelList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/member-level/list`)
}
