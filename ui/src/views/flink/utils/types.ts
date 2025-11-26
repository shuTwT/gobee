export type FlinkGroupFormItemProps = {
  id?:number
  name:string
}
export type FlinkFormItemProps = {
  id?:number
  name:string
  url:string
  avatar_url:string
  description:string
  cover_url:string
  snapshot_url:string
  email:string
  enable_friend_circle:boolean
  friend_circle_rule_id?:number
}

export type FlinkGroupFormPorps = {
  formInline:FlinkGroupFormItemProps
}

export type FlinkFormProps={
  formInline:FlinkFormItemProps
}
