export type FlinkGroupFormItemProps = {
  id?:number
  name:string
}
export type FlinkFormItemProps = {
  id?:number
  name:string
  url:string
  logo:string
  description:string
}

export type FlinkGroupFormPorps = {
  formInline:FlinkGroupFormItemProps
}

export type FlinkFormProps={
  formInline:FlinkFormItemProps
}
