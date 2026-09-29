package fractions

import "errors"

func Add(a, b Fraction) (Fraction, error) {
	return New(a.Num*b.Den+b.Num*a.Den, a.Den*b.Den)
}

func Sub(a, b Fraction) (Fraction, error) {
	return New(a.Num*b.Den-b.Num*a.Den, a.Den*b.Den)
}

func Mul(a, b Fraction) (Fraction, error) {
	return New(a.Num*b.Num, a.Den*b.Den)
}

func Div(a, b Fraction) (Fraction, error) {
	if b.Num == 0 {
		return Fraction{}, errors.New("деление на ноль")
	}
	return New(a.Num*b.Den, a.Den*b.Num)
}
