package distanceUnit

import (
	"testing"
)

func TestDistanceUnitCreation(t *testing.T) {
	_, err := NewDistanceUnit(1, "km")
	if err != nil {
		t.Errorf("Distance unit Not created")
	}

}

func TestDistanceUnitNotCreatedWithNegativeDistance(t *testing.T) {
	_, err := NewDistanceUnit(-1, "km")
	if err == nil {
		t.Errorf("Distance unit should not be created with negative distance")
	}

}

func TestEqualityforMeters(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, "meters")
	d2, _ := NewDistanceUnit(1000, "meters")

	if EqualityForMeters(*d1, *d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforMetersAndkilometers(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, "meters")
	d2, _ := NewDistanceUnit(1, "kilometer")

	d2 = d2.KilometerToMeter()
	if EqualityForMeters(*d1, *d2) != true {
		t.Errorf("want both distance equal")
	}

}
