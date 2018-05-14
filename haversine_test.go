package haversine_test

import (
	"testing"

	haversine "github.com/speterson-zoll/go-haversine"
)

var testsCases = []struct {
	from           haversine.Point
	to             haversine.Point
	expectedMeters float64
}{
	{
		haversine.Point{Lat: 22.55, Lon: 43.12},
		haversine.Point{Lat: 13.45, Lon: 100.28},
		6094544.408786774,
	},
	{
		haversine.Point{Lat: 51.510357, Lon: -0.116773},
		haversine.Point{Lat: 38.889931, Lon: -77.009003},
		5897658.288856054,
	},
}

func TestDistance(t *testing.T) {
	for _, input := range testsCases {
		meters := float64(input.from.MetresTo(input.to))

		if input.expectedMeters != meters {
			t.Errorf("FAIL: Want distance from %v to %v to be: %v but we got %v",
				input.from,
				input.to,
				input.expectedMeters,
				meters,
			)
		}
	}
}
