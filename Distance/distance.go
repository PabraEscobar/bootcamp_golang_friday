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

func (d1 distance) toString() string { //for displaying structure in a readable format
	return fmt.Sprintf("value %f %s", d1.value, d1.unit)
}

func (d1 *distance) Add(d2 *distance) *distance {
	m := distanceMap()

	baseResult := d1.InBase().value + d2.InBase().value

	resultInSelfUnit := baseResult / m[d1.unit]

	return &distance{value: resultInSelfUnit, unit: d1.unit} //converting both distance values to  meter and back to the left operand's unit
}

func NewDistanceUnit(value float64, unit Unit) (*distance, error) { //creating new Distance struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &distance{value: value, unit: unit}, nil
}

func (d1 *distance) equals(d2 *distance) bool { //Checking equality between the distances
	return d1.InBase().value == d2.InBase().value
}

func (d *distance) InBase() *distance { //Converting all distances to meter
	m := distanceMap()
	return &distance{value: d.value * m[d.unit], unit: meters}
}

func distanceMap() map[Unit]float64 { //Distance map for conversion of units to meter
	m := make(map[Unit]float64)
	m[kilometers] = 1000
	m[meters] = 1
	m[centimeters] = 0.01
	return m
}
