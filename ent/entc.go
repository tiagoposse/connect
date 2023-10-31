//go:build ignore

package main

import (
	"log"

	"github.com/tiagoposse/connect/filter"
	"github.com/tiagoposse/connect/internal/types"
	"github.com/tiagoposse/connect/templates"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"

	exclusion "github.com/tiagoposse/entoas-fields-exclusion"
	authz "github.com/tiagoposse/go-auth/controller"
	ogauth "github.com/tiagoposse/ogent-auth"
)

func main() {
	spec := new(ogen.Spec)

	auth, err := ogauth.NewAuthExtension(
		ogauth.WithApiKeySecurity(),
		ogauth.WithCookieSecurity(),
	)
	if err != nil {
		log.Fatalf("creating ogent-auth extension: %v", err)
	}
	// filter.Mutator(
	// 	filter.WithPage(filter.Name("x-page"), filter.In("header")),
	// 	filter.WithItemsPerPage(filter.Name("x-items-per-page"), filter.In("header")),
	// 	filter.WithReturnTotal(filter.Name("x-total"), filter.In("header")),
	// 	filter.WithSort(),
	// ),)
	listExt := filter.NewOperationExtension(
		filter.WithPage(filter.Name("x-page"), filter.In("header")),
		filter.WithItemsPerPage(filter.Name("x-items-per-page"), filter.In("header")),
		filter.WithReturnTotal(filter.Name("x-total"), filter.In("header")),
		filter.WithSort(),
	)

	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(
			auth.SecurityMutation(),
			addCustomPaths,
			exclusion.Mutator,
			listExt.Mutator,
		),
	)

	ogent, err := ogent.NewExtension(spec, ogent.Templates(templates.NewTemplates("../templates")))
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureNamedEdges,
			gen.FeatureUpsert,
			gen.FeatureIntercept,
		},
	}, entc.Extensions(auth, ogent, oas, listExt))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func addCustomPaths(graph *gen.Graph, spec *ogen.Spec) error {
	secReqs := []ogen.SecurityRequirement{map[string][]string{"ApiKeyAuth": {}}, map[string][]string{"CookieAuth": {}}}
	authHeader := map[string]*ogen.Parameter{
		"Set-Cookie": {Schema: ogen.String(), Required: true},

	}

	locationHeader := map[string]*ogen.Parameter{
		"Location": {Schema: &ogen.Schema{Type: "string", Format: "uri"}, Required: true},
	}

	samlAuthHeader := map[string]*ogen.Parameter{
		"Set-Cookie": {Schema: ogen.String(), Required: true},
		"Location":   {Schema: &ogen.Schema{Type: "string", Format: "uri"}, Required: true},
	}

	spec.Paths["/status"] = new(ogen.PathItem).
		SetDescription("Check auth status").
		SetGet(ogen.NewOperation().
			SetTags([]string{"Auth"}).
			SetOperationID("status").
			SetSummary("Check authentication status").
			AddResponse("200", ogen.NewResponse().SetDescription("User is valid").AddContent("application/json", &ogen.Schema{
				Properties: []ogen.Property{
					{ Name: "id", Schema: ogen.String() },
					{ Name: "photo_url", Schema: ogen.String() },
					{ Name: "provider", Schema: ogen.String() },
					{ Name: "email", Schema: ogen.String() },
					{ Name: "lastname", Schema: ogen.String() },
					{ Name: "firstname", Schema: ogen.String() },
					{ Name: "group", Schema: ogen.String() },
					{ Name: "scopes", Schema: ogen.String().AsEnum(nil, types.AllScopes.ToRaw()...).AsArray() },
				},
				Required: []string{"id", "photo_url", "provider", "email", "lastname", "firstname", "group", "scopes" },
			})).
			AddResponse("401", ogen.NewResponse().SetDescription("User is unauthorized")).
			AddResponse("400", ogen.NewResponse().SetDescription("Bad request")),
		)
	spec.Paths["/status"].Get.Security = secReqs

	spec.Paths["/logout"] = new(ogen.PathItem).
		SetDescription("Logout").
		SetGet(ogen.NewOperation().
			SetTags([]string{"Auth"}).
			SetOperationID("logout").
			SetSummary("logout").
			AddResponse("200", ogen.NewResponse().SetHeaders(authHeader).SetDescription("Logout successful")).
			AddResponse("401", ogen.NewResponse().SetDescription("User is unauthorized")),
		)
	spec.Paths["/logout"].Get.Security = secReqs

	spec.Paths["/auth/userpass/login"] = new(ogen.PathItem).
		SetDescription("Login using username and password combination").
		SetPost(ogen.NewOperation().
			SetTags([]string{"Auth"}).
			SetOperationID("userpassLogin").
			SetSummary("Login with a user and password").
			SetRequestBody(&ogen.RequestBody{
				Description: "Username and password to login",
				Content: map[string]ogen.Media{
					"application/json": {
						Schema: &ogen.Schema{
							Type: "object",
							Properties: ogen.Properties{
								ogen.Property{Name: "username", Schema: ogen.String()},
								ogen.Property{Name: "password", Schema: ogen.String()},
							},
							Required: []string{"username", "password"},
						},
					},
				},
			}).
			AddResponse("200", ogen.NewResponse().SetHeaders(authHeader).SetDescription("Authentication Successful")).
			AddResponse("401", ogen.NewResponse().SetDescription("Username/Password combination incorrect")).
			AddResponse("400", ogen.NewResponse().SetDescription("Bad request")),
		)

	spec.Paths["/auth/google/start"] = new(ogen.PathItem).
		SetDescription("Start SAML login via Google provider").
		SetGet(ogen.NewOperation().
			SetOperationID("googleAuthStart").
			SetTags([]string{"Auth"}).
			AddResponse("301", ogen.NewResponse().SetHeaders(locationHeader).SetDescription("Starting Authentication")).
			AddResponse("400", ogen.NewResponse().SetDescription("Bad request")),
		)

	spec.Paths["/auth/google/callback"] = new(ogen.PathItem).
		SetDescription("Callback endpoint for google SAML login process").
		SetPost(ogen.NewOperation().
			SetTags([]string{"Auth"}).
			SetOperationID("googleAuthCallback").
			SetRequestBody(&ogen.RequestBody{
				Description: "Username and password to login",
				Content: map[string]ogen.Media{
					"application/x-www-form-urlencoded": {
						Schema: &ogen.Schema{
							Type: "object",
							Properties: ogen.Properties{
								ogen.Property{Name: "SAMLResponse", Schema: ogen.String()},
								ogen.Property{Name: "RelayState", Schema: ogen.String()},
							},
							Required: []string{"SAMLResponse", "RelayState"},
						},
					},
				},
			}).
			AddResponse("301", ogen.NewResponse().SetHeaders(samlAuthHeader).SetDescription("Authentication Successful")).
			AddResponse("401", ogen.NewResponse().SetDescription("Login unsuccessful")).
			AddResponse("400", ogen.NewResponse().SetDescription("Bad request")).
			AddResponse("500", ogen.NewResponse().SetDescription("Internal Server Error")),
		)

	spec.Paths["/auth/google/sync"] = new(ogen.PathItem).
		SetDescription("Synchronize users with google workspace").
		SetGet(ogen.NewOperation().
			SetTags([]string{"Auth"}).
			SetOperationID("googleAuthSync").
			SetSummary("Synchronize users for the google provider").
			AddResponse("200", ogen.NewResponse().SetDescription("Authentication Successful")).
			AddResponse("500", ogen.NewResponse().SetDescription("Internal Server Error")),
		)
	spec.Paths["/auth/google/sync"].Get.Security = secReqs

	scopes := make(map[string]authz.Scopes)
	scopes["googleAuthSync"] = authz.Scopes{types.AdminAll, types.AdminUsersWrite}
	ann := ogauth.AuthGraphAnnotation{
		Scopes: scopes,
	}
	if _, ok := graph.Annotations[ogauth.AuthAnnotationKey]; !ok {
		graph.Annotations[ogauth.AuthAnnotationKey] = ann
	} else {
		graph.Annotations[ogauth.AuthAnnotationKey].(ogauth.AuthGraphAnnotation).Merge(ann)
	}

	return nil
}
