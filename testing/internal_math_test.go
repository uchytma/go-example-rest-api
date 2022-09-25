package testing

import (
	"testing"

	"github.com/uchytma/go-example-rest-api/internal/math"
)

func testMathFunction(t *testing.T, mathFunction math.MathFunctionMultipleInputs, input []uint, expectedResult uint, shouldError bool) {
	got, err := mathFunction.Calculate(input...)
	if shouldError {
		if err == nil {
			t.Errorf("err should be nil")
		}
		return
	}
	if err != nil {
		t.Errorf("err is not nil")
		return
	}
	if got != expectedResult {
		t.Errorf("input data: %v, result: %d; should be: %d", input, got, expectedResult)
		return
	}
}

func TestGcd(t *testing.T) {
	var m = math.Gcd{}
	testMathFunction(t, m, []uint{0}, 0, false)
	testMathFunction(t, m, []uint{1}, 1, false)
	testMathFunction(t, m, []uint{5, 10, 100}, 5, false)
	testMathFunction(t, m, []uint{100, 5, 10}, 5, false)
	testMathFunction(t, m, []uint{33, 58145, 145123}, 1, false)
	testMathFunction(t, m, []uint{33, 66, 11}, 11, false)
	testMathFunction(t, m, nil, 0, true)
	testMathFunction(t, m, []uint{}, 0, true)
}

func TestLcm(t *testing.T) {
	var m = math.Lcm{}
	testMathFunction(t, m, []uint{0}, 0, false)
	testMathFunction(t, m, []uint{1}, 1, false)
	testMathFunction(t, m, []uint{5, 10, 100}, 100, false)
	testMathFunction(t, m, []uint{33, 58145, 145123}, 25314530505, false)
	testMathFunction(t, m, []uint{33, 66, 11}, 66, false)
	testMathFunction(t, m, nil, 0, true)
	testMathFunction(t, m, []uint{}, 0, true)
}
