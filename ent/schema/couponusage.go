package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 优惠券使用记录
type CouponUsage struct {
	ent.Schema
}

func (CouponUsage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the CouponUsage.
func (CouponUsage) Fields() []ent.Field {
	return []ent.Field{
		field.String("coupon_code").
			Comment("优惠券代码"),
		field.Int("user_id").
			Comment("用户ID"),
		field.Int("order_id").
			Optional().
			Comment("订单ID"),
		field.Int("status").
			Default(0).
			Comment("状态: 0-未使用 1-已使用 2-已过期"),
		field.Time("used_at").
			Optional().
			Comment("使用时间"),
		field.Int("discount_amount").
			Default(0).
			Comment("优惠金额(分)"),
		field.Time("expire_at").
			Default(time.Now).
			Comment("过期时间"),
		field.String("remark").
			Optional().
			MaxLen(500).
			Comment("备注"),
	}
}

// Edges of the CouponUsage.
func (CouponUsage) Edges() []ent.Edge {
	return nil
}
