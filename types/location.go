package types

// Location type models the coordinates of the venue [longitude, latitude] as a float64 array
type Location struct {
	Coordinates []float64 `json:"coordinates"`
}

// Lon enables us to access the longitude from the coordinate array
func (l *Location) Lon() float64 {
	if len(l.Coordinates) > 0 {
		return l.Coordinates[0]
	}
	return 0
}

// Lat enables us to access the latitude from the coordinate array
func (l *Location) Lat() float64 {
	if len(l.Coordinates) > 1 {
		return l.Coordinates[1]
	}
	return 0
}
