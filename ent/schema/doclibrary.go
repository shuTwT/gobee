package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 文档库
type DocLibrary struct {
	ent.Schema
}

func (DocLibrary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the DocLibrary.
func (DocLibrary) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("文档库名称"),
		field.String("alias").
			Comment("文档库别名"),
		field.String("description").
			Optional().
			Comment("文档库描述"),
		field.String("url").
			Comment("文档库URL"),
	}
}

// Edges of the DocLibrary.
func (DocLibrary) Edges() []ent.Edge {
	return nil
}
