package schema

import "entgo.io/ent"

// 优惠券
type Coupon struct {
	ent.Schema
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
	return nil
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
	return nil
}
