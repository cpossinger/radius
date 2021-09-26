// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/mgmt/keyvault"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/radius/pkg/azure/armauth"
	"github.com/Azure/radius/pkg/azure/azresources"
	"github.com/Azure/radius/pkg/azure/clients"
	"github.com/Azure/radius/pkg/healthcontract"
	"github.com/Azure/radius/pkg/radlogger"
	"github.com/Azure/radius/pkg/radrp/outputresource"
)

const (
	KeyVaultSecretNameKey  = "keyvaultsecretname"
	KeyVaultSecretValueKey = "keyvaultsecretvalue"
)

// NewAzureKeyVaultSecretHandler initializes a new handler for resources of kind Azure KeyVault Secret
func NewAzureKeyVaultSecretHandler(arm armauth.ArmConfig) ResourceHandler {
	return &azureKeyVaultSecretHandler{arm: arm}
}

type azureKeyVaultSecretHandler struct {
	arm armauth.ArmConfig
}

func (handler *azureKeyVaultSecretHandler) Put(ctx context.Context, options *PutOptions) (map[string]string, error) {
	logger := radlogger.GetLogger(ctx)
	properties := mergeProperties(*options.Resource, options.Existing, options.ExistingOutputResource)

	secretName := properties[KeyVaultSecretNameKey]
	secretValue := properties[KeyVaultSecretValueKey]
	keyVaultName := properties[KeyVaultNameKey]
	keyVaultSecretsResourceType := azresources.KeyVaultVaults + "/" + azresources.KeyVaultVaultsSecrets

	// UserAgent() returns a string of format: Azure-SDK-For-Go/v52.2.0 keyvault/2019-09-01 profiles/latest
	keyVaultAPIVersion := strings.Split(strings.Split(keyvault.UserAgent(), "keyvault/")[1], " ")[0]

	// KeyVault URI has the format: "https://<kv name>.vault.azure.net"
	secretFullName := keyVaultName + "/" + secretName
	template := map[string]interface{}{
		"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"parameters":     map[string]interface{}{},
		"resources": []interface{}{
			map[string]interface{}{
				"type":       keyVaultSecretsResourceType,
				"name":       secretFullName,
				"apiVersion": keyVaultAPIVersion,
				"properties": map[string]interface{}{
					"contentType": "text/plain",
					"value":       secretValue,
				},
			},
		},
	}

	dc := clients.NewDeploymentsClient(handler.arm.SubscriptionID, handler.arm.Auth)
	parameters := map[string]interface{}{}
	deploymentProperties := &resources.DeploymentProperties{
		Parameters: parameters,
		Mode:       resources.DeploymentModeIncremental,
		Template:   template,
	}
	deploymentName := "create-secret-" + keyVaultName + "-" + secretName
	resultFuture, err := dc.CreateOrUpdate(context.Background(), handler.arm.ResourceGroup, deploymentName, resources.Deployment{
		Properties: deploymentProperties,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create key vault secret: %w", err)
	}

	err = resultFuture.WaitForCompletionRef(context.Background(), dc.Client)
	if err != nil {
		return nil, fmt.Errorf("could not create key vault secret: %w", err)
	}

	_, err = resultFuture.Result(dc)
	if err != nil {
		return nil, fmt.Errorf("could not create key vault secret: %w", err)
	}

	logger.WithValues(radlogger.LogFieldLocalID, outputresource.LocalIDKeyVaultSecret).Info(fmt.Sprintf("Created secret: %s in Key Vault: %s successfully", secretName, keyVaultName))

	secretResource := azure.Resource{
		SubscriptionID: handler.arm.SubscriptionID,
		ResourceGroup:  handler.arm.ResourceGroup,
		Provider:       "Microsoft.KeyVault",
		ResourceType:   keyVaultSecretsResourceType,
		ResourceName:   secretFullName,
	}

	options.Resource.Info = outputresource.ARMInfo{
		ID:           secretResource.String(),
		ResourceType: keyVaultSecretsResourceType,
		APIVersion:   keyVaultAPIVersion,
	}

	return properties, nil
}

func (handler *azureKeyVaultSecretHandler) Delete(ctx context.Context, options DeleteOptions) error {
	// TODO: right now this resource is deleted in a different handler :(
	// this should be done here instead when we have built a more mature system.

	return nil
}

func NewAzureKeyVaultSecretHealthHandler(arm armauth.ArmConfig) HealthHandler {
	return &azureKeyVaultSecretHealthHandler{arm: arm}
}

type azureKeyVaultSecretHealthHandler struct {
	arm armauth.ArmConfig
}

func (handler *azureKeyVaultSecretHealthHandler) GetHealthOptions(ctx context.Context) healthcontract.HealthCheckOptions {
	return healthcontract.HealthCheckOptions{}
}