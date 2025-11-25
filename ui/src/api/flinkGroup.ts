import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

export const getFlinkGroupList = ()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/flink-group/list`);
}

export const createFlinkGroup = (data:any) =>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/flink-group/create`,{data})
}

export const updateFlinkGroup = (id:number,data:any) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/flink-group/update/${id}`,{data})
}

export const deleteFlinkGroup = () =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/flink-group/delete`)
}
