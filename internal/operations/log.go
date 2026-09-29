package operations

import "math"

func Ln(a float64) (float64, error) {
	if a <= 0 {
		return 0, ErrLogNonPositive
	}
	return math.Log(a), nil
}

func Log10(a float64) (float64, error) {
	if a <= 0 {
		return 0, ErrLogNonPositive
	}
	return math.Log10(a), nil
}
