package main

import "fmt"

// The type of it must be defined in here (global var I assume?)
var theAnsToUniverse int

func init() {
	theAnsToUniverse = 42
}

func init() {
	// One single file can have two `init`, don't do this if it's not a big func
	theAnsToUniverse = 42 * 42
}

func main() {
	fmt.Println(theAnsToUniverse)
}
