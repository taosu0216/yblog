package schema

import (
	"blug/internal/pkg"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("task_id").Default(""),
		field.String("task_name").Default(""),
		field.String("task_type").Default(""),
		field.String("status").Default(""),
		field.String("reason").Default(""),
		field.String("create_time").Default(pkg.NowTimeStr()).Immutable(),
		field.String("finish_time").Default(""),
	}
}

func (Task) Edges() []ent.Edge {
	return nil
}
