import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

export const getPostList = (params:any)=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/post/list`,{params});
}

export const createPost = (data:any) =>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/post/create`,{data})
}

export const updatePostContent = (id:number|string,data:any) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/post/update/content/${id}`,{data})
}

export const updatePostSetting = (id:number|string,data:any) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/post/update/setting/${id}`,{data})
}

export const publishPost = (id:number|string) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/post/publish/${id}`)
}

export const unpublishPost = (id:number|string) =>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/post/unpublish/${id}`)
}

export const queryPost = (id:number|string) =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/post/query/${id}`)
}

export const deletePost = (id:number|string) =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/post/delete/${id}`)
}
