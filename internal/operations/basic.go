package operations

func Add(a, b float64) float64 { return a + b }
func Sub(a, b float64) float64 { return a - b }
func Mul(a, b float64) float64 { return a * b }

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func Mod(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return float64(int64(a) % int64(b)), nil
}
