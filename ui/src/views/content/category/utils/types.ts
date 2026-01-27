export type CategoryItemProps = {
  id?: number
  name: string
  slug?: string
  description?: string
  sort_order: number
  active: boolean
}

export type FormProps = {
  formInline: CategoryItemProps
}
