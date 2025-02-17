package main

import (
	"Pluralsight-Deep-Dive-Into-Go-Functions/simplemath"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
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
	raiseam(&set)

	func() {
		fmt.Println("This service calls itself.")
	}()
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
