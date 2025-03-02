package main

import (
	"Pluralsight-Deep-Dive-Into-Go-Functions/simplemath"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strings"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	answer, err := simplemath.Divide(6, 2)
	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}

	numbers := []float64{12.2, 14, 16, 22.4}
	total := simplemath.Sum(numbers...)
	fmt.Printf("total of sum: %f\n", total)

	sv := simplemath.NewSemanticVersion(23, 10, 5)
	sv2 := &simplemath.SemanticVersion{}
	sv2.IncrementMajor()
	sv2.IncrementMajor()
	sv2.IncrementMinor()
	sv2.IncrementPatch()
	sv2.IncrementPatch()
	sv2.IncrementPatch()
	fmt.Println(sv2.String())
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	fmt.Println(sv.String())

	var tripper http.RoundTripper = &RoundTripCounter{}
	// had to use &RoundTripCounter cause the method is a pointer based receiver
	r, _ := http.NewRequest(http.MethodGet, "https://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)

	// Anonymous functions
	func(name string) {
		fmt.Printf("%s first anonymous function\n", name)
		fmt.Println(rand.Float64())
	}("Sulaimon's")

	//setting it to variable name
	set := 1
	raiseam := func(val *int) {
		*val++
		fmt.Printf("The value of val is now : %v\n", *val)
	}
	raiseam(&set)
	raiseam(&set)

	// returning functions from functions
	xadder := mathExpression(AddExpr)
	xsuber := mathExpression(SubtractExpr)
	xmultiplier := mathExpression(MultiplyExpr)
	fmt.Println(xadder(5.0, 3.0))
	fmt.Println(xsuber(5.0, 3.0))
	fmt.Println(xmultiplier(5.0, 3.0))

	// passing functions as parameters
	fmt.Printf("%f\n", double(3.0, 4.0, mathExpression(AddExpr)))

	// stateful functions
	pot := powerOfTwo()
	fmt.Println("The Value of x square is now : ", pot())
	fmt.Println("The Value of x square is now : ", pot())
	fmt.Println("The Value of x square is now : ", pot())
}

// returning functions from functions
func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return nil
	}
}

// passing functions as parameteters
func double(f, f1 float64, MathExpr func(float64, float64) float64) float64 {
	return 2 * MathExpr(f, f1)
}

// stateful functions
func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

// interfaces & pointer based method receivers

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(r *http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"-"`
}
