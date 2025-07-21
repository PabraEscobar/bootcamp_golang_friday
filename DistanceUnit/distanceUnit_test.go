package distanceUnit

import (
	"testing"
)

func TestDistanceUnitCreation(t *testing.T) {
	_, err := NewDistanceUnit(1, kilometers)
	if err != nil {
		t.Errorf("Distance unit Not created")
	}

}

func TestDistanceUnitNotCreatedWithNegativeDistance(t *testing.T) {
	_, err := NewDistanceUnit(-1, kilometers)
	if err == nil {
		t.Errorf("Distance unit should not be created with negative distance")
	}

}

func TestEqualityforMeters(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, meters)
	d2, _ := NewDistanceUnit(1000, meters)

	if d1.EqualityOfDistance(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforMetersAndkilometers(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, meters)
	d2, _ := NewDistanceUnit(1, kilometers)
	if d1.EqualityOfDistance(d2) != true {
		t.Errorf("want both distance equal")
	}

}

func TestEqualitForKilometerAndMeter(t *testing.T) {
	d1, _ := NewDistanceUnit(1, kilometers)
	d2, _ := NewDistanceUnit(1000, meters)
	if d1.EqualityOfDistance(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualityforKiloMeters(t *testing.T) {
	d1, _ := NewDistanceUnit(1, kilometers)
	d2, _ := NewDistanceUnit(1, kilometers)

	if d1.EqualityOfDistance(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestCentimeterUnitCreation(t *testing.T) {
	_, err := NewDistanceUnit(10, centimeters)
	if err != nil {
		t.Errorf("unit not created with centimeter")
	}
}

func TestEqualityforCentiMeters(t *testing.T) {
	d1, _ := NewDistanceUnit(10, centimeters)
	d2, _ := NewDistanceUnit(10, centimeters)

	if d1.EqualityOfDistance(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualitForCentimeterAndMeter(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, centimeters)
	d2, _ := NewDistanceUnit(10, meters)
	if d1.EqualityOfDistance(d2) != true {
		t.Errorf("want both distance equal")
	}
}
