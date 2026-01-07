package schema

import "entgo.io/ent"

// 会员等级
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
