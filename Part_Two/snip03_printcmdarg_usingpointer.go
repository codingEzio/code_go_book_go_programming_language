// Print command-line arguments, options supported (`-n`, `-s SEP`)
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	/*
		Usage explained
		>> go run THIS_FILE --help
		>> go run THIS FILE -n
		>> go run THIS FILE -s _ hello a b c
	*/
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
