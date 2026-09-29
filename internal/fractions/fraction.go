package fractions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Fraction struct {
	Num int64
	Den int64
}

func New(num, den int64) (Fraction, error) {
	if den == 0 {
		return Fraction{}, errors.New("знаменатель не может быть 0")
	}
	if den < 0 {
		num, den = -num, -den
	}
	g := gcd(num, den)
	num, den = num/g, den/g

	return Fraction{Num: num, Den: den}, nil
}

// Можно проще
// a%b сам поменяет местами младший и старший аргумент

func gcd(a, b int64) int64 {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (f Fraction) String() string {
	if f.Den == 1 {
		return strconv.FormatInt(f.Num, 10)
	}
	return fmt.Sprintf("%d/%d", f.Num, f.Den)
}

func Parse(s string) (Fraction, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return Fraction{}, errors.New("пустая строка")
	}
	if strings.Contains(s, "/") {
		p := strings.SplitN(s, "/", 2)
		n, e1 := strconv.ParseInt(strings.TrimSpace(p[0]), 10, 64)
		d, e2 := strconv.ParseInt(strings.TrimSpace(p[1]), 10, 64)
		if e1 != nil || e2 != nil {
			return Fraction{}, errors.New("неверный формат дроби")
		}
		return New(n, d)
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return Fraction{}, errors.New("неверный формат числа")
	}
	return New(n, 1)
}
