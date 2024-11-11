/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PolicyAssociationRequest struct {
	NotificationUri string `json:"notificationUri" yaml:"notificationUri" bson:"notificationUri" mapstructure:"NotificationUri"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty" yaml:"altNotifIpv4Addrs" bson:"altNotifIpv4Addrs" mapstructure:"AltNotifIpv4Addrs"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []string                `json:"altNotifIpv6Addrs,omitempty" yaml:"altNotifIpv6Addrs" bson:"altNotifIpv6Addrs" mapstructure:"AltNotifIpv6Addrs"`
	AltNotifFqdns     []string                `json:"altNotifFqdns,omitempty" yaml:"altNotifFqdns" bson:"altNotifFqdns" mapstructure:"altNotifFqdns"`
	Supi              string                  `json:"supi" yaml:"supi" bson:"supi" mapstructure:"Supi"`
	Gpsi              string                  `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi"`
	AccessType        AccessType              `json:"accessType,omitempty" yaml:"accessType" bson:"accessType" mapstructure:"AccessType"`
	Pei               string                  `json:"pei,omitempty" yaml:"pei" bson:"pei" mapstructure:"Pei"`
	UserLoc           *UserLocation           `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc" mapstructure:"UserLoc"`
	TimeZone          string                  `json:"timeZone,omitempty" yaml:"timeZone" bson:"timeZone" mapstructure:"TimeZone"`
	ServingPlmn       *NetworkId              `json:"servingPlmn,omitempty" yaml:"servingPlmn" bson:"servingPlmn" mapstructure:"ServingPlmn"`
	RatType           RatType                 `json:"ratType,omitempty" yaml:"ratType" bson:"ratType" mapstructure:"RatType"`
	RatTypes          []RatType                 `json:"ratTypes,omitempty" yaml:"ratTypes" bson:"ratTypes" mapstructure:"RatTypes"`
	GroupIds          []string                `json:"groupIds,omitempty" yaml:"groupIds" bson:"groupIds" mapstructure:"GroupIds"`
	ServAreaRes       *ServiceAreaRestriction `json:"servAreaRes,omitempty" yaml:"servAreaRes" bson:"servAreaRes" mapstructure:"ServAreaRes"`
	Rfsp              int32                   `json:"rfsp,omitempty" yaml:"rfsp" bson:"rfsp" mapstructure:"Rfsp"`
	Guami             *Guami                  `json:"guami,omitempty" yaml:"guami" bson:"guami" mapstructure:"Guami"`
	// If the NF service consumer is an AMF, it should provide the name of a service produced by the AMF that makes use of information received within the Npcf_AMPolicyControl_UpdateNotify service operation.
	ServiveName string     `json:"serviveName,omitempty" yaml:"serviveName" bson:"serviveName" mapstructure:"ServiveName"`
	TraceReq    *TraceData `json:"traceReq,omitempty" yaml:"traceReq" bson:"traceReq" mapstructure:"TraceReq"`
	SuppFeat    string     `json:"suppFeat" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat"`
	AllowedNssai      []string   `json:"allowedNssai,omitempty"`
	//AllowedSnssais    []string   `json:"allowedSnssais,omitempty" yaml:"allowedSnssais" bson:"allowedSnssais" mapstructure:"AllowedSnssais"`
	UeAmbr            *UeAmbr    `json:"ueAmbr,omitempty" yaml:"ueAmbr" bson:"ueAmbr" mapstructure:"UeAmbr"`
	N3gAllowedSnssais []string   `json:"n3gAllowedSnssais,omitempty" yaml:"n3gAllowedSnssais" bson:"n3gAllowedSnssais" mapstructure:"N3gAllowedSnssais"`
	//MappingSnssais    []string   `json:"mappingSnssais,omitempty" yaml:"mappingSnssais" bson:"mappingSnssais" mapstructure:"MappingSnssais"`
        MappingSnssais []MappingOfSnssai `json:"mappingSnssais,omitempty" yaml:"mappingSnssais" bson:"mappingSnssais" mapstructure:"mappingSnssais"` //RQIMPL-A
	ServiceName    ServiceName       `json:"serviceName,omitempty" yaml:"serviceName" bson:"serviceName" mapstructure:"serviceName"`             //RQIMPL-A
        AllowedSnssais []Snssai          `json:"allowedSnssais,omitempty" yaml:"allowedSnssais" bson:"allowedSnssais" mapstructure:"allowedSnssais"` //RQIMPL-A
	
}
type UeAmbr struct {
	Uplink   int64 `json:"uplink,omitempty"`   // Uplink maximum bit rate in bps
	Downlink int64 `json:"downlink,omitempty"` // Downlink maximum bit rate in bps
}
