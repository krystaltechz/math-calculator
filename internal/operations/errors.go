package operations

import "errors"

var (
	ErrDivByZero      = errors.New("деление на ноль")
	ErrNegativeSqrt   = errors.New("корень из отрицательного числа")
	ErrLogNonPositive = errors.New("логарифм только для положительных чисел")
	ErrNotInteger     = errors.New("требуется целое неотрицательное число")
	ErrTooBig         = errors.New("слишком большое число (макс. 20!)")
)
