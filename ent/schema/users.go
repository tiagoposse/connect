package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationGroups("create")),
		entoas.ListOperation(entoas.OperationGroups("list")),
		entoas.ReadOperation(entoas.OperationGroups("read")),
		entoas.UpdateOperation(entoas.OperationGroups("update")),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			// An example of a dumb ID generator - use a production-ready alternative instead.
			uuid, _ := uuid.NewUUID()
			return uuid.String()
		}),
		field.String("email").NotEmpty().Unique(),
		field.String("firstname").NotEmpty(),
		field.String("lastname").NotEmpty(),
		field.String("provider").Immutable().Annotations(
			entoas.Groups("read", "list"),
		),
		field.String("password").Optional().NotEmpty().Sensitive(),
		field.String("salt").Optional().NotEmpty().Sensitive(),
		field.String("photo_url").Optional(),
		field.Bool("disabled").Default(false),
		field.String("disabled_reason").Optional(),
	}
}

// Hooks of the User.
// func (User) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		// Validate password and salt not empty when userpass create.
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
// 					if prov, _ := m.Provider(); prov == "userpass" {
// 						if _, exists := m.Password(); !exists {
// 							return nil, fmt.Errorf("password must provided when user provider is userpass")
// 						} else if _, exists := m.Salt(); !exists {
// 							return nil, fmt.Errorf("salt must provided when user provider is userpass")
// 						}
// 					}
// 					return next.Mutate(ctx, m)
// 				})
// 			},
// 			// Limit the hook only for these operations.
// 			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
// 		),
// 	}
// }

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Unique().Ref("users").Required(),
		edge.To("devices", Device.Type).Annotations(
			entoas.CreateOperation(entoas.OperationGroups("create")),
			entoas.ListOperation(entoas.OperationGroups("list")),
			entoas.ReadOperation(entoas.OperationGroups("read")),
			entoas.UpdateOperation(entoas.OperationGroups("update")),
		),
		edge.To("keys", ApiKey.Type),
	}
}
