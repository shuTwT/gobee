export interface StorageStrategy {
  id: number
  name: string
  type: 'local' | 's3'
  domain: string
  endpoint?: string
  base_path?: string
  node_id?: string
  access_key?: string
  secret_key?: string
  bucket?: string
  region?: string
  master: boolean
  created_at: string
  updated_at: string
}

export type FormItemProps = {
  id?: number
  name: string
  type: 'local' | 's3'
  domain: string
  endpoint?: string
  base_path?: string
  node_id?: string
  access_key?: string
  secret_key?: string
  bucket?: string
  region?: string
  master: boolean
}

export type FormProps = {
  formInline: FormItemProps
}
