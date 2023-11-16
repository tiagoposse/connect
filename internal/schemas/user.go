package schemas

import (
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/schemast"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tiagoposse/connect/filter"
)

var UserSchemaMutator = &schemast.UpsertSchema{
	Name: "User",
	Fields: []ent.Field{
		field.String("id").DefaultFunc(func() string {
			// An example of a dumb ID generator - use a production-ready alternative instead.
			uuid, _ := uuid.NewUUID()
			return uuid.String()
		}),
		field.String("email").NotEmpty().Unique(),
		field.String("firstname").NotEmpty(),
		field.String("lastname").NotEmpty(),
		field.String("provider").Immutable().NotEmpty(),
		field.String("password").Optional().NotEmpty().Sensitive(),
		field.String("salt").Optional().NotEmpty().Sensitive(),
		field.String("photo_url").Optional(),
		field.Bool("disabled").Default(false),
		field.String("disabled_reason").Optional(),
		field.String("group_id"),
	},
	Edges: []ent.Edge{
		withType(edge.To("devices", placeholder.Type), "Device"),
		withType(edge.To("keys", placeholder.Type), "ApiKey"),
		withType(edge.To("audit", placeholder.Type).Annotations(
			entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		), "Audit"),
		withType(edge.From("group", placeholder.Type).Unique().Ref("users").Required().Field("group_id"), "Group"),
	},
	Annotations: []schema.Annotation{
		entoas.CreateOperation(entoas.OperationGroups("create")),
		entoas.ListOperation(entoas.OperationGroups("list")),
		entoas.ReadOperation(entoas.OperationGroups("read")),
		filter.WithFieldFilter("id", "firstname", "lastname", "provider", "email", "disabled", "disabled_reason"),
	},
}
