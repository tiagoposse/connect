package schemas

import (
	"entgo.io/ent"
)

type placeholder struct {
	ent.Schema
}

func withType(e ent.Edge, typeName string) ent.Edge {
	e.Descriptor().Type = typeName
	return e
}
