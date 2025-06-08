package models

import (
	"dopc-service/types"
)

type DeliveryPricing struct {
	BasePrice      int                   `json:"base_price"`
	DistanceRanges []types.DistanceRange `json:"distance_ranges"`
}
type DeliverySpecs struct {
	OrderMinNoSurcharge int             `json:"order_minimum_no_surcharge"`
	DeliveryPricing     DeliveryPricing `json:"delivery_pricing"`
}
type VenueRawDynamic struct {
	DeliverySpecs *DeliverySpecs `json:"delivery_specs"`
}
type VenueDynamicInfo struct {
	VenueRaw VenueRawDynamic `json:"venue_raw"`
}
