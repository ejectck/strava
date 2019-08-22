/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package strava

type SummaryClub struct {
	// The club's unique identifier.
	Id int32 `json:"id,omitempty"`
	// Resource state, indicates level of detail. Possible values: 1 -> \"meta\", 2 -> \"summary\", 3 -> \"detail\"
	ResourceState int32 `json:"resource_state,omitempty"`
	// The club's name.
	Name string `json:"name,omitempty"`
	// URL to a 60x60 pixel profile picture.
	ProfileMedium string `json:"profile_medium,omitempty"`
	// URL to a ~1185x580 pixel cover photo.
	CoverPhoto string `json:"cover_photo,omitempty"`
	// URL to a ~360x176  pixel cover photo.
	CoverPhotoSmall string `json:"cover_photo_small,omitempty"`
	SportType       string `json:"sport_type,omitempty"`
	// The club's city.
	City string `json:"city,omitempty"`
	// The club's state or geographical region.
	State string `json:"state,omitempty"`
	// The club's country.
	Country string `json:"country,omitempty"`
	// Whether the club is private.
	Private bool `json:"private,omitempty"`
	// The club's member count.
	MemberCount int32 `json:"member_count,omitempty"`
	// Whether the club is featured or not.
	Featured bool `json:"featured,omitempty"`
	// Whether the club is verified or not.
	Verified bool `json:"verified,omitempty"`
	// The club's vanity URL.
	Url string `json:"url,omitempty"`
}
