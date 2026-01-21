package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type FLinkApplication struct {
	ent.Schema
}

func (FLinkApplication) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (FLinkApplication) Fields() []ent.Field {
	return []ent.Field{
		field.String("website_url").NotEmpty().Comment("网站链接"),
		field.String("application_type").NotEmpty().Comment("申请类型: create/update"),
		field.String("website_name").NotEmpty().Comment("网站名称"),
		field.String("website_logo").NotEmpty().Comment("网站logo"),
		field.String("website_description").NotEmpty().Comment("网站描述"),
		field.String("contact_email").NotEmpty().Comment("联系邮箱"),
		field.String("snapshot_url").Optional().Comment("网页快照"),
		field.String("original_website_url").Optional().Comment("原网站链接"),
		field.String("modification_reason").Optional().Comment("修改原因"),
		field.Int("status").Default(0).Comment("审批状态: 0-待审批, 1-已通过, 2-已拒绝"),
		field.String("reject_reason").Optional().Comment("拒绝原因"),
	}
}
