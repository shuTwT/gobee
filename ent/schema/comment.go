package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("post_id").Comment("评论的帖子ID"),
		field.Int("page_id").Comment("评论的页面ID"),
		field.String("content").NotEmpty().MaxLen(1024).Comment("评论内容"),
		field.Int("user_id").Comment("评论者用户ID"),
		field.Int("status").Comment("状态(1未审核,2已审核)"),
		field.String("user_agent").Optional().Nillable().Comment("评论者用户代理"),
		field.String("ip_address").MaxLen(45).Comment("评论者IP地址"),
		field.String("ip_location").Optional().Nillable().MaxLen(255).Comment("评论者IP位置"),
		field.Bool("pinned").Default(false).Comment("是否置顶"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
