import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse} from "../utils"

export const getAsyncRoutes=()=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/routes`)
}

export const getAllRoutes=(params:any)=>{
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/api-interface/page`,{params})
}
