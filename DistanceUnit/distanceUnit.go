package distance

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

type distance struct {
	value float64
	unit  Unit
}

func (d1 distance) toString() string {
	return fmt.Sprintf("value %f %s", d1.value, d1.unit)
}

func (d1 *distance) Add(d2 *distance) *distance {
	if d1.unit == kilometers {
		return &distance{value: (d1.Inmeter().value + d2.Inmeter().value) * 0.001, unit: kilometers}
	}
	return &distance{value: d1.Inmeter().value + d2.Inmeter().value, unit: meters}
}

func NewDistanceUnit(value float64, unit Unit) (*distance, error) {
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &distance{value: value, unit: unit}, nil
}

func (d1 *distance) equals(d2 *distance) bool {
	return d1.Inmeter().value == d2.Inmeter().value
}

func (d *distance) Inmeter() *distance {
	m := distanceMap()
	return &distance{value: d.value * m[d.unit], unit: meters}
}

func distanceMap() map[Unit]float64 {
	m := make(map[Unit]float64)
	m[kilometers] = 1000
	m[meters] = 1
	m[centimeters] = 0.01
	return m
}
