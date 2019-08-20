// Testing stuff (it might not be able to run 😅since there's too much used vars)
package main

import (
	"fmt"
	"os"
)

// It'll be available to all the files in this pkg
const PI = 3.14

func main() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("[freeze] %g°F or %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("[boilin] %g°F or %g°C\n", boilingF, fToC(boilingF))

	// Different init val for different type (e.g. "", 0, nil, false etc.)
	var InitialValue int

	// Declare (& initialize) a set of variables
	var one, two, three int
	var One, Two, Three = 1, 2, 3

	// Initialize by calling functions
	var file, err = os.Open("hello")

	// Do replace the 2nd line with `=` (all assignment no declaration)
	a, b := 10, 20
	a, b := 20, 30
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
