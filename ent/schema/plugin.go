package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Plugin holds the schema definition for the User entity.
type Plugin struct {
	ent.Schema
}

// Fields of the Plugin.
func (Plugin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
		field.String("name").MaxLen(32).Unique(),
		field.String("description").MaxLen(512).Optional(),
		field.String("repo").Optional(),
		field.String("homepage").Optional(),
		field.UUID("owner_id", uuid.UUID{}),
		field.Strings("contributors"),
	}
}

// Edges of the Plugin.
func (Plugin) Edges() []ent.Edge {
	return nil
}
