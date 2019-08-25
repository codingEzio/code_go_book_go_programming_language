// Counts the occurrences of each distinct UnicodeCodePoint in its input :)
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // e.g. ❖: 1
	var utflen [utf8.UTFMax + 1]int // count of lens of UTF8 encodings
	invalid := 0                    // count of invalid characters

	in := bufio.NewReader(os.Stdin)
	for {
		rne, nby, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if rne == unicode.ReplacementChar && nby == 1 {
			invalid++
			continue
		}
		counts[rne]++
		utflen[nby]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
