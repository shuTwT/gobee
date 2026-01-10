import type { Component } from 'vue'

export type SourceType = 'rss' | 'atom' | 'md'

export type ImportType = 'file' | 'url' | 'folder'

export interface SourceTypeItem {
  value: SourceType
  label: string
  description: string
  icon: Component
}

export interface MigrationResult {
  /**
   * 
   */
  status: 'success' | 'failed'| null
  message: string
  stats: {
    total: number
    success: number
    failed: number
  }
  errors?: string[]
}

export interface PreviewData {
  type: 'file' | 'url'
  count?: number
  files?: string[]
  url?: string
}
