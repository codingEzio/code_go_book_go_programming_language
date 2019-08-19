// Print the text of each line from stdin/File that appears more than once
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		This solution wasn't written by me, here's the link:
		https://github.com/torbiak/gopl/blob/master/ex1.4/main.go
	*/

	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines_(os.Stdin, counts, foundIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines_(f, counts, foundIn)

			_ = f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, line, foundIn[line])
		}
	}
}

func in(needle string, strings []string) bool {
	for _, s := range strings {
		if needle == s {
			return true
		}
	}
	return false
}

func countLines_(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()
		counts[line] = counts[line] + 1

		if !in(f.Name(), foundIn[line]) {
			foundIn[line] = append(foundIn[line], f.Name())
		}
	}

	// Optional
	// fmt.Println("counts =>", counts)
}
