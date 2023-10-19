package schema

import (
	"encoding/json"
	"fmt"

	"github.com/tiagoposse/connect/internal/types"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"

	authz "github.com/tiagoposse/go-auth/controller"
	ogauth "github.com/tiagoposse/ogent-auth"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Annotations of the Group.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		ogauth.WithCreateScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithUpdateScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithDeleteScopes(types.AdminAll, types.AdminGroupsWrite),
		ogauth.WithListScopes(types.AdminAll, types.AdminGroupsWrite, types.AdminGroupsReadOnly),
		ogauth.WithReadScopes(types.AdminAll, types.AdminGroupsWrite, types.AdminGroupsReadOnly),
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Other("scopes", authz.Scopes{}).
			SchemaType(map[string]string{dialect.Postgres: "varchar"}).
			Default(authz.Scopes{types.UserAll}).
			Annotations(
				entoas.Schema(ogen.String().
					AsEnum(nil,
						json.RawMessage(fmt.Sprintf(`"%s"`, types.UserAll)),
						json.RawMessage(fmt.Sprintf(`"%s"`, types.AdminAll)),
					).
					AsArray(),
				),
			),

		field.String("cidr").GoType(types.Cidr{}).SchemaType(map[string]string{
			dialect.Postgres: "cidr",
		}).Annotations(entoas.Schema(&ogen.Schema{Type: "string"})),
		field.JSON("rules", []types.Rule{}).Annotations(entoas.Schema(
			ogen.NewSchema().AddRequiredProperties(
				ogen.String().ToProperty("id"),
				ogen.String().AsEnum(nil, json.RawMessage(`"allow"`), json.RawMessage(`"deny"`)).ToProperty("type"),
				ogen.String().ToProperty("target"),
			).AsArray(),
		)),
	}
}

// Edges of the user.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Annotations(
			entoas.CreateOperation(entoas.OperationGroups("create")),
			entoas.ListOperation(entoas.OperationGroups("list")),
			entoas.ReadOperation(entoas.OperationGroups("read")),
			entoas.UpdateOperation(entoas.OperationGroups("update")),
		),
	}
}
