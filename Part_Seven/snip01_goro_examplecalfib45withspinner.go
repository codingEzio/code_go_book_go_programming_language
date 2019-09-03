package main

import (
	"fmt"
	"time"
)

func main() {
	// The `spinner` and the `fib` are running together (concurrently)
	go spinner(100 * time.Millisecond)

	fibN := fib(45)
	fmt.Printf("\rFibonacci(%d) = %d\n", 45, fibN)
}

// Display an animated textual "spinner"
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
