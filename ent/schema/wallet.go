package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 钱包
type Wallet struct {
	ent.Schema
}

func (Wallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Optional().
			Comment("用户ID"),
		field.Int("balance").
			Default(0).
			Comment("余额(分)"),
		field.Int("frozen_amount").
			Default(0).
			Comment("冻结金额(分)"),
		field.Int("total_income").
			Default(0).
			Comment("总收入(分)"),
		field.Int("total_expense").
			Default(0).
			Comment("总支出(分)"),
		field.String("password").
			Optional().
			Sensitive().
			MaxLen(255).
			Comment("支付密码"),
		field.Bool("active").
			Default(true).
			Comment("是否激活"),
		field.String("remark").
			Optional().
			MaxLen(500).
			Comment("备注"),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{}
}
