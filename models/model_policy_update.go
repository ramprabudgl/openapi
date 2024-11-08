/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PolicyUpdate struct {
	ResourceUri string `json:"resourceUri" yaml:"resourceUri" bson:"resourceUri" mapstructure:"ResourceUri"`
	// Request Triggers that the PCF subscribes. Only values \"LOC_CH\" and \"PRA_CH\" are permitted.
	Triggers    []RequestTrigger        `json:"triggers,omitempty" yaml:"triggers" bson:"triggers" mapstructure:"Triggers"`
	ServAreaRes *ServiceAreaRestriction `json:"servAreaRes,omitempty" yaml:"servAreaRes" bson:"servAreaRes" mapstructure:"ServAreaRes"`
	Rfsp        int32                   `json:"rfsp,omitempty" yaml:"rfsp" bson:"rfsp" mapstructure:"Rfsp"`
	// Map of PRA information.
	Pras map[string]PresenceInfoRm `json:"pras,omitempty" yaml:"pras" bson:"pras" mapstructure:"Pras"`
        // extra parameters -Hemanth
	AltNotifFqdns     []string          `json:"altNotifFqdns,omitempty" yaml:"altNotifFqdns" bson:"altNotifFqdns" mapstructure:"AltNotifFqdns"`
	SmfSelInfo        *SmfSelectionData `json:"smfSelInfo,omitempty" yaml:"smfSelInfo" bson:"smfSelInfo" mapstructure:"SmfSelInfo"`
	UeAmbr            *Ambr             `json:"ueAmbr,omitempty" yaml:"ueAmbr" bson:"ueAmbr" mapstructure:"UeAmbr"`
	AllowedSnssais    []Snssai          `json:"allowedSnssais,omitempty" yaml:"allowedSnssais" bson:"allowedSnssais" mapstructure:"AllowedSnssais"`
	MappingSnssais    []MappingOfSnssai `json:"mappingSnssais,omitempty" yaml:"mappingSnssais" bson:"mappingSnssais" mapstructure:"MappingSnssais"`
	N3gAllowedSnssais []Snssai          `json:"n3gAllowedSnssais,omitempty" yaml:"n3gAllowedSnssais" bson:"n3gAllowedSnssais" mapstructure:"N3gAllowedSnssais"`
	AccessTypes       []AccessType      `json:"accessTypes,omitempty" yaml:"accessTypes" bson:"accessTypes" mapstructure:"AccessTypes"`
	Guami             *Guami            `json:"guami,omitempty" yaml:"guami" bson:"guami" mapstructure:"Guami"`
	UserLoc           *UserLocation     `json:"userLoc,omitempty" yaml:"userLoc" bson:"userLoc" mapstructure:"UserLoc"`
	TraceReq          *TraceData        `json:"traceReq,omitempty" yaml:"traceReq" bson:"traceReq" mapstructure:"TraceReq"`

}
