package schema

import "entgo.io/ent"

// ModelSchema holds the schema definition for the ModelSchema entity.
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
