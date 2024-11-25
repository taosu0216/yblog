package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ent.Schema
}

// Fields of the Bank.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("username").Unique(),
		field.String("password"),
		field.Bool("is_root"),
		field.Time("create_time").Default(time.Now).Immutable(),
	}
}

// Edges of the Bank.
func (User) Edges() []ent.Edge {
	return nil
}
