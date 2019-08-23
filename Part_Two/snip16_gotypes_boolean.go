package main

import "fmt"

func main() {
	/*
		Go does not have implicit conversion from a bool val to num 0,1
		it means u can't directly use `if number { .. }` (doesn't work that way)

		Do note that we're actually TALKING ABOUT `ternary opt` ( i ? 1 : 2 )
		official statement: https://golang.org/doc/faq#Does_Go_have_a_ternary_form
	*/

	i := 0

	if itob(i) {
		fmt.Println("Meh")
	} else {
		fmt.Println("Hello!")
	}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// This one is quite
func itob(i int) bool {
	return i != 0
}
