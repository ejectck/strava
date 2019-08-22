/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package strava

import (
	"time"
)

type DetailedSegment struct {
	// The unique identifier of this segment
	Id int64 `json:"id,omitempty"`
	// The name of this segment
	Name         string `json:"name,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	// The segment's distance, in meters
	Distance float32 `json:"distance,omitempty"`
	// The segment's average grade, in percents
	AverageGrade float32 `json:"average_grade,omitempty"`
	// The segments's maximum grade, in percents
	MaximumGrade float32 `json:"maximum_grade,omitempty"`
	// The segments's highest elevation, in meters
	ElevationHigh float32 `json:"elevation_high,omitempty"`
	// The segments's lowest elevation, in meters
	ElevationLow float32 `json:"elevation_low,omitempty"`
	StartLatlng  *LatLng `json:"start_latlng,omitempty"`
	EndLatlng    *LatLng `json:"end_latlng,omitempty"`
	// The category of the climb
	ClimbCategory int32 `json:"climb_category,omitempty"`
	// The segments's city.
	City string `json:"city,omitempty"`
	// The segments's state or geographical region.
	State string `json:"state,omitempty"`
	// The segment's country.
	Country string `json:"country,omitempty"`
	// Whether this segment is private.
	Private         bool                  `json:"private,omitempty"`
	AthletePrEffort *SummarySegmentEffort `json:"athlete_pr_effort,omitempty"`
	// The time at which the segment was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The time at which the segment was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The segment's total elevation gain.
	TotalElevationGain float32      `json:"total_elevation_gain,omitempty"`
	Map_               *PolylineMap `json:"map,omitempty"`
	// The total number of efforts for this segment
	EffortCount int32 `json:"effort_count,omitempty"`
	// The number of unique athletes who have an effort for this segment
	AthleteCount int32 `json:"athlete_count,omitempty"`
	// Whether this segment is considered hazardous
	Hazardous bool `json:"hazardous,omitempty"`
	// The number of stars for this segment
	StarCount int32 `json:"star_count,omitempty"`
}
