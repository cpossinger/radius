//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DaprSecretStoreClient contains the methods for the DaprSecretStore group.
// Don't use this type directly, use NewDaprSecretStoreClient() instead.
type DaprSecretStoreClient struct {
	host string
	rootScope string
	pl runtime.Pipeline
}

// NewDaprSecretStoreClient creates a new instance of DaprSecretStoreClient with the specified values.
// rootScope - The scope in which the resource is present. For Azure resource this would be /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDaprSecretStoreClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DaprSecretStoreClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DaprSecretStoreClient{
		rootScope: rootScope,
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a DaprSecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprSecretStoreName - Dapr SecretStore name
// resource - Resource create parameters.
// options - DaprSecretStoreClientCreateOrUpdateOptions contains the optional parameters for the DaprSecretStoreClient.CreateOrUpdate
// method.
func (client *DaprSecretStoreClient) CreateOrUpdate(ctx context.Context, daprSecretStoreName string, resource DaprSecretStoreResource, options *DaprSecretStoreClientCreateOrUpdateOptions) (DaprSecretStoreClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, daprSecretStoreName, resource, options)
	if err != nil {
		return DaprSecretStoreClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprSecretStoreClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DaprSecretStoreClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprSecretStoreClient) createOrUpdateCreateRequest(ctx context.Context, daprSecretStoreName string, resource DaprSecretStoreResource, options *DaprSecretStoreClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Dapr/daprSecretStores/{daprSecretStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprSecretStoreName == "" {
		return nil, errors.New("parameter daprSecretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprSecretStoreName}", url.PathEscape(daprSecretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprSecretStoreClient) createOrUpdateHandleResponse(resp *http.Response) (DaprSecretStoreClientCreateOrUpdateResponse, error) {
	result := DaprSecretStoreClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return DaprSecretStoreClientCreateOrUpdateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSecretStoreResource); err != nil {
		return DaprSecretStoreClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing DaprSecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprSecretStoreName - Dapr SecretStore name
// options - DaprSecretStoreClientDeleteOptions contains the optional parameters for the DaprSecretStoreClient.Delete method.
func (client *DaprSecretStoreClient) Delete(ctx context.Context, daprSecretStoreName string, options *DaprSecretStoreClientDeleteOptions) (DaprSecretStoreClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, daprSecretStoreName, options)
	if err != nil {
		return DaprSecretStoreClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprSecretStoreClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return DaprSecretStoreClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *DaprSecretStoreClient) deleteCreateRequest(ctx context.Context, daprSecretStoreName string, options *DaprSecretStoreClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Dapr/daprSecretStores/{daprSecretStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprSecretStoreName == "" {
		return nil, errors.New("parameter daprSecretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprSecretStoreName}", url.PathEscape(daprSecretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *DaprSecretStoreClient) deleteHandleResponse(resp *http.Response) (DaprSecretStoreClientDeleteResponse, error) {
	result := DaprSecretStoreClientDeleteResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return DaprSecretStoreClientDeleteResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	return result, nil
}

// Get - Retrieves information about a DaprSecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprSecretStoreName - Dapr SecretStore name
// options - DaprSecretStoreClientGetOptions contains the optional parameters for the DaprSecretStoreClient.Get method.
func (client *DaprSecretStoreClient) Get(ctx context.Context, daprSecretStoreName string, options *DaprSecretStoreClientGetOptions) (DaprSecretStoreClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, daprSecretStoreName, options)
	if err != nil {
		return DaprSecretStoreClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprSecretStoreClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprSecretStoreClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprSecretStoreClient) getCreateRequest(ctx context.Context, daprSecretStoreName string, options *DaprSecretStoreClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Dapr/daprSecretStores/{daprSecretStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprSecretStoreName == "" {
		return nil, errors.New("parameter daprSecretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprSecretStoreName}", url.PathEscape(daprSecretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprSecretStoreClient) getHandleResponse(resp *http.Response) (DaprSecretStoreClientGetResponse, error) {
	result := DaprSecretStoreClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSecretStoreResource); err != nil {
		return DaprSecretStoreClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRootScopePager - Lists information about all DaprSecretStoreResources in the given root scope
// Generated from API version 2022-03-15-privatepreview
// options - DaprSecretStoreClientListByRootScopeOptions contains the optional parameters for the DaprSecretStoreClient.ListByRootScope
// method.
func (client *DaprSecretStoreClient) NewListByRootScopePager(options *DaprSecretStoreClientListByRootScopeOptions) (*runtime.Pager[DaprSecretStoreClientListByRootScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[DaprSecretStoreClientListByRootScopeResponse]{
		More: func(page DaprSecretStoreClientListByRootScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DaprSecretStoreClientListByRootScopeResponse) (DaprSecretStoreClientListByRootScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRootScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DaprSecretStoreClientListByRootScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DaprSecretStoreClientListByRootScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DaprSecretStoreClientListByRootScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRootScopeHandleResponse(resp)
		},
	})
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *DaprSecretStoreClient) listByRootScopeCreateRequest(ctx context.Context, options *DaprSecretStoreClientListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Dapr/daprSecretStores"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *DaprSecretStoreClient) listByRootScopeHandleResponse(resp *http.Response) (DaprSecretStoreClientListByRootScopeResponse, error) {
	result := DaprSecretStoreClientListByRootScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSecretStoreResourceListResult); err != nil {
		return DaprSecretStoreClientListByRootScopeResponse{}, err
	}
	return result, nil
}

