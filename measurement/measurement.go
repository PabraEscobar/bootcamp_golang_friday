package measurement

import (
	"errors"
	"math"
)

type Adder interface {
	Add(measurement Adder) (Adder, error)
}

type TemperatureUnit struct {
	name                 string
	baseConversionFactor float64
	baseAdditionFactor   float64
}
type DistanceUnit struct {
	name                 string
	baseConversionFactor float64
}
type WeightUnit struct {
	name                 string
	baseConversionFactor float64
}

var (
	Meter      = DistanceUnit{name: "meter", baseConversionFactor: 1}
	Kilometer  = DistanceUnit{name: "kilometer", baseConversionFactor: 1000}
	Centimeter = DistanceUnit{name: "centimeter", baseConversionFactor: 0.01}
	Celsius    = TemperatureUnit{name: "celsius", baseConversionFactor: 1, baseAdditionFactor: 0}
	Kelvin     = TemperatureUnit{name: "fahrenheit", baseConversionFactor: math.Round(float64(5.0 / 9.0)), baseAdditionFactor: -273.15}
	Fahrenheit = TemperatureUnit{name: "kelvin", baseConversionFactor: 1, baseAdditionFactor: math.Round(-32 * (math.Round(float64(5.0 / 9.0))))}
	Gram       = WeightUnit{name: "gram", baseConversionFactor: 1}
	Kilogram   = WeightUnit{name: "kilogram", baseConversionFactor: 1000}
	Milligram  = WeightUnit{name: "milligram", baseConversionFactor: 0.001}
)

type Temperature struct {
	value float64
	unit  TemperatureUnit
}

type Distance struct {
	value float64
	unit  DistanceUnit
}

type Weight struct {
	value float64
	unit  WeightUnit
}

func (u *TemperatureUnit) toCelsius(value float64) float64 {
	return value*u.baseConversionFactor + u.baseAdditionFactor
}
func (d *DistanceUnit) toMeter(value float64) float64 {
	return value * d.baseConversionFactor
}
func (w *WeightUnit) toGram(value float64) float64 {
	return value * w.baseConversionFactor
}
func NewTemperature(value float64, unit TemperatureUnit) (*Temperature, error) {
	celsiusValue := unit.toCelsius(value)
	if celsiusValue < (-273.15) {
		return nil, errors.New("Temperarture cannot be creted with this value")
	}
	return &Temperature{value: value, unit: unit}, nil

}

func NewDistance(value float64, unit DistanceUnit) (*Distance, error) { //creating new Distance struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &Distance{value: value, unit: unit}, nil

}

func NewWeight(value float64, unit WeightUnit) (*Weight, error) { //creating new Weight struct
	if value < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &Weight{value: value, unit: unit}, nil

}

func (t1 *Temperature) equals(t2 *Temperature) bool {
	return (t1.inBase()).value == (t2.inBase()).value
}

func (d1 *Distance) equals(d2 *Distance) bool { //Checking equality between the distances
	valueOfd1 := d1.unit.toMeter(d1.value)
	valueOfd2 := d2.unit.toMeter(d2.value)
	return valueOfd1 == valueOfd2
}
func (w1 *Weight) equals(w2 *Weight) bool { //Checking equality between the wegihts
	valueOfw1 := w1.unit.toGram(w1.value)
	valueOfW2 := w2.unit.toGram(w2.value)
	return valueOfw1 == valueOfW2
}

func (t *Temperature) inBase() *Temperature {
	valueInCelsius := t.unit.toCelsius(t.value)
	return &Temperature{value: valueInCelsius, unit: Celsius}
}

func (d1 *Distance) Add(m Adder) (Adder, error) {
	m1 := d1.unit.toMeter(d1.value)
	d2, flag := m.(*Distance)
	if flag == true {
		m2 := d2.unit.toMeter(d2.value)
		baseResult := m1 + m2
		resultInSelfUnit := baseResult / d1.unit.baseConversionFactor
		return &Distance{value: resultInSelfUnit, unit: d1.unit}, nil
	}
	return nil, errors.New("want adder of type *Distance")

}

func (w1 *Weight) Add(m Adder) (Adder, error) {
	d1 := w1.unit.toGram(w1.value)
	w2, flag := m.(*Weight)
	if flag == true {
		d2 := w2.unit.toGram(w2.value)
		baseResult := d1 + d2
		resultInSelfUnit := baseResult / w1.unit.baseConversionFactor
		return &Weight{value: resultInSelfUnit, unit: w1.unit}, nil
	}
	return nil, errors.New("want adder of type *Weight")
}
