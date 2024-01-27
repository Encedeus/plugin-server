package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/pluginServer/ent/publication"
	"github.com/google/uuid"
	"time"
)

// Publication holds the schema definition for the Publication entity.
type Publication struct {
	ent.Schema
}

// Fields of the Source.
func (Publication) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Bool("is_deprecated").Default(false),
		field.String("name").MaxLen(32),
		field.String("tag"),
		field.String("uri_to_file"),
		field.UUID("plugin_id", uuid.UUID{}),
	}
}

func (Publication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plugin", Plugin.Type).
			Ref("publications").
			Field("plugin_id").
			Unique().
			Required(),
	}
}

func OrderByHook(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {

		if query, ok := mutation.(interface {
			Order(o ...publication.OrderOption)
		}); ok {
			query.Order(publication.ByCreatedAt(sql.OrderDesc()))
		}

		value, err := next.Mutate(ctx, mutation)

		return value, err
	})
}
