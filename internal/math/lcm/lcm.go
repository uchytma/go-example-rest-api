package lcm

import (
	"errors"

	"github.com/uchytma/go-example-rest-api/internal/math/gcd"
)

// Calculate least common multiple (LCM) of specified numbers.
// More about used alghorithm: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func CalculateLcm(numbers ...uint) (computedResult uint, error error) {
	if numbers == nil {
		return 0, errors.New("Numbers parameter cannot be null.")
	}

	numsCount := len(numbers)
	if numsCount == 1 {
		return numbers[0], nil
	}

	resp := numbers[0]

	for i := 1; i < numsCount; i++ {
		var gcdResponse, err = gcd.CalculateGcd(resp, numbers[i])
		if err != nil {
			return 0, err
		}
		resp = resp * numbers[i] / gcdResponse
	}

	return resp, nil
}
