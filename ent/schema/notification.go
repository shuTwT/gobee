package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 通知
type Notification struct {
	ent.Schema
}

func (Notification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().MaxLen(255).Comment("通知标题"),
		field.String("content").Comment("通知内容"),
		field.Time("publish_time").Optional().Comment("发布时间"),
		field.Bool("is_read").Default(false).Comment("是否已读"),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return nil
}
