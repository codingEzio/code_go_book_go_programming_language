package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// Array vs Slices
	// * [array]	fixed-length
	// * [slice]	can grow & shrink
	var a [3]int
	fmt.Println(a[0], a[len(a)-1])

	for idx, elem := range a {
		fmt.Println(idx, elem)
	}

	var b [3]int = [3]int{1, 2, 3}
	var d [3]int = [3]int{1, 2} // last one is `0`
	_ = fmt.Sprint(b, d)

	e := [...]int{1, 2, 3}
	fmt.Printf("%T -_- %v\n", e, e)

	// size of an array it part of its type, so `[3]int` is NOT the same as `[4]int`
	// >> y := [3]int{1, 2, 3}
	// >> y := [4]int{1, 2, 0, 0}

	type Currency int64
	const (
		USD Currency = iota // 0 -> 1 -> 2 -> 3
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	r := [...]int{5: -2} // 5-1 with initial 0, the last one is `-2`
	fmt.Println(r)

	// array with different sizes cannot be compared with each other
	fmt.Println([2]int{1, 2} == [...]int{1, 3})
	fmt.Println([2]int{2} == [2]int{1, 3})

	c1 := sha256.Sum256([]byte("abc"))
	c2 := sha256.Sum256([]byte("abc"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
