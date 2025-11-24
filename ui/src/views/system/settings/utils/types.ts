export type basicSetting = {
  siteName: string
  siteDescription: string
  siteLogo: string
  siteFavicon: string
  author: string
  language: string
  timezone: string
  dateFormat: string
  timeFormat: string
}
export type AiSetting = {
  openaiApiKey: string
  openaiApiUrl: string
  aiModel: string
  aiTemperature: number
  aiMaxTokens: number
  aiTopP: number
  aiFrequencyPenalty: number
  aiPresencePenalty: number
}
export type EmailSetting = {
  smtpHost: string
  smtpPort: number
  smtpUsername: string
  smtpPassword: string
  smtpEncryption: string
  senderEmail: string
  senderName: string
}
export type SiteSetting = {
  maintenanceMode: boolean
  allowRegistration: boolean
  emailVerification: boolean
  commentModeration: boolean
  uploadLimit: number
  allowedFileTypes: Array<'jpg' | 'jpeg' | 'png' | 'gif' | 'pdf' | 'doc' | 'docx'>
  enableCache: boolean
  cacheTime: number
  enableSSL: boolean
  enableCDN: boolean
  cdnUrl: string
}

export type SecuritySetting = {
  enableTwoFactor: boolean
  maxLoginAttempts: number
  lockoutDuration: number
  sessionTimeout: number
  minPasswordLength: number
  passwordComplexity: ('uppercase' | 'lowercase' | 'number')[]
  apiRateLimit: number
  ipWhitelist: string[]
  ipBlacklist: string[]
}

export type PaymentSetting = {
  enableAlipay: boolean
  alipayAppId: string
  alipayPrivateKey: string
  alipayPublicKey: string
  enableWechatPay: boolean
  wechatMchId: string
  wechatApiKey: string
  wechatAppId: string
  paymentNotifyUrl: string
  orderTimeout: number
  refundReview: boolean
}

export type NotifySetting = {
  enableEmailNotification: boolean
  enableSmsNotification: boolean
  enableInAppNotification: boolean
  notifyNewUserRegistration: boolean
  notifyLoginAbnormal: boolean
  notifyPasswordChange: boolean
  notifyNewOrder: boolean
  notifyOrderPaid: boolean
  notifyOrderShipped: boolean
  notifyOrderCompleted: boolean
  notifySystemError: boolean
  notifyMaintenance: boolean
  notificationEmail: string
}

export type SeoSetting = {
  enableSEO: boolean
  siteTitleSuffix: string
  siteKeywords: string
  siteAuthor: string
  siteCopyright: string
  siteIcp: string
  sitePoliceIcp: string
  analyticsCode: string
  robotsContent: string
  autoGenerateSitemap: boolean
  urlRewrite: boolean
  custom404Content: string
}

export type BackupSetting = {
  enableAutoBackup: boolean
  backupFrequency: string
  backupRetentionDays: number
  backupStorageLocation: string
  backupDatabase: boolean
  backupUploads: boolean
  backupConfig: boolean
  enableRemoteBackup: boolean
  remoteStorageType: string
  remoteStorageConfig: string
}

export type LogSetting = {
  enableSystemLogs: boolean
  logLevel: string
  logRetentionDays: number
  maxLogFileSize: number
  logUserActions: boolean
  logSystemErrors: boolean
  logDatabaseQueries: false
  logApiCalls: boolean
  logPaymentActions: boolean
  logStorageType: string
  externalLogService: string
  externalLogConfig: string
}
export type SettingsProps = Partial<{}>
