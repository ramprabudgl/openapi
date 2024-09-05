/*
 * 3gpp-pfd-management
 *
 * API for PFD management. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

import (
	"github.com/ramprabudgl/openapi/models"
)

type LocationArea struct {
	// Indicates a list of Cell Global Identities of the user which identifies the cell the UE is registered.
	CellIds []string `json:"cellIds,omitempty" bson:"cellIds"`

	// Indicates a list of eNodeB identities in which the UE is currently located.
	EnodeBIds []string `json:"enodeBIds,omitempty" bson:"enodeBIds"`

	// Identifies a list of Routing Area Identities of the user where the UE is located.
	RoutingAreaIds []string `json:"routingAreaIds,omitempty" bson:"routingAreaIds"`

	// Identifies a list of Tracking Area Identities of the user where the UE is located.
	TrackingAreaIds []string `json:"trackingAreaIds,omitempty" bson:"trackingAreaIds"`

	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []models.GeographicArea `json:"geographicAreas,omitempty" bson:"geographicAreas"`

	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []models.CivicAddress `json:"civicAddresses,omitempty" bson:"civicAddresses"`
}
