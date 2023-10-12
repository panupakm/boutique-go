package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.UUID("id", uuid.UUID{}).Immutable().Unique().Default(uuid.New).StorageKey("oid"),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password_hash").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}
