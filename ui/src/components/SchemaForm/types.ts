/** 主题设置表单支持的字段类型 */
export type SettingFieldType =
  | 'text'
  | 'textarea'
  | 'number'
  | 'select'
  | 'radio'
  | 'switch'
  | 'color'
  | 'date'
  | 'secret'

/** select/radio 字段的候选项 */
export type SettingOptionProps = {
  label: string
  value: string
}

/** 设置表单中的单个字段定义 */
export type SettingFieldProps = {
  type: SettingFieldType
  name: string
  label: string
  placeholder?: string
  help?: string
  default?: any
  required?: boolean
  min?: number
  max?: number
  options?: SettingOptionProps[]
}

/** 设置表单分组，一个分组对应一个 Tab */
export type SettingFormGroupProps = {
  group: string
  label: string
  formSchema?: SettingFieldProps[]
}
