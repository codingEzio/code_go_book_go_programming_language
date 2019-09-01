package main

import (
	"fmt"
	"strconv"
)

type MyStringer interface {
	String() string
}

type Temp int

type Pointtt struct {
	x, y int
}

func main() {
	var z MyStringer
	fmt.Printf("%v %T\n", z, z) // nil, nil

	z = Temp(20)
	fmt.Printf("%v %T\n", z, z) // 20°C main.Temp

	z = &Pointtt{1, 2}
	fmt.Printf("%v %T\n", z, z) // (1,2) *main.Pointtt

	z = (*Pointtt)(nil)
	fmt.Printf("%v %T\n", z, z) // <nil> *main.Pointtt
}

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + "°C"
}

func (p *Pointtt) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
