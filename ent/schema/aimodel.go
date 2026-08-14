package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AIModel is a model entry under an AIProvider, used by chat and summary.
type AIModel struct {
	ent.Schema
}

func (AIModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AIModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int("provider_id").
			Positive().
			Comment("所属提供商 ID"),
		field.String("model_name").
			NotEmpty().
			MaxLen(255).
			Comment("模型名称（如 gpt-4o、deepseek-chat）"),
		field.String("display_name").
			Optional().
			MaxLen(255).
			Comment("显示名称"),
		field.Bool("is_enabled").
			Default(true).
			Comment("是否启用"),
		field.Int("sort").
			Default(0).
			Comment("排序值，越小越靠前"),
	}
}

func (AIModel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", AIProvider.Type).
			Ref("models").
			Field("provider_id").
			Unique().
			Required().
			Comment("所属提供商"),
	}
}
