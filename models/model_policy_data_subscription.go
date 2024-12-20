/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies a subscription to policy data change notification.
type PolicyDataSubscription struct {
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationUri       string   `json:"notificationUri" bson:"notificationUri"`
	NotifId              string              `json:"notifId" bson:"notifId"`
	Expiry              string           `json:"expiry" bson:"expiry"` 
	MonitoredResourceUris []string `json:"monitoredResourceUris" bson:"monitoredResourceUris"`
	MonResItems          []string      `json:"monResItems" bson:"monResItems"`         // Array of ResourceItems
	SupportedFeatures     string   `json:"supportedFeatures,omitempty" bson:"supportedFeatures"`
	
}
