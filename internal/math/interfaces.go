package math

type MathFunctionMultipleInputs interface {
	Calculate(numbers ...uint) (computedResult uint, error error)
}
