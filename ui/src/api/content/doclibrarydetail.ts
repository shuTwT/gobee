import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

export const getDocLibraryDetailPage = (params: any) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/doclibrarydetail/page`, { params });
}

export const getDocLibraryDetail = (id: number | string) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/doclibrarydetail/query/${id}`);
}

export const createDocLibraryDetail = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/doclibrarydetail/create`, { data });
}

export const updateDocLibraryDetail = (id: number | string, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/doclibrarydetail/update/${id}`, { data });
}

export const deleteDocLibraryDetail = (id: number | string) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/doclibrarydetail/delete/${id}`);
}

export const getDocLibraryDetailTree = (params: any) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/doclibrarydetail/tree`, { params });
}
