package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// 文章购买记录
type PostPurchase struct {
	ent.Schema
}

func (PostPurchase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the PostPurchase.
func (PostPurchase) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Comment("购买用户"),
		field.Int("post_id").
			Comment("购买文章"),
		field.Int("order_id").
			Optional().
			Comment("关联支付订单"),
	}
}

// Edges of the PostPurchase.
func (PostPurchase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id").
			Comment("购买用户"),
		edge.To("post", Post.Type).
			Unique().
			Required().
			Field("post_id").
			Comment("购买文章"),
		edge.To("order", PayOrder.Type).
			Unique().
			Field("order_id").
			Comment("关联支付订单"),
	}
}

// Indexes of the PostPurchase.
func (PostPurchase) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "post_id").Unique(),
	}
}
