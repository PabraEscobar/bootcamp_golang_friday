package distanceUnit

import (
	"errors"
)

type Unit int

const (
	meters Unit = iota
	kilometers
	centimeters
)

type distanceUnit struct {
	distance float64
	unit     Unit
}

func NewDistanceUnit(distance float64, unit Unit) (*distanceUnit, error) {
	if distance < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &distanceUnit{distance: distance, unit: unit}, nil
}

func (d1 *distanceUnit) EqualityOfDistance(d2 *distanceUnit) bool {
	return d1.Inmeter().distance == d2.Inmeter().distance
}

func (d *distanceUnit) Inmeter() *distanceUnit {
	if d.unit == centimeters {
		d = &distanceUnit{distance: d.distance / 100, unit: meters}
	} else if d.unit == kilometers {
		d = &distanceUnit{distance: d.distance * 1000, unit: meters}
	}
	return d
}

func (d1 *distanceUnit) TotalDistanceInMeters(d2 *distanceUnit) *distanceUnit {
	return &distanceUnit{distance: d1.Inmeter().distance + d2.Inmeter().distance, unit: meters}
}
