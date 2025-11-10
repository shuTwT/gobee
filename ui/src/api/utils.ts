export const BASE_URL = import.meta.env.VITE_BASE_API

export type ApiResponse<T = any> = {
  code: number
  msg: string
  data: T
}

export type TableResult<T = any> = {
  records:T[]
  total:number
}

export type TableResponse<T = any> = ApiResponse<TableResult<T>>
