package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 会员
type Member struct {
	ent.Schema
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Optional().
			Comment("用户ID"),
		field.Int("member_level").
			Comment("会员等级ID"),
		field.String("member_no").
			Unique().
			NotEmpty().
			Comment("会员编号"),
		field.Time("join_time").
			Default(time.Now).
			Comment("入会时间"),
		field.Time("expire_time").
			Optional().
			Comment("会员到期时间"),
		field.Int("points").
			Default(0).
			Comment("会员积分"),
		field.Int("total_spent").
			Default(0).
			Comment("累计消费金额(分)"),
		field.Int("order_count").
			Default(0).
			Comment("订单数量"),
		field.Bool("active").
			Default(true).
			Comment("是否激活"),
		field.String("remark").
			Optional().
			MaxLen(500).
			Comment("备注"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("member").
			Field("user_id").
			Unique(),
	}
}
