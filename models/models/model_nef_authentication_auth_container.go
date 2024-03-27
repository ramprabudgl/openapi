/*
 * Nnef_Authentication
 *
 * NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.256 V17.3.0; 5G System;Uncrewed Aerial Systems Network Function (UAS-NF); Aerial Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.256/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Authentication/Authorization data
type NefAuthenticationAuthContainer struct {
	// string with format 'bytes' as defined in OpenAPI
	AuthMsgType    string                      `json:"authMsgType,omitempty" yaml:"authMsgType" bson:"authMsgType,omitempty"`
	AuthMsgPayload *RefToBinaryData            `json:"authMsgPayload,omitempty" yaml:"authMsgPayload" bson:"authMsgPayload,omitempty"`
	AuthResult     NefAuthenticationAuthResult `json:"authResult,omitempty" yaml:"authResult" bson:"authResult,omitempty"`
}