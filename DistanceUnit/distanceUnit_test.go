package distanceUnit

import "testing"

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
