package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/ogen-go/ogen/middleware"
	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/ent/ogent"
	"github.com/tiagoposse/connect/ent/user"
	"github.com/tiagoposse/connect/internal/sessions"

	log "github.com/sirupsen/logrus"
	authz "github.com/tiagoposse/go-auth/sessions"
)

func buildAuthFromUser(e *ent.User) sessions.SessionInfo {

	return sessions.SessionInfo{
		ID: e.ID,
		PhotoURL: e.PhotoURL,
		Provider: e.Provider,
		Email: e.Email,
		Firstname: e.Firstname,
		Lastname: e.Lastname,
		Group: e.Edges.Group.ID,
		Scopes: e.Edges.Group.Scopes,
	}
}

func (c *Controller) Status(ctx context.Context) (ogent.StatusRes, error) {
	ctxSess := ctx.Value(authz.ContextSessionKey{})
	sess := ctxSess.(*authz.Session)

	bs, err := json.Marshal(sess.SessionInfo)
	if err != nil {
		return &ogent.StatusBadRequest{}, errors.New("could not marshal session info")
	}
	var resp ogent.StatusOKApplicationJSON = bs

	return &resp, nil
}

func (c *Controller) Logout(ctx context.Context) (ogent.LogoutRes, error) {
	return &ogent.LogoutOK{
		SetCookie: "Authorization=; Same-Site=Lax; HttpOnly; Secure; Path=/",
	}, nil
}

func (c *Controller) GoogleAuthStart(ctx context.Context) (ogent.GoogleAuthStartRes, error) {
	authURL, err := c.google.GetAuthUrl(ctx.Value(AfterUrlContextKey{}).(string))
	if err != nil {
		return &ogent.GoogleAuthStartBadRequest{}, err
	}

	return &ogent.GoogleAuthStartMovedPermanently{Location: *authURL}, nil
}

func (c *Controller) GoogleAuthSync(ctx context.Context) (ogent.GoogleAuthSyncRes, error) {
	if err := c.SyncUsers(ctx, c.google); err != nil {
		return &ogent.GoogleAuthSyncInternalServerError{}, err
	}
	return &ogent.GoogleAuthSyncInternalServerError{}, nil
}


func (c *Controller) UserpassLogin(ctx context.Context, req ogent.OptUserpassLoginReq) (ogent.UserpassLoginRes, error) {
	e, err := c.client.User.Query().Where(user.EmailEQ(req.Value.Username)).First(ctx)
	if err != nil {
		return &ogent.UserpassLoginUnauthorized{}, err
	}
	
	info := buildAuthFromUser(e)

	token, err := c.auth.CreateSessionToken(ctx, info)
	if err != nil {
		return &ogent.UserpassLoginBadRequest{}, err
	}
	auditCtx := context.WithValue(ctx, authz.ContextSessionKey{}, &authz.Session{
		SessionInfo: info,
	})

	c.audit.AuditAction(auditCtx, "userpassLogin")
	return &ogent.UserpassLoginOK{
		SetCookie: fmt.Sprintf("Authorization=%s; Same-Site=Lax; HttpOnly; Secure; Path=/", token),
	}, nil
}

func (c *Controller) GoogleAuthCallback(ctx context.Context, req ogent.OptGoogleAuthCallbackReq) (ogent.GoogleAuthCallbackRes, error) {
	email, err := c.google.ParseSamlResponse(req.Value.SAMLResponse)
	if err != nil {
		log.Debugf("parsing saml response: %s", err.Error())

		return &ogent.GoogleAuthCallbackBadRequest{}, nil
	}

	e, err := c.client.User.Query().Where(user.EmailEQ(email)).WithGroup().Only(ctx)
	if err != nil {
		return &ogent.GoogleAuthCallbackUnauthorized{}, err
	}
	
	info := buildAuthFromUser(e)
	
	token, err := c.auth.CreateSessionToken(ctx, info)
	if err != nil {
		return &ogent.GoogleAuthCallbackInternalServerError{}, err
	}

	auditCtx := context.WithValue(ctx, authz.ContextSessionKey{}, &authz.Session{
		SessionInfo: info,
	})
	c.audit.AuditAction(auditCtx, "ssoLoginGoogle")

	location, _ := url.Parse(req.Value.RelayState)
	return &ogent.GoogleAuthCallbackMovedPermanently{
		Location:  *location,
		SetCookie: fmt.Sprintf("Authorization=%s; Same-Site=Lax; HttpOnly; Secure; Path=/", token),
	}, nil
}

type AfterUrlContextKey struct {}

func GetAuthAfterUrl(req middleware.Request, next middleware.Next) (middleware.Response, error) {
	if req.OperationID == "googleAuthStart" {
		req.Context = context.WithValue(req.Context, AfterUrlContextKey{}, req.Raw.Header.Get("Referer"))
	}

	return next(req)
}
