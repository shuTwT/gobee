/**
 * 用户
 */
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			NotEmpty().MaxLen(255).
			Immutable(),
		field.Bool("email_verified").
			Default(false),
		field.String("name").
			Default("unknown"),
		field.String("phone_number").
			Optional().
			Unique().
			MaxLen(20),
		field.Bool("phone_number_verified").
			Default(false),
		field.String("password").
			NotEmpty().MaxLen(255).
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
