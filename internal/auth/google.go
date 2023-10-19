package auth

import (
	"context"
	"fmt"

	"github.com/tiagoposse/connect/ent"
	"github.com/tiagoposse/connect/internal/config"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

type GoogleAuthController struct {
	client   *admin.Service
	domain   string
	mappings map[string]string
	ignore   *IgnoringCheck
	*SamlAuthController
}

func NewGoogleAuthController(externalUrl string, conf *config.GoogleAuth) (*GoogleAuthController, error) {
	ctx := context.Background()

	// Configure the JWT config
	cfg, err := google.JWTConfigFromJSON(
		[]byte(*conf.ServiceAccountKey.Value),
		admin.AdminDirectoryUserReadonlyScope,
		admin.AdminDirectoryGroupMemberReadonlyScope,
		admin.AdminDirectoryGroupReadonlyScope,
	)

	cfg.Subject = *conf.UserToImpersonate
	if err != nil {
		return nil, fmt.Errorf("creating google config: %w", err)
	}

	adminService, err := admin.NewService(ctx, option.WithHTTPClient(cfg.Client(ctx)))
	if err != nil {
		return nil, fmt.Errorf("creating google admin client: %w", err)
	}

	serviceProvider, err := NewSamlServiceProvider(fmt.Sprintf("%s/api/v1/auth/google", externalUrl), conf.SamlAuth)
	if err != nil {
		return nil, fmt.Errorf("creating service provider: %w", err)
	}

	return &GoogleAuthController{
		client:   adminService,
		domain:   conf.Domain,
		mappings: conf.GroupMappings,
		ignore: &IgnoringCheck{
			ignoreUsers:  conf.IgnoreUsers,
			ignoreGroups: conf.IgnoreGroups,
			userFilter:   conf.UserFilters,
			groupFilter:  conf.GroupFilters,
		},
		SamlAuthController: serviceProvider,
	}, nil
}

func (gac *GoogleAuthController) GetUsersAndMemberships() ([]*ent.User, map[string][]string, error) {
	res, err := gac.client.Users.List().Domain(gac.domain).Do()
	if err != nil {
		return nil, nil, fmt.Errorf("getting users: %w", err)
	}

	gusers := make([]*ent.User, 0)
	groupMemberships := make(map[string][]string)

	for _, user := range res.Users {
		res, err := gac.client.Groups.List().Domain(gac.domain).UserKey(user.Id).Do()
		if err != nil {
			return nil, nil, fmt.Errorf("getting groups for user %s: %w", user.PrimaryEmail, err)
		}

		group := ""
		if val, ok := gac.mappings[user.Id]; ok {
			group = val
		}
		groupMemberships[user.Id] = make([]string, 0)
		for _, g := range res.Groups {
			groupMemberships[user.Id] = append(groupMemberships[user.Id], g.Email)
		}

		if group == "" {
			group = "default"
		}

		gusers = append(gusers, &ent.User{
			ID:        user.Id,
			Email:     user.PrimaryEmail,
			Provider:  "google",
			Disabled:  user.Suspended,
			Firstname: user.Name.GivenName,
			Lastname:  user.Name.FamilyName,
			PhotoURL:  user.ThumbnailPhotoUrl,
		})
	}

	return gusers, groupMemberships, nil
}

func (gac *GoogleAuthController) GetGroupMappings() map[string]string {
	return gac.mappings
}

func (gac *GoogleAuthController) GetIgnores() *IgnoringCheck {
	return gac.ignore
}

func (gac *GoogleAuthController) GetName() string {
	return "google"
}
