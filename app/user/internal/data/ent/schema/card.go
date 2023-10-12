package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the User.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique().Default(uuid.New).StorageKey("oid"),
		field.String("number").NotEmpty().MaxLen(16).MinLen(16),
		field.String("name").NotEmpty().MaxLen(128),
		field.Int("ccv").Min(100).Max(999),
		field.Int("expiration_year").Positive(),
		field.Int("expiration_month").Positive(),
	}
}

// Edges of the User.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("cards").Unique(),
	}
}
