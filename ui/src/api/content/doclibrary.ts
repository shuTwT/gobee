import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

export const getDocLibraryPage = (params: any) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/doclibrary/page`, { params });
}

export const getDocLibrary = (id: number | string) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/doclibrary/query/${id}`);
}

export const createDocLibrary = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/doclibrary/create`, { data });
}

export const updateDocLibrary = (id: number | string, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/doclibrary/update/${id}`, { data });
}

export const deleteDocLibrary = (id: number | string) => {
  return http.request<ApiResponse<any>>('delete', `${BASE_URL}/v1/doclibrary/delete/${id}`);
}

export const getDocLibraryList = () => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/doclibrary/list`);
}
