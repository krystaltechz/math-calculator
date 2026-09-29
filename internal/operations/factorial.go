package operations

import "math"

func Factorial(n float64) (int64, error) {
	if n < 0 || n != math.Trunc(n) {
		return 0, ErrNotInteger
	}
	if n > 20 {
		return 0, ErrTooBig
	}
	r := int64(1)
	for i := int64(2); i <= int64(n); i++ {
		r *= i
	}
	return r, nil
}
