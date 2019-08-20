package main

import "fmt"

var global *int

func main() {
	// The escaped `x`
	// * the `x` havn't been recycled because of function `ff()`
	// * it's still bad even though `global` doesn't have a value (`nil`)
	fmt.Println("global =", *global)

	// The return value is the same as `return 1`
	// it was garbage-collected without any problems 😄
	fmt.Println("y =", y)
}

func ff() {
	var x int
	x = 1
	global = &x
}

func gg() {
	y := new(int)
	*y = 1
}
