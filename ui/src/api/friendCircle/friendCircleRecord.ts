import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "../utils";

export const getFriendCircleRecordPage = ()=>{
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/friend-circle-record/page`);
}

export const createFriendCircleRecord = (data:any) =>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/friend-circle-record/create`,{data})
}

export const updateFriendCircleRecord = (id:number,data:any) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/friend-circle-record/update/${id}`,{data})
}

export const queryFriendCircleRecord = () =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/friend-circle-record/query`)
}

export const deleteFriendCircleRecord = (id:number) =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/friend-circle-record/delete/${id}`)
}
