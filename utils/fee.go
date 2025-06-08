package utils

import (
	"dopc-service/types"
	"fmt"
)

// CalcDeliveryFee calculates the delivery fee using the formula: base_price + a + b * distance / 10
func CalcDeliveryFee(basePrice int, distanceMeters int, distanceRanges *[]types.DistanceRange) (int, error) {
	deliveryFee := basePrice
	for i := 0; i < len(*distanceRanges); i++ {
		if (*distanceRanges)[i].Max == 0 {
			return -1, fmt.Errorf("delivery is not available for delivery distances equal to or longer than %v meters", (*distanceRanges)[i].Min)
		}
		if distanceMeters >= int((*distanceRanges)[i].Min) && distanceMeters < int((*distanceRanges)[i].Max) {
			deliveryFee += int((*distanceRanges)[i].A) + int((*distanceRanges)[i].B)*distanceMeters/10
			break
		}
	}
	return deliveryFee, nil
}
