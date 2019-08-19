// Print its cmd arguments (short description goes here by convention)
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		Symbol `:=`
		>> var score int = 74	(var VAR[,VAR] TYPE [= VAL])
		>> score 		:= 74	(it's the same)
	*/

	// Method 1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "_"
	}
	fmt.Println(s)

	// Method 2
	fmt.Println(strings.Join(os.Args[1:], "_"))

	// Method 3
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = "_"
	}
	fmt.Println(s2)

	// Method 4
	s3, sep3 := "", ""
	for idx, arg := range os.Args[1:] {
		s3 += sep3 + arg
		sep3 = "_"

		fmt.Println(idx, s3)
	}

	// Exercises
	fmt.Println("---- Exercises ----")
	fmt.Println(os.Args[0])

}
