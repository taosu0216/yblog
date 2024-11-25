package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Friend struct {
	ent.Schema
}

func (Friend) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("Title").Default(""),
		field.String("Desc").Default(""),
		field.String("Link").Default(""),
		field.String("Avatar").Default(""),
		field.Time("create_time").Default(time.Now).
			Immutable(),
	}
}

func (Friend) Edges() []ent.Edge {
	return nil
}
