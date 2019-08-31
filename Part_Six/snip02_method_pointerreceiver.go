package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)

	fmt.Println(v.Abs())    // reason1: using pointer receiver, you can modify it
	fmt.Println((&v).Abs()) // reason2: avoid copying the val on each call, therefore efficient

	// more on this
	// https://go101.org/article/nil.html
	nilIsAValidReceiverValue()
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f // you won't be able to change the val of inst(sort of)
	v.Y = v.Y * f // unless you pass into a Pointer (receiver) of the method
}

func nilIsAValidReceiverValue() {
	/*
		Since I can't wrap my head around this, I'll just leave this here

		type IntList struct {
			Value int
			Tail  *IntList
		}

		func (list *IntList) Sum() int {
			if list == nil {
				return 0
			}
			return list.Value + list.Tail.Sum()
		}
	*/
}
