import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils"

export const getAsyncRoutes=()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/routes`)
}
