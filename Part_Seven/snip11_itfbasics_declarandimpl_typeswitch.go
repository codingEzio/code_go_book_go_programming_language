package main

import (
	"fmt"
)

func main() {
	findType(10)
	findType("Hi")
	findType(true)
	findType(99.99)
}

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("I am an int and my val is %d\n", i.(int)) // i.(int) -> val, bool
	case string:
		fmt.Printf("I am a string and my val is %s\n", i.(string))
	case bool:
		fmt.Printf("I am a boolean and my val is %t\n", i.(bool))
	default:
		fmt.Printf("Unknown type\n")
	}
}
