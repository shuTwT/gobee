package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// 访问日志
type VisitLog struct {
	ent.Schema
}

func (VisitLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the VisitLog.
func (VisitLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Comment("访问ip"),
		field.String("user_agent").Optional().Nillable().Comment("访问用户代理"),
		field.String("path").Comment("访问路径"),
		field.String("os").Optional().Nillable().Comment("操作系统"),
		field.String("browser").Optional().Nillable().Comment("浏览器"),
		field.String("device").Optional().Nillable().Comment("设备"),
	}
}

// Edges of the VisitLog.
func (VisitLog) Edges() []ent.Edge {
	return nil
}

func (VisitLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ip"),
		index.Fields("path"),
	}
}
