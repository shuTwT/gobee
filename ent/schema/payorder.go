/**
 * 支付订单
 */
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PayOrder holds the schema definition for the PayOrder entity.
type PayOrder struct {
	ent.Schema
}

func (PayOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the PayOrder.
func (PayOrder) Fields() []ent.Field {
	return []ent.Field{
		field.String("channel_id").
			Comment("支付渠道ID"),
		field.String("order_id").
			Comment("支付订单ID"),
		field.String("out_trade_no").
			Comment("商户订单号"),
		field.String("total_fee").
			Comment("订单金额"),
		field.String("subject").
			Comment("订单标题"),
		field.String("body").
			Comment("订单描述"),
		field.String("notify_url").
			Comment("异步通知地址"),
		field.String("return_url").
			Comment("同步通知地址"),
		field.String("extra").
			Comment("额外参数"),
		field.String("pay_url").
			Optional().
			Comment("支付链接"),
		field.String("state").
			Default("1").
			Comment("支付状态"),
		field.String("error_msg").
			Optional().
			Comment("错误信息"),
		field.String("raw").
			Optional().
			Comment("原始响应"),
	}
}

// Edges of the PayOrder.
func (PayOrder) Edges() []ent.Edge {
	return nil
}
