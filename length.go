package unitconvert

var Meter Unit = Unit{Name: "meter", Type: LengthUnit, System: "metric"}
var Feet Unit = Unit{Name: "feet", Type: LengthUnit, System: "english"}

// Convert length quantities.
//
// Args:
// - v (float64): The value to be converted.
// - from (Unit): The unit of the value.
// - to (Unit): The destination unit.
//
// Returns:
// - The converted value (float64).
// - An error if there is one (error).
func Length(v float64, from, to Unit) (float64, error) {
	if from.Type != LengthUnit || to.Type != LengthUnit {
		return 0, invalidInputTypes
	}

	if from == to {
		return v, nil
	}

	factor, ok := conversionFactors[from][to]
	if !ok {
		return 0, unsupportedDestinationError
	}

	return v * factor, nil
}
