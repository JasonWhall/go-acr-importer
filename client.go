package main

import (
	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
)

var (
	environment     azure.Environment
	environmentName = azure.PublicCloud.Name
)

// Uses helpers to authenticate to Azure
func NewContainerRegistryClient() (*containerregistry.RegistriesClient, error) {

	builder := authentication.Builder{
		TenantID:       *tenantId,
		SubscriptionID: *subscriptionID,
		ClientID:       *clientId,
		ClientSecret:   *clientSecret,
		Environment:    environmentName,

		SupportsClientSecretAuth:       true,
		SupportsAzureCliToken:          true,
		SupportsManagedServiceIdentity: *useMsi,
	}

	environment, err := authentication.DetermineEnvironment(environmentName)
	if err != nil {
		return nil, err
	}

	client, err := builder.Build()
	if err != nil {
		return nil, err
	}

	sender := sender.BuildSender("ACR Importer")

	oauthConfig, err := client.BuildOAuthConfig(environment.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	authorizer, err := client.GetAuthorizationToken(sender, oauthConfig, environment.TokenAudience)
	if err != nil {
		return nil, err
	}

	acrClient := containerregistry.NewRegistriesClient(*subscriptionID)
	acrClient.Authorizer = authorizer

	return &acrClient, nil
}
