/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ParameterProvisionProfileDataFor5GVNGroupDocumentApiService service

/*
ParameterProvisionProfileDataFor5GVNGroupDocumentApiService Retrieves the parameter provision profile data for 5G VN Group
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ExtGroupIds - List of external VN group identifiers
 * @param SupportedFeatures - Supported Features

@return Query5GVNGroupPPDataResponse
*/

// Query5GVNGroupPPDataRequest
type Query5GVNGroupPPDataRequest struct {
	ExtGroupIds       []string
	SupportedFeatures *string
}

func (r *Query5GVNGroupPPDataRequest) SetExtGroupIds(ExtGroupIds []string) {
	r.ExtGroupIds = ExtGroupIds
}
func (r *Query5GVNGroupPPDataRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type Query5GVNGroupPPDataResponse struct {
	Pp5gVnGroupProfileData models.Pp5gVnGroupProfileData
}

type Query5GVNGroupPPDataError struct {
}

func (a *ParameterProvisionProfileDataFor5GVNGroupDocumentApiService) Query5GVNGroupPPData(ctx context.Context, request *Query5GVNGroupPPDataRequest) (*Query5GVNGroupPPDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Query5GVNGroupPPDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/5g-vn-groups/pp-profile-data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.ExtGroupIds != nil {
		if len(request.ExtGroupIds) < 1 {
			return &localVarReturnValue, openapi.ReportError("ExtGroupIds must have at least 1 elements")
		}
		localVarQueryParams.Add("ext-group-ids", openapi.ParameterToString(request.ExtGroupIds, "csv"))
	}
	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.Pp5gVnGroupProfileData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}