/*
 * Nudm_SSAU
 *
 * Nudm Service Specific Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.9.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents authorization update information.
type AuthUpdateInfo struct {
	AuthorizationData *ServiceSpecificAuthorizationData `json:"authorizationData" yaml:"authorizationData" bson:"authorizationData,omitempty"`
	InvalidityInd     bool                              `json:"invalidityInd,omitempty" yaml:"invalidityInd" bson:"invalidityInd,omitempty"`
	InvalidCause      InvalidCause                      `json:"invalidCause,omitempty" yaml:"invalidCause" bson:"invalidCause,omitempty"`
}