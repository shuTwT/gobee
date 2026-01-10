export type FormItemProps = {
  id?: number
  title:string,
  alias:string,
  content:string,
  md_content:string,
  html_content:string,
  content_type:string,
  status:string,
  categories:any[],
  tags:any[],
  is_autogen_summary:boolean,
  is_visible:boolean,
  is_pin_to_top:boolean,
  is_allow_comment:boolean,
  is_visible_after_comment:boolean,
  is_visible_after_pay:boolean,
  price:number,
  author:string,
  cover:string,
  keywords:string,
  copyright:string,
  summary:string,
  view_count:number,
  comment_count:number
}

export type FormProps={
  formInline: FormItemProps
}
