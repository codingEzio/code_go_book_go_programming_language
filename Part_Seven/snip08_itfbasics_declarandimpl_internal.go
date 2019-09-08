package main

import (
	"fmt"
)

type Tester interface {
	Test()
}

type MyFloat float64

func main() {
	var tester Tester
	mf := MyFloat(90.0)
	tester = mf

	describe(tester)
	tester.Test()
}

// Now we've implemented the `Test` in the interface,
// at the same time, it got its own (concrete) type, namely `float64` (<= MyFloat)
func (mf MyFloat) Test() {
	fmt.Println(mf)
}

func describe(tester Tester) {
	fmt.Printf("Interface type: %T, value: %v\n", tester, tester)
}
