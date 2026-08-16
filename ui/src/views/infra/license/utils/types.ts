export interface License {
  id: number
  created_at: number
  updated_at: number
  domain: string
  license_key: string
  customer_name: string
  expire_date: number
  status: number
}

export interface LicensePageParams {
  page: number
  page_size: number
  domain?: string
  customer_name?: string
  status?: number
}

export type FormItemProps = {
  id?: number
  domain: string
  license_key?: string
  customer_name: string
  expire_date: number
  status?: number
}

export type FormProps = {
  formInline: FormItemProps
}
