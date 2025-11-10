export type FormItemProps = {
  id?: number
  title:string,
  alias:string,
  categorys:any[],
  tags:any[],
  is_autogen_summary:boolean,
  author:string,
  is_allow_comments:boolean,
  is_pin_to_top:boolean,
  is_visible:boolean,
  cover:string
}

export type FormProps={
  formInline: FormItemProps
}
