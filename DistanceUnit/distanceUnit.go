package distanceUnit

import (
	"errors"
)

type Unit int

const (
	meters Unit = iota
	kilometers
	centimeters
)

type distanceUnit struct {
	distance int
	unit     Unit
}

func NewDistanceUnit(distance int, unit Unit) (*distanceUnit, error) {
	if distance < 0 {
		return nil, errors.New("distance cannot be negative")
	}
	return &distanceUnit{distance: distance, unit: unit}, nil
}

func (d1 *distanceUnit) EqualityOfDistance(d2 *distanceUnit) bool {
	// if d1.distance == d2.distance && d1.unit == d2.unit {
	// 	return true
	// } else if d1.unit == kilometers {
	// 	d1 = d1.KilometerToMeter()
	// 	if d2.unit == meters {
	// 		if d1.distance == d2.distance && d1.unit == d2.unit {
	// 			return true
	// 		}
	// 	} else if d2.unit == centimeters {
	// 		d2 = d2.CentimeterToMeter()
	// 		if d1.distance == d2.distance && d1.unit == d2.unit {
	// 			return true
	// 		}
	// 	}

	// } else if d2.unit == kilometers {
	// 	d2 = d2.KilometerToMeter()
	// 	if d1.unit == meters {
	// 		if d1.distance == d2.distance && d1.unit == d2.unit {
	// 			return true
	// 		}
	// 	} else if d1.unit == centimeters {
	// 		d1 = d1.CentimeterToMeter()
	// 		if d1.distance == d2.distance && d1.unit == d2.unit {
	// 			return true
	// 		}
	// 	}
	// } else if d1.unit == centimeters {
	// 	d1 = d1.CentimeterToMeter()
	// 	if d2.unit == meters {
	// 		if d1.distance == d2.distance && d1.unit == d2.unit {
	// 			return true
	// 		}
	// 	}
	// } else if d1.unit == meters {
	// 	if d2.unit == centimeters {
	// 		d2 = d2.CentimeterToMeter()
	// 		if d1.distance == d2.distance && d1.unit == d2.unit {
	// 			return true
	// 		}
	// 	}
	// }
	// return false
	return d1.Inmeter().distance == d2.Inmeter().distance
}

func (d *distanceUnit) Inmeter() *distanceUnit {
	if d.unit == centimeters {
		d = &distanceUnit{distance: d.distance / 100, unit: meters}
	} else if d.unit == kilometers {
		d = &distanceUnit{distance: d.distance * 1000, unit: meters}
	}
	return d
}

func (d *distanceUnit) KilometerToMeter() *distanceUnit {
	return &distanceUnit{distance: (d.distance * 1000), unit: meters}
}

func (d *distanceUnit) CentimeterToMeter() *distanceUnit {
	return &distanceUnit{distance: (d.distance / 100), unit: meters}
}
