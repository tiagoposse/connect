package config

import (
	"errors"
	"fmt"
	"net/netip"
	"os"
	"path/filepath"
	"time"

	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/internal/types"
	"github.com/tiagoposse/connect/internal/utils"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v3"
)

type Config struct {
	General   *General   `yaml:"general"`
	Web       *Web       `yaml:"web"`
	Wireguard *Wireguard `yaml:"wireguard"`
	Auth      *Auth      `yaml:"auth"`
	Database  *Database  `yaml:"database"`
}

type General struct {
	LogLevel string
}

type WebTls struct {
	Certificate string `yaml:"certificate"`
	Key         string `yaml:"key"`
}

type Web struct {
	CsrfKey       *utils.ResolverField `yaml:"csrfKey"`
	ExternalUrl   string               `yaml:"externalUrl"`
	Tls           WebTls               `yaml:"tls"`
	ServeFrontend bool                 `yaml:"serveFrontend"`
	FrontendUrl   string               `yaml:"frontendUrl"`
}

type Wireguard struct {
	Interface        string                `yaml:"interface"`
	DnsServers       []string              `yaml:"dnsServers"`
	Cidr             string                `yaml:"cidr"`
	SelfProvisioning bool                  `yaml:"selfProvisioning"`
}

type Auth struct {
	Google       *GoogleAuth `yaml:"google"`
	UserPassword bool        `yaml:"userpass"`
	Session      *Session    `yaml:"session"`
	Admin        *Admin      `yaml:"admin"`
	Groups       map[string]*ent.Group `yaml:"groups"`
}

type Session struct {
	SessionKey    *utils.ResolverField `yaml:"key"`
	JwtExpiration utils.Duration       `yaml:"expiration"`
}

type Database struct {
	Host     string               `yaml:"host"`
	Port     int                  `yaml:"port"`
	Username string               `yaml:"username"`
	Password *utils.ResolverField `yaml:"password"`
	Type     DatabaseType         `yaml:"type"`
	Ssl      bool                 `yaml:"ssl"`
	Database string               `yaml:"database"`
}

type Admin struct {
	Group string `yaml:"group"`
}

type DatabaseType string

const (
	Postgres DatabaseType = "postgres"
	Mysql    DatabaseType = "mysql"
)

type SamlAuth struct {
	IdpMetadata *utils.ResolverField `yaml:"idpMetadata"`
}

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

func NewConfig() (*Config, error) {
	config := &Config{
		Wireguard: &Wireguard{
			Interface:        "wg0",
			DnsServers:       []string{"1.1.1.1"},
			Cidr:             "10.254.0.0/16",
			SelfProvisioning: false,
		},
		General: &General{
			LogLevel: log.InfoLevel.String(),
		},
		Web: &Web{
			ExternalUrl: "https://127.0.0.1:8888",
			Tls: WebTls{
				Certificate: "",
				Key:         "",
			},
		},
		Auth: &Auth{
			UserPassword: true,
			Session: &Session{
				JwtExpiration: utils.Duration(8 * time.Hour),
			},
			Groups: map[string]*ent.Group{
				"super-admins": {
					Name: "Super Admins",
					Cidr: types.Cidr{Prefix: netip.MustParsePrefix("0.0.0.0/32")},
					Scopes: types.AllScopes,
					Rules: []types.Rule{},
				},
			},
			Admin: &Admin{
				Group: "super-admins",
			},
		},
		Database: &Database{
			Host:     "localhost",
			Port:     5432,
			Username: "wireguard",
			Database: "wireguard",
			Password: &utils.ResolverField{},
			Ssl:      false,
		},
	}

	err := envconfig.Process("wg", config)
	if err != nil {
		return nil, err
	}

	confPath := os.Getenv("WG_CONFIG_PATH")
	if confPath == "" {
		confPath = filepath.Join(os.Getenv("HOME"), ".wgportal", "conf.yaml")
	}

	bs, err := os.ReadFile(confPath)
	if errors.Is(err, os.ErrNotExist) {
		log.Warnf("config file not found at %s\n", confPath)
		return config, nil
	} else if err != nil {
		return nil, fmt.Errorf("checking config path: %w", err)
	}

	if err := yaml.Unmarshal(bs, &config); err != nil {
		return nil, fmt.Errorf("decoding read config: %w", err)
	}

	return config, nil
}

func (c *Config) Validate(resolver *utils.Resolver) []error {
	errs := make([]error, 0)

	if c.Database.Type == "" {
		errs = append(errs, errors.New("database type not set, either postgres or mysql need to be set"))
	}

	if err := resolver.Resolve(c.Database.Password); err != nil {
		errs = append(errs, fmt.Errorf("database password not set: %w", err))
	}

	if c.Auth.Google != nil {
		errs = append(errs, c.Auth.Google.Validate(resolver)...)
	}

	return errs
}
