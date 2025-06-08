package services

import (
	"dopc-service/constants"
	models2 "dopc-service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func fetchAPI(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d from API", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}
	return nil
}

// defining API fetcher functions as variables so that they can be mocked in UTs
var (
	GetVenueStaticInfo = func(venueSlug string) (*models2.VenueStaticInfo, error) {
		url := fmt.Sprintf(constants.StaticEndpoint, venueSlug)
		var info models2.VenueStaticInfo
		if err := fetchAPI(url, &info); err != nil {
			return nil, err
		}
		return &info, nil
	}

	GetVenueDynamicInfo = func(venueSlug string) (*models2.VenueDynamicInfo, error) {
		url := fmt.Sprintf(constants.DynamicEndpoint, venueSlug)
		var info models2.VenueDynamicInfo
		if err := fetchAPI(url, &info); err != nil {
			return nil, err
		}
		return &info, nil
	}
)
