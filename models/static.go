package models

import (
	"dopc-service/types"
)

type VenueRawStatic struct {
	Location *types.Location `json:"location"`
}

type VenueStaticInfo struct {
	VenueRaw VenueRawStatic `json:"venue_raw"`
}
