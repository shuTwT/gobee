package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DocLibraryDetail struct {
	ent.Schema
}

func (DocLibraryDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (DocLibraryDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("文档标题"),
		field.String("version").
			Optional().
			Comment("文档版本"),
		field.Text("content").
			Comment("文档内容"),
		field.Int("parent_id").
			Optional().
			Comment("父文档ID，用于构建树形结构"),
		field.String("path").
			Optional().
			Comment("文档路径"),
		field.String("url").
			Optional().
			Comment("文档URL"),
		field.String("language").
			Optional().
			Default("zh").
			Comment("文档语言"),
		field.Int("library_id").
			Optional().
			Comment("文档库ID"),
	}
}

func (DocLibraryDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("library", DocLibrary.Type).
			Ref("details").
			Field("library_id").
			Unique(),
	}
}
