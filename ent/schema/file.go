package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(255).Comment("文件名称"),
		field.String("path").Default("").MaxLen(512).Comment("文件路径"),
		field.String("url").NotEmpty().MaxLen(1024).Comment("文件地址"),
		field.String("type").Comment("文件类型"),
		field.String("size").Comment("文件大小"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
