/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contains conditions of applicability for a rule.
type ConditionData struct {
	// Uniquely identifies the condition data within a PDU session.
	CondId string `json:"condId" yaml:"condId" bson:"condId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.
	ActivationTime *time.Time `json:"activationTime,omitempty" yaml:"activationTime" bson:"activationTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.
	DeactivationTime *time.Time `json:"deactivationTime,omitempty" yaml:"deactivationTime" bson:"deactivationTime,omitempty"`
	AccessType       AccessType `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	RatType          RatType    `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
}