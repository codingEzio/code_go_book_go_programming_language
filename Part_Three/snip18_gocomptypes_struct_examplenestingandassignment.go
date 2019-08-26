package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point  // anonymous field
	Radius int
}
type Wheel struct {
	Circle // anonymous field
	Spokes int
}

func main() {
	// There are still more to talk about, such as, methods 😅.
	nestedStructOldWay()
	nestedStructNewWay()
	nestedStructHowToAssign()
}

func nestedStructOldWay() {
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Center Point
		Radius int
	}
	type Wheel struct {
		Circle Circle
		Spokes int
	}

	// The ordinary & kinda intuitive way
	var w Wheel
	w.Circle.Center.X = 10
	w.Circle.Center.Y = 10
	w.Circle.Radius = 5
	w.Spokes = 20
}

func nestedStructNewWay() {
	var w Wheel
	w.X = 8      // equivalent to `w.Circle.Point.X`
	w.Y = 8      // equivalent to `w.Circle.Point.Y`
	w.Radius = 5 // equivalent to `w.Circle.Radius`
	w.Spokes = 20
}

func nestedStructHowToAssign() {
	/*
		You CAN NOT do something like this (so-called "shorthand")
		>> var w = Wheel{8, 8, 5, 20}
		>> var w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}
	*/

	// They're the same
	var y = Wheel{Circle{Point{8, 8}, 5}, 20}
	var z = Wheel{
		Circle: Circle{
			Point:  Point{8, 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	// %#v => Go-syntax format (field+val)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%#v\n", z)
}
