import { http } from "@/utils/http";
import { BASE_URL, type ApiResponse } from "./utils";

type LoginRequest = {
  email: string;
  password: string;
};

export const passwordLogin = (data: LoginRequest) => {
  return http.post<ApiResponse<any>, LoginRequest>(`${BASE_URL}/auth/login/password`, {data});
};
