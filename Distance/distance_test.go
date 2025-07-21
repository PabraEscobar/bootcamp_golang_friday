package distance

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
		t.Errorf("Distance unit should not be created with negative value")
	}

}

func TestEqualityforMeters(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, meters)
	d2, _ := NewDistanceUnit(1000, meters)

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualityforMetersAndkilometers(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, meters)
	d2, _ := NewDistanceUnit(1, kilometers)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}

}

func TestEqualitForKilometerAndMeter(t *testing.T) {
	d1, _ := NewDistanceUnit(1, kilometers)
	d2, _ := NewDistanceUnit(1000, meters)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualityforKiloMeters(t *testing.T) {
	d1, _ := NewDistanceUnit(1, kilometers)
	d2, _ := NewDistanceUnit(1, kilometers)

	if d1.equals(d2) != true {
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

	if d1.equals(d2) != true {
		t.Errorf("unequal distance")
	}
}

func TestEqualitForCentimeterAndMeter(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, centimeters)
	d2, _ := NewDistanceUnit(10, meters)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForMeterAndCentimeter(t *testing.T) {
	d2, _ := NewDistanceUnit(1000, centimeters)
	d1, _ := NewDistanceUnit(10, meters)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForCentimeterAndKilometer(t *testing.T) {
	d1, _ := NewDistanceUnit(1000000, centimeters)
	d2, _ := NewDistanceUnit(10, kilometers)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestEqualitForKilometerAndCentimeter(t *testing.T) {
	d2, _ := NewDistanceUnit(1000000, centimeters)
	d1, _ := NewDistanceUnit(10, kilometers)
	if d1.equals(d2) != true {
		t.Errorf("want both distance equal")
	}
}

func TestAddDistance(t *testing.T) {
	d1, _ := NewDistanceUnit(1000, meters)
	d2, _ := NewDistanceUnit(1000, meters)
	d3 := d1.Add(d2)
	expectedDistance := distance{value: 2000, unit: meters}
	if d3.equals(&expectedDistance) != true {
		t.Errorf("Wanted  %v but got  %v", expectedDistance.toString(), (*d3).toString())

	}

}

func TestAddDistance1(t *testing.T) {
	d1, _ := NewDistanceUnit(100, meters)
	d2, _ := NewDistanceUnit(1, kilometers)
	d3 := d1.Add(d2)
	expectedDistance := distance{value: 1100, unit: meters}
	if d3.equals(&expectedDistance) != true {
		t.Errorf("Wanted  %v but got  %v", expectedDistance.toString(), (*d3).toString())
	}

}

func TestAddDistance2(t *testing.T) {
	d1, _ := NewDistanceUnit(100, kilometers)
	d2, _ := NewDistanceUnit(1, meters)
	d3 := d1.Add(d2)
	expectedDistance := distance{value: 100.001, unit: kilometers}
	if d3.equals(&expectedDistance) != true {
		t.Errorf("Wanted  %v but got  %v", expectedDistance.toString(), (*d3).toString())
	}

}
