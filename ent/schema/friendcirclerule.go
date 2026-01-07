package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FriendCircleRule 好友圈规则
type FriendCircleRule struct {
	ent.Schema
}

func (FriendCircleRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the FriendCircleRule.
func (FriendCircleRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("规则名称"),
		field.String("title_selector").Optional().Comment("标题选择器"),
		field.String("link_selector").Optional().Comment("链接选择器"),
		field.String("created_selector").Optional().Comment("创建时间选择器"),
		field.String("updated_selector").Optional().Comment("更新时间选择器"),
	}
}

// Edges of the FriendCircleRule.
func (FriendCircleRule) Edges() []ent.Edge {
	return nil
}
