import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

// 系统设置接口返回类型
export interface SettingResponse {
  settings: Record<string,string>;
  initialized: boolean;
}

// 获取系统设置和初始化状态
export const getSettings = () => {
  return http.request<ApiResponse<SettingResponse>>('get',`${BASE_URL}/v1/settings`);
};

export const getSettingsMap = ()=>{
  return http.request<ApiResponse<Record<string,string>>>('get',`${BASE_URL}/v1/settings/map`);
}
