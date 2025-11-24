import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

// 系统设置接口返回类型
export interface SettingResponse {
  settings: Record<string,string>;
  initialized: boolean;
}

// 获取系统设置和初始化状态
export const getSettings = () => {
  return http.request<ApiResponse<SettingResponse>>('get',`${BASE_URL}/v1/settings`);
};

export const getSettingsMap = (key:string)=>{
  return http.request<ApiResponse<any>>('get',`${BASE_URL}/v1/settings/json/${key}`);
}

export const saveSettings = (key:string,data:Record<string,any>)=>{
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/v1/settings/json/save/${key}`,{data})
}
