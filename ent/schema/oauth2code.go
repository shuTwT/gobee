package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 授权码
type Oauth2Code struct {
	ent.Schema
}

func (Oauth2Code) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Oauth2Code.
func (Oauth2Code) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户ID"),
		field.String("code").NotEmpty().MaxLen(32).Comment("授权码"),
		field.Time("expire_at").Comment("过期时间"),
		field.String("client_id").NotEmpty().MaxLen(255).Comment("客户端ID"),
		field.String("redirect_uri").NotEmpty().MaxLen(255).Comment("重定向URI"),
		field.String("scope").NotEmpty().MaxLen(255).Comment("授权范围"),
	}
}

// Edges of the Oauth2Code.
func (Oauth2Code) Edges() []ent.Edge {
	return nil
}
