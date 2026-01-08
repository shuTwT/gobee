package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 会员等级
type MemberLevel struct {
	ent.Schema
}

func (MemberLevel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the MemberLevel.
func (MemberLevel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(50).
			Comment("等级名称"),
		field.String("description").
			Optional().
			MaxLen(500).
			Comment("等级描述"),
		field.Int("level").
			Unique().
			Comment("等级级别"),
		field.Int("min_points").
			Default(0).
			Comment("最低积分要求"),
		field.Int("discount_rate").
			Default(100).
			Comment("折扣率(百分比)"),
		field.JSON("privileges", []string{}).
			Optional().
			Comment("会员特权"),
		field.String("icon").
			Optional().
			MaxLen(255).
			Comment("等级图标"),
		field.Bool("active").
			Default(true).
			Comment("是否启用"),
		field.Int("sort_order").
			Default(0).
			Comment("排序"),
	}
}

// Edges of the MemberLevel.
func (MemberLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", Member.Type),
	}
}
