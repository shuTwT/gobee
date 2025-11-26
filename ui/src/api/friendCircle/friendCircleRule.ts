import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "../utils";

export const getFriendCircleRuleList = ()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/friend-circle-rule/list`);
}

export const getFriendCircleRulePage = ()=>{
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/friend-circle-rule/page`);
}

export const createFriendCircleRule = (data:any) =>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/friend-circle-rule/create`,{data})
}

export const updateFriendCircleRule = (id:number,data:any) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/friend-circle-rule/update/${id}`,{data})
}

export const queryFriendCircleRule = () =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/friend-circle-rule/query`)
}

export const deleteFriendCircleRule = (id:number) =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/friend-circle-rule/delete/${id}`)
}
