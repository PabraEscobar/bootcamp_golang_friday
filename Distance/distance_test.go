package distance

import (
	"testing"
)

func TestDistanceUnitCreation(t *testing.T) {
	_, err := NewDistance(1, kilometer)
	if err != nil {
		t.Errorf("Distance unit Not created")
	}

}

func TestDistanceUnitNotCreatedWithNegativeDistance(t *testing.T) {
	_, err := NewDistance(-1, kilometer)
	if err == nil {
		t.Errorf("Distance unit should not be created with negative value")
	}

}

func TestEqualityforMeters(t *testing.T) {
	d1, _ := NewDistance(1000, meter)
	d2, _ := NewDistance(1000, meter)

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforMetersAndkilometers(t *testing.T) {
	d1, _ := NewDistance(1000, meter)
	d2, _ := NewDistance(1, kilometer)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}

}

func TestEqualitForKilometerAndMeter(t *testing.T) {
	d1, _ := NewDistance(1, kilometer)
	d2, _ := NewDistance(1000, meter)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualityforKiloMeters(t *testing.T) {
	d1, _ := NewDistance(1, kilometer)
	d2, _ := NewDistance(1, kilometer)

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestCentimeterUnitCreation(t *testing.T) {
	_, err := NewDistance(10, centimeter)
	if err != nil {
		t.Errorf("unit not created with centimeter")
	}
}

func TestEqualityforCentiMeters(t *testing.T) {
	d1, _ := NewDistance(10, centimeter)
	d2, _ := NewDistance(10, centimeter)

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualitForCentimeterAndMeter(t *testing.T) {
	d1, _ := NewDistance(1000, centimeter)
	d2, _ := NewDistance(10, meter)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForMeterAndCentimeter(t *testing.T) {
	d2, _ := NewDistance(1000, centimeter)
	d1, _ := NewDistance(10, meter)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForCentimeterAndKilometer(t *testing.T) {
	d1, _ := NewDistance(1000000, centimeter)
	d2, _ := NewDistance(10, kilometer)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForKilometerAndCentimeter(t *testing.T) {
	d2, _ := NewDistance(1000000, centimeter)
	d1, _ := NewDistance(10, kilometer)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestAddDistance(t *testing.T) {
	d1, _ := NewDistance(1000, meter)
	d2, _ := NewDistance(1000, meter)
	d3 := d1.Add(d2)
	expectedDistance := Distance{measurement{value: 2000, unit: meter}}
	if d3.equals(&expectedDistance) != true {
		t.Errorf("Wanted  %v but got  %v", expectedDistance.toString(), (*d3).toString())

	}

}

func TestAddDistance1(t *testing.T) {
	d1, _ := NewDistance(100, meter)
	d2, _ := NewDistance(1, kilometer)
	d3 := d1.Add(d2)
	expectedDistance := Distance{measurement{value: 1100, unit: meter}}
	if d3.equals(&expectedDistance) != true {
		t.Errorf("Wanted  %v but got  %v", expectedDistance.toString(), (*d3).toString())
	}

}

func TestAddDistance2(t *testing.T) {
	d1, _ := NewDistance(100, kilometer)
	d2, _ := NewDistance(1, meter)
	d3 := d1.Add(d2)
	expectedDistance := Distance{measurement{value: 100.001, unit: kilometer}}
	if d3.equals(&expectedDistance) != true {
		t.Errorf("Wanted  %v but got  %v", expectedDistance.toString(), (*d3).toString())
	}

}

func TestNewDistance(t *testing.T) {
	_, err := NewDistance(12, meter)

	if err != nil {
		t.Errorf("Wanted distance but not created")
	}
}

func TestWeightCreation(t *testing.T) {
	_, err := NewWeight(1, kilogram)
	if err != nil {
		t.Errorf("Weight Not created")
	}
}

func TestEqualityforgrams(t *testing.T) {
	w1, _ := NewWeight(1000, gram)
	w2, _ := NewWeight(1000, gram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforGramsAndKilograms(t *testing.T) {
	w1, _ := NewWeight(1000, gram)
	w2, _ := NewWeight(1, kilogram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}

}

func TestEqualitForKilogramsAndGrams(t *testing.T) {
	w1, _ := NewWeight(1, kilogram)
	w2, _ := NewWeight(1000, gram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforKiloGrams(t *testing.T) {
	w1, _ := NewWeight(1000, kilogram)
	w2, _ := NewWeight(1000, kilogram)

	if w1.equals(w2) != true {
		t.Errorf("unequal distance")
	}
}

func TestAddWeight(t *testing.T) {
	w1, _ := NewWeight(1000, kilogram)
	w2, _ := NewWeight(1, gram)
	w3 := w1.Add(w2)
	expectedWeight := Weight{measurement{value: 1000.001, unit: kilogram}}
	if w3.equals(&expectedWeight) != true {
		t.Errorf("Wanted  %v but got  %v", expectedWeight.toString(), (*w3).toString())

	}

}

func TestAddWeight1(t *testing.T) {
	w1, _ := NewWeight(10, kilogram)
	w2, _ := NewWeight(1000000, milligram)
	w3 := w1.Add(w2)
	expectedWeight := Weight{measurement{value: 11, unit: kilogram}}
	if w3.equals(&expectedWeight) != true {
		t.Errorf("Wanted  %v but got  %v", expectedWeight.toString(), (*w3).toString())

	}

}

func TestAddWeight2(t *testing.T) {
	w1, _ := NewWeight(1000, gram)
	w2, _ := NewWeight(1000000, milligram)
	w3 := w1.Add(w2)
	expectedWeight := Weight{measurement{value: 2000, unit: gram}}
	if w3.equals(&expectedWeight) != true {
		t.Errorf("Wanted  %v but got  %v", expectedWeight.toString(), (*w3).toString())

	}
}

func TestDistanceWithgram(t *testing.T) {
	_, err := NewDistance(12, gram)
	if err == nil {
		t.Errorf("distance is not in correct unit")
	}
}

func TestWeightWithmeter(t *testing.T) {
	_, err := NewWeight(12, meter)
	if err == nil {
		t.Errorf("weight can't measured in this unit")
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
