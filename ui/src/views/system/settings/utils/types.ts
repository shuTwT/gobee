export type SiteSetting = {
  maintenance_mode: boolean
  allow_registration: boolean
  email_verification: boolean
  comment_moderation: boolean
  enable_cdn: boolean
  cdn_url: string
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
