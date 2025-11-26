import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "./utils";

export const getFlinkPage = ()=>{
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/flink/page`);
}

export const createFlink = (data:any) =>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/flink/create`,{data})
}

export const updateFlink = (id:number,data:any) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/flink/update/${id}`,{data})
}

export const queryFlink = () =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/flink/query`)
}

export const deleteFlink = (id:number) =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/flink/delete/${id}`)
}
