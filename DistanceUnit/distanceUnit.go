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

func (d1 *distanceUnit) EqualityOfDistance(d2 distanceUnit) bool {
	if (*d1).distance == d2.distance && (*d1).unit == d2.unit {
		return true
	}
	return false
}

func (d *distanceUnit) KilometerToMeter() *distanceUnit {
	return &distanceUnit{distance: (d.distance * 1000), unit: "meters"}
}
