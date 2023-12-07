package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("transaction_id"),
		field.String("title"),
		field.Float("amount"),
		field.String("description").Optional().Nillable(),
		field.Time("date"),
		field.Int("categoryId"),
		field.Int("accountId"),
		field.String("userId"),
		field.String("type"),
		field.String("status").Optional().Nillable(),
		field.String("currency").Optional().Nillable(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return nil
}
