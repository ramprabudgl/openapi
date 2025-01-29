/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains changed policy data for which notification was requested.
type PolicyDataChangeNotification struct {
	AmPolicyData            *AmPolicyData            `json:"amPolicyData,omitempty" yaml:"amPolicyData" bson:"amPolicyData" mapstructure:"AmPolicyData"`
	UePolicySet             *UePolicySet             `json:"uePolicySet,omitempty" yaml:"uePolicySet" bson:"uePolicySet" mapstructure:"UePolicySet"`
	SmPolicyData            *SmPolicyData            `json:"smPolicyData,omitempty" yaml:"smPolicyData" bson:"smPolicyData" mapstructure:"SmPolicyData"`
	UsageMonData            *UsageMonData            `json:"usageMonData,omitempty" yaml:"usageMonData" bson:"usageMonData" mapstructure:"UsageMonData"`
	SponsorConnectivityData *SponsorConnectivityData `json:"SponsorConnectivityData,omitempty" yaml:"SponsorConnectivityData" bson:"SponsorConnectivityData" mapstructure:"SponsorConnectivityData"`
	BdtData                 *BdtData                 `json:"bdtData,omitempty" yaml:"bdtData" bson:"bdtData" mapstructure:"BdtData"`
	UeId                    string                   `json:"ueId,omitempty" yaml:"ueId" bson:"ueId" mapstructure:"UeId"`
	SponsorId               string                   `json:"sponsorId,omitempty" yaml:"sponsorId" bson:"sponsorId" mapstructure:"SponsorId"`
	BdtRefId                string                   `json:"bdtRefId,omitempty" yaml:"bdtRefId" bson:"bdtRefId" mapstructure:"BdtRefId"`
	UsageMonId              string                   `json:"usageMonId,omitempty" yaml:"usageMonId" bson:"usageMonId" mapstructure:"UsageMonId"`
	DelResources 		[]string    		 `json:"delResources"`
	DelData                 string                 `json:"delData"`
}

