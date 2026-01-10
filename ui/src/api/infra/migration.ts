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
