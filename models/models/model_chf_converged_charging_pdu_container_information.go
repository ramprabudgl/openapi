/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type ChfConvergedChargingPduContainerInformation struct {
	// string with format 'date-time' as defined in OpenAPI.
	TimeofFirstUsage *time.Time `json:"timeofFirstUsage,omitempty" yaml:"timeofFirstUsage" bson:"timeofFirstUsage,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeofLastUsage    *time.Time          `json:"timeofLastUsage,omitempty" yaml:"timeofLastUsage" bson:"timeofLastUsage,omitempty"`
	QoSInformation     *QosData            `json:"qoSInformation,omitempty" yaml:"qoSInformation" bson:"qoSInformation,omitempty"`
	QoSCharacteristics *QosCharacteristics `json:"qoSCharacteristics,omitempty" yaml:"qoSCharacteristics" bson:"qoSCharacteristics,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	AfChargingIdentifier int32 `json:"afChargingIdentifier,omitempty" yaml:"afChargingIdentifier" bson:"afChargingIdentifier,omitempty"`
	// Application provided charging identifier allowing correlation of charging information.
	AfChargingIdString      string        `json:"afChargingIdString,omitempty" yaml:"afChargingIdString" bson:"afChargingIdString,omitempty"`
	UserLocationInformation *UserLocation `json:"userLocationInformation,omitempty" yaml:"userLocationInformation" bson:"userLocationInformation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UetimeZone                         string                                         `json:"uetimeZone,omitempty" yaml:"uetimeZone" bson:"uetimeZone,omitempty"`
	RATType                            RatType                                        `json:"rATType,omitempty" yaml:"rATType" bson:"rATType,omitempty"`
	ServingNodeID                      []ChfConvergedChargingServingNetworkFunctionId `json:"servingNodeID,omitempty" yaml:"servingNodeID" bson:"servingNodeID,omitempty"`
	PresenceReportingAreaInformation   map[string]PresenceInfo                        `json:"presenceReportingAreaInformation,omitempty" yaml:"presenceReportingAreaInformation" bson:"presenceReportingAreaInformation,omitempty"`
	Var3gppPSDataOffStatus             Model3GpppsDataOffStatus                       `json:"3gppPSDataOffStatus,omitempty" yaml:"3gppPSDataOffStatus" bson:"3gppPSDataOffStatus,omitempty"`
	SponsorIdentity                    string                                         `json:"sponsorIdentity,omitempty" yaml:"sponsorIdentity" bson:"sponsorIdentity,omitempty"`
	ApplicationserviceProviderIdentity string                                         `json:"applicationserviceProviderIdentity,omitempty" yaml:"applicationserviceProviderIdentity" bson:"applicationserviceProviderIdentity,omitempty"`
	ChargingRuleBaseName               string                                         `json:"chargingRuleBaseName,omitempty" yaml:"chargingRuleBaseName" bson:"chargingRuleBaseName,omitempty"`
	MAPDUSteeringFunctionality         SteeringFunctionality                          `json:"mAPDUSteeringFunctionality,omitempty" yaml:"mAPDUSteeringFunctionality" bson:"mAPDUSteeringFunctionality,omitempty"`
	MAPDUSteeringMode                  *SteeringMode                                  `json:"mAPDUSteeringMode,omitempty" yaml:"mAPDUSteeringMode" bson:"mAPDUSteeringMode,omitempty"`
	TrafficForwardingWay               TrafficForwardingWay                           `json:"trafficForwardingWay,omitempty" yaml:"trafficForwardingWay" bson:"trafficForwardingWay,omitempty"`
	QosMonitoringReport                []ChfConvergedChargingQosMonitoringReport      `json:"qosMonitoringReport,omitempty" yaml:"qosMonitoringReport" bson:"qosMonitoringReport,omitempty"`
}