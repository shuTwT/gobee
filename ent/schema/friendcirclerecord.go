package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FriendCircleRecord holds the schema definition for the FriendCircleRecord entity.
type FriendCircleRecord struct {
	ent.Schema
}

func (FriendCircleRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the FriendCircleRecord.
func (FriendCircleRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("author").Comment("作者"),
		field.String("title").Comment("标题"),
		field.String("link_url").Comment("链接"),
		field.String("avatar_url").Comment("头像地址"),
		field.String("published_at").Optional().Comment("发布时间"),
	}
}

// Edges of the FriendCircleRecord.
func (FriendCircleRecord) Edges() []ent.Edge {
	return nil
}
