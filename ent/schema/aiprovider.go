package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AIProvider stores an OpenAI-compatible provider configuration.
// Multiple providers can coexist; the provider key is always encrypted
// before it reaches this table.
type AIProvider struct {
	ent.Schema
}

func (AIProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AIProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(100).
			Comment("提供商名称"),
		field.String("provider_type").
			NotEmpty().
			MaxLen(50).
			Comment("提供商类型（openai/deepseek/moonshot/zhipu/volcengine/siliconflow/ollama/custom 等）"),
		field.String("base_url").
			NotEmpty().
			MaxLen(2048).
			Comment("OpenAI 兼容接口基础地址"),
		field.String("api_key_ciphertext").
			Optional().
			Sensitive().
			Comment("AES-GCM 加密后的供应商 API Key，本地服务可留空"),
		field.Float("temperature").
			Default(0.7).
			Comment("温度"),
		field.Int("max_tokens").
			Positive().
			Default(2048).
			Comment("最大输出 Token"),
		field.Float("top_p").
			Default(1).
			Comment("Top P"),
		field.Float("frequency_penalty").
			Default(0).
			Comment("频率惩罚"),
		field.Float("presence_penalty").
			Default(0).
			Comment("存在惩罚"),
		field.Bool("is_default").
			Default(false).
			Comment("是否默认提供商（AI 聊天与文章摘要使用）"),
		field.Bool("is_enabled").
			Default(true).
			Comment("是否启用"),
		field.Int("sort").
			Default(0).
			Comment("排序值，越小越靠前"),
		field.String("remark").
			Optional().
			MaxLen(255).
			Comment("备注"),
	}
}

func (AIProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("models", AIModel.Type).
			Comment("该提供商下的模型列表"),
	}
}
