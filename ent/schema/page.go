package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

func (Page) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().MaxLen(255).Comment("页面标题"),
		field.Text("content").NotEmpty().Comment("页面内容"),
		field.String("description").Optional().Comment("页面描述"),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
