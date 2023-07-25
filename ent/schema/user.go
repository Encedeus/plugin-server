package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
		field.Time("deleted_at").Optional(),
		field.Time("auth_updated_at").Optional().Default(time.Now),
		field.String("email").MaxLen(32).Unique(),
		field.String("name").MaxLen(32).Unique(),
		field.String("password"),
		field.Bool("verified").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
