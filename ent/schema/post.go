package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().MaxLen(255).Comment("文章标题"),
		field.String("alias").NotEmpty().Optional().MaxLen(255).Comment("文章别名"),
		field.Text("content").NotEmpty().Comment("文章内容"),
		field.Text("md_content").Optional().Nillable().Comment("md文章内容"),
		field.Text("html_content").Optional().Nillable().Comment("html文章内容"),
		field.Enum("content_type").Values("markdown", "html").Default("html").Comment("内容类型"),
		field.Enum("status").Values("draft", "published", "archived").Default("draft").Comment("状态"),
		field.Bool("is_autogen_summary").Default(false).Comment("是否自动生成摘要"),
		field.Bool("is_visible").Default(true).Comment("是否可见"),
		field.Bool("is_tip_to_top").Default(false).Comment("是否置顶"),
		field.Bool("is_allow_comment").Default(true).Comment("是否允许评论"),
		field.Bool("is_visible_after_comment").Default(false).Comment("是否评论后可见"),
		field.Bool("is_visible_after_pay").Default(false).Comment("是否支付后可见"),
		field.Int("money").Default(0).NonNegative().Comment("支付金额"),
		field.Time("published_at").Optional().Nillable().Comment("发布时间"),
		field.Int("view_count").Default(0).NonNegative().Comment("浏览次数"),
		field.Int("comment_count").Default(0).NonNegative().Comment("评论次数"),
		field.String("cover").Optional().MaxLen(512).Comment("文章封面"),
		field.String("keywords").Optional().MaxLen(255).Comment("文章关键词"),
		field.String("copyright").Optional().MaxLen(512).Comment("文章版权"),
		field.String("author").Default("匿名作者").Comment("作者"),
		field.String("summary").Optional().MaxLen(512).Comment("文章摘要"),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).Ref("posts"),
		edge.From("tags", Tag.Type).Ref("posts"),
	}
}
