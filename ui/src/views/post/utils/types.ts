export type FormItemProps = {
  id?: number
  title:string,
  alias:string,
  categorys:any[],
  tags:any[],
  autogenSummary:boolean,
  author:string,
  allowComments:boolean,
  pinToTop:boolean,
  visible:boolean,
  cover:string
}

export type FormProps={
  formInline: FormItemProps
}
