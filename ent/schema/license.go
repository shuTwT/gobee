package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type License struct {
	ent.Schema
}

func (License) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (License) Fields() []ent.Field {
	return []ent.Field{
		field.String("machine_code").
			NotEmpty().
			MaxLen(255).
			Comment("机器码"),
		field.String("license_key").
			NotEmpty().
			MaxLen(512).
			Comment("授权密钥"),
		field.String("customer_name").
			Optional().
			MaxLen(255).
			Comment("客户名称"),
		field.Time("expire_date").
			Comment("过期日期"),
		field.Int("status").
			Default(1).
			Comment("状态: 1-有效, 2-过期, 3-禁用"),
	}
}
