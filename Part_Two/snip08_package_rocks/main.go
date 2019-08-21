// Convert numeric arguments to Celsius and Fahrenheit
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/code_go_book_go_programming_language/Part_Two/snip08_package_rocks/conv_lib"
)

/*
	Let's talk about package (in a practical perspective)
	* Put the "main entry" and rest of the code in a separate location
	* The package name for "main entry" CAN BE `package main`
	* The package name could be anything (`haha` => `haha.Celsius(100)`)
	* The package path for import is based on `$GOPATH` (under src, `src` excluded)
	* Any code except the `main` one should be put into sub-folders (used for "importing")
*/
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64) // float64
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
			os.Exit(1)
		}

		f := conv_lib.Fahrenheit(t)
		c := conv_lib.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, conv_lib.FToC(f), c, conv_lib.CToF(c))
	}
}
