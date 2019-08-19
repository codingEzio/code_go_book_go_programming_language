// Detailed examples of 'composition literal' (technique)
package main

import (
	"fmt"
)

func main() {
	/*
		Materials
		- https://medium.com/golangspec/composite-literals-in-go-10dc62eec06a
		- https://medium.com/@fishnux/golang-what-the-heck-is-a-composite-literal-6dd194480a75
	*/

	// #TODO finish the examples of the 2nd article (later)
	manualCreateOrCompLiteral()
	comparisonBetweenPythonAndGo()
	keyCanOnlyBeLitConstOrExpression()
}

func manualCreateOrCompLiteral() {
	var myArr [5]int
	for i := 0; i < len(myArr); i++ {
		myArr[i] = 10
	}

	myArrCompLitShort := [5]int{20, 20, 20, 20, 20}
	var myArrCompLitLong [5]int = [5]int{20, 20, 20, 20, 20}

	fmt.Println(myArr)
	fmt.Println(myArrCompLitShort)
	fmt.Println(myArrCompLitLong)
}

func comparisonBetweenPythonAndGo() {

	// elements = [1, 2, 3, 4]
	elements := []int{1, 2, 3, 4}

	// thing = { KEY: VAL, KEY: VAL .. }
	type Thing struct {
		name  string
		gen   int
		model string
	}
	thing := Thing{}
	thing = Thing{name: "iPhone", gen: 11, model: "128GB"}

	fmt.Println(elements)
	fmt.Println(thing)
}

func keyCanOnlyBeLitConstOrExpression() {
	// f := func() int { return 1 }
	// elem_wrong := []string{0: "zero", f(): "one"}

	elem_right := []string{0: "zero", 3 / 2: "one"}
	fmt.Println(elem_right)
}
