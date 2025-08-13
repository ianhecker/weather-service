package coordinates

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Coordinates struct {
	Latitude, Longitude float64
}

func NewCoordinates(lat, long string) (*Coordinates, error) {

	if lat == "" || long == "" {
		return nil, fmt.Errorf("given empty coordinates. lat='%s' long='%s'", lat, long)
	}

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, err
	}

	if math.IsNaN(latitude) {
		return nil, errors.New("latitude is NaN")
	}

	if math.IsInf(latitude, 0) {
		return nil, errors.New("latitude is infinity")
	}

	if -90 > latitude || latitude > 90 {
		return nil, fmt.Errorf("latitude is outside bounds: -90<lat>90: lat=%f", latitude)
	}

	longitude, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return nil, err
	}

	if math.IsNaN(longitude) {
		return nil, errors.New("longitude is NaN")
	}

	if math.IsInf(longitude, 0) {
		return nil, errors.New("longitude is infinity")
	}

	if -180 > longitude || longitude > 180 {
		return nil, fmt.Errorf("longitude is outside bounds: -180<lat>180: lat=%f", longitude)
	}

	return &Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
