package schema

import (
	"github.com/google/uuid"
	"github.com/tiagoposse/connect/filter"
	"github.com/tiagoposse/connect/internal/types"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"

	authz "github.com/tiagoposse/go-auth/authorization"
	ogauth "github.com/tiagoposse/ogent-auth/authorization"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Annotations of the Group.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ogauth.WithCreateScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithUpdateScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithDeleteScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithListScopes(types.AdminAll, types.AdminGroupsWrite, types.AdminGroupsReadOnly),
		ogauth.WithReadScopes(types.AdminAll, types.AdminGroupsWrite, types.AdminGroupsReadOnly),
		filter.WithFieldFilter("id", "name", "scopes", "cidr"),
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
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
	}
}

// Edges of the user.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
