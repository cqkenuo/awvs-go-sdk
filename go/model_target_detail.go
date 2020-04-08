/*
 * AWVS12 client api
 *
 * Awvs12 client api [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TargetDetail struct for TargetDetail
type TargetDetail struct {
	// url
	Address               string `json:"address"`
	Description           string `json:"description,omitempty"`
	Type                  string `json:"type"`
	Criticality           int64  `json:"criticality"`
	TargetId              string `json:"target_id,omitempty"`
	ContinuousMode        bool   `json:"continuous_mode,omitempty"`
	LastScanDate          string `json:"last_scan_date,omitempty"`
	LastScanId            string `json:"last_scan_id,omitempty"`
	LastScanSessionId     string `json:"last_scan_session_id,omitempty"`
	LastScanSessionStatus string `json:"last_scan_session_status,omitempty"`
}
