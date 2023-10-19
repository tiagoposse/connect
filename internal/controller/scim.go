package controller

import (
	"context"
	"fmt"

	"github.com/tiagoposse/connect/ent/user"
	"github.com/tiagoposse/connect/internal/auth"
	"github.com/tiagoposse/connect/internal/types"

	log "github.com/sirupsen/logrus"
)

func (c *Controller) SyncUsers(ctx context.Context, idp auth.IdentityProvider) error {
	dbUsers, err := c.client.User.Query().Where(user.ProviderEQ(idp.GetName())).All(ctx)
	if err != nil {
		return fmt.Errorf("fetching db users: %w", err)
	}

	users, idpMemberships, err := idp.GetUsersAndMemberships()
	if err != nil {
		return fmt.Errorf("fetching db users: %w", err)
	}

	for _, u := range users {
		if ignore, err := idp.GetIgnores().CheckIgnoreUser(u.Email, idpMemberships[u.ID]); err != nil {
			return err
		} else if ignore {
			log.Debugln("ignoring " + u.Email)
			continue
		}

		userDisabledReason := ""
		if u.Disabled {
			userDisabledReason = "Disabled in provider"
		}

		userGroup := "default"
		for _, group := range idpMemberships[u.ID] {
			if val, ok := idp.GetGroupMappings()[group]; ok && val != "default" && group == "default" {
				userGroup = val
			}
		}

		if err := c.client.User.Create().
			SetID(u.ID).
			SetEmail(u.Email).
			SetFirstname(u.Firstname).
			SetLastname(u.Lastname).
			SetProvider(idp.GetName()).
			SetPhotoURL(u.PhotoURL).
			SetDisabled(u.Disabled).
			SetGroupID(userGroup).
			SetDisabledReason(userDisabledReason).
			OnConflictColumns(user.FieldID).UpdateNewValues().Exec(ctx); err != nil {
			return err
		}
	}

	for _, dbUser := range dbUsers {
		found := true
		for _, guser := range users {
			if dbUser.ID == guser.ID {
				found = true
				break
			}
		}

		if !found {
			if err := c.client.User.UpdateOne(dbUser).
				SetDisabled(true).
				SetDisabledReason(types.ReasonProviderDeleted.String()).Exec(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
