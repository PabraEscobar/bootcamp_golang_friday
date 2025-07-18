package distanceUnit

import (
	"reflect"
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

	if !reflect.DeepEqual(d1, d2) {
		t.Errorf("unequal distance")
	}
}
