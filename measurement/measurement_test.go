package measurement

import (
	"testing"
)

func TestDistanceUnitCreation(t *testing.T) {
	_, err := NewDistance(1, Kilometer)
	if err != nil {
		t.Errorf("Distance unit Not created")
	}

}

func TestDistanceUnitNotCreatedWithNegativeDistance(t *testing.T) {
	_, err := NewDistance(-1, Kilometer)
	if err == nil {
		t.Errorf("Distance unit should not be created with negative value")
	}

}

func TestEqualityforMeters(t *testing.T) {
	d1, _ := NewDistance(1000, Meter)
	d2, _ := NewDistance(1000, Meter)

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforMetersAndkilometers(t *testing.T) {
	d1, _ := NewDistance(1000, Meter)
	d2, _ := NewDistance(1, Kilometer)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}

}

func TestEqualitForKilometerAndMeter(t *testing.T) {
	d1, _ := NewDistance(1, Kilometer)
	d2, _ := NewDistance(1000, Meter)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualityforKiloMeters(t *testing.T) {
	d1, _ := NewDistance(1, Kilometer)
	d2, _ := NewDistance(1, Kilometer)

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestCentimeterUnitCreation(t *testing.T) {
	_, err := NewDistance(10, Centimeter)
	if err != nil {
		t.Errorf("unit not created with Centimeter")
	}
}

func TestEqualityforCentiMeters(t *testing.T) {
	d1, _ := NewDistance(10, Centimeter)
	d2, _ := NewDistance(10, Centimeter)

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualitForCentimeterAndMeter(t *testing.T) {
	d1, _ := NewDistance(1000, Centimeter)
	d2, _ := NewDistance(10, Meter)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForMeterAndCentimeter(t *testing.T) {
	d2, _ := NewDistance(1000, Centimeter)
	d1, _ := NewDistance(10, Meter)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForCentimeterAndKilometer(t *testing.T) {
	d1, _ := NewDistance(1000000, Centimeter)
	d2, _ := NewDistance(10, Kilometer)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForKilometerAndCentimeter(t *testing.T) {
	d2, _ := NewDistance(1000000, Centimeter)
	d1, _ := NewDistance(10, Kilometer)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestNewDistance(t *testing.T) {
	_, err := NewDistance(12, Meter)

	if err != nil {
		t.Errorf("Wanted distance but not created")
	}
}

func TestWeightCreation(t *testing.T) {
	_, err := NewWeight(1, Kilogram)
	if err != nil {
		t.Errorf("Weight Not created")
	}
}

func TestEqualityforgrams(t *testing.T) {
	w1, _ := NewWeight(1000, Gram)
	w2, _ := NewWeight(1000, Gram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforGramsAndKilograms(t *testing.T) {
	w1, _ := NewWeight(1000, Gram)
	w2, _ := NewWeight(1, Kilogram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}

}

func TestEqualitForKilogramsAndGrams(t *testing.T) {
	w1, _ := NewWeight(1, Kilogram)
	w2, _ := NewWeight(1000, Gram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforKiloGrams(t *testing.T) {
	w1, _ := NewWeight(1000, Kilogram)
	w2, _ := NewWeight(1000, Kilogram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}
}

func TestAddWeight(t *testing.T) {
	w1, _ := NewWeight(1000, Kilogram)
	w2, _ := NewWeight(1, Gram)
	w3 := w1.Add(w2)
	expectedWeight := Weight{value: 1000.001, unit: Kilogram}
	if w3.equals(&expectedWeight) != true {
		t.Errorf("Wanted  %v but got  %v", expectedWeight.value, w3.value)

	}

}

func TestAddWeight1(t *testing.T) {
	w1, _ := NewWeight(10, Kilogram)
	w2, _ := NewWeight(1000000, Milligram)
	w3 := w1.Add(w2)
	expectedWeight := Weight{value: 11, unit: Kilogram}
	if w3.equals(&expectedWeight) != true {
		t.Errorf("Wanted  %v but got  %v", expectedWeight.value, w3.value)

	}

}

func TestAddWeight2(t *testing.T) {
	w1, _ := NewWeight(1000, Gram)
	w2, _ := NewWeight(1000000, Milligram)
	w3 := w1.Add(w2)
	expectedWeight := Weight{value: 2000, unit: Gram}
	if w3.equals(&expectedWeight) != true {
		t.Errorf("Wanted  %v but got  %v", expectedWeight.value, w3.value)

	}
}

func TestTemperatureCreationWithCelsius(t *testing.T) {
	_, err := NewTemperature(32, Celsius)
	if err != nil {
		t.Errorf("temperature not created")
	}
}

func TestTemperatureCreationWithKelvin(t *testing.T) {
	_, err := NewTemperature(32, Kelvin)
	if err != nil {
		t.Errorf("temperature not created")
	}
}

func TestTemperatureCreationWithFahreinheit(t *testing.T) {
	_, err := NewTemperature(32, Fahrenheit)
	if err != nil {
		t.Errorf("temperature not created")
	}
}

func TestEqualitForCelsiusAndKelvin(t *testing.T) {
	oneCelsius, _ := NewTemperature(1, Celsius)
	twoSeventyFourKelvin, _ := NewTemperature(274.15, Kelvin)

	if oneCelsius.equals(twoSeventyFourKelvin) != true {
		t.Errorf("one celsius should be equal to 274.15 kelvin")
	}
}

func TestEqualitForCelsiusAndCelsius(t *testing.T) {
	oneCelsius, _ := NewTemperature(1, Celsius)
	twoCelsius, _ := NewTemperature(2, Celsius)

	if oneCelsius.equals(twoCelsius) == true {
		t.Errorf("one celsius should not be equal to two celsius")
	}
}

func TestEqualitForCelsiusAndFahrenheit(t *testing.T) {
	zeroCelsius, _ := NewTemperature(0, Celsius)
	thirtyTwoFahrenheit, _ := NewTemperature(32, Fahrenheit)
	if zeroCelsius.equals(thirtyTwoFahrenheit) != true {
		t.Errorf("zero celsius should be equal to thirty two fahrenheit")
	}
}

func TestCannotCreateTempBelowNegativeLimitOfCelsius(t *testing.T) {
	_, err := NewTemperature(-473, Celsius)
	if err == nil {
		t.Errorf("Must not create temperature below -273.15 celsius")
	}
}

func TestKilometerToMeter(t *testing.T) {
	oneKilometer, _ := NewDistance(1, Kilometer)
	if oneKilometer.unit.toMeter(oneKilometer.value) != 1000 {
		t.Errorf("one Kilometer should be equal to 1000 meter")
	}
}

func TestCentimeterToMeter(t *testing.T) {
	hundredCentimeter, _ := NewDistance(100, Centimeter)
	if hundredCentimeter.unit.toMeter(hundredCentimeter.value) != 1 {
		t.Errorf("hundred centimeter should be equal to 1 meter")
	}
}

func TestKilogramToGram(t *testing.T) {
	oneKilogram, _ := NewWeight(1, Kilogram)
	if oneKilogram.unit.toGram(oneKilogram.value) != 1000 {
		t.Errorf("one kilogram should be equal to 1000 gram")
	}
}

func TestMilligramToGram(t *testing.T) {
	thousandMilligram, _ := NewWeight(1000, Milligram)
	if thousandMilligram.unit.toGram(thousandMilligram.value) != 1 {
		t.Errorf("thousand milligram should be equal to 1 gram")
	}
}

func TestAdderInterface(t *testing.T) {
	oneKilometer, _ := NewDistance(1, Kilometer)
	var a Adder
	a = oneKilometer
	_, err := oneKilometer.Add(a)
	if err != nil {
		t.Errorf("Distance type implements Adder interface so error should be nil")
	}
}

func TestAdd1000MeterWithOneKiloMeterWithAdderInterface(t *testing.T) {
	thousandMeter, _ := NewDistance(1000, Meter)
	oneKilometer, _ := NewDistance(1, Kilometer)
	var a Adder
	a = oneKilometer
	res, err := thousandMeter.Add(a)
	if err != nil {
		t.Errorf("Distance type implements Adder interface so error should be nil")
	}
	if res.(*Distance).value != 2000 {
		t.Errorf("1000 meter and 1 kilometer should be equal to 2000 meter")
	}
}
