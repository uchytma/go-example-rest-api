package math

import (
	"errors"
)

type Gcd struct{}

// Calculate greater common divisor (GCD) of specified numbers.
func (gcd Gcd) Calculate(numbers ...uint) (computedResult uint, error error) {
	if numbers == nil {
		return 0, errors.New("Numbers parameter cannot be null.")
	}
	numsCount := len(numbers)
	if numsCount == 0 {
		return 0, errors.New("At least one number is required.")
	}

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
