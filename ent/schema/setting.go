package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 设置
type Setting struct {
	ent.Schema
}

func (Setting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().MaxLen(255).Comment("设置键"),
		field.String("value").Comment("设置值"),
		field.String("comment").MaxLen(512).Optional().Comment("设置注释"),
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}
