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
// Profile struct for Profile
type Profile struct {
	Custom bool `json:"custom,omitempty"`
	Name string `json:"name,omitempty"`
	ProfileId string `json:"profile_id,omitempty"`
	SortOrder int64 `json:"sort_order,omitempty"`
}