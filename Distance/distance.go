package distance

import (
	"errors"
	"fmt"
)

type Unit struct {
	name                 string
	baseConversionFactor float64
}

var (
	meter      = Unit{name: "meter", baseConversionFactor: 1}
	kilometer  = Unit{name: "kilometer", baseConversionFactor: 1000}
	centimeter = Unit{name: "centimeter", baseConversionFactor: 0.01}
	gram       = Unit{name: "gram", baseConversionFactor: 1}
	kilogram   = Unit{name: "kilogram", baseConversionFactor: 1000}
	milligram  = Unit{name: "milligram", baseConversionFactor: 0.001}
)

type measurement struct {
	value float64
	unit  Unit
}

type Distance struct {
	measurement
}

type Weight struct {
	measurement
}

func (d1 measurement) toString() string { //for displaying structure in a readable format
	return fmt.Sprintf("value %f %v", d1.value, d1.unit)
}

func NewDistanceUnit(value float64, unit Unit) (*Distance, error) { //creating new Distance struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	if unit == meter || unit == kilometer || unit == centimeter {
		return &Distance{measurement{value: value, unit: unit}}, nil
	}
	return nil, errors.New("invalid unit")
}

func NewWeightUnit(value float64, unit Unit) (*Weight, error) { //creating new Weight struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	if unit == gram || unit == kilogram || unit == milligram {
		return &Weight{measurement{value: value, unit: unit}}, nil
	}
	return nil, errors.New("invalid unit")
}

func (d1 *Distance) equals(d2 *Distance) bool { //Checking equality between the distances
	return d1.measurement.equals(&d2.measurement)
}
func (w1 *Weight) equals(w2 *Weight) bool { //Checking equality between the wegihts
	return w1.measurement.equals(&w2.measurement)
}
func (d1 *measurement) equals(d2 *measurement) bool { //Checking equality between the measurement
	return d1.InBase().value == d2.InBase().value
}

func (d *measurement) InBase() *measurement { //Converting measurement to base unit
	return &measurement{value: d.value * d.unit.baseConversionFactor, unit: d.unit}
}

func (d1 *measurement) Add(d2 *measurement) *measurement {

	baseResult := d1.InBase().value + d2.InBase().value

	resultInSelfUnit := baseResult / d1.unit.baseConversionFactor

	return &measurement{value: resultInSelfUnit, unit: d1.unit} //converting measurement back to the left operand's unit
}

func (d1 *Distance) Add(d2 *Distance) *Distance {
	return &Distance{*(d1.measurement.Add(&d2.measurement))}
}

func (w1 *Weight) Add(w2 *Weight) *Weight {
	return &Weight{*(w1.measurement.Add(&w2.measurement))}
}
