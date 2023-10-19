package controller

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/ogent"
	"github.com/tiagoposse/connect/internal/auth"
	"github.com/tiagoposse/connect/internal/config"
	"github.com/tiagoposse/connect/internal/sessions"
	"github.com/tiagoposse/connect/internal/types"
	"github.com/tiagoposse/connect/internal/utils"
	"github.com/tiagoposse/connect/internal/wireguard"

	log "github.com/sirupsen/logrus"

	ogauth "github.com/tiagoposse/connect/ent/ogentauth"

	authz "github.com/tiagoposse/go-auth/authorization"

	"entgo.io/ent/dialect/sql/sqlgraph"
)

type ControllerOption func (c *Controller)

func NewController(client *ent.Client, cfg *config.Config, opts ...ControllerOption) (*Controller, error) {
	google, err := auth.NewGoogleAuthController(cfg.Web.ExternalUrl, cfg.Auth.Google)
	if err != nil {
		return nil, err
	}

	ctrl := &Controller{
		OgentHandler: ogent.NewOgentHandler(client),
		wg: wireguard.NewWireGuardManager(cfg.Wireguard),
		client: client,
		cfg: cfg,
		google: google,
		auth: ogauth.NewSecurityHandler(sessions.NewSecurityHandler(cfg.Auth.Session, client)),
	}

	for _, o := range opts {
		o(ctrl)
	}

	return ctrl, nil
}

type Controller struct {
	*ogent.OgentHandler
	client *ent.Client
	wg     *wireguard.WireGuardManager
	google *auth.GoogleAuthController
	cfg    *config.Config
	auth   *ogauth.SecurityHandler
}

func (c *Controller) GetAuthorizationHandler() *ogauth.SecurityHandler {
	return c.auth
}

func (c *Controller) Init(ctx context.Context) error {
	wires := &ogent.OgentHandlerWiring{
		DeleteDeviceWire: c.DeleteDeviceWire,
		CreateGroupWire: c.CreateGroupWire,
		UpdateGroupWire: c.UpdateGroupWire,
		DeleteGroupWire: c.DeleteGroupWire,
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
		SetGroupID("super-admins").
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

	currentGroups, err := c.client.Group.Query().All(ctx)
	if err != nil {
		return err
	}

	iptables := make([]string, 0)
	for _, g := range currentGroups {
		hasAllowRules := false
		for _, rule := range g.Rules {
			if rule.Type == types.RuleAllow {
				hasAllowRules = true
				iptables = append(iptables, fmt.Sprintf("-A FORWARD -s %s -d %s -j ACCEPT", g.Cidr, rule.Target))
			} else {
				iptables = append(iptables, fmt.Sprintf("-A FORWARD -s %s -d %s -j DROP", g.Cidr, rule.Target))
			}
		}

		if hasAllowRules {
			iptables = append(iptables, fmt.Sprintf("-A FORWARD -s %s -d 0.0.0.0/0 -j DROP", g.Cidr))
		}
	}

	for _, rule := range iptables {
		cmd := exec.Command("iptables", strings.Split(rule, " ")...)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
