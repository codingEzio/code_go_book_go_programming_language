package main

import (
	"fmt"
)

func main() {
	/*
		Expected errors 	use `error`
		Unexpected errors	use `panic`
	*/

	f(3)
}

func f(x int) {
	// calls three times in one row
	// * there's defer statement (so it won't exec immediately)
	// * there's recur calls 	 (calls another `f` before the `defer` being called)
	// * the `defer`'ll be called before "the end of the world" (aka. panic) (IMPORTANT)
	fmt.Printf("f(%d)\n", x+0/x)

	// Since "we" "doesn't have time" to call this (cuz the recursion)
	// those calls're kinda like a 'stack', the most recent one'll be called first (REV)
	defer fmt.Printf("defer %d\n", x)

	// 3 2 1 (Printf)
	// 1 2 3 (defer ..)
	// panic (integer divide by zero)
	f(x - 1)
}
