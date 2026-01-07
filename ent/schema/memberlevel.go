package schema

import "entgo.io/ent"

// MemberLevel holds the schema definition for the MemberLevel entity.
type MemberLevel struct {
	ent.Schema
}

// Fields of the MemberLevel.
func (MemberLevel) Fields() []ent.Field {
	return nil
}

// Edges of the MemberLevel.
func (MemberLevel) Edges() []ent.Edge {
	return nil
}
