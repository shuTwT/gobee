import { http } from "@/utils/http";
import { BASE_URL, type TableResponse } from "@/api/utils";

export const getFriendCircleRecordPage = ()=>{
  return http.request<TableResponse<any>>('get',`${BASE_URL}/v1/friend-circle/page`);
}