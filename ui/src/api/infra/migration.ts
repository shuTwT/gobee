import {http} from '@/utils/http'
import { BASE_URL, type ApiResponse } from '@/api/utils'

type MigrationResult = {
  status: string
  message: string
  total: number
  success: number
  failed: number
  errors?: string[]
}

type DuplicateFile = {
  filename: string
  title: string
  post_id?: number
}

type MigrationCheckResult = {
  has_duplicates: boolean
  duplicates: DuplicateFile[]
}

export const importMarkdown = (files: File[]) => {
  const formData = new FormData()
  files.forEach(file => {
    formData.append('files', file)
  })
  return http.request<ApiResponse<MigrationResult>>('post', `${BASE_URL}/v1/migration/md`, {
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

export const checkDuplicate = (files: File[]) => {
  const formData = new FormData()
  files.forEach(file => {
    formData.append('files', file)
  })
  return http.request<ApiResponse<MigrationCheckResult>>('post', `${BASE_URL}/v1/migration/check-duplicate`, {
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
