package main

import "fmt"

func main() {
	vals := [...]int{1, 2, 3}

	PrintElemUsingArrayAsParam(vals)    // param's size are fixed, passing array directly
	PrintElemUsingSliceAsParam(vals[:]) // param's size is flexible, passing slices (VAL[IDX])
}

func GoDoesntWorkThatWay() {
	/*
		Go treats arrays like any other type
		- but its behavior is different from languages that
		- implicitly pass arrays by reference (in Go, it copies everything)
	*/
}

func PrintElemUsingArrayAsParam(values [3]int) {
	for _, elem := range values {
		_ = fmt.Sprint(elem)
	}
}

func PrintElemUsingSliceAsParam(values []int) {
	for _, elem := range values {
		_ = fmt.Sprint(elem)
	}
}
