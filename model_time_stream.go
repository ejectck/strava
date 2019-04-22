/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package strava

type TimeStream struct {
	// The number of data points in this stream
	OriginalSize int32 `json:"original_size,omitempty"`
	// The level of detail (sampling) in which this stream was returned
	Resolution string `json:"resolution,omitempty"`
	// The base series used in the case the stream was downsampled
	SeriesType string `json:"series_type,omitempty"`
	// The sequence of time values for this stream, in seconds
	Data []int32 `json:"data,omitempty"`
}
