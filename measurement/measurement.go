package measurement

import (
	"errors"
	"math"
)

type Unit struct {
	name                 string
	baseConversionFactor float64
}
type TemperatureUnit struct {
	name                 string
	baseConversionFactor float64
	baseAdditionFactor   float64
}

var (
	Celsius    = TemperatureUnit{name: "celsius", baseConversionFactor: 1, baseAdditionFactor: 0}
	Kelvin     = TemperatureUnit{name: "fahrenheit", baseConversionFactor: math.Round(float64(5.0 / 9.0)), baseAdditionFactor: -273.15}
	Fahrenheit = TemperatureUnit{name: "kelvin", baseConversionFactor: 1, baseAdditionFactor: math.Round(-32 * (math.Round(float64(5.0 / 9.0))))}
)
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

type Temperature struct {
	value float64
	unit  TemperatureUnit
}

type Distance struct {
	value float64
	unit  Unit
}

type Weight struct {
	measurement
}

func (u *TemperatureUnit) toCelsius(value float64) float64 {
	return value*u.baseConversionFactor + u.baseAdditionFactor
}
func NewTemperature(value float64, unit TemperatureUnit) (*Temperature, error) {
	celsiusValue := unit.toCelsius(value)
	if celsiusValue < (-273.15) {
		return nil, errors.New("Temperarture cannot be creted with this value")
	}
	return &Temperature{value: value, unit: unit}, nil

}

func NewDistance(value float64, unit Unit) (*Distance, error) { //creating new Distance struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	if unit == meter || unit == kilometer || unit == centimeter {
		return &Distance{value: value, unit: unit}, nil
	}
	return nil, errors.New("invalid unit")
}

func NewWeight(value float64, unit Unit) (*Weight, error) { //creating new Weight struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	if unit == gram || unit == kilogram || unit == milligram {
		return &Weight{measurement{value: value, unit: unit}}, nil
	}
	return nil, errors.New("invalid unit")
}

func (t1 *Temperature) equals(t2 *Temperature) bool {
	return (t1.inBase()).value == (t2.inBase()).value
}

func (d1 *Distance) equals(d2 *Distance) bool { //Checking equality between the distances
	d1 = &Distance{value: (d1.value * d1.unit.baseConversionFactor)}
	d2 = &Distance{value: (d2.value * d2.unit.baseConversionFactor)}
	return d1.value == d2.value
}
func (w1 *Weight) equals(w2 *Weight) bool { //Checking equality between the wegihts
	w1 = &Weight{measurement{value: (w1.value * w1.unit.baseConversionFactor)}}
	w2 = &Weight{measurement{value: (w2.value * w2.unit.baseConversionFactor)}}
	return w1.value == w2.value
}

func (t *Temperature) inBase() *Temperature {
	valueInCelsius := t.unit.toCelsius(t.value)
	return &Temperature{value: valueInCelsius, unit: Celsius}
}

func (d1 *Distance) Add(d2 *Distance) *Distance {
	m1 := d1
	m2 := d2

	baseResult := (m1.value * m1.unit.baseConversionFactor) + (m2.value * m2.unit.baseConversionFactor)

	resultInSelfUnit := baseResult / m1.unit.baseConversionFactor

	return &Distance{value: resultInSelfUnit, unit: m1.unit}
}

func (w1 *Weight) Add(w2 *Weight) *Weight {
	d1 := w1.measurement
	d2 := w2.measurement
	baseResult := (d1.value * d1.unit.baseConversionFactor) + (d2.value * d2.unit.baseConversionFactor)

	resultInSelfUnit := baseResult / d1.unit.baseConversionFactor

	return &Weight{measurement{value: resultInSelfUnit, unit: d1.unit}}
}
