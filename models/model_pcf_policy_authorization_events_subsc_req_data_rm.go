/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V17.9.0; 5G System; Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// This data type is defined in the same way as the EventsSubscReqData data type, but with the OpenAPI nullable property set to true.
type PcfPolicyAuthorizationEventsSubscReqDataRm struct {
	Events []AfEventSubscription `json:"events" yaml:"events" bson:"events,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri        string                                            `json:"notifUri,omitempty" yaml:"notifUri" bson:"notifUri,omitempty"`
	ReqQosMonParams []RequestedQosMonitoringParameter                 `json:"reqQosMonParams,omitempty" yaml:"reqQosMonParams" bson:"reqQosMonParams,omitempty"`
	QosMon          *PcfPolicyAuthorizationQosMonitoringInformationRm `json:"qosMon,omitempty" yaml:"qosMon" bson:"qosMon,omitempty"`
	ReqAnis         []RequiredAccessInfo                              `json:"reqAnis,omitempty" yaml:"reqAnis" bson:"reqAnis,omitempty"`
	UsgThres        *UsageThresholdRm                                 `json:"usgThres,omitempty" yaml:"usgThres" bson:"usgThres,omitempty"`
	NotifCorreId    string                                            `json:"notifCorreId,omitempty" yaml:"notifCorreId" bson:"notifCorreId,omitempty"`
	DirectNotifInd  bool                                              `json:"directNotifInd,omitempty" yaml:"directNotifInd" bson:"directNotifInd,omitempty"`
}