package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AIChatMessage is one persisted user or completed assistant message.
type AIChatMessage struct {
	ent.Schema
}

func (AIChatMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AIChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("session_id").Positive().Comment("会话 ID"),
		field.Enum("role").Values("user", "assistant").Comment("消息角色"),
		field.Text("content").NotEmpty().Comment("消息内容"),
		field.String("model").Optional().MaxLen(255).Comment("生成该回复的模型"),
	}
}

func (AIChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("session", AIChatSession.Type).
			Ref("messages").
			Field("session_id").
			Unique().
			Required(),
	}
}
