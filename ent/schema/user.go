package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/pluginServer/ent/hook"
	"github.com/Encedeus/pluginServer/ent/user"
	"github.com/google/uuid"
	"time"
)

// Source holds the schema definition for the Source entity.
type User struct {
	ent.Schema
}

// Fields of the Source.
func (User) Fields() []ent.Field {

	//config.InitConfig()

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
		field.Time("auth_updated_at").Default(time.Now),
		field.Time("deleted_at").Optional(),
		field.String("email").MaxLen(32 /*config.Config.Validation.MaxEmailLen*/).Unique(),
		field.String("password"),
		field.String("name").MaxLen(32 /*config.Config.Validation.MaxNameLen*/).Unique(),
		field.Bool("email_verified").Default(false),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugins", Plugin.Type),
		edge.From("verification_session", VerificationSession.Type).
			Ref("session"),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(
			AuthUpdateHook,
			hook.And(
				hook.Or(
					hook.HasOp(ent.OpUpdateOne),
					hook.HasOp(ent.OpUpdateOne),
				),
				hook.Or(
					hook.HasFields(user.FieldPassword),
					hook.HasFields(user.FieldDeletedAt),
				),
			),
		),

		hook.If(
			UnVerifyEmailHook,
			hook.And(
				hook.Or(
					hook.HasOp(ent.OpUpdateOne),
					hook.HasOp(ent.OpUpdateOne),
				),
				hook.HasFields(user.FieldEmail),
			),
		),
	}
}

func AuthUpdateHook(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {

		if query, ok := mutation.(interface{ SetAuthUpdatedAt(time.Time) }); ok {
			query.SetAuthUpdatedAt(time.Now())
		}

		value, err := next.Mutate(ctx, mutation)

		return value, err
	})
}

func UnVerifyEmailHook(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {

		if query, ok := mutation.(interface{ SetEmailVerified(bool) }); ok {
			query.SetEmailVerified(false)
		}

		value, err := next.Mutate(ctx, mutation)

		return value, err
	})
}
