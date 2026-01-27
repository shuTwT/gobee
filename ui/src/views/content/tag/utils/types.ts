export type TagItemProps = {
  id?: number
  name: string
  slug?: string
  description?: string
  color?: string
  sort_order: number
  active: boolean
}

export type FormProps = {
  formInline: TagItemProps
}
