package schema

import (
	"blug/internal/pkg"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Article struct {
	ent.Schema
}

// Fields of the Bank.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("title").Unique(),
		field.String("desc"),
		field.String("category"),
		field.String("tags"),
		field.String("url"),
		field.String("create_time").Default(pkg.NowTimeStr()).Immutable(),
		field.Bool("is_show").Default(true),
		field.Text("content").Default(""),
	}
}

// Edges of the Bank.
func (Article) Edges() []ent.Edge {
	return nil
}
