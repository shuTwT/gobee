import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse, type TableResponse } from "./utils";

// 获取首页统计信息
export const getHomeStatistic = () => {
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/statistic`);
};
