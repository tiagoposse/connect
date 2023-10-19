package auth

import (
	"github.com/tiagoposse/connect/ent"
)

type IdentityProvider interface {
	GetUsersAndMemberships() ([]*ent.User, map[string][]string, error)
	GetName() string
	GetGroupMappings() map[string]string
	GetIgnores() *IgnoringCheck
}
