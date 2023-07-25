package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VerifySession holds the schema definition for the User entity.
type VerifySession struct {
	ent.Schema
}

// Fields of the VerifySession.
func (VerifySession) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("sid", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the VerifySession.
func (VerifySession) Edges() []ent.Edge {
	return nil
}
