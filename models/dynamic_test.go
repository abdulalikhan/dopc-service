package models

import (
	"dopc-service/types"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVenueDynamicInfo_UnmarshalJSON(t *testing.T) {
	// Test case 1: valid JSON
	jsonData := `{
		"venue_raw": {
			"delivery_specs": {
				"order_minimum_no_surcharge": 500,
				"delivery_pricing": {
					"base_price": 100,
					"distance_ranges": [
						{"min": 0, "max": 10, "a": 0, "b": 0},
						{"min": 10, "max": 20, "a": 2, "b": 0}
					]
				}
			}
		}
	}`

	// Expected output
	expectedDistanceRanges := []types.DistanceRange{
		{Min: 0, Max: 10, A: 0, B: 0},
		{Min: 10, Max: 20, A: 2, B: 0},
	}

	var dynamicInfo VenueDynamicInfo
	err := json.Unmarshal([]byte(jsonData), &dynamicInfo)
	assert.NoError(t, err)
	assert.Equal(t, 500, dynamicInfo.VenueRaw.DeliverySpecs.OrderMinNoSurcharge)
	assert.Equal(t, 100, dynamicInfo.VenueRaw.DeliverySpecs.DeliveryPricing.BasePrice)
	assert.NotNil(t, dynamicInfo.VenueRaw.DeliverySpecs.DeliveryPricing.DistanceRanges)
	assert.Equal(t, expectedDistanceRanges, dynamicInfo.VenueRaw.DeliverySpecs.DeliveryPricing.DistanceRanges)

	// Test case 2: invalid JSON
	jsonData = `{
		"venue_raw": {
			"delivery_specs": {
				"order_minimum_no_surcharge": "invalid_value",
				"delivery_pricing": {
					"distance_ranges": 0
				}
			}
		}
	}`

	err = json.Unmarshal([]byte(jsonData), &dynamicInfo)
	assert.Error(t, err)
}
