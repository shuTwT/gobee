import {http} from '@/utils/http'
import { BASE_URL, type ApiResponse, type TableResponse } from '@/api/utils'

export interface ScheduleJob {
  id: number
  created_at: string
  updated_at: string
  name: string
  type: 'cron' | 'interval' | 'once'
  expression: string
  description: string
  enabled: boolean
  next_run_time: string
  last_run_time: string
  execution_type: 'http' | 'internal' | 'command' | 'mq'
  http_method?: string
  http_url?: string
  http_headers?: Record<string, string>
  http_body?: string
  http_timeout: number
  max_retries: number
  failure_notification: boolean
}

export interface CreateScheduleJobParams {
  name: string
  type: ScheduleJob['type']
  expression: string
  description?: string
  enabled: boolean
  execution_type: ScheduleJob['execution_type']
  http_method?: string
  http_url?: string
  http_headers?: Record<string, string>
  http_body?: string
  http_timeout?: number
  max_retries: number
  failure_notification: boolean
}

export interface UpdateScheduleJobParams extends Partial<CreateScheduleJobParams> {
  id: number
}

export interface ScheduleJobPageParams {
  page: number
  size: number
}

export const getScheduleJobPage = (params: ScheduleJobPageParams) => {
  return http.request<TableResponse<ScheduleJob>>('get', `${BASE_URL}/v1/schedule-job/page`, { params })
}

export const queryScheduleJob = (id: number) => {
  return http.request<ApiResponse<ScheduleJob>>('get', `${BASE_URL}/v1/schedule-job/query/${id}`)
}

export const createScheduleJob = (data: CreateScheduleJobParams) => {
  return http.request<ApiResponse<ScheduleJob>>('post', `${BASE_URL}/v1/schedule-job/create`, { data })
}

export const updateScheduleJob = (data: UpdateScheduleJobParams) => {
  return http.request<ApiResponse<ScheduleJob>>('put', `${BASE_URL}/v1/schedule-job/update/${data.id}`, { data })
}

export const deleteScheduleJob = (id: number) => {
  return http.request<ApiResponse<null>>('delete', `${BASE_URL}/v1/schedule-job/delete/${id}`)
}
