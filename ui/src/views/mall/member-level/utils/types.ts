export type MemberLevelItemProps = {
  id?: number
  name: string
  level: number
  min_points: number
  discount_rate: number
  description?: string
  privileges?: string[]
  icon?: string
  active: boolean
  sort_order: number
}

export type FormProps = {
  formInline: MemberLevelItemProps
}
