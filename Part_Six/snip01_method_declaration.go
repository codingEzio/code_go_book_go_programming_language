package main

import (
	"fmt"
	"math"
)

// Note most of the examples here demonstrate by using a struct
// actually, it can also works on non-struct types as well
type Point struct {
	X, Y float64
}
type Path []Point

// You can also declare methods on non-struct types :)
type MyFloat float64

func main() {
	/*
		Terms
		- receiver	the extra parameter	(core diff between func & method)
		- selector	STRUCT.METHOD 		(call: p.Distance, access: p.X)

		Also, some quotes from official doc/tutorial
		* "Remember, a method is just a function with a receiver argument."
	*/

	p := Point{1, 2}
	q := Point{3, 4}

	fmt.Println(Distance(p, q))

	fmt.Println(p.Distance(q))
	fmt.Println(q.Distance(p)) // it's the same anyway

	perimeter := Path{
		{0, 0}, // origin
		{0, 3}, // 3 (↑️Y)
		{4, 0}, // 5 (↘︎X)
		{0, 0}, // 4 (←X)
	}
	fmt.Println(perimeter.Distance())

	fmt.Println(MyFloat(2).Abs())
	fmt.Println(MyFloat(-2).Abs())
}

// There's no conflict between the two declarations of funcs called `Distance` here
// * this one is a "pkg-level" func	->	snip01_method_declaration.Distance
// * that one is a "Point's method"	->	Point.Distance
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// The extra param attaches the func to the type of that param (e.g. p Point)
// it was called as "receiver" (being used just like any other params; no `this`, `self` BTW)
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// For methods, each type has its own name space
// so we can use the name `Distance` for other methods (like this one)
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i]) // calls 'Point.Distance'
		}
	}
	return sum
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
