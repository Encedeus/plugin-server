package schema

import (
	"crypto/rand"
	"encoding/base64"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/google/uuid"
)

// Source holds the schema definition for the Source entity.
type VerificationSession struct {
	ent.Schema
}

// Fields of the Source.
func (VerificationSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			randomId := make([]byte, 12)
			rand.Read(randomId)
			randomStr := base64.StdEncoding.EncodeToString(randomId)

			fmt.Println(randomStr)

			return randomStr
		}),
		field.UUID("user_id", uuid.UUID{}).Unique(),
	}
}

func (VerificationSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("session", User.Type).
			Field("user_id").
			Unique().
			Required(),
	}
}
