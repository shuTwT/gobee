export interface FormItemProps {
  tagId?: number
  tagData: {
    name: string
    slug: string
    description: string
    color: string
    sort_order: number
    active: boolean
  }
}

export interface FormProps {
  tagId?: number
  tagData: any
}
