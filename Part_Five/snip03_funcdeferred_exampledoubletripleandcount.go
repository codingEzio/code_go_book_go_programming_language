package main

import (
	"fmt"
)

func main() {
	double(20) // not just return but also print (access anno-func val)
	triple(30) // not just return but also print (access anno-func val)

	fmt.Println("Count:", count(30))
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("ARG [%d], RESULT [%d]\n", x, result)
	}()

	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += x
		fmt.Printf("ARG [%d], RESULT [%d]\n", x, result)
	}()

	return x + x
}

// https://blog.learngoprogramming.com/golang-defer-simplified-77d3b2b817ff
func count(i int) (result int) {

	// [1] This `i` was evaluated immediately (i = 30)
	defer func(i int) {
		// [1] Now `i` is 30

		result = result + i
	}(i)

	i = i * 2  // i=60, but this one won't affect the `i` in the anno-func
	result = i // result=60

	// Executes `defer func(..)()`
	// i = 30
	// result = 60

	// 90 (instead of 120(60+60))
	return
}
