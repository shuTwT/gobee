import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

export const getPostList = (params:any)=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/post/list`,{params});
}

export const createPost = (data:any) =>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/post/create`,{data})
}

export const updatePost = (data:any) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/post/update`,{data})
}

export const queryPost = (id:number|string) =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/post/query/${id}`)
}

export const deletePost = (id:number|string) =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/post/delete/${id}`)
}
