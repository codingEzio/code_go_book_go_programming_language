// IntsToString is like `fmt.Sprint(values)` but adds commas (oh)
package main

import (
	"bytes"
	"fmt"
)

/*
	Since this is the first example of "byte slices", time for a brief intro!
	1. The elements of a byte slice can be freely modified.
	  - more efficient than string, possibly
	  - lots of the same API impl_ed for 'byte slice'
	2. Strings can be converted to byte slices and back again ☺️
	  >> s  := "abc"
	  >> b  := []byte(s)	// byte slice
	  >> s2 := string(b)	// string (back again)

*/
func main() {
	fmt.Println(IntsToString([]int{1}))
	fmt.Println(IntsToString([]int{1, 2, 3}))
}

func IntsToString(values []int) string {

	var buf bytes.Buffer

	// Just saying, the `bytes.Buffer` also provides similar funcs
	// like `WriteRune`, `WriteString` and else, do use the right one
	buf.WriteByte('[')

	for idx, v := range values {
		if idx > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}
