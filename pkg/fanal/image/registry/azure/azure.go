package azure

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/preview/preview/containerregistry/runtime/containerregistry"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"golang.org/x/xerrors"

	"github.com/aquasecurity/trivy/pkg/fanal/types"
)

type Registry struct {
	domain string
	scope  string
	cloud  cloud.Configuration
}

const (
	azureURL      = ".azurecr.io"
	chinaAzureURL = ".azurecr.cn"
	scope         = "https://management.azure.com/.default"
	chinaScope    = "https://management.chinacloudapi.cn/.default"
	scheme        = "https"
)

func (r *Registry) CheckOptions(domain string, _ types.RegistryOptions) error {
	if strings.HasSuffix(domain, azureURL) {
		r.domain = domain
		r.scope = scope
		r.cloud = cloud.AzurePublic
		return nil
	} else if strings.HasSuffix(domain, chinaAzureURL) {
		r.domain = domain
		r.scope = chinaScope
		r.cloud = cloud.AzureChina
		return nil
	}

	return xerrors.Errorf("Azure registry: %w", types.InvalidURLPattern)
}

func (r *Registry) GetCredential(ctx context.Context) (string, string, error) {
	opts := azcore.ClientOptions{Cloud: r.cloud}
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{ClientOptions: opts})
	if err != nil {
		return "", "", xerrors.Errorf("unable to generate acr credential error: %w", err)
	}
	aadToken, err := cred.GetToken(ctx, policy.TokenRequestOptions{Scopes: []string{r.scope}})
	if err != nil {
		return "", "", xerrors.Errorf("unable to get an access token: %w", err)
	}
	rt, err := refreshToken(ctx, aadToken.Token, r.domain)
	if err != nil {
		return "", "", xerrors.Errorf("unable to refresh token: %w", err)
	}
	return "00000000-0000-0000-0000-000000000000", *rt.RefreshToken, err
}

func refreshToken(ctx context.Context, accessToken, domain string) (containerregistry.RefreshToken, error) {
	tenantID := os.Getenv("AZURE_TENANT_ID")
	if tenantID == "" {
		return containerregistry.RefreshToken{}, errors.New("missing environment variable AZURE_TENANT_ID")
	}
	repoClient := containerregistry.NewRefreshTokensClient(fmt.Sprintf("%s://%s", scheme, domain))
	return repoClient.GetFromExchange(ctx, "access_token", domain, tenantID, "", accessToken)
}
