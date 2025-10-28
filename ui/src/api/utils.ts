export const BASE_URL = import.meta.env.VITE_BASE_API

export type ApiResponse<T> = {
  code: number
  msg: string
  data: T
}
