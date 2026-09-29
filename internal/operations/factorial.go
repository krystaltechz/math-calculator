package operations

// через рекурсию проще
// почему факториал float?

func Factorial(n int64) int64 {
	if n == 1 || n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}
