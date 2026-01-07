package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 访问令牌
type Oauth2AccessToken struct {
	ent.Schema
}

func (Oauth2AccessToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Oauth2AccessToken.
func (Oauth2AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户ID"),
		field.String("access_token").NotEmpty().MaxLen(255).Comment("访问令牌"),
		field.String("refresh_token").NotEmpty().MaxLen(32).Comment("刷新令牌"),
		field.String("client_id").NotEmpty().MaxLen(255).Comment("客户端ID"),
		field.String("scope").NotEmpty().MaxLen(255).Comment("授权范围"),
		field.Time("expire_at").Comment("过期时间"),
	}
}

// Edges of the Oauth2AccessToken.
func (Oauth2AccessToken) Edges() []ent.Edge {
	return nil
}
