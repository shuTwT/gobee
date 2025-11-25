package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FLink holds the schema definition for the FLink entity.
type FLink struct {
	ent.Schema
}

func (FLink) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the FLink.
func (FLink) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("名称"),
		field.String("url").NotEmpty().Comment("链接"),
		field.String("logo").Optional().Comment("logo"),
		field.String("description").Optional().Comment("简介"),
		field.Int("status").Default(1).Comment("状态"),
		field.String("snapshot").Optional().Comment("快照"),
		field.String("email").Optional().Comment("邮箱"),
		field.Int("group_id").Optional().Comment("友链组"),
	}
}

// Edges of the FLink.
func (FLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", FLinkGroup.Type).Ref("links").Unique().Field("group_id"),
	}
}
