package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RefreshToken 刷新令牌
type RefreshToken struct {
	ent.Schema
}

func (RefreshToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the RefreshToken.
func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		// 仅存储 refresh token 的 SHA256 哈希，不存明文
		field.String("token_hash").NotEmpty().Unique().Comment("refresh token 的SHA256哈希"),
		field.Int("user_id").Comment("所属用户ID"),
		field.Time("expires_at").Comment("过期时间"),
		field.Bool("revoked").Default(false).Comment("是否已吊销"),
		field.String("user_agent").Optional().Comment("签发时的用户代理"),
		field.String("ip").Optional().Comment("签发时的IP"),
	}
}

// Edges of the RefreshToken.
func (RefreshToken) Edges() []ent.Edge {
	return nil
}

func (RefreshToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// GetRefreshTokenExpiration 返回刷新令牌的过期时间（7天）
func GetRefreshTokenExpiration() time.Time {
	return time.Now().Add(time.Hour * 24 * 7)
}
