/**
 * 相册
 */
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

func (Album) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(255),
		field.String("description").Optional().MaxLen(255),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return nil
}
