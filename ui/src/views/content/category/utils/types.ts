export interface FormItemProps {
  categoryId?: number
  categoryData: {
    name: string
    slug: string
    description: string
    sort_order: number
    active: boolean
  }
}

export interface FormProps {
  categoryId?: number
  categoryData: any
}
