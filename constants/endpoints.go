package constants

// Base URL for the Venues API
const BaseURL = "https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues"

var (
	// API Endpoints
	DynamicEndpoint = BaseURL + "/%s/dynamic"
	StaticEndpoint  = BaseURL + "/%s/static"
)
