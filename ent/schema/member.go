package schema

import "entgo.io/ent"

// 会员
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return nil
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
