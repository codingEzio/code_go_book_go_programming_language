// Counts the number of bits that are different in two SHA256 hashes
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	t1 := []byte{1, 2, 3}
	t2 := []byte{4, 5, 6}

	// Do note the number is based on BITS!! (%b)
	fmt.Println(ShaBitDiff(t1, t2))

	fmt.Printf("[1] %x\n", sha256.Sum256(t1))
	fmt.Printf("[2] %x\n", sha256.Sum256(t2))
}

func popCount(bt byte) int {
	count := 0
	for ; bt != 0; count++ {
		bt &= bt - 1
	}
	return count
}

func bitDiff(e1, e2 []byte) int {
	count := 0
	for i := 0; i < len(e1) || i < len(e2); i++ {
		switch {
		case i >= len(e1):
			count += popCount(e2[i])
		case i >= len(e2):
			count += popCount(e1[i])
		default:
			count += popCount(e1[i] ^ e2[i])
		}
	}
	return count
}

func ShaBitDiff(e1, e2 []byte) int {
	shaA := sha256.Sum256(e1)
	shaB := sha256.Sum256(e2)
	return bitDiff(shaA[:], shaB[:])
}
