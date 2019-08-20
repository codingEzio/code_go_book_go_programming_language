package main

import (
	"fmt"
	"os"
)

func main() {
	// Nice 😏
	x, y := 10, 20
	x, y = y, x

	fmt.Println(gcd(10, 20))
	fmt.Println(fib(10))

	f, err := os.Open("./foo.txt")
	fmt.Println(f, err)

	// Two more examples
	// >> _, ok = x.(T)
	// >> _, err = io.Copy(dst, src)

	languages := []string{"Golang", "Python", "JavaScript"}

	var langs [3]string
	langs[0] = "Golang"
	langs[1] = "Python"
	langs[2] = "JavaScript"

	fmt.Println(languages)
	fmt.Println(langs)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
