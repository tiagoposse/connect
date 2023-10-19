package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"
	"github.com/tiagoposse/connect/filter"
	"github.com/tiagoposse/connect/internal/types"
	exclusion "github.com/tiagoposse/entoas-fields-exclusion"
	ogauth "github.com/tiagoposse/ogent-auth"
)

// Todo holds the schema definition for the Todo entity.
type Device struct {
	ent.Schema
}

// Fields of the Todo.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").Optional(),
		field.String("type"),
		field.String("public_key").Immutable().Unique(),
		field.String("preshared_key").
			Immutable().Unique().
			Annotations(
				entoas.Groups("create"),
				exclusion.SkipCreate(),
				exclusion.SkipUpdate(),
			),
		field.String("endpoint").
			Unique().GoType(types.Inet{}).
			Annotations(
				exclusion.SkipCreate(),
				entoas.Schema(&ogen.Schema{Type: "string"}),
			).
			SchemaType(map[string]string{
			dialect.Postgres: "inet",
		}).Unique(),
		field.String("allowed_ips").Annotations(exclusion.SkipCreate()),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("devices").Unique(),
	}
}

func (Device) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationGroups("create")),
		entoas.ListOperation(entoas.OperationGroups("list")),
		entoas.ReadOperation(entoas.OperationGroups("read")),
		entoas.UpdateOperation(entoas.OperationGroups("update")),
		filter.Annotation{FilterFields: []*filter.Opt{{Name: "user", In: "query"}}},
		ogauth.WithCreateScopes(types.AdminAll, types.AdminDevicesWrite),
		ogauth.WithUpdateScopes(types.AdminAll, types.AdminDevicesWrite),
		ogauth.WithDeleteScopes(types.AdminAll, types.AdminDevicesWrite),
		ogauth.WithListScopes(types.AdminAll, types.AdminDevicesWrite, types.AdminDevicesReadOnly),
		ogauth.WithReadScopes(types.AdminAll, types.AdminDevicesWrite, types.AdminDevicesReadOnly),
	}
}
