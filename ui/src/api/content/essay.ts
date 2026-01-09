import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

export const getEssayPage = (params: any) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/essay/page`, { params });
}

export const createEssay = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/essay/create`, { data });
}

export const updateEssay = (id: number | string, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/essay/update/${id}`, { data });
}

export const deleteEssay = (id: number | string) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/essay/delete/${id}`);
}
