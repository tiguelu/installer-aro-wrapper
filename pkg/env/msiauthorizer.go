package env

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"fmt"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

type MSIContext string

const (
	MSIContextRP      MSIContext = "RP"
	MSIContextGateway MSIContext = "GATEWAY"
)

func (c *core) NewMSIAuthorizer(msiContext MSIContext, resource string) (autorest.Authorizer, error) {
	if !c.IsLocalDevelopmentMode() {
		return auth.NewAuthorizerFromEnvironmentWithResource(resource)
	}

	for _, key := range []string{
		"ARO_AZURE_" + string(msiContext) + "_CLIENT_ID",
		"ARO_AZURE_" + string(msiContext) + "_CLIENT_SECRET",
		"ARO_AZURE_TENANT_ID",
	} {
		if _, found := os.LookupEnv(key); !found {
			return nil, fmt.Errorf("environment variable %q unset (development mode)", key)
		}
	}

	config := &auth.ClientCredentialsConfig{
		ClientID:     os.Getenv("ARO_AZURE_" + string(msiContext) + "_CLIENT_ID"),
		ClientSecret: os.Getenv("ARO_AZURE_" + string(msiContext) + "_CLIENT_SECRET"),
		TenantID:     os.Getenv("ARO_AZURE_TENANT_ID"),
		Resource:     resource,
		AADEndpoint:  c.Environment().ActiveDirectoryEndpoint,
	}

	return config.Authorizer()
}
