package main

import (
	"fmt"
)

type ByteCounter int

func main() {
	var b ByteCounter
	b.Write([]byte("Hello"))
	fmt.Println(b)

	b = 0
	var name = "dolly"
	fmt.Fprintf(&b, "hello %s", name)
	fmt.Println(b)
}

// Implements interface 'https://golang.org/pkg/io/#Writer'
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
