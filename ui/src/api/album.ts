import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "./utils";

export const getAlbumList =(params?:any)=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/album/list`,{params});
}

export const getAlbumPage =(params?:any)=>{
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/album/page`,{params});
}

export const createAlbum = (data:any)=>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/album/create`,{data})
}

export const updateAlbum = (data:any)=>{
  return http.request<ApiResponse<any>>('put',`${BASE_URL}/v1/album/update`,{data})
}

export const queryAlbum = (id:number|string) =>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/album/query/${id}`)
}

export const deleteAlbum = (id:number|string) =>{
  return http.request<ApiResponse<any>>('delete',`${BASE_URL}/v1/album/delete/${id}`)
}
