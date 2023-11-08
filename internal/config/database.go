package config

import (
	"errors"
	"fmt"

	"github.com/tiagoposse/connect/internal/utils"
)

type DatabaseType string

const (
	Postgres DatabaseType = "postgres"
	Mysql    DatabaseType = "mysql"
)

type Database struct {
	Host     string               `yaml:"host"`
	Port     int                  `yaml:"port"`
	Username string               `yaml:"username"`
	Password *utils.ResolverField `yaml:"password"`
	Type     DatabaseType         `yaml:"type"`
	Ssl      bool                 `yaml:"ssl"`
	Database string               `yaml:"database"`
}

func (db *Database) Validate(resolver *utils.Resolver) []error {
	errs := make([]error, 0)
	if db.Type == "" {
		errs = append(errs, errors.New("database type not set, either postgres or mysql need to be set"))
	}

	if err := resolver.Resolve(db.Password); err != nil {
		errs = append(errs, fmt.Errorf("database password not set: %w", err))
	}

	return errs
}
