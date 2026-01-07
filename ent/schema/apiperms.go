package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// 接口权限
type ApiPerms struct {
	ent.Schema
}

func (ApiPerms) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ApiPerms.
func (ApiPerms) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("path").Comment("请求路径"),
		field.String("method").Comment("请求方法"),
		field.String("desc").Comment("描述"),
		field.String("permission_type").Default("private").Comment("权限类型"),
		field.JSON("roles", []string{}).Default([]string{}).Comment("角色"),
		field.String("status").Default("active").Comment("状态"),
	}
}

func (ApiPerms) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("path", "method").Unique(),
	}
}

// Edges of the ApiPerms.
func (ApiPerms) Edges() []ent.Edge {
	return nil
}
