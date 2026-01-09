package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(50).
			Comment("标签名称"),
		field.String("description").
			Optional().
			MaxLen(500).
			Comment("标签描述"),
		field.String("slug").
			Optional().
			MaxLen(50).
			Comment("标签别名"),
		field.String("color").
			Optional().
			MaxLen(20).
			Comment("标签颜色"),
		field.Int("sort_order").
			Default(0).
			Comment("排序"),
		field.Bool("active").
			Default(true).
			Comment("是否启用"),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
