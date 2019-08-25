package main

import (
	"bytes"
	"fmt"
)

func main() {
	traitAndDifferencesOfSlice()
	sliceCannotBeComparedAutomatically()
}

func traitAndDifferencesOfSlice() {
	// Using make to create a slice
	var m1 = make([]byte, 3, 10)
	var m2 = make([]byte, 10)[:3] // fancier but the same (len: 3, cap: 10)
	fmt.Println(m1, len(m1), cap(m1))
	fmt.Println(m2, len(m2), cap(m2))

	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun"}
	spring := months[1:3]
	summer := months[3:]

	// Multiple slices can share/refer the same underlying array (& the overlapping parts)
	fmt.Println(len(months), cap(months))
	fmt.Printf("%T, %[1]v, %d, %d\n", spring, len(spring), cap(spring))
	fmt.Printf("%T, %[1]v, %d, %d\n", summer, len(summer), cap(summer))

	// The length of a slice could be extended (within cap(capacity) of course)
	expandedLenOfSpring := spring[0 : len(months)-1]
	fmt.Printf("%T, %[1]v, %d\n", expandedLenOfSpring, len(expandedLenOfSpring))

	// A slice contains a pointer to an element of an array,
	// so a slice being passed to a func might affect the underlying array elements
	zeroToFive := [...]int{0, 1, 2, 3, 4, 5}
	reverse(zeroToFive[:])
	fmt.Printf("Reversed: %v (type: %[1]T)\n", zeroToFive)

	// Like an array but the size is not given ( "[...]" != "[]" )
	zeroToFiveSlice := []int{0, 1, 2, 3, 4, 5}
	reverse(zeroToFiveSlice)
	fmt.Printf("Reversed: %v (type: %[1]T)\n", zeroToFiveSlice)

	// This version of `reverse` use "array pointer" as argument
	// so it's necessary to pass its address instead of value (z -> &z)
	// and also, the array you passed in must be the same size as the func ( [5]int )
	zeroToFiveSlice2 := [5]int{0, 1, 2, 3, 4}
	reverseUsingArrayPointer(&zeroToFiveSlice2)
	fmt.Printf("Reversed: %v (type: %[1]T)\n", zeroToFiveSlice2)
}

func sliceCannotBeComparedAutomatically() {
	/* Generally, slices are NOT comparable (but u CAN find a way) */

	// For bytes, simply using `bytes.Equal`
	cmpByte1 := []byte{1, 2, 3}
	cmpByte2 := []byte{1, 2, 3}
	fmt.Printf("%t\n", bytes.Equal(cmpByte1, cmpByte2))

	// For anything else, use(write) our own (LOL)
	o1 := []string{"day", "night"}
	o2 := []string{"day", "night"}
	o3 := []string{"night", "day"}
	fmt.Printf("%t, %t\n", equal(o1, o2), equal(o2, o3))

	baseArr := [2]string{"night", "day"}        // an array	(type: [2]string)
	baseArrSlice := []string{"night", "day"}[:] // a slice	(slice CANNOT compared with array either)

	// How to test whether a slice is empty anyway?
	// * simply use `len(s) == 0` instead of `== nil`
	// * other than cmp-eq to `nil`, it behaves like any other zero-len slice!
	var nilSlice []string    // len: 0, s is nil
	nilSlice = nil           // len: 0, s is nil
	nilSlice = []string(nil) // len: 0, s is nil
	nilSlice = []string{}    // len: 0, s is NOT nil

	// Unnecessary stuff
	_ = fmt.Sprint(baseArr, baseArrSlice)
	_ = fmt.Sprint(nilSlice)
}

// ----- ----- ----- Helper ----- ----- -----

func reverse(slice []int) {
	/* Reverse elements in an array or a slice */
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
func reverseUsingArrayPointer(ints *[5]int) {
	for i := 0; i < len(ints)/2; i++ {
		end := len(ints) - i - 1
		ints[i], ints[end] = ints[end], ints[i]
	}
}

func equal(x, y []string) bool {
	/* Compare whether two slices have the same elements */
	if len(x) != len(y) {
		return false
	}

	for idx := range x {
		if x[idx] != y[idx] {
			return false
		}
	}

	return true
}
