package distanceUnit

import "errors"

type distanceUnit struct {
	distance int
	unit     string
}

func NewDistanceUnit(distance int, unit string) (*distanceUnit, error) {
	if distance < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &distanceUnit{distance: distance, unit: unit}, nil
}
