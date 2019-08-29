package main

import (
	"fmt"
)

func main() {
	// x -> x++ -> x*x
	f := squares()

	fmt.Println(f(), &f) // 0 -> .. -> 1
	fmt.Println(f(), &f) // 1 -> .. -> 4
	fmt.Println(f(), &f) // 2 -> .. -> 9
}

func squares() func() int {
	// It'll only be initialized once
	fmt.Println("before init x")
	var x int
	fmt.Println("after  init x")

	// The `x` still exists when the 2nd/else call happens
	return func() int {
		fmt.Println("before incr x")
		x++
		fmt.Println("after  incr x")
		return x * x
	}
}
