import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const createFlinkApplication = (data: any) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/v1/flink-application/create`, { data });
};

export const getFlinkApplicationPage = (params: any) => {
  return http.request<TableResponse<any>>('get', `${BASE_URL}/v1/flink-application/page`, { params });
};

export const queryFlinkApplication = (id: number | string) => {
  return http.request<ApiResponse<any>>('get', `${BASE_URL}/v1/flink-application/query/${id}`);
};

export const approveFlinkApplication = (id: number | string, data: any) => {
  return http.request<ApiResponse<any>>('put', `${BASE_URL}/v1/flink-application/update/${id}`, { data });
};
