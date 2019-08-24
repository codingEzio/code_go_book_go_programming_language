package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// ConstantBasics()
	// ConstantIota()
	ConstantUntyped()
}

func ConstantBasics() {
	// convention
	const (
		x_axis float64 = 1.5
		y_axis         = 3.0
	)

	// compile time | error
	const (
		c_dividend = 10 // if you try to do calc like `10 / 0` (0 itself is a const)
		c_divisor  = 0  // the error would detected at compile time (go build)
	)
	var (
		v_dividend int = 10 // but if you try do calc `10/0` while using variables
		v_divisor  int = 0  // the error won't be detected until at run time 🤨
	)
	_ = fmt.Sprint(c_dividend, c_divisor, v_dividend, v_divisor)

	// type
	const noDelay time.Duration = 0 // explicit type specified (time.Duration)
	const timeOut = time.Minute * 5 // the type'd be inferred from the right side
	fmt.Printf("noDealy	  : %T %[1]v\n", noDelay)
	fmt.Printf("timeout	  : %T %[1]v\n", timeOut)
	fmt.Printf("time.Minute: %T %[1]v\n", time.Minute)

	// re-use expression
	const (
		ax = 1.0 // ax == 1.0
		bx       // bx == 1.0 (hmm)
		ay = 2.5 // ay == 2.5
		by       // by == 2.5 (hmm)
	)
}

func ConstantIota() {
	type Ranking int
	const (
		First Ranking = iota // 0 -> 1 -> 2
		Second
		Third
	)

	type Flags uint
	const (
		FlagUp           Flags = 1 << iota // 2**0
		FlagBroadcast                      // 2**1
		FlagLoopBack                       // 2**2
		FlagPointToPoint                   // 2**3
	)

	type Fruit int
	const (
		Apple, Avocado         = iota + 1, iota + 2 // [0]+1, [0]+2 => 1,2
		Banana, Blueberry                           // [1]+1, [1]+2 => 2,3
		Cantaloupe, Clementine                      // [2]+1, [2]+2 => 3,4
	)

	// KiB | KB
	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776              	(exceeds 1 << 32)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620717411303424     	(exceeds 1 << 64)
		YiB // 1208925819614629174706176
	)
	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
		TB = 1000 * GB
	)

	const (
		x         = iota * 42 // untyped int constant
		y float64 = iota * 42
		z         = iota * 42 // untyped int constant
	)
}

func ConstantUntyped() {
	/*
		Constants in Go are a bit unusual
		- be able NOT committed to a particular type (inferred)
		- untyped constant would be held with much higher precision (at least 256 bits)
		- be able to participate in many exprs without requiring conv (e.g. "too big")
	*/

	// Able to participate in "BIG NUMBER CALC" without conversion
	const ZiB = 1180591620717411303424    // > 2**64
	const YiB = 1208925819614629174706176 // > 2**64
	fmt.Printf("YiB/ZiB = %d\n", YiB/ZiB)

	// Precision loss
	var PiInF64 float64 = math.Pi
	var PiInCx128 complex128 = math.Pi
	fmt.Println(PiInF64, PiInCx128)

	// Type inferring with constants
	c1 := 0      // untyped int		=>	implicit int(0)			size is NOT guaranteed!
	c2 := '\000' // untyped rune	=>	implicit rune('\000')	or `int32` (alias)
	c3 := 0.0    // untyped float	=>	implicit float64(0.0)
	c4 := 0i     // untyped complex	=>	implicit complex128(0i)

	// Type inferring with variables
	var v float64 = 3 + 0i // => float64	it's actually `float64(3+0i)`
	v = 2                  // => float64 	it's actually `float64(2)`
	v = 1e123              // => float64 	it's actually `float64(1e123)`
	v = 'a'                // => float64 	it's actually `float64('a')`

	_ = fmt.Sprint(c1, c2, c3, c4)
	_ = fmt.Sprint(v)
}
