// Insert commas in a non-negative floating point string (no recursion)
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Run the other test file by `go test -v`
	fmt.Printf("%s\n", CommaForFloatingNumber("22.00"))
	fmt.Printf("%s\n", CommaForFloatingNumber("1245.43213"))
	fmt.Printf("%s\n", CommaForFloatingNumber("6543210.54321"))
}

func CommaForFloatingNumber(str string) string {
	/*
		For non-native English speakers:
		* mantissa: 尾數, 底數 (3.123->0.123 | 1.2e3 -> 1.2)
	*/

	bt := bytes.Buffer{}

	mantissaStart := 0
	if str[0] == '+' || str[0] == '-' {
		bt.WriteByte(str[0])
		mantissaStart = 1
	}

	mantissaEnd := strings.Index(str, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(str)
	}

	mantissa := str[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3
	if pre > 0 {
		bt.Write([]byte(mantissa[:pre]))
		if len(mantissa) > pre {
			bt.WriteString(",")
		}
	}

	for idx, c := range mantissa[pre:] {
		if idx%3 == 0 && idx != 0 {
			bt.WriteString(",")
		}
		bt.WriteRune(c)
	}

	bt.WriteString(str[mantissaEnd:])
	return bt.String()
}
