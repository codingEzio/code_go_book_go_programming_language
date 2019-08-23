package main

import "fmt"

func main() {
	/*
		Types
		- Basic types 		number, string, boolean				etc.
		- Aggregate types	array, struct						etc.
		- Reference types	pointer, slice, map, func, channel	etc.
		- Interface types	leave blank for now

		Unsigned numbers
		* it tends to be used only when their bitwise/else operators are required
		* as when impl bit sets, parsing binary file fmts, or for hashing and crypto
	*/

	IntType()
	BitOperation()

	// `len` returns a signed `int`
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	// Type conversion
	var apples int32 = 1
	var lemons int64 = 2
	var _ = int(apples) + int(lemons)

	// Formatting for 0o and 0x
	var num int64 = 1024
	fmt.Printf("%d, %[1]o, %#[1]o, %[1]x, %#[1]x\n", num)

	ascii := 'a'
	unicode := '你'
	newline := '\n'
	fmt.Printf("%d, %[1]c %[1]q\n", ascii)
	fmt.Printf("%d, %[1]c %[1]q\n", unicode)
	fmt.Printf("%d, %[1]c %[1]q\n", newline)
}

func IntType() {

	var (
		a1 int  = 3 // 32/64, do it explicitly
		a2 uint = 3

		b1 int8
		b2 uint8
		b3 byte // a synonym for `uint8` (emphasize "raw-data" instead of num shit)

		c1 int16 // int/uint OR int/int32/int64 is NOT the same
		c2 int64 // that means <an explicit conversion> is always needed!

		d1 int32 // int doesn't always equal to i32
		d2 rune  // a synonym for `int32` (indicates a val is a Unicode Code Point)
	)
	_ = fmt.Sprint(a1, a2, b1, b2, b3, c1, c2, d1, d2)

	{
		// Remainders
		// * only works in integers (Python could do floating num)
		// * the remainder's sign is the same as the dividend (in Py it doesn't)
		fmt.Printf("Should be -2,-2 \t:: %d, %d\n", -5%-3, -5%3)
		fmt.Printf("Should be 1.25/1 \t:: %.2f %d\n", 5.0/4.0, 5/4)
	}

	{
		var o1u uint8 = 255
		var o1 int8 = 127

		// Overflow
		fmt.Printf("Should be 255/0/1 \t:: %d, %d, %d\n", o1u, o1u+1, o1u*o1u)
		fmt.Printf("Should be 127/-128/1 \t:: %d, %d, %d\n\n", o1, o1+1, o1*o1)
	}
}

func BitOperation() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("     x :: %08b\n", x) // 0010,0010
	fmt.Printf("     y :: %08b\n", y) // 0000,0110

	/*
		不要多想, just follow the rules. 此种运算不存在"进位"!
		- &		1,1	  ->true	else ->false	(false:0, true:1)
		- |		0,0	  ->false	else ->true
		- ^		01,10 ->true	else ->false
		- &^	x ^ (x&y)
	*/
	fmt.Printf("x &  y :: %08b\n", x&y)  // 0000,0010
	fmt.Printf("x |  y :: %08b\n", x|y)  // 0010,0110
	fmt.Printf("x ^  y :: %08b\n", x^y)  // 0010,0100
	fmt.Printf("x &^ y :: %08b\n", x&^y) // 0010,0000

	fmt.Printf("x << 1 :: %08b\n", x<<1)
	fmt.Printf("   x   :: %08b\n", x<<0)
	fmt.Printf("x >> 1 :: %08b\n", x>>1)
}
