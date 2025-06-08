package models

import (
	"dopc-service/types"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVenueStaticInfo_UnmarshalJSON(t *testing.T) {
	// Test case 1: valid JSON
	jsonData := `{
		"venue_raw": {
			"location": {
				"coordinates": [24.92813512, 60.17012143]
			}
		}
	}`

	// Expected output
	expectedLocation := &types.Location{
		Coordinates: []float64{24.92813512, 60.17012143},
	}

	var staticInfo VenueStaticInfo
	err := json.Unmarshal([]byte(jsonData), &staticInfo)

	assert.NoError(t, err)
	assert.NotNil(t, staticInfo.VenueRaw.Location)
	assert.Equal(t, expectedLocation, staticInfo.VenueRaw.Location)

	// Test case 2: invalid JSON
	jsonData = `{
		"venue_raw": {
			"location": {
				"coordinates": "INVALID VALUE"
			}
		}
	}`

	err = json.Unmarshal([]byte(jsonData), &staticInfo)
	assert.Error(t, err)
}
