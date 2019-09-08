package main

import (
	"fmt"
)

func main() {
	var s interface{} = 50
	assert(s)

	var i interface{} = "unknown"
	assert(i)
}

func assert(i interface{}) {
	v, ok := i.(int) // check whether the underlying type is int or not
	fmt.Println(v, ok)
}
