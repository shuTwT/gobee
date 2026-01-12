import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

export const getKnowledgeBasePage = (params: any) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/knowledgebase/page`, { params });
}

export const getKnowledgeBase = (id: number | string) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/knowledgebase/query/${id}`);
}

export const createKnowledgeBase = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/knowledgebase/create`, { data });
}

export const updateKnowledgeBase = (id: number | string, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/knowledgebase/update/${id}`, { data });
}

export const deleteKnowledgeBase = (id: number | string) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/knowledgebase/delete/${id}`);
}

export const getKnowledgeBaseList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/knowledgebase/list`);
}
