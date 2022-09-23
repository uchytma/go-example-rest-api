package gcd

import (
	"errors"
)

// Calculate greater common divisor (GCD) of specified numbers.
func CalculateGcd(numbers *[]uint) (uint, error) {
	if numbers == nil {
		return 0, errors.New("Numbers parameter cannot be null.")
	}

	nums := *numbers
	numsCount := len(nums)
	if numsCount == 1 {
		return nums[0], nil
	}

	resp := nums[0]

	for index, _ := range nums {
		if index == numsCount-1 {
			return resp, nil
		}
		resp = calculateGcdEuclid(resp, nums[index+1])
	}

	panic("Gcd algorithm failed.")
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
