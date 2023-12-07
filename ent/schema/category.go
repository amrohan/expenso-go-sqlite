package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("icon"),
		field.Int("userId"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
