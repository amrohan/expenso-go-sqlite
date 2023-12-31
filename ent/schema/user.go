package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// field.Int("id").Unique().Immutable().StorageKey("user_id"),
		field.UUID("id", uuid.New()).Unique().Immutable().StorageKey("user_id"),
		field.String("username").Unique(),
		field.String("password"),
		field.String("email").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
