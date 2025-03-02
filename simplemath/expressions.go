package simplemath

import (
	"errors"
	"math"
)

func Divide(o1, o2 float64) (answer float64, err error) {
	if o2 == 0 {
		answer, err = math.NaN(), errors.New("Cannot divide by zero")
	}

	answer, err = o1/o2, nil
	return
}

func Add(o1, o2 float64) float64 {
	return o1 + o2
}

func Sum(values ...float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total
}

func Subtract(o1, o2 float64) float64 {
	return o1 - o2
}

func Multiply(o1, o2 float64) float64 {
	return o1 * o2
}
