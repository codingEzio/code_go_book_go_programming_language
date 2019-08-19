// Print the text of each line from stdin that appears more than once
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		Call `ctrl + D` to end this program (& get the result)
	*/

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// Each call to `input.Scan` reads the next line and remove
	// the newline character from the end (get text by `input.Text()`)
	for input.Scan() {
		// counts[input.Text()]++
		line := input.Text()
		counts[line] = counts[line] + 1
	}

	// Actually, it also catches "\t" and "\n"
	fmt.Println("counts =>", counts)

	// line :: string
	// n	:: int		how many time it appears
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
