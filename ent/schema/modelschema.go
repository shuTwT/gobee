package schema

import "entgo.io/ent"

// ModelSchema 模型架构
type ModelSchema struct {
	ent.Schema
}

// Fields of the ModelSchema.
func (ModelSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ModelSchema.
func (ModelSchema) Edges() []ent.Edge {
	return nil
}
