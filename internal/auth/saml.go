package auth

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/tiagoposse/connect/internal/config"

	saml2 "github.com/russellhaering/gosaml2"
	samltypes "github.com/russellhaering/gosaml2/types"
	dsig "github.com/russellhaering/goxmldsig"
)

type SamlAuthMethod interface {
	ParseSamlResponse(samlResponse string) (string, error)
	GetAuthUrl(relayState string) (string, error)
}

type SamlAuthController struct {
	sp *saml2.SAMLServiceProvider
}

func NewSamlServiceProvider(externalUrl string, config *config.SamlAuth) (*SamlAuthController, error) {
	metadata := &samltypes.EntityDescriptor{}
	err := xml.Unmarshal([]byte(*config.IdpMetadata.Value), metadata)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling idp metadata: %w", err)
	}

	certStore := dsig.MemoryX509CertificateStore{
		Roots: []*x509.Certificate{},
	}

	for _, kd := range metadata.IDPSSODescriptor.KeyDescriptors {
		for idx, xcert := range kd.KeyInfo.X509Data.X509Certificates {
			if xcert.Data == "" {
				return nil, fmt.Errorf("metadata certificate(%d) must not be empty", idx)
			}
			certData, err := base64.StdEncoding.DecodeString(xcert.Data)
			if err != nil {
				return nil, err
			}

			idpCert, err := x509.ParseCertificate(certData)
			if err != nil {
				return nil, err
			}

			certStore.Roots = append(certStore.Roots, idpCert)
		}
	}

	callbackUrl := fmt.Sprintf("%s/callback", externalUrl)

	randomKeyStore := dsig.RandomKeyStoreForTest()
	return &SamlAuthController{
		sp: &saml2.SAMLServiceProvider{
			IdentityProviderSSOURL:      metadata.IDPSSODescriptor.SingleSignOnServices[0].Location,
			IdentityProviderIssuer:      metadata.EntityID,
			ServiceProviderIssuer:       externalUrl,
			AssertionConsumerServiceURL: callbackUrl,
			SignAuthnRequests:           true,
			AudienceURI:                 externalUrl,
			IDPCertificateStore:         &certStore,
			SPKeyStore:                  randomKeyStore,
		},
	}, nil
}

func (sac *SamlAuthController) GetAuthUrl(relayState string) (*url.URL, error) {
	rawUrl, err := sac.sp.BuildAuthURL(relayState)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (sac *SamlAuthController) ParseSamlResponse(samlResponse string) (string, error) {
	assertionInfo, err := sac.sp.RetrieveAssertionInfo(samlResponse)
	if err != nil {
		return "", fmt.Errorf("retrieving assertion: %s", err.Error())
	}

	if assertionInfo.WarningInfo.InvalidTime {
		return "", fmt.Errorf("invalid assertion time: %s", err.Error())
	}

	if assertionInfo.WarningInfo.NotInAudience {
		return "", fmt.Errorf("assertion not in info: %s", err.Error())
	}

	return assertionInfo.NameID, nil
}
