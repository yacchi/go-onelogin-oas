// Code generated by "openapi-codegen"; DO NOT EDIT.
/*
 * OneLogin API
 *
 * This is an administrative API for OneLogin customers
 *
 * API version: 1.1.0-oas3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onelogin

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/yacchi/go-onelogin-oas/models"
)

// Linger please
var (
	_ _context.Context
)

type apiAuthenticateUserRequest struct {
	ctx                        _context.Context
	client                     *APIClient
	customAllowedOriginHeader1 *string
	inlineObject6              *models.InlineObject6
}

func (r apiAuthenticateUserRequest) CustomAllowedOriginHeader1(customAllowedOriginHeader1 string) apiAuthenticateUserRequest {
	r.customAllowedOriginHeader1 = &customAllowedOriginHeader1
	return r
}

func (r apiAuthenticateUserRequest) InlineObject6(inlineObject6 models.InlineObject6) apiAuthenticateUserRequest {
	r.inlineObject6 = &inlineObject6
	return r
}

/*
AuthenticateUser Authenticate a user
Use this API to generate a session login token in scenarios in which MFA may or may not be required. Both scenarios are supported. A session login token expires two minutes after creation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiAuthenticateUserRequest
*/
func (c *APIClient) AuthenticateUser(ctx _context.Context) apiAuthenticateUserRequest {
	return apiAuthenticateUserRequest{
		client: c,
		ctx:    ctx,
	}
}

/*
Execute executes the request
 @return models.UserLoginResponse
*/
func (r apiAuthenticateUserRequest) Execute() (models.UserLoginResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.UserLoginResponse
	)

	localBasePath, err := r.client.cfg.ServerURLWithContext(r.ctx, "LoginApiService.AuthenticateUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/login/auth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.customAllowedOriginHeader1 != nil {
		localVarHeaderParams["Custom-Allowed-Origin-Header-1"] = parameterToString(*r.customAllowedOriginHeader1, "")
	}
	// body params
	localVarPostBody = r.inlineObject6
	req, err := r.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v models.UserLoginResponse
			err = r.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.Status
			err = r.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.Status
			err = r.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiVerifyLoginMFATokenRequest struct {
	ctx                        _context.Context
	client                     *APIClient
	customAllowedOriginHeader1 *string
	inlineObject7              *models.InlineObject7
}

func (r apiVerifyLoginMFATokenRequest) CustomAllowedOriginHeader1(customAllowedOriginHeader1 string) apiVerifyLoginMFATokenRequest {
	r.customAllowedOriginHeader1 = &customAllowedOriginHeader1
	return r
}

func (r apiVerifyLoginMFATokenRequest) InlineObject7(inlineObject7 models.InlineObject7) apiVerifyLoginMFATokenRequest {
	r.inlineObject7 = &inlineObject7
	return r
}

/*
VerifyLoginMFAToken Verify an MFA token
Verify a one-time password (OTP) value, provided for a second factor, when multi-factor authentication (MFA) is required.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiVerifyLoginMFATokenRequest
*/
func (c *APIClient) VerifyLoginMFAToken(ctx _context.Context) apiVerifyLoginMFATokenRequest {
	return apiVerifyLoginMFATokenRequest{
		client: c,
		ctx:    ctx,
	}
}

/*
Execute executes the request
 @return models.LoginVerifyMfaResponse
*/
func (r apiVerifyLoginMFATokenRequest) Execute() (models.LoginVerifyMfaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.LoginVerifyMfaResponse
	)

	localBasePath, err := r.client.cfg.ServerURLWithContext(r.ctx, "LoginApiService.VerifyLoginMFAToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/1/login/verify_factor"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.customAllowedOriginHeader1 != nil {
		localVarHeaderParams["Custom-Allowed-Origin-Header-1"] = parameterToString(*r.customAllowedOriginHeader1, "")
	}
	// body params
	localVarPostBody = r.inlineObject7
	req, err := r.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v models.LoginVerifyMfaResponse
			err = r.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.Status
			err = r.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.Status
			err = r.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
