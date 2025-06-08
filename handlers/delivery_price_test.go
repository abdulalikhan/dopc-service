package handlers

import (
	models2 "dopc-service/models"
	"dopc-service/services"
	types2 "dopc-service/types"
	"encoding/json"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeliveryPriceHandler_Get(t *testing.T) {
	// mocking the functions that fetch venue info from the API
	services.GetVenueStaticInfo = func(venueSlug string) (*models2.VenueStaticInfo, error) {
		return &models2.VenueStaticInfo{
			VenueRaw: models2.VenueRawStatic{
				Location: &types2.Location{
					Coordinates: []float64{24.92813512, 60.17012143},
				},
			},
		}, nil
	}
	services.GetVenueDynamicInfo = func(venueSlug string) (*models2.VenueDynamicInfo, error) {
		return &models2.VenueDynamicInfo{
			VenueRaw: models2.VenueRawDynamic{
				DeliverySpecs: &models2.DeliverySpecs{
					OrderMinNoSurcharge: 1000,
					DeliveryPricing: models2.DeliveryPricing{
						BasePrice: 190,
						DistanceRanges: []types2.DistanceRange{
							{Min: 0, Max: 500, A: 0, B: 0},
							{Min: 500, Max: 1000, A: 100, B: 0},
							{Min: 1000, Max: 1500, A: 200, B: 0},
							{Min: 1500, Max: 2000, A: 200, B: 1},
							{Min: 2000, Max: 0, A: 0, B: 0},
						},
					},
				},
			},
		}, nil
	}

	// creating a request and response recorder
	req, err := http.NewRequest("GET", "/?venue_slug=test-venue&cart_value=1000&user_lat=60.17094&user_lon=24.93087", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	// initializing context for the DeliveryPriceHandler
	ctx := context.NewContext()
	ctx.Reset(rr, req)
	handler := &DeliveryPriceHandler{}
	handler.Ctx = ctx

	// invoke the handler's GET method and assert the response
	handler.Get()
	assert.Equal(t, http.StatusOK, rr.Code)

	// match the expectations with the actual response from the handler's GET method
	var response PriceResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, 1190, response.TotalPrice)
	assert.Equal(t, 0, response.SmallOrderSurcharge)
	assert.Equal(t, 1000, response.CartValue)
	assert.Equal(t, 190, response.Delivery.Fee)
	assert.Equal(t, 177, response.Delivery.Distance)
}
