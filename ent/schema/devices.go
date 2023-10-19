package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen"
	"github.com/tiagoposse/connect/filter"
	"github.com/tiagoposse/connect/internal/types"
	exclusion "github.com/tiagoposse/entoas-fields-exclusion"
	ogauth "github.com/tiagoposse/ogent-auth/authorization"
)

// Todo holds the schema definition for the Todo entity.
type Device struct {
	ent.Schema
}

// Fields of the Todo.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("description").Optional(),
		field.String("type"),
		field.Other("dns", types.InetSlice{}).
			SchemaType(map[string]string{dialect.Postgres: "varchar"}).
			Annotations(
				entoas.Schema(ogen.String().AsArray()),
			),
		field.String("public_key").Immutable().Unique(),
		field.String("preshared_key").
			Immutable().Unique().
			Annotations(
				entoas.Groups("create"),
				exclusion.SkipCreate(),
			),
		field.Bool("keep_alive"),
		field.String("endpoint").
			Unique().GoType(types.Inet{}).
			Annotations(
				entoas.Schema(ogen.String()),
			).
			SchemaType(map[string]string{
			dialect.Postgres: "inet",
		}).Unique(),
		field.Other("allowed_ips", types.CidrSlice{}).
		SchemaType(map[string]string{dialect.Postgres: "varchar"}).
		Annotations(
			entoas.Schema(ogen.String().AsArray()),
		),
		field.String("user_id").Immutable(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("devices").Unique().Immutable().Required().Field("user_id"),
	}
}

func (Device) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationGroups("create")),
		entoas.ListOperation(entoas.OperationGroups("list")),
		entoas.ReadOperation(entoas.OperationGroups("read")),
		entoas.UpdateOperation(entoas.OperationGroups("update")),
		filter.WithFieldFilter("id", "user", "name", "type", "endpoint", "allowed_ips", "public_key"),
		ogauth.WithCreateScopes(types.AdminAll, types.AdminDevicesWrite, types.UserAll, types.UserDevicesWrite),
		ogauth.WithUpdateScopes(types.AdminAll, types.AdminDevicesWrite, types.UserAll, types.UserDevicesWrite),
		ogauth.WithDeleteScopes(types.AdminAll, types.AdminDevicesWrite, types.UserAll, types.UserDevicesWrite),
		ogauth.WithListScopes(types.AdminAll, types.AdminDevicesWrite, types.AdminDevicesReadOnly, types.UserAll, types.UserDevicesWrite, types.UserDevicesReadOnly),
		ogauth.WithReadScopes(types.AdminAll, types.AdminDevicesWrite, types.AdminDevicesReadOnly, types.UserAll, types.UserDevicesWrite, types.UserDevicesReadOnly),
	}
}
