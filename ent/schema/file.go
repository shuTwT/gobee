package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 文件
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
		field.Int("storage_strategy_id").Optional().Comment("存储策略ID"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("storage_strategy", StorageStrategy.Type).
			Field("storage_strategy_id").
			Unique(),
	}
}
