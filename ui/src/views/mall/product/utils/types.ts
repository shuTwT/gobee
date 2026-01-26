export type ProductItemProps = {
  id?: number
  name: string
  sku: string
  price: number
  original_price?: number
  cost_price?: number
  stock: number
  min_stock?: number
  category_id?: number | null
  brand?: string
  unit?: string
  weight?: number
  volume?: number
  description?: string
  short_description?: string
  images?: any[]
  attributes?: Record<string, any>
  tags?: string[]
  active?: boolean
  featured?: boolean
  digital?: boolean
  meta_title?: string
  meta_description?: string
  meta_keywords?: string
  sort_order?: number
}

export type FormProps = {
  formInline: ProductItemProps
}
