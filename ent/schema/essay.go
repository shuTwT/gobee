package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 说说
type Essay struct {
	ent.Schema
}

func (Essay) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Essay.
func (Essay) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Comment("用户ID"),
		field.String("content").
			NotEmpty().
			MaxLen(1000).
			Comment("说说内容"),
		field.Bool("draft").
			Default(false).
			Comment("是否草稿"),
		field.JSON("images", []string{}).
			Optional().
			Comment("图片列表"),
		field.Int("like_count").
			Default(0).
			Comment("点赞数"),
		field.Int("comment_count").
			Default(0).
			Comment("评论数"),
		field.Int("share_count").
			Default(0).
			Comment("分享数"),
		field.Bool("public").
			Default(true).
			Comment("是否公开"),
		field.String("location").
			Optional().
			MaxLen(255).
			Comment("位置信息"),
		field.JSON("tags", []string{}).
			Optional().
			Comment("标签"),
	}
}

// Edges of the Essay.
func (Essay) Edges() []ent.Edge {
	return nil
}
