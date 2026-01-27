export type ThemeItemProps = {
  id?: number
  name: string
  display_name: string
  description?: string
  author_name?: string
  author_email?: string
  logo?: string
  homepage?: string
  repo?: string
  issue?: string
  setting_name?: string
  config_map_name?: string
  version: string
  require?: string
  license?: string
  path: string
  enabled: boolean
}

export type FormProps = {
  formInline: ThemeItemProps
}
