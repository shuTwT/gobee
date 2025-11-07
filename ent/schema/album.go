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
		field.String("name").NotEmpty().Unique().Comment("相册分类名称"),
		field.String("description").Optional().MaxLen(255).Comment("相册分类描述"),
		field.Int("sort").Optional().Default(0).Comment("相册分类排序"),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return nil
}
