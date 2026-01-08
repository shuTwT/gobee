export type basicSetting = {
  site_name: string
  site_description: string
  site_logo: string
  site_favicon: string
  author: string
  language: string
  timezone: string
  date_format: string
  time_format: string
}
export type AiSetting = {
  openai_api_key: string
  openai_api_url: string
  ai_model: string
  ai_memperature: number
  ai_maxTokens: number
  ai_topP: number
  ai_frequency_penalty: number
  ai_presence_penalty: number
}
export type EmailSetting = {
  smtp_host: string
  smtp_port: number
  smtp_username: string
  smtp_password: string
  smtp_encryption: string
  sender_email: string
  sender_name: string
}
export type SiteSetting = {
  maintenance_mode: boolean
  allow_registration: boolean
  email_verification: boolean
  comment_moderation: boolean
  upload_limit: number
  allowed_file_types: Array<'jpg' | 'jpeg' | 'png' | 'gif' | 'pdf' | 'doc' | 'docx'>
  enable_cache: boolean
  cache_time: number
  enable_ssl: boolean
  enable_cdn: boolean
  cdn_url: string
}

export type SecuritySetting = {
  enable_two_factor: boolean
  max_login_attempts: number
  lockout_duration: number
  session_timeout: number
  min_password_length: number
  password_complexity: ('uppercase' | 'lowercase' | 'number')[]
  api_rate_limit: number
  ip_whitelist: string[]
  ip_blacklist: string[]
}

export type PaymentSetting = {
  enable_alipay: boolean
  alipay_app_id: string
  alipay_private_key: string
  alipay_public_key: string
  enable_wechat_pay: boolean
  wechat_mch_id: string
  wechat_api_key: string
  wechat_app_id: string
  payment_notify_url: string
  order_timeout: number
  refund_review: boolean
}

export type NotifySetting = {
  enable_email_notification: boolean
  enable_sms_notification: boolean
  enable_in_app_notification: boolean
  notify_new_user_registration: boolean
  notify_login_abnormal: boolean
  notify_password_change: boolean
  notify_new_order: boolean
  notify_order_paid: boolean
  notify_order_shipped: boolean
  notify_order_completed: boolean
  notify_system_error: boolean
  notify_maintenance: boolean
  notification_email: string
}

export type SeoSetting = {
  enable_seo: boolean
  site_title_suffix: string
  site_keywords: string
  site_author: string
  site_copyright: string
  site_icp: string
  site_police_icp: string
  analytics_code: string
  robots_content: string
  auto_generate_sitemap: boolean
  url_rewrite: boolean
  custom_404_content: string
}

export type BackupSetting = {
  enable_auto_backup: boolean
  backup_frequency: string
  backup_retention_days: number
  backup_storage_location: string
  backup_database: boolean
  backup_uploads: boolean
  backup_config: boolean
  enable_remote_backup: boolean
  remote_storage_type: string
  remote_storage_config: string
}

export type LogSetting = {
  enable_system_logs: boolean
  log_level: string
  log_retention_days: number
  max_log_file_size: number
  log_user_actions: boolean
  log_system_errors: boolean
  log_database_queries: boolean
  log_api_calls: boolean
  log_payment_actions: boolean
  log_storage_type: string
  external_log_service: string
  external_log_config: string
}
export type SettingsProps = Partial<{}>
