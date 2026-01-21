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
