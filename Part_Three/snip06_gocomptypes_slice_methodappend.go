package main

import "fmt"

func main() {
	// appendBasics()
	// appendIntBreakdown()
	appendAndappendIntImprovedFinalArg()
}

func appendBasics() {
	var runes []rune
	for _, r := range "你好 Tom" {
		runes = append(runes, r)
	}

	fmt.Println(runes)
	fmt.Printf("%q\n", runes)
}

func appendIntBreakdown() {
	/*
		Well, most of the concepts were commented in the original func part,
		what we'll do here is to point out some important points.

		[1]
		Since we don't know whether a given call will cause a reallocation or else,
		we can't assume (based on our test program) the original slices refer to the
		same array as the resulting slice , nor that it refers to a different one :)

		[2]
		Similarly, we must not assume that assignments to elements of the old slice
		will (or will not) be reflected in the new slice.

		[3]
		Updating* the slice var is required not just when calling append (a = append(a, i)),
		but for any function that may change the length/capacity of a slice
		or make it refer to a different underlying array (p.s. Python doesn't require it)
	*/

	var dest, returnedVal []int
	for i := 0; i < 10; i++ {
		returnedVal = appendInt(dest, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(returnedVal), returnedVal)
		dest = returnedVal
	}
}

func appendAndappendIntImprovedFinalArg() {
	var ap []int
	ap = append(ap, 1)
	ap = append(ap, 2, 3, 4)
	ap = append(ap, ap...)
	fmt.Println("[ AP]:", ap)

	var apf []int
	apf = append(apf, 1)
	apf = append(apf, 2, 3, 4)
	apf = append(apf, apf...)
	fmt.Println("[APF]:", apf)
}

// ----- ----- ----- Helper ----- ----- -----

func appendInt(slice []int, elem int) []int {
	var z []int
	zLen := len(slice) + 1 // len of slice plus 1 as end-index

	if zLen <= cap(slice) { // the slice must still have space to hold elements
		z = slice[:zLen] // z shares the same underlying array with `slice`
	} else {
		zCap := zLen // zCap now is the same as zLen (for end-index)
		if zCap < 2*len(slice) {
			zCap = 2 * len(slice) // expand by two times
		}
		z = make([]int, zLen, zCap) // same len as prev, two times bigger capacity
		copy(z, slice)              // a NEW array of (doubled) size in cmp to `slice`
	}

	z[len(slice)] = elem // add elem to the last (enough capacity to hold that)
	return z             // return a temporary slice (that means no real changes to the orig slice)
}

func appendIntImprovedFinalArg(slice []int, elems ...int) []int {
	var z []int
	zLen := len(slice) + len(elems)

	if zLen <= cap(slice) {
		z = slice[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(slice) {
			zCap = 2 * len(slice)
		}
		z = make([]int, zLen, zCap)
		copy(z, slice)
	}

	copy(z[len(slice):], elems)
	return z
}
