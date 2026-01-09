package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(100).
			Comment("分类名称"),
		field.String("description").
			Optional().
			MaxLen(500).
			Comment("分类描述"),
		field.String("slug").
			Optional().
			MaxLen(100).
			Comment("分类别名"),
		field.Int("sort_order").
			Default(0).
			Comment("排序"),
		field.Bool("active").
			Default(true).
			Comment("是否启用"),
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
