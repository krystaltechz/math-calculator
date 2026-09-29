package operations

import "math"

func Pow(a, b float64) float64 { return math.Pow(a, b) }

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSqrt
	}
	return math.Sqrt(a), nil
}

func Abs(a float64) float64 { return math.Abs(a) }
