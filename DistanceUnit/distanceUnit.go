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
	value float64
	unit  Unit
}

func (d1 distanceUnit) toString() string {
	return fmt.Sprintf("value %f %s", d1.value, d1.unit)
}

func (d1 *distanceUnit) Add(d2 *distanceUnit) *distanceUnit {
	if d1.unit == kilometers {
		return &distanceUnit{value: (d1.Inmeter().value + d2.Inmeter().value) * 0.001, unit: kilometers}
	}
	return &distanceUnit{value: d1.Inmeter().value + d2.Inmeter().value, unit: meters}
}

func NewDistanceUnit(value float64, unit Unit) (*distanceUnit, error) {
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &distanceUnit{value: value, unit: unit}, nil
}

func (d1 *distanceUnit) equals(d2 *distanceUnit) bool {
	return d1.Inmeter().value == d2.Inmeter().value
}

func (d *distanceUnit) Inmeter() *distanceUnit {
	m := distanceMap()
	return &distanceUnit{value: d.value * m[d.unit], unit: meters}
}

func distanceMap() map[Unit]float64 {
	m := make(map[Unit]float64)
	m[kilometers] = 1000
	m[meters] = 1
	m[centimeters] = 0.01
	return m
}
