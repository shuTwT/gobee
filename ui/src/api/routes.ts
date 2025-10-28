import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils"

export const getAsyncRoutes=()=>{
  return http.get<ApiResponse<any>, any>(`${BASE_URL}/v1/routes`)
}
