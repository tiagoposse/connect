package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tiagoposse/connect/filter"
)

// Audit holds the schema definition for the Audit entity.
type Audit struct {
	ent.Schema
}

// Annotations of the User.
func (Audit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		filter.WithFieldFilter("id", "action", "author"),
	}
}

// Fields of the User.
func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			// An example of a dumb ID generator - use a production-ready alternative instead.
			uuid, _ := uuid.NewUUID()
			return uuid.String()
		}),
		field.String("action").NotEmpty().Immutable(),
		field.String("author").NotEmpty().Immutable(),
	}
}

// Edges of the Audit.
func (Audit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("audit").Unique().Immutable().Required(),
	}
}
