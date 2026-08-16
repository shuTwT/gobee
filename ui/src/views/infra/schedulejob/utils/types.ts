export interface ScheduleJob {
  id: number
  created_at: string
  updated_at: string
  name: string
  type: 'cron' | 'interval'
  expression: string
  description: string
  enabled: boolean
  next_run_time: string
  last_run_time: string
  job_name: string
  max_retries: number
  failure_notification: boolean
}
