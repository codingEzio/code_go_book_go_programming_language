package main

import (
	"flag"
	"fmt"
	"time"
)

// All of these dynamic values are stored in the `Value` interface,
// the default val is represented as a string (String() string, Set(string) error)
var period = flag.Duration("period", 1*time.Second, "sleep period")

// WHERE	go run THISCODE -FLAG FLAG_VAL
// HOW		go run THISCODE -period [10ms, 200s, 30m, 40h]
func main() {
	flag.Parse()

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)

	fmt.Println()
}
