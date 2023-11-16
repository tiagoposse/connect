package controller

import (
	"context"
	"slices"
	"time"

	"github.com/ogen-go/ogen/middleware"
	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/internal/sessions"

	authz "github.com/tiagoposse/go-auth/sessions"
)

type AuditController struct {
	client *ent.Client
	skipOperations []string
}

func NewAuditController(client *ent.Client, skipOperations []string) *AuditController {
	return &AuditController{
		client: client,
		skipOperations: skipOperations,
	}
}

func (ac *AuditController) OgentMiddleware(req middleware.Request, next middleware.Next) (resp middleware.Response, err error) {
	resp, err = next(req)

	if err != nil {
		return resp, err
	}

	if slices.Index(ac.skipOperations, req.OperationID) == -1 {
		err = ac.AuditAction(req.Context, req.OperationID)
	}

	return resp, err
}

func (c *AuditController) AuditAction(ctx context.Context, op string) error {
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
