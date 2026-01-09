export interface FormItemProps {
  memberLevelId?: number
  memberLevelData: {
    name: string
    level: number
    min_points: number
    discount_rate: number
    description: string
    privileges: string[]
    icon: string
    active: boolean
    sort_order: number
  }
}

export interface FormProps {
  memberLevelId?: number
  memberLevelData: any
}
