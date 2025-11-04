package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().MaxLen(255).Comment("文章标题"),
		field.Text("content").NotEmpty().Comment("文章内容"),
		field.Bool("is_published").Default(false).Comment("是否已发布"),
		field.Time("published_at").Optional().Comment("发布时间"),
		field.Int("view_count").Default(0).NonNegative().Comment("浏览次数"),
		field.Int("comment_count").Default(0).NonNegative().Comment("评论次数"),
		field.String("cover").Optional().MaxLen(512).Comment("文章封面"),
		field.String("keywords").Optional().MaxLen(255).Comment("文章关键词"),
		field.String("copyright").Optional().MaxLen(512).Comment("文章版权"),
		field.String("author").Default("匿名作者").Comment("作者"),
		field.String("summary").Optional().MaxLen(512).Comment("文章摘要"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
