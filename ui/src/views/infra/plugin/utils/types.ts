export interface Plugin {
  id: number
  created_at: string
  updated_at: string
  key: string
  name: string
  version: string
  description: string
  bin_path: string
  protocol_version: string
  magic_cookie_key: string
  magic_cookie_value: string
  dependencies: string[]
  config: string
  enabled: boolean
  auto_start: boolean
  status: 'stopped' | 'running' | 'error' | 'loading'
  last_error: string
  last_started_at?: string
  last_stopped_at?: string
}
