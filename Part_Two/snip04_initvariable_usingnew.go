package main

import "fmt"

func main() {
	// Return a pointer to a newly allocated zero val of that type
	p := new(int)
	var q int

	fmt.Printf("p is %T / q is %T\n", p, q)
	fmt.Printf("p = %d / q = %d\n", *p, q)
	fmt.Printf("(address) p != &q -> %t\n", p != &q)

	// There's exceptions though (addr == addr)
	// * specifically, for types like struct or array which
	// * carries no information and is therefore of size zero
	x := new(int)
	y := new(int)
	fmt.Printf("(address) x != y -> %t\n", x != y)
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
