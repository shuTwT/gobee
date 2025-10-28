import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

export const initialize = (data:any)=>{
  return http.post<ApiResponse<any>, any>(`${BASE_URL}/initialize`, {data});
}
