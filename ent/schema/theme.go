package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Theme struct {
	ent.Schema
}

func (Theme) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Theme) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MaxLen(100).Comment("主题名称，唯一标识"),
		field.String("display_name").NotEmpty().MaxLen(100).Comment("显示名称"),
		field.String("description").MaxLen(500).Optional().Comment("主题描述"),
		field.String("author_name").MaxLen(100).Optional().Comment("作者姓名"),
		field.String("author_email").MaxLen(100).Optional().Comment("作者邮箱"),
		field.String("logo").MaxLen(255).Optional().Comment("主题logo"),
		field.String("homepage").MaxLen(255).Optional().Comment("主页"),
		field.String("repo").MaxLen(255).Optional().Comment("仓库地址"),
		field.String("issue").MaxLen(255).Optional().Comment("问题反馈地址"),
		field.String("setting_name").MaxLen(100).Optional().Comment("设置名称"),
		field.String("config_map_name").MaxLen(100).Optional().Comment("配置映射名称"),
		field.String("version").NotEmpty().MaxLen(20).Comment("版本号"),
		field.String("require").MaxLen(50).Default("*").Comment("要求的gobee版本"),
		field.String("license").MaxLen(50).Optional().Comment("许可证"),
		field.String("path").NotEmpty().MaxLen(255).Comment("主题文件路径"),
		field.Bool("enabled").Default(false).Comment("是否启用"),
	}
}

func (Theme) Edges() []ent.Edge {
	return nil
}
