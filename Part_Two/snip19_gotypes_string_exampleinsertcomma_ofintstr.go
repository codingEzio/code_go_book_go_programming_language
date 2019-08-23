// Insert commas in a non-negative decimal integer string
package main

import "fmt"

func main() {
	fmt.Println(CommaForInteger("200"))
	fmt.Println(CommaForInteger("2000"))
	fmt.Println(CommaForInteger("2000000") + "\n")
}

func CommaForInteger(str string) string {
	ln := len(str)
	if ln <= 3 {
		return str
	}

	// Keyword: recursion, concatenation
	return CommaForInteger(str[:ln-3]) + "," + str[ln-3:]
}
