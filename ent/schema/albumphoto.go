package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AlbumPhoto holds the schema definition for the AlbumPhoto entity.
type AlbumPhoto struct {
	ent.Schema
}

// Fields of the AlbumPhoto.
func (AlbumPhoto) Fields() []ent.Field {
	return []ent.Field{
		field.String("image_url").Comment("图片地址"),
		field.Int("view_count").Default(0).Comment("查看次数"),
		field.Int("album_id").Comment("相册ID"),
	}
}

// Edges of the AlbumPhoto.
func (AlbumPhoto) Edges() []ent.Edge {
	return nil
}
