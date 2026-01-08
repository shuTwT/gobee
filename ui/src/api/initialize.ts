import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

export const initialize = (data:any)=>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/initialize`, {data});
}
