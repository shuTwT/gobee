package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 定时任务
type ScheduleJob struct {
	ent.Schema
}

func (ScheduleJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ScheduleJob.
func (ScheduleJob) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("任务名称"),
		field.String("type").Comment("任务类型: cron, interval, once"),
		field.String("expression").Comment("调度表达式: cron表达式或时间间隔"),
		field.Text("description").Optional().Comment("任务描述"),
		field.Bool("enabled").Default(true).Comment("是否启用"),
		field.Time("next_run_time").Optional().Comment("下次执行时间"),
		field.Time("last_run_time").Optional().Comment("上次执行时间"),
		field.String("execution_type").Default("http").Comment("执行类型: http, internal, command, mq"),
		field.String("http_method").Optional().Comment("HTTP方法: GET, POST, PUT, DELETE"),
		field.String("http_url").Optional().Comment("HTTP URL"),
		field.JSON("http_headers", map[string]string{}).Optional().Comment("HTTP请求头"),
		field.Text("http_body").Optional().Comment("HTTP请求体"),
		field.Int("http_timeout").Default(30).Comment("HTTP超时时间(秒)"),
		field.Int("max_retries").Default(3).Comment("最大重试次数"),
		field.Bool("failure_notification").Default(false).Comment("失败是否通知"),
	}
}

// Edges of the ScheduleJob.
func (ScheduleJob) Edges() []ent.Edge {
	return nil
}
