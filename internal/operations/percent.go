package operations

func PercentOf(p, n float64) float64 {
	return n * p / 100
}

func NumberByPercent(p, part float64) (float64, error) {
	if p == 0 {
		return 0, ErrDivByZero
	}
	return part * 100 / p, nil
}

func WhatPercent(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b * 100, nil
}

func AddPercent(n, p float64) float64 {
	return n * (1 + p/100)
}

func SubPercent(n, p float64) float64 {
	return n * (1 - p/100)
}

func ChangePercent(old, newV float64) (float64, error) {
	if old == 0 {
		return 0, ErrDivByZero
	}
	return (newV - old) / old * 100, nil
}
