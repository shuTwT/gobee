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
		field.String("avatar_url").Optional().Comment("头像"),
		field.String("description").Optional().Comment("简介"),
		field.Int("status").Default(1).Comment("状态"),
		field.String("snapshot_url").Optional().Comment("网站截图"),
		field.String("cover_url").Optional().Comment("网站封面"),
		field.String("email").Optional().Comment("邮箱"),
		field.Int("group_id").Optional().Comment("友链组"),
		field.Bool("enable_friend_circle").Default(true).Comment("是否开启朋友圈"),
		field.Int("friend_circle_rule_id").Optional().Nillable().Comment("朋友圈解析规则"),
	}
}

// Edges of the FLink.
func (FLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", FLinkGroup.Type).Ref("links").Unique().Field("group_id"),
	}
}
