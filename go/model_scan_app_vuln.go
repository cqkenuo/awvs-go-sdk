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
// ScanAppVuln struct for ScanAppVuln
type ScanAppVuln struct {
	Time string `json:"time,omitempty"`
	Name string `json:"name,omitempty"`
	VulnId string `json:"vuln_id,omitempty"`
	Serverity int64 `json:"serverity,omitempty"`
	TargetInfo ScanAppVulnTargetInfo `json:"target_info,omitempty"`
}
