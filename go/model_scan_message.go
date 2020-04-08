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
// ScanMessage struct for ScanMessage
type ScanMessage struct {
	Time string `json:"time,omitempty"`
	TargetInfo map[string]interface{} `json:"target_info,omitempty"`
	Kind string `json:"kind,omitempty"`
	Data string `json:"data,omitempty"`
	Level int64 `json:"level,omitempty"`
}