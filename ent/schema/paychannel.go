package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 支付渠道
type PayChannel struct {
	ent.Schema
}

func (PayChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the PayChannel.
func (PayChannel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(255).Comment("支付渠道名称"),
		field.String("code").NotEmpty().MaxLen(255).Comment("支付渠道代码"),
		field.String("type").NotEmpty().MaxLen(255).Comment("支付渠道类型"),
		field.String("config").Comment("支付渠道配置"),
	}
}

// Edges of the PayChannel.
func (PayChannel) Edges() []ent.Edge {
	return nil
}
