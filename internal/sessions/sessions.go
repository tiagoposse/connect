package sessions

import (
	"context"
	"errors"
	"time"

	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/ent/apikey"
	"github.com/tiagoposse/connect/internal/config"

	authz "github.com/tiagoposse/go-auth/controller"
	sessions "github.com/tiagoposse/go-auth/sessions"
)

type SessionInfo struct {
	ID string `json:"id"`
	PhotoURL string `json:"photo_url"`
	Provider string `json:"provider"`
	Email string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname string  `json:"lastname"`
	Group string  `json:"group"`
	Scopes authz.Scopes `json:"scopes"`
}

const (
	AdminUsersWrite       authz.Scope = "admin.users.write"
	AdminUsersReadOnly    authz.Scope = "admin.users.readonly"
	AdminDevicesWrite     authz.Scope = "admin.groups.write"
	AdminDevicesReadOnly  authz.Scope = "admin.groups.readonly"
	AdminGroupsWrite      authz.Scope = "admin.devices.write"
	AdminGroupsReadOnly   authz.Scope = "admin.devices.readonly"
	AdminSettingsWrite    authz.Scope = "admin.settings.write"
	AdminSettingsReadOnly authz.Scope = "admin.settings.readonly"
	AdminAll              authz.Scope = "admin.*"
	UserDevicesWrite      authz.Scope = "user.devices.write"
	UserApiKeyWrite       authz.Scope = "user.apikey.write"
	UserDevicesReadOnly   authz.Scope = "user.devices.readonly"
	UserAll               authz.Scope = "user.*"
)

type SecurityHandler struct {
	*sessions.SessionsController
	client *ent.Client
}

func NewSecurityHandler(cfg *config.Session, client *ent.Client) *SecurityHandler {
	return &SecurityHandler{
		client: client,
		SessionsController: sessions.NewSessionsController(*cfg.SessionKey.Value, time.Duration(cfg.JwtExpiration)),
	}
}

func (h *SecurityHandler) ValidateApiKeyAuth(ctx context.Context, key string) (authz.ISession, error) {
	ak, err := h.client.ApiKey.Query().Where(apikey.KeyEQ(key)).WithUser(func(uq *ent.UserQuery) { uq.WithGroup() }).First(ctx)
	if err != nil {
		return nil, err
	}
	if ak.Edges.User == nil {
		return nil, errors.New("user not found")
	}

	return &sessions.Session{
		SessionInfo: SessionInfo{
			ID: ak.Edges.User.ID,
			PhotoURL: ak.Edges.User.PhotoURL,
			Provider: ak.Edges.User.Provider,
			Email: ak.Edges.User.Email,
			Firstname: ak.Edges.User.Firstname,
			Lastname: ak.Edges.User.Lastname,
			Group: ak.Edges.User.Edges.Group.ID,
			Scopes: ak.Scopes,
		},
	}, nil
}

func (h *SecurityHandler) ValidateCookieAuth(ctx context.Context, key string) (authz.ISession, error) {
	isess, err := h.ValidateSessionToken(ctx, key)
	info := isess.(*sessions.Session).SessionInfo.(map[string]interface{})

	scopes := authz.Scopes{}
	for _, s := range info["scopes"].([]interface{}) {
		scopes = append(scopes, authz.Scope(s.(string)))
	}

	return &sessions.Session{SessionInfo: SessionInfo{
		ID: info["id"].(string),
		Email: info["email"].(string),
		Provider: info["provider"].(string),
		Firstname: info["firstname"].(string),
		Lastname: info["lastname"].(string),
		PhotoURL: info["photo_url"].(string),
		Scopes: scopes,
		Group: info["group"].(string),
	}}, err
	// sess := &sessions.Session{
	// 	SessionInfo: SessionInfo{
	// 		ID: info["id"],
	// 	},
	// }
}
