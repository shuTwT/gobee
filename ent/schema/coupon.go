package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 优惠券
type Coupon struct {
	ent.Schema
}

func (Coupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(100).
			Comment("优惠券名称"),
		field.String("code").
			Unique().
			NotEmpty().
			MaxLen(50).
			Comment("优惠券代码"),
		field.String("description").
			Optional().
			MaxLen(500).
			Comment("优惠券描述"),
		field.Int("type").
			Default(0).
			Comment("优惠券类型: 0-满减 1-折扣 2-无门槛"),
		field.Int("value").
			Default(0).
			Comment("优惠券值"),
		field.Int("min_amount").
			Default(0).
			Comment("最低消费金额(分)"),
		field.Int("max_discount").
			Default(0).
			Comment("最大优惠金额(分)"),
		field.Int("total_count").
			Default(0).
			Comment("发放总数"),
		field.Int("used_count").
			Default(0).
			Comment("已使用数量"),
		field.Int("per_user_limit").
			Default(1).
			Comment("每用户限领数量"),
		field.Time("start_time").
			Default(time.Now).
			Comment("开始时间"),
		field.Time("end_time").
			Default(time.Now).
			Comment("结束时间"),
		field.Bool("active").
			Default(true).
			Comment("是否启用"),
		field.String("image").
			Optional().
			MaxLen(255).
			Comment("优惠券图片"),
		field.JSON("product_ids", []int{}).
			Optional().
			Comment("适用商品ID列表"),
		field.JSON("category_ids", []int{}).
			Optional().
			Comment("适用分类ID列表"),
	}
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
	return []ent.Edge{}
}
