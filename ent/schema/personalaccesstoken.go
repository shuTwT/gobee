package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PersonalAccessToken holds the schema definition for the PersonalAccessToken entity.
type PersonalAccessToken struct {
	ent.Schema
}

func (PersonalAccessToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the PersonalAccessToken.
func (PersonalAccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().Comment("名称"),
		field.Time("expires").Optional().Comment("过期时间"),
		field.String("description").Optional().Comment("描述"),
		field.String("token"),
		field.Int("user_id"),
	}
}

// Edges of the PersonalAccessToken.
func (PersonalAccessToken) Edges() []ent.Edge {
	return nil
}
