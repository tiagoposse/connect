package controller

import (
	"context"
	"fmt"

	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/ogent"
	"github.com/tiagoposse/connect/internal/auth"
	"github.com/tiagoposse/connect/internal/config"
	"github.com/tiagoposse/connect/internal/sessions"
	"github.com/tiagoposse/connect/internal/utils"
	"github.com/tiagoposse/connect/internal/wireguard"

	log "github.com/sirupsen/logrus"

	ogauth "github.com/tiagoposse/connect/ent/ogentauth"

	authz "github.com/tiagoposse/go-auth/controller"

	"entgo.io/ent/dialect/sql/sqlgraph"
)

func NewController(client *ent.Client, cfg *config.Config) (*Controller, error) {
	google, err := auth.NewGoogleAuthController(cfg.Web.ExternalUrl, cfg.Auth.Google)
	if err != nil {
		return nil, err
	}

	secHandler, err := ogauth.NewOgentAuthHandler(sessions.NewSecurityHandler(cfg.Auth.Session, client))
	if err != nil {
		log.Fatal(err)
	}

	return &Controller{
		OgentHandler: ogent.NewOgentHandler(client),
		wg: wireguard.NewWireGuardManager(cfg.Wireguard),
		client: client,
		cfg: cfg,
		google: google,
		auth: secHandler,
	}, nil
}

type Controller struct {
	*ogent.OgentHandler
	client *ent.Client
	wg     *wireguard.WireGuardManager
	google *auth.GoogleAuthController
	cfg    *config.Config
	auth   *ogauth.OgentAuthHandler
}

func (c *Controller) Init(ctx context.Context) error {
	wires := &ogent.OgentHandlerWiring{
		CreateDeviceWire: c.CreateDeviceWire,
		DeleteDeviceWire: c.DeleteDeviceWire,
	}
	c.SetWiring(wires)

	for gID, g := range c.cfg.Auth.Groups {
		if err := c.client.Group.Create().
			SetID(gID).
			SetCidr(g.Cidr).
			SetName(g.Name).
			SetRules(g.Rules).
			OnConflictColumns(group.FieldID).
			UpdateNewValues().Exec(ctx); err != nil {
			return err
		}
	}

	password := utils.GenerateRandomString(20, true, true, true)
	salt := utils.GenerateRandomString(16, true, true, true)
	apiKey := utils.GenerateRandomString(32, true, false, true)

	if _, err := c.client.User.Create().
		SetID("1").
		SetFirstname("Admin").
		SetLastname("User").
		SetEmail("admin@local.com").
		SetPassword(fmt.Sprintf("%s.%s", password, salt)).
		SetSalt(salt).
		SetGroupID("0").
		SetProvider("userpass").
		Save(ctx); err == nil {
		log.Printf("Admin password is %s", password)
	} else if !sqlgraph.IsUniqueConstraintError(err) {
		return err
	}

	if _, err := c.client.ApiKey.Create().
		SetUserID("1").
		SetKey(apiKey).
		SetScopes(authz.Scopes{
			sessions.AdminAll,
		}).
		SetName("default").
		Save(ctx); err == nil {
		log.Printf("Admin API Key is %s", apiKey)
	} else if !utils.IsMaxItemsError(err) {
		return err
	}

	return nil
}
