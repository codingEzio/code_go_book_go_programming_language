package main

import (
	"fmt"
)

func main() {
	// defAndModifyByRefAndAddr()
	// structLiteral()
	// structAsArguments()
	structComparableAndAsKey()
}

func defAndModifyByRefAndAddr() {
	type Human struct {
		Name, Gender string
		Age          int
	}

	var nat Human
	nat.Name = "Natalie"
	nat.Gender = "Female"
	nat.Age = 30

	bbCopy := &nat // same as `var X *Human = &nat`
	bbCopy.Age = 21

	bbCopyName := &nat.Name
	*bbCopyName = *bbCopyName + " Wood"

	fmt.Println(nat)
}

func structLiteral() {
	type Point struct{ X, Y int }

	// This form requires that a value be specified for every field,
	// and must be in the right order (=> fragile and mental-effort+)
	// this should be only used for which there's an obvious field
	// ordering convention, like `Color.RGBA{red, green, blue, alpha}`
	p1 := Point{1, 2}

	// In this form
	// * you can set just some of them or all of them
	// * if a filed is omitted in this, it's set to zero-val for its type
	// * two forms cannot be mixed ( `Point{X: 0, Y}` (filed Y without val) )
	p2 := Point{Y: 2, X: 1}
	p3 := Point{
		X: 1,
	}

	_ = fmt.Sprint(p1, p2, p3)
}

func structAsArguments() {
	/*
		For efficiency, larger struct types are usually
		> passed to or returned from function indirectly using a pointer,
		> and this is REQUIRED if the function must modify its argument(!!)

		Since in a call-by-val languages like Go
		- the called function receives only a copy of an argument
		- not a reference to the original argument (so.. use a pointer!)
	*/

	s := Scale(OrigPoint{1, 2}, 5)
	sp := ScalePassPointer(&OrigPoint{1, 2}, 5)
	spc := OrigEmployee{
		Name:   "Joe",
		Salary: 7100,
	}

	fmt.Println(s)
	fmt.Println(sp)
	fmt.Println(spc)

	// 7100 => 8165
	AwardRaisePassPointerAndModify(&spc)
	fmt.Println(spc)

	// One last example
	// pp := &OrigPoint{1, 2}						// method 1
	// pp := new(OrigPoint)  *pp = OrigPoint{1, 2}	// method 2
}

func structComparableAndAsKey() {
	type P struct {
		X, Y int
	}

	pt1 := P{1, 2}
	pt2 := P{Y: 2, X: 1}
	q := P{2, 1}

	_ = fmt.Sprint(pt1 == pt2, pt1 != q)         // true, true
	_ = fmt.Sprint(pt1.X == pt2.X, pt1.X != q.X) // true, true

	type address struct {
		hostname string
		port     int
	}

	// As the key of a map
	hits := make(map[address]int)
	hits[address{"golang", 443}]++
	hits[address{"golang", 443}]++

	fmt.Println(hits) // => 2
}

// ----- ----- ----- Helper ----- ----- -----

type OrigPoint struct {
	X, Y int
}

type OrigEmployee struct {
	Name   string
	Salary int
}

func Scale(p OrigPoint, factor int) OrigPoint {
	return OrigPoint{p.X * factor, p.Y * factor}
}

func ScalePassPointer(p *OrigPoint, factor int) OrigPoint {
	return OrigPoint{p.X * factor, p.Y * factor}
}

func AwardRaisePassPointerAndModify(e *OrigEmployee) {
	e.Salary = e.Salary * 115 / 100
}
