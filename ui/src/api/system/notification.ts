import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "@/api/utils";

export const getNotificationPage = (params?:any) => {
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/notifications/page`,{params})
}

export const queryNotification = (id:number) => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/notifications/query/${id}`)
}

export const deleteNotification = (id:number) => {
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/notifications/delete/${id}`)
}

export const batchMarkAsRead = (data:any) => {
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/notifications/batch/read`,{data})
}
