package distanceUnit

type distanceUnit struct {
	distance int
	unit     string
}

func NewDistanceUnit(distance int, unit string) (*distanceUnit, error) {
	return &distanceUnit{distance: distance, unit: unit}, nil
}
