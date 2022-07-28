//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MongoDatabasesClient contains the methods for the MongoDatabases group.
// Don't use this type directly, use NewMongoDatabasesClient() instead.
type MongoDatabasesClient struct {
	ep string
	pl runtime.Pipeline
	rootScope string
}

// NewMongoDatabasesClient creates a new instance of MongoDatabasesClient with the specified values.
func NewMongoDatabasesClient(con *arm.Connection, rootScope string) *MongoDatabasesClient {
	return &MongoDatabasesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), rootScope: rootScope}
}

// CreateOrUpdate - Creates or updates a MongoDatabase resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoDatabasesClient) CreateOrUpdate(ctx context.Context, mongoDatabaseName string, mongoDatabaseParameters MongoDatabaseResource, options *MongoDatabasesCreateOrUpdateOptions) (MongoDatabasesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, mongoDatabaseName, mongoDatabaseParameters, options)
	if err != nil {
		return MongoDatabasesCreateOrUpdateResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return MongoDatabasesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return MongoDatabasesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MongoDatabasesClient) createOrUpdateCreateRequest(ctx context.Context, mongoDatabaseName string, mongoDatabaseParameters MongoDatabaseResource, options *MongoDatabasesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/mongoDatabases/{mongoDatabaseName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if mongoDatabaseName == "" {
		return nil, errors.New("parameter mongoDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoDatabaseName}", url.PathEscape(mongoDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, mongoDatabaseParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MongoDatabasesClient) createOrUpdateHandleResponse(resp *http.Response) (MongoDatabasesCreateOrUpdateResponse, error) {
	result := MongoDatabasesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MongoDatabaseResponseResource); err != nil {
		return MongoDatabasesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *MongoDatabasesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an existing mongoDatabase resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoDatabasesClient) Delete(ctx context.Context, mongoDatabaseName string, options *MongoDatabasesDeleteOptions) (MongoDatabasesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, mongoDatabaseName, options)
	if err != nil {
		return MongoDatabasesDeleteResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return MongoDatabasesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return MongoDatabasesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return MongoDatabasesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MongoDatabasesClient) deleteCreateRequest(ctx context.Context, mongoDatabaseName string, options *MongoDatabasesDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/mongoDatabases/{mongoDatabaseName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if mongoDatabaseName == "" {
		return nil, errors.New("parameter mongoDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoDatabaseName}", url.PathEscape(mongoDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *MongoDatabasesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves information about a mongoDatabases resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoDatabasesClient) Get(ctx context.Context, mongoDatabaseName string, options *MongoDatabasesGetOptions) (MongoDatabasesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, mongoDatabaseName, options)
	if err != nil {
		return MongoDatabasesGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return MongoDatabasesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MongoDatabasesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MongoDatabasesClient) getCreateRequest(ctx context.Context, mongoDatabaseName string, options *MongoDatabasesGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/mongoDatabases/{mongoDatabaseName}"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if mongoDatabaseName == "" {
		return nil, errors.New("parameter mongoDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoDatabaseName}", url.PathEscape(mongoDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MongoDatabasesClient) getHandleResponse(resp *http.Response) (MongoDatabasesGetResponse, error) {
	result := MongoDatabasesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MongoDatabaseResponseResource); err != nil {
		return MongoDatabasesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MongoDatabasesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByRootScope - Lists information about all mongoDatabases resources in the given root scope
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoDatabasesClient) ListByRootScope(options *MongoDatabasesListByRootScopeOptions) (*MongoDatabasesListByRootScopePager) {
	return &MongoDatabasesListByRootScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByRootScopeCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp MongoDatabasesListByRootScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MongoDatabaseList.NextLink)
		},
	}
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *MongoDatabasesClient) listByRootScopeCreateRequest(ctx context.Context, options *MongoDatabasesListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/mongoDatabases"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *MongoDatabasesClient) listByRootScopeHandleResponse(resp *http.Response) (MongoDatabasesListByRootScopeResponse, error) {
	result := MongoDatabasesListByRootScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MongoDatabaseList); err != nil {
		return MongoDatabasesListByRootScopeResponse{}, err
	}
	return result, nil
}

// listByRootScopeHandleError handles the ListByRootScope error response.
func (client *MongoDatabasesClient) listByRootScopeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListSecrets - Lists secrets values for the specified MongoDatabase resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoDatabasesClient) ListSecrets(ctx context.Context, mongoDatabaseName string, options *MongoDatabasesListSecretsOptions) (MongoDatabasesListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, mongoDatabaseName, options)
	if err != nil {
		return MongoDatabasesListSecretsResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return MongoDatabasesListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MongoDatabasesListSecretsResponse{}, client.listSecretsHandleError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *MongoDatabasesClient) listSecretsCreateRequest(ctx context.Context, mongoDatabaseName string, options *MongoDatabasesListSecretsOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Connector/mongoDatabases/{mongoDatabaseName}/listSecrets"
	if client.rootScope == "" {
		return nil, errors.New("parameter client.rootScope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if mongoDatabaseName == "" {
		return nil, errors.New("parameter mongoDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoDatabaseName}", url.PathEscape(mongoDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *MongoDatabasesClient) listSecretsHandleResponse(resp *http.Response) (MongoDatabasesListSecretsResponse, error) {
	result := MongoDatabasesListSecretsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MongoDatabaseSecrets); err != nil {
		return MongoDatabasesListSecretsResponse{}, err
	}
	return result, nil
}

// listSecretsHandleError handles the ListSecrets error response.
func (client *MongoDatabasesClient) listSecretsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

