package gcd

import (
	"errors"
)

// Calculate greater common divisor (GCD) of specified numbers.
func CalculateGcd(numbers ...uint) (computedResult uint, error error) {
	if numbers == nil {
		return 0, errors.New("Numbers parameter cannot be null.")
	}

	numsCount := len(numbers)
	if numsCount == 1 {
		return numbers[0], nil
	}

	resp := numbers[0]

	for i := 1; i < numsCount; i++ {
		resp = calculateGcdEuclid(resp, numbers[i])
	}

	return resp, nil
}

// https://siongui.github.io/2017/05/14/go-gcd-via-euclidean-algorithm/
func calculateGcdEuclid(a uint, b uint) uint {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
