/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.537 V17.3.0; 5G System; Multicast/Broadcast Policy Control Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.537/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the parameters to request the modification of an existing MBS Policy Association.
type MbsPolicyCtxtDataUpdate struct {
	MbsServInfo    *MbsServiceInfo `json:"mbsServInfo,omitempty" yaml:"mbsServInfo" bson:"mbsServInfo,omitempty"`
	MbsPcrts       []MbsPcrt       `json:"mbsPcrts,omitempty" yaml:"mbsPcrts" bson:"mbsPcrts,omitempty"`
	MbsErrorReport *MbsErrorReport `json:"mbsErrorReport,omitempty" yaml:"mbsErrorReport" bson:"mbsErrorReport,omitempty"`
}