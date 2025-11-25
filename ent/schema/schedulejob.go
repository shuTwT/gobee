package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ScheduleJob holds the schema definition for the ScheduleJob entity.
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
		field.String("name").Comment("名称"),
	}
}

// Edges of the ScheduleJob.
func (ScheduleJob) Edges() []ent.Edge {
	return nil
}
