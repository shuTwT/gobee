import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

export const getAlbumPhotoList =(params:any)=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/album-photo/list`,{params});
}

export const createAlbumPhoto = (data:any)=>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/album-photo/create`,{data})
}

export const updateAlbumPhoto = (data:any)=>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/album-photo/update`,{data})
}

export const queryAlbumPhoto = (id:number|string) =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/album-photo/query/${id}`)
}

export const deleteAlbumPhoto = (id:number|string) =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/album-photo/delete/${id}`)
}
