package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3)) // []int{1, 2, 3, 4} -> X...

	notTheSameAsSlice()

	// Another example of using `...ARGS`
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}

func sum(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}

	return total
}

func notTheSameAsSlice() {
	f := func(...int) {}
	g := func([]int) {}

	fmt.Printf("%T, %T\n", f, g)
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
