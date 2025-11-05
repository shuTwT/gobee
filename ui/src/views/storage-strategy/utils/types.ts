export type FormItemProps = {
  id?: number
  name: string
  type: string
  node_id?: string
  endpoint?:string
  domain: string
  access_key?: string
  secret_key?: string
  bucket?: string
  region?:string
  base_path?: string
  master?: boolean
}

export type FormProps={
  formInline: FormItemProps
}
