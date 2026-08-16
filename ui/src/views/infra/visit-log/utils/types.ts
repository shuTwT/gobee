export type VisitLogResponse = {
  id: number
  ip: string
  user_agent: string
  path: string
  os: string
  browser: string
  device: string
  created_at: string
  updated_at: string
}

export type LogDetailProps = {
  data: VisitLogResponse
}
