package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AIChatSession belongs to exactly one authenticated user.
type AIChatSession struct {
	ent.Schema
}

func (AIChatSession) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AIChatSession) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Positive().Comment("所属用户 ID"),
		field.String("title").NotEmpty().MaxLen(255).Default("新对话").Comment("会话标题"),
	}
}

func (AIChatSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("ai_chat_sessions").
			Field("user_id").
			Unique().
			Required(),
		edge.To("messages", AIChatMessage.Type),
	}
}
