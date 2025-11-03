package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Oauth2RefreshToken holds the schema definition for the Oauth2RefreshToken entity.
type Oauth2RefreshToken struct {
	ent.Schema
}

func (Oauth2RefreshToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Oauth2RefreshToken.
func (Oauth2RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户ID"),
		field.String("refresh_token").NotEmpty().MaxLen(32).Comment("刷新令牌"),
		field.String("client_id").NotEmpty().MaxLen(255).Comment("客户端ID"),
		field.String("scope").NotEmpty().MaxLen(255).Comment("授权范围"),
		field.Time("expire_at").Comment("过期时间"),
	}
}

// Edges of the Oauth2RefreshToken.
func (Oauth2RefreshToken) Edges() []ent.Edge {
	return nil
}
