package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FLinkCategory holds the schema definition for the FLinkCategory entity.
type FLinkGroup struct {
	ent.Schema
}

func (FLinkGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the FLinkCategory.
func (FLinkGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
		field.String("description").Optional().Comment("描述"),
	}
}

// Edges of the FLinkCategory.
func (FLinkGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("links", FLink.Type),
	}
}
