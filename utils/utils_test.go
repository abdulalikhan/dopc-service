package utils

import (
	"dopc-service/types"
	"testing"
)

func TestCalcDeliveryFee(t *testing.T) {
	testCases := []struct {
		label          string
		basePrice      int
		distanceMeters int
		distanceRanges *[]types.DistanceRange
		expectedFee    int
		error          string
	}{
		{
			label:          "Test #1: params within valid range",
			basePrice:      199,
			distanceMeters: 600,
			distanceRanges: &[]types.DistanceRange{
				{Min: 0, Max: 500, A: 0, B: 0},
				{Min: 500, Max: 1000, A: 100, B: 1},
				{Min: 1000, Max: 0, A: 0, B: 0},
			},
			expectedFee: 359, // 199 + 100 + 1 * 600 / 10 == 359
			error:       "",
		},
		{
			label:          "Test #2: Delivery not available for distances >= 2000 meters",
			basePrice:      100,
			distanceMeters: 3000,
			distanceRanges: &[]types.DistanceRange{
				{Min: 0, Max: 1000, A: 50, B: 5},
				{Min: 1000, Max: 2000, A: 60, B: 10},
				{Min: 2000, Max: 0, A: 70, B: 15},
			},
			expectedFee: -1,
			error:       "delivery is not available for delivery distances equal to or longer than 2000 meters",
		},
		{
			label:          "Test #3: Delivery fee when no distance ranges are provided",
			basePrice:      100,
			distanceMeters: 500,
			distanceRanges: &[]types.DistanceRange{},
			expectedFee:    100,
			error:          "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.label, func(t *testing.T) {
			fee, err := CalcDeliveryFee(testCase.basePrice, testCase.distanceMeters, testCase.distanceRanges)
			if err != nil && err.Error() != testCase.error {
				t.Errorf("unexpected error `%v` occurred - expected `%v`", err, testCase.error)
			}
			if err == nil && testCase.error != "" {
				t.Errorf("expected error `%v` - got none.", testCase.error)
			}
			if fee != testCase.expectedFee {
				t.Errorf("unexpected delivery fee `%v` - expected `%v`", fee, testCase.expectedFee)
			}
		})
	}
}
