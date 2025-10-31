import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

type LoginRequest = {
  email: string;
  password: string;
};

export const passwordLogin = (data: LoginRequest) => {
  return http.request<ApiResponse<any>>('post',`${BASE_URL}/auth/login/password`, {data});
};
