import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

type LoginRequest = {
  username: string;
  password: string;
};

export const login = (data: LoginRequest) => {
  return http.post<ApiResponse<any>, LoginRequest>(`${BASE_URL}/v1/auth/login/password`, {data});
};
