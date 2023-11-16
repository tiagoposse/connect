package schemas

import (
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/schemast"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen"
	"github.com/tiagoposse/connect/filter"
	"github.com/tiagoposse/connect/internal/types"
	authz "github.com/tiagoposse/go-auth/controller"
	ogauth "github.com/tiagoposse/ogent-auth"
)

var GroupSchemaMutator = &schemast.UpsertSchema{
	Name: "Group",
	Fields: []ent.Field{
		field.String("id").DefaultFunc(func() string {
			// An example of a dumb ID generator - use a production-ready alternative instead.
			uuid, _ := uuid.NewUUID()
			return uuid.String()
		}),
		field.String("name"),
		field.Other("scopes", authz.Scopes{}).
			SchemaType(map[string]string{dialect.Postgres: "varchar"}).
			Default(authz.Scopes{types.UserAll}).
			Annotations(
				entoas.Schema(ogen.String().
					AsEnum(nil, types.AllScopes.ToRaw()...).
					AsArray(),
				),
			),

		field.String("cidr").GoType(types.Cidr{}).SchemaType(map[string]string{
			dialect.Postgres: "cidr",
		}).Annotations(entoas.Schema(&ogen.Schema{Type: "string"})),
		field.JSON("rules", []types.Rule{}).Annotations(entoas.Schema(
			ogen.NewSchema().AddRequiredProperties(
				ogen.String().AsEnum(nil,
					types.RuleAllow.ToRawMessage(),
					types.RuleDeny.ToRawMessage(),
				).ToProperty("type"),
				ogen.String().ToProperty("target"),
			).AsArray(),
		)),
	},
	Edges: []ent.Edge{
		withType(edge.To("users", placeholder.Type), "User"),
	},
	Annotations: []schema.Annotation{
		ogauth.WithCreateScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithUpdateScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithDeleteScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithListScopes(types.AdminAll, types.AdminGroupsWrite, types.AdminGroupsReadOnly),
		ogauth.WithReadScopes(types.AdminAll, types.AdminGroupsWrite, types.AdminGroupsReadOnly),
		filter.WithFieldFilter("id", "name", "scopes", "cidr"),
	},
}
