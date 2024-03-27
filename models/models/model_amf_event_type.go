/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmfEventType string

// List of AmfEventType
const (
	AmfEventType_LOCATION_REPORT                       AmfEventType = "LOCATION_REPORT"
	AmfEventType_PRESENCE_IN_AOI_REPORT                AmfEventType = "PRESENCE_IN_AOI_REPORT"
	AmfEventType_TIMEZONE_REPORT                       AmfEventType = "TIMEZONE_REPORT"
	AmfEventType_ACCESS_TYPE_REPORT                    AmfEventType = "ACCESS_TYPE_REPORT"
	AmfEventType_REGISTRATION_STATE_REPORT             AmfEventType = "REGISTRATION_STATE_REPORT"
	AmfEventType_CONNECTIVITY_STATE_REPORT             AmfEventType = "CONNECTIVITY_STATE_REPORT"
	AmfEventType_REACHABILITY_REPORT                   AmfEventType = "REACHABILITY_REPORT"
	AmfEventType_COMMUNICATION_FAILURE_REPORT          AmfEventType = "COMMUNICATION_FAILURE_REPORT"
	AmfEventType_UES_IN_AREA_REPORT                    AmfEventType = "UES_IN_AREA_REPORT"
	AmfEventType_SUBSCRIPTION_ID_CHANGE                AmfEventType = "SUBSCRIPTION_ID_CHANGE"
	AmfEventType_SUBSCRIPTION_ID_ADDITION              AmfEventType = "SUBSCRIPTION_ID_ADDITION"
	AmfEventType_LOSS_OF_CONNECTIVITY                  AmfEventType = "LOSS_OF_CONNECTIVITY"
	AmfEventType__5_GS_USER_STATE_REPORT               AmfEventType = "5GS_USER_STATE_REPORT"
	AmfEventType_AVAILABILITY_AFTER_DDN_FAILURE        AmfEventType = "AVAILABILITY_AFTER_DDN_FAILURE"
	AmfEventType_TYPE_ALLOCATION_CODE_REPORT           AmfEventType = "TYPE_ALLOCATION_CODE_REPORT"
	AmfEventType_FREQUENT_MOBILITY_REGISTRATION_REPORT AmfEventType = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	AmfEventType_SNSSAI_TA_MAPPING_REPORT              AmfEventType = "SNSSAI_TA_MAPPING_REPORT"
	AmfEventType_UE_LOCATION_TRENDS                    AmfEventType = "UE_LOCATION_TRENDS"
	AmfEventType_UE_ACCESS_BEHAVIOR_TRENDS             AmfEventType = "UE_ACCESS_BEHAVIOR_TRENDS"
	AmfEventType_UE_MM_TRANSACTION_REPORT              AmfEventType = "UE_MM_TRANSACTION_REPORT"
)