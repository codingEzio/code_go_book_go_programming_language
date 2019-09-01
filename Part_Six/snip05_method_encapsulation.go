package main

import (
	"fmt"
)

type Counter struct {
	n int
}

/*
	encapsulated: 封裝, also called as 'information hiding'

	To encapsulate an object, you MUST make it a struct.
*/
func main() {
	cnt := Counter{1}

	fmt.Println(cnt.N()) // 1

	cnt.Increment()
	cnt.Increment()
	cnt.Increment()
	fmt.Println(cnt.N()) // 4

	cnt.Reset()
	fmt.Println(cnt.N()) // 0
}

func (c *Counter) N() int {
	return c.n
}

func (c *Counter) Increment() {
	c.n++
}
func (c *Counter) Reset() {
	c.n = 0
}
