package schema

import "entgo.io/ent"

// 钱包
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return nil
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return nil
}
