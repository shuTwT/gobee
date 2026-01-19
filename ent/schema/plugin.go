package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Plugin struct {
	ent.Schema
}

func (Plugin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Plugin) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Unique().MaxLen(100).Comment("插件标识，用于pluginMap的key"),
		field.String("name").NotEmpty().MaxLen(100).Comment("插件名称"),
		field.String("version").NotEmpty().MaxLen(20).Comment("插件版本号"),
		field.String("description").MaxLen(500).Optional().Comment("插件描述"),
		field.String("bin_path").NotEmpty().MaxLen(255).Comment("插件二进制文件路径"),
		field.Uint("protocol_version").Default(1).Comment("协议版本"),
		field.String("magic_cookie_key").Default("GO_PLUGIN").MaxLen(100).Comment("Magic Cookie Key"),
		field.String("magic_cookie_value").NotEmpty().MaxLen(255).Comment("Magic Cookie Value"),
		field.JSON("dependencies", []string{}).Optional().Comment("依赖的其他插件的key列表"),
		field.Text("config").Optional().Comment("插件配置JSON"),
		field.Bool("enabled").Default(true).Comment("是否启用"),
		field.Bool("auto_start").Default(false).Comment("是否自动启动"),
		field.Enum("status").Values("stopped", "running", "error", "loading").Default("stopped").Comment("插件状态"),
		field.String("last_error").MaxLen(1000).Optional().Comment("最后错误信息"),
		field.Time("last_started_at").Optional().Comment("最后启动时间"),
		field.Time("last_stopped_at").Optional().Comment("最后停止时间"),
	}
}

func (Plugin) Edges() []ent.Edge {
	return nil
}
