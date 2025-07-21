package distanceUnit

import (
	"errors"
	"fmt"
)

type Unit string

const (
	meters      Unit = "meters"
	kilometers  Unit = "kilometers"
	centimeters Unit = "centimeters"
)

type distanceUnit struct {
	distance float64
	unit     Unit
}

func (d1 distanceUnit) toString() string {
	return fmt.Sprintf("distance %f %s", d1.distance, d1.unit)
}

func (d1 *distanceUnit) Add(d2 *distanceUnit) *distanceUnit {
	if d1.unit == kilometers {
		return &distanceUnit{distance: (d1.Inmeter().distance + d2.Inmeter().distance) * 0.001, unit: kilometers}
	}
	return &distanceUnit{distance: d1.Inmeter().distance + d2.Inmeter().distance, unit: meters}
}

func NewDistanceUnit(distance float64, unit Unit) (*distanceUnit, error) {
	if distance < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &distanceUnit{distance: distance, unit: unit}, nil
}

func (d1 *distanceUnit) equals(d2 *distanceUnit) bool {
	return d1.Inmeter().distance == d2.Inmeter().distance
}

func (d *distanceUnit) Inmeter() *distanceUnit {
	switch d.unit {
	case centimeters:
		return &distanceUnit{distance: d.distance * 0.01, unit: meters}
	case kilometers:
		return &distanceUnit{distance: d.distance * 1000, unit: meters}
	default:
		return d
	}
}
