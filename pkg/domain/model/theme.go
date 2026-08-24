package model

type ThemeConfig struct {
	Type          string       `yaml:"type" validate:"required"`
	Name          string       `yaml:"name" validate:"required"`
	DisplayName   string       `yaml:"display-name" validate:"required"`
	Description   string       `yaml:"description,omitempty"`
	Author        *ThemeAuthor `yaml:"author,omitempty"`
	Logo          string       `yaml:"logo,omitempty"`
	Homepage      string       `yaml:"homepage,omitempty"`
	Repo          string       `yaml:"repo,omitempty"`
	Issue         string       `yaml:"issue,omitempty"`
	SettingName   string       `yaml:"setting-name,omitempty"`
	ConfigMapName string       `yaml:"config-map-name,omitempty"`
	Version       string       `yaml:"version" validate:"required"`
	Require       string       `yaml:"require,omitempty"`
	License       string       `yaml:"license,omitempty"`
}

type ThemeAuthor struct {
	Name  string `yaml:"name,omitempty"`
	Email string `yaml:"email,omitempty"`
}

type CreateThemeReq struct {
	Type        string `json:"type" validate:"required,oneof=internal external static"`
	FilePath    string `json:"file_path,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Description string `json:"description,omitempty"`
	ExternalURL string `json:"external_url,omitempty"`
	Version     string `json:"version,omitempty"`
}

type CreateExternalThemeReq struct {
	Name        string `json:"name" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
	Description string `json:"description,omitempty"`
	ExternalURL string `json:"external_url" validate:"required,url"`
	Version     string `json:"version" validate:"required"`
}

// 主题设置表单支持的字段类型
const (
	SettingFieldTypeText     = "text"
	SettingFieldTypeTextarea = "textarea"
	SettingFieldTypeNumber   = "number"
	SettingFieldTypeSelect   = "select"
	SettingFieldTypeRadio    = "radio"
	SettingFieldTypeSwitch   = "switch"
	SettingFieldTypeColor    = "color"
	SettingFieldTypeDate     = "date"
	SettingFieldTypeSecret   = "secret"
)

// ValidSettingFieldTypes 允许的主题设置字段类型集合
var ValidSettingFieldTypes = map[string]struct{}{
	SettingFieldTypeText:     {},
	SettingFieldTypeTextarea: {},
	SettingFieldTypeNumber:   {},
	SettingFieldTypeSelect:   {},
	SettingFieldTypeRadio:    {},
	SettingFieldTypeSwitch:   {},
	SettingFieldTypeColor:    {},
	SettingFieldTypeDate:     {},
	SettingFieldTypeSecret:   {},
}

// SettingSchema 主题包 setting.yaml 的顶层结构
type SettingSchema struct {
	Type  string        `yaml:"type" json:"type"`
	Forms []SettingForm `yaml:"forms" json:"forms"`
}

// SettingForm 设置表单分组，一个分组对应后台设置页的一个 Tab
type SettingForm struct {
	Group      string         `yaml:"group" json:"group"`
	Label      string         `yaml:"label" json:"label"`
	FormSchema []SettingField `yaml:"formSchema" json:"formSchema"`
}

// SettingField 设置表单中的单个字段定义
type SettingField struct {
	Type        string          `yaml:"type" json:"type"`
	Name        string          `yaml:"name" json:"name"`
	Label       string          `yaml:"label" json:"label"`
	Placeholder string          `yaml:"placeholder,omitempty" json:"placeholder,omitempty"`
	Help        string          `yaml:"help,omitempty" json:"help,omitempty"`
	Default     any             `yaml:"default,omitempty" json:"default,omitempty"`
	Required    bool            `yaml:"required,omitempty" json:"required,omitempty"`
	Min         *float64        `yaml:"min,omitempty" json:"min,omitempty"`
	Max         *float64        `yaml:"max,omitempty" json:"max,omitempty"`
	Options     []SettingOption `yaml:"options,omitempty" json:"options,omitempty"`
}

// SettingOption select/radio 字段的候选项
type SettingOption struct {
	Label string `yaml:"label" json:"label"`
	Value string `yaml:"value" json:"value"`
}

// ThemeSettingResp 主题设置接口响应：表单定义 + 当前配置值（已合并默认值）
type ThemeSettingResp struct {
	Forms  []SettingForm  `json:"forms"`
	Values map[string]any `json:"values"`
}

// SaveThemeSettingReq 保存主题设置请求
type SaveThemeSettingReq struct {
	Values map[string]any `json:"values" validate:"required"`
}

type ThemeResp struct {
	ID            int       `json:"id"`
	CreatedAt     LocalTime `json:"created_at"`
	UpdatedAt     LocalTime `json:"updated_at"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	DisplayName   string    `json:"display_name"`
	Description   string    `json:"description"`
	AuthorName    string    `json:"author_name"`
	AuthorEmail   string    `json:"author_email"`
	Logo          string    `json:"logo"`
	Homepage      string    `json:"homepage"`
	Repo          string    `json:"repo"`
	Issue         string    `json:"issue"`
	SettingName   string    `json:"setting_name"`
	ConfigMapName string    `json:"config_map_name"`
	Version       string    `json:"version"`
	Require       string    `json:"require"`
	License       string    `json:"license"`
	Path          string    `json:"path"`
	ExternalURL   string    `json:"external_url"`
	Enabled       bool      `json:"enabled"`
}
