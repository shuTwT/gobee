package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 说说
type Essay struct {
	ent.Schema
}

func (Essay) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Essay.
func (Essay) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Comment("说说内容"),
		field.Bool("draft").Default(false).Comment("是否草稿"),
		field.JSON("images", []string{}).Optional(),
	}
}

// Edges of the Essay.
func (Essay) Edges() []ent.Edge {
	return nil
}
