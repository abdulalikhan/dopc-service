package handlers

import (
	"dopc-service/constants"
	"dopc-service/services"
	"dopc-service/utils"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/beego/beego/v2/server/web"
	"github.com/jftuga/geodist"
)

type DeliveryInfo struct {
	Fee      int `json:"fee"`
	Distance int `json:"distance"`
}

type PriceResponse struct {
	TotalPrice          int          `json:"total_price"`
	SmallOrderSurcharge int          `json:"small_order_surcharge"`
	CartValue           int          `json:"cart_value"`
	Delivery            DeliveryInfo `json:"delivery"`
}

type DeliveryPriceHandler struct {
	web.Controller
}

// @Title Get Delivery Price
// @Description Get the price for a delivery order based on the provided venue slug, cart value, and user's location.
// @Param   venue_slug  query   string     true    "The venue slug"
// @Param   cart_value  query   int        true    "The value of the shopping cart"
// @Param   user_lat    query   float64    true    "The latitude of the user"
// @Param   user_lon    query   float64    true    "The longitude of the user"
// @Success 200 {object}  PriceResponse  "The calculated price details"
// @Failure 400 {string}  "Invalid input parameters"
// @Failure 500 {string}  "Internal server error"
// @BasePath /api/v1
// @router / [get]
func (p *DeliveryPriceHandler) Get() {
	venueSlug := p.Ctx.Input.Query(constants.QueryParamVenueSlug)
	cartValueRaw := p.Ctx.Input.Query(constants.QueryParamCartValue)
	userLatRaw := p.Ctx.Input.Query(constants.QueryParamUserLat)
	userLonRaw := p.Ctx.Input.Query(constants.QueryParamUserLon)

	// convert cart_value, user_lat, and user_lon to proper types
	cartValue, err := strconv.Atoi(cartValueRaw)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(400) // Bad Request
		p.Ctx.WriteString("Invalid cart_value")
		return
	}
	userLat, err := strconv.ParseFloat(userLatRaw, 64)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(400)
		p.Ctx.WriteString("Invalid user_lat")
		return
	}
	userLon, err := strconv.ParseFloat(userLonRaw, 64)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(400)
		p.Ctx.WriteString("Invalid user_lon")
		return
	}

	// fetch venue information from API's static endpoint
	staticInfo, err := services.GetVenueStaticInfo(venueSlug)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(500) // Internal Server Error
		p.Ctx.WriteString(fmt.Sprintf("Error fetching static venue info: %v", err))
		return
	}

	// Fetch venue information from API's dynamic endpoint
	dynamicInfo, err := services.GetVenueDynamicInfo(venueSlug)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(500)
		p.Ctx.WriteString(fmt.Sprintf("Error fetching dynamic venue info: %v", err))
		return
	}

	// small_order_surcharge is the difference between order_minimum_no_surcharge (as received from the API) and the cart value
	smallOrderSurcharge := int(math.Abs(float64(dynamicInfo.VenueRaw.DeliverySpecs.OrderMinNoSurcharge - cartValue)))

	// calculate the distance between the venue and the user in meters
	var userCoordinates = geodist.Coord{Lat: userLat, Lon: userLon}
	var venueCoordinates = geodist.Coord{Lat: staticInfo.VenueRaw.Location.Lat(), Lon: staticInfo.VenueRaw.Location.Lon()}
	_, distanceKm := geodist.HaversineDistance(venueCoordinates, userCoordinates)
	distanceMeters := int(math.Round(distanceKm * 1000.0))

	// use the calculated distance to calculate the delivery fee
	deliveryFee, err := utils.CalcDeliveryFee(dynamicInfo.VenueRaw.DeliverySpecs.DeliveryPricing.BasePrice, distanceMeters, &dynamicInfo.VenueRaw.DeliverySpecs.DeliveryPricing.DistanceRanges)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(400)
		p.Ctx.WriteString(fmt.Sprintf("Error calculating delivery fee: %v", err))
		return
	}

	// prepare response object for JSON response
	response := PriceResponse{
		TotalPrice:          cartValue + smallOrderSurcharge + deliveryFee,
		SmallOrderSurcharge: smallOrderSurcharge,
		CartValue:           cartValue,
		Delivery: DeliveryInfo{
			Fee:      deliveryFee,
			Distance: distanceMeters,
		},
	}

	p.Ctx.Output.Header("Content-Type", "application/json")
	if err := json.NewEncoder(p.Ctx.ResponseWriter).Encode(response); err != nil {
		p.Ctx.ResponseWriter.WriteHeader(500)
		p.Ctx.WriteString("Error encoding JSON")
	}
}
