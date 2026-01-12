package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 知识库
type KnowledgeBase struct {
	ent.Schema
}

func (KnowledgeBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the KnowledgeBase.
func (KnowledgeBase) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Comment("知识库名称"),
		field.Enum("model_provider").
			Values("openai", "anthropic", "google", "azure", "cohere", "huggingface", "local").
			Comment("模型供应商"),
		field.String("model").
			NotEmpty().
			Comment("模型名称"),
		field.Int("vector_dimension").
			Positive().
			Comment("向量维度"),
		field.Int("max_batch_document_count").
			Positive().
			Comment("最大批量处理文档数量"),
	}
}
