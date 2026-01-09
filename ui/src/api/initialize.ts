import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

export const preInit = ()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/preinit`);
}

export const initialize = (data:any)=>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/initialize`, {data});
}
