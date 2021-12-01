package utils

import (
	"errors"
)

// Abs returns the absolute value of an integer.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Mod calculates non-negative remainder of dividing two integers.
// It returns -1 and an error if the modulus is zero.
// Otherwise it returns the remainder and nil.
func Mod(d, m int) (int, error) {
	if m == 0 {
		return -1, errors.New("cannot divide by 0")
	}
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		res += m
	}
	return Abs(res), nil
}
