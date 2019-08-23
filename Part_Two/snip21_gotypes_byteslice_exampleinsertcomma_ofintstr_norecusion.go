// Insert commas in a non-negative decimal integer string (no recursion)
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(CommaForIntegerNoRecursion("200"))
	fmt.Println(CommaForIntegerNoRecursion("2000"))
	fmt.Println(CommaForIntegerNoRecursion("2000000"))
}

func CommaForIntegerNoRecursion(str string) string {
	byt := &bytes.Buffer{}

	// Write the 1st group of up to 3 digits (copied from someone's exercise code)
	pre := len(str) % 3
	if pre == 0 {
		pre = 3
	}
	byt.WriteString(str[:pre])

	for idx := pre; idx < len(str); idx += 3 {
		byt.WriteByte(',')
		byt.WriteString(str[idx : idx+3])
	}

	return byt.String()
}
