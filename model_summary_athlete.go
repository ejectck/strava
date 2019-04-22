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

type SummaryAthlete struct {
	// The unique identifier of the athlete
	Id int32 `json:"id,omitempty"`
	// Resource state, indicates level of detail. Possible values: 1 -> \"meta\", 2 -> \"summary\", 3 -> \"detail\"
	ResourceState int32 `json:"resource_state,omitempty"`
	// The athlete's first name.
	Firstname string `json:"firstname,omitempty"`
	// The athlete's last name.
	Lastname string `json:"lastname,omitempty"`
	// URL to a 62x62 pixel profile picture.
	ProfileMedium string `json:"profile_medium,omitempty"`
	// URL to a 124x124 pixel profile picture.
	Profile string `json:"profile,omitempty"`
	// The athlete's city.
	City string `json:"city,omitempty"`
	// The athlete's state or geographical region.
	State string `json:"state,omitempty"`
	// The athlete's country.
	Country string `json:"country,omitempty"`
	// The athlete's sex.
	Sex string `json:"sex,omitempty"`
	// Whether the currently logged-in athlete follows this athlete.
	Friend string `json:"friend,omitempty"`
	// Whether this athlete follows the currently logged-in athlete.
	Follower string `json:"follower,omitempty"`
	// Deprecated.  Use summit field instead. Whether the athlete has any Summit subscription.
	Premium bool `json:"premium,omitempty"`
	// Whether the athlete has any Summit subscription.
	Summit bool `json:"summit,omitempty"`
	// The time at which the athlete was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The time at which the athlete was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
