package main

import (
	"fmt"
)

func main() {
	describe("Hello World")
	describe(true)
	describe(55)

	describe(struct {
		name string
	}{
		name: "Natasha",
	})
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
