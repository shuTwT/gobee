package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 相册相片
type AlbumPhoto struct {
	ent.Schema
}

func (AlbumPhoto) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the AlbumPhoto.
func (AlbumPhoto) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("相片名称"),
		field.String("image_url").Comment("图片地址"),
		field.String("description").Comment("相片描述"),
		field.Int("view_count").Default(0).Comment("查看次数"),
		field.Int("album_id").Comment("相册ID"),
	}
}

// Edges of the AlbumPhoto.
func (AlbumPhoto) Edges() []ent.Edge {
	return nil
}
