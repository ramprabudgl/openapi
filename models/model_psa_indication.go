/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PsaIndication string

// List of PsaIndication
const (
	PsaIndication_INSERTED      PsaIndication = "PSA_INSERTED"
	PsaIndication_REMOVED       PsaIndication = "PSA_REMOVED"
	PsaIndication_INSERTED_ONLY PsaIndication = "PSA_INSERTED_ONLY"
	PsaIndication_REMOVED_ONLY  PsaIndication = "PSA_REMOVED_ONLY"
)