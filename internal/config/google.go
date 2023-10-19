package config

import (
	"errors"
	"fmt"

	"github.com/tiagoposse/connect/internal/utils"
)

type GoogleAuth struct {
	*SamlAuth         `yaml:",inline"`
	Domain            string               `yaml:"domain"`
	ServiceAccountKey *utils.ResolverField `yaml:"serviceAccountKey"`
	UserToImpersonate *string              `yaml:"userToImpersonate"`
	AdminGroups       []string             `yaml:"adminGroups"`
	IgnoreUsers       []string             `yaml:"ignoreUsers"`
	IgnoreGroups      []string             `yaml:"ignoreGroups"`
	GroupFilters      []string             `yaml:"groupFilters"`
	UserFilters       []string             `yaml:"userFilters"`
	GroupMappings     map[string]string    `yaml:"groupMappings"`
}

func (ga *GoogleAuth) Validate(resolver *utils.Resolver) []error {
	errs := make([]error, 0)

	if err := resolver.Resolve(ga.IdpMetadata); err != nil {
		errs = append(errs, fmt.Errorf("google idp metadata not set: %w", err))
	}

	if ga.UserToImpersonate == nil {
		errs = append(errs, errors.New("google user to impersonate not set"))
	}

	if err := resolver.Resolve(ga.ServiceAccountKey); err != nil {
		errs = append(errs, fmt.Errorf("google service account not set: %w", err))
	}

	if ga.Domain == "" {
		errs = append(errs, errors.New("google domain not set"))
	}

	return errs
}
