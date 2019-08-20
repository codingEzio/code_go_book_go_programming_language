package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	p2 := &x

	fmt.Printf("value: %d, address: %x\n", *p, p)
	fmt.Println("Pointers p equals to p2:", p == p2, *p == *p2)

	*p = 1024
	fmt.Println("Variable x changed from 1 to", x)

	// y == z		initial value (for int it's 0)
	// &y == &z		different address indeed
	// &y &z nil	won't be nil unless it doesn't point to a variable 😢
	var y, z int
	fmt.Println(y == z, &y == &z, &y == nil, &z == nil)

	// A few things need to notice
	// * the addr of `v` will still exist even after the call has returned
	// * the return value will change every time when u call it (huh 😅)
	var whoa = f()
	fmt.Printf("\nType of `whoa` is %T\n", whoa)
	fmt.Printf("whoa = %x, f() == f() = %t\n", whoa, f() == f())

	// The final value of `inc` would be 3
	inc := 1
	incr(&inc)
	fmt.Println(incr(&inc))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
