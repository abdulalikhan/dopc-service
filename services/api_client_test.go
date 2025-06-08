package services

import (
	"dopc-service/constants"
	models2 "dopc-service/models"
	types2 "dopc-service/types"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVenueStaticInfo(t *testing.T) {
	venueSlug := "test-venue"
	mockStaticInfo := models2.VenueStaticInfo{
		VenueRaw: models2.VenueRawStatic{
			Location: &types2.Location{
				Coordinates: []float64{24.92813512, 60.17012143},
			},
		},
	}

	// Mock the API for this UT
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := fmt.Sprintf("/%s/static", venueSlug)
		assert.Equal(t, expectedURL, r.URL.Path) // compare only the path
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockStaticInfo)
	}))
	defer mockServer.Close()

	constants.StaticEndpoint = mockServer.URL + "/%s/static"
	info, err := GetVenueStaticInfo(venueSlug)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, mockStaticInfo, *info)
}

func TestGetVenueDynamicInfo(t *testing.T) {
	venueSlug := "test-venue"
	mockDynamicInfo := models2.VenueDynamicInfo{
		VenueRaw: models2.VenueRawDynamic{
			DeliverySpecs: &models2.DeliverySpecs{
				OrderMinNoSurcharge: 500,
				DeliveryPricing: models2.DeliveryPricing{
					BasePrice: 100,
					DistanceRanges: []types2.DistanceRange{
						{Min: 0, Max: 10, A: 0, B: 0},
						{Min: 10, Max: 20, A: 2, B: 0},
					},
				},
			},
		},
	}

	// Mock the API for this UT
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := fmt.Sprintf("/%s/dynamic", venueSlug)
		assert.Equal(t, expectedURL, r.URL.Path) // compare only the path
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockDynamicInfo)
	}))
	defer mockServer.Close()

	constants.DynamicEndpoint = mockServer.URL + "/%s/dynamic"
	info, err := GetVenueDynamicInfo(venueSlug)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, mockDynamicInfo, *info)
}
