package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// WebHook
type WebHook struct {
	ent.Schema
}

func (WebHook) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the WebHook.
func (WebHook) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(255),
		field.String("url").NotEmpty().MaxLen(255),
		field.String("event").NotEmpty().MaxLen(255),
	}
}

// Edges of the WebHook.
func (WebHook) Edges() []ent.Edge {
	return nil
}
