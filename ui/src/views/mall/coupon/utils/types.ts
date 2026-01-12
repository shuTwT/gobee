export type FormItemProps = {
  id?:number
  name:string
  code:string
  description?:string
  type:number
  value:number
  min_amount:number
  max_discount:number
  total_count:number
  per_user_limit:number
  start_time:number
  end_time:number
  active:boolean
  image:string
  product_ids:number[]
  category_ids:number[]
}

export type FormProps = {
  formInline: FormItemProps
}
