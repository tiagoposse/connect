package controller

import (
	"context"
	"time"

	"github.com/tiagoposse/connect/internal/sessions"
	authz "github.com/tiagoposse/go-auth/sessions"
)

func (c *Controller) AuditAction(ctx context.Context, op string) error {
	ctxSess := ctx.Value(authz.ContextSessionKey{})
	sess := ctxSess.(*authz.Session)

	user := sess.SessionInfo.(sessions.SessionInfo).ID
	_, err := c.client.Audit.Create().
		SetAction(op).
		SetAuthor(user).
		SetTimestamp(time.Now()).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}
