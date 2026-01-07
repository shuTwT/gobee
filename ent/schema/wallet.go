package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 钱包
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户ID"),
		field.Int("balance").Default(0).Comment("余额 单位分"),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return nil
}
