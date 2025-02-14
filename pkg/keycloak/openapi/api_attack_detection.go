/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AttackDetectionAPIService AttackDetectionAPI service
type AttackDetectionAPIService service

type ApiAdminRealmsRealmAttackDetectionBruteForceUsersDeleteRequest struct {
	ctx context.Context
	ApiService *AttackDetectionAPIService
	realm string
}

func (r ApiAdminRealmsRealmAttackDetectionBruteForceUsersDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AdminRealmsRealmAttackDetectionBruteForceUsersDeleteExecute(r)
}

/*
AdminRealmsRealmAttackDetectionBruteForceUsersDelete Clear any user login failures for all users This can release temporary disabled users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @return ApiAdminRealmsRealmAttackDetectionBruteForceUsersDeleteRequest
*/
func (a *AttackDetectionAPIService) AdminRealmsRealmAttackDetectionBruteForceUsersDelete(ctx context.Context, realm string) ApiAdminRealmsRealmAttackDetectionBruteForceUsersDeleteRequest {
	return ApiAdminRealmsRealmAttackDetectionBruteForceUsersDeleteRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
	}
}

// Execute executes the request
func (a *AttackDetectionAPIService) AdminRealmsRealmAttackDetectionBruteForceUsersDeleteExecute(r ApiAdminRealmsRealmAttackDetectionBruteForceUsersDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttackDetectionAPIService.AdminRealmsRealmAttackDetectionBruteForceUsersDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/attack-detection/brute-force/users"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteRequest struct {
	ctx context.Context
	ApiService *AttackDetectionAPIService
	realm string
	userId string
}

func (r ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteExecute(r)
}

/*
AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete Clear any user login failures for the user This can release temporary disabled user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param userId
 @return ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteRequest
*/
func (a *AttackDetectionAPIService) AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete(ctx context.Context, realm string, userId string) ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteRequest {
	return ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		userId: userId,
	}
}

// Execute executes the request
func (a *AttackDetectionAPIService) AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteExecute(r ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttackDetectionAPIService.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/attack-detection/brute-force/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetRequest struct {
	ctx context.Context
	ApiService *AttackDetectionAPIService
	realm string
	userId string
}

func (r ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetExecute(r)
}

/*
AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet Get status of a username in brute force detection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param userId
 @return ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetRequest
*/
func (a *AttackDetectionAPIService) AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet(ctx context.Context, realm string, userId string) ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetRequest {
	return ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		userId: userId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AttackDetectionAPIService) AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetExecute(r ApiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttackDetectionAPIService.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/attack-detection/brute-force/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
