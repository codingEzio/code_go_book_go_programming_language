// Rotates a slice of ints by one position to the left.
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	fmt.Println(s) // 1,2,3,4

	rotate_ints(s)
	fmt.Println(s) // 4,1,2,3
}

func rotate_ints(ints []int) {
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
}
