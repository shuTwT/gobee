export type FormItemProps = {
  id?: number
  name: string
  title_selector: string
  link_selector: string
  created_selector: string
  updated_selector: string
}

export type FormProps = {
  formInline: FormItemProps
}
