import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "@/api/utils";

type LoginRequest = {
  email: string;
  password: string;
};

export const passwordLogin = (data: LoginRequest) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/auth/login/password`, {data});
};

export const socialLogin = (code: string) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/auth/login/social`, {data: {code}})
}

// 刷新令牌（该接口在请求拦截器白名单中，不会附带 access token）
export const refreshToken = (data: { refreshToken: string }) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/auth/refresh-token`, {data})
}

// 登出并吊销 refresh token（公开接口，不走刷新逻辑）
export const logout = (refreshToken: string) => {
  return http.request<ApiResponse<any>>('post', `${BASE_URL}/auth/logout`, {data: {refreshToken}})
}
