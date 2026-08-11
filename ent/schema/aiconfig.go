package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AIConfig stores the single, server-side OpenAI-compatible provider configuration.
// The provider key is always encrypted before it reaches this table.
type AIConfig struct {
	ent.Schema
}

func (AIConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AIConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("config_key").
			Immutable().
			Unique().
			Comment("单例配置标识"),
		field.String("base_url").
			NotEmpty().
			MaxLen(2048).
			Comment("OpenAI 兼容接口基础地址"),
		field.String("model").
			NotEmpty().
			MaxLen(255).
			Comment("默认模型"),
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
		field.String("api_key_ciphertext").
			NotEmpty().
			Sensitive().
			Comment("AES-GCM 加密后的供应商 API Key"),
	}
}
