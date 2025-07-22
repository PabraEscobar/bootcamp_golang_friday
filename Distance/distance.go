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
	meters      = Unit{name: "meters", baseConversionFactor: 1}
	kilometers  = Unit{name: "kilometers", baseConversionFactor: 1000}
	centimeters = Unit{name: "centimeters", baseConversionFactor: 0.01}
)

type measurement struct {
	value float64
	unit  Unit
}

type Distance struct {
	measurement
}

// type Weight struct{
// 	measurement
// }

func (d1 measurement) toString() string { //for displaying structure in a readable format
	return fmt.Sprintf("value %f %v", d1.value, d1.unit)
}

func NewDistanceUnit(value float64, unit Unit) (*Distance, error) { //creating new Distance struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &Distance{measurement{value: value, unit: unit}}, nil
}

func (d1 *measurement) equals(d2 *measurement) bool { //Checking equality between the distances
	return d1.InBase().value == d2.InBase().value
}

func (d *measurement) InBase() *measurement { //Converting all distances to meter
	return &measurement{value: d.value * d.unit.baseConversionFactor, unit: d.unit}
}

func (d1 *measurement) Add(d2 *measurement) *measurement {

	baseResult := d1.InBase().value + d2.InBase().value

	resultInSelfUnit := baseResult / d1.unit.baseConversionFactor

	return &measurement{value: resultInSelfUnit, unit: d1.unit} //converting both distance values to  meter and back to the left operand's unit
}
