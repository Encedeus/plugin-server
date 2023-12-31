package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Source holds the schema definition for the Source entity.
type Source struct {
	ent.Schema
}

// Fields of the Source.
func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.String("repository"),
	}
}

func (Source) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plugin", Plugin.Type).
			Ref("source"),
	}
}
