package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Notes holds the schema definition for the Notes entity.
type Notes struct {
	ent.Schema
}
func (Notes) Config() ent.Config {
	return ent.Config{
		Table: "notes",
	}
}
// Fields of the Notes.
func (Notes) Fields() []ent.Field {
	return []ent.Field{
		field.String("Title").
			Unique(),
		field.String("Content"),
		field.Bool("Private").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Notes.
func (Notes) Edges() []ent.Edge {
	return nil
}
