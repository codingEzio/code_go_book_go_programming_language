// Two functions (i.e. remove & remove2) remove a specific element from given index.
package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}
	data2 := []int{1, 2, 3, 4, 5}

	data = remove(data, 1)    // 1,3,4,5
	data2 = remove2(data2, 1) // 1,5,3,4

	fmt.Println(data)  // order was preserved
	fmt.Println(data2) // the last elem was moved to the gap
}

func remove(slice []int, idx int) []int {
	/* Order matters ( 1,2,3,4 del[0] -> 2,3,4 ) */
	copy(slice[idx:], slice[idx+1:])

	return slice[:len(slice)-1]
}

func remove2(slice []int, idx int) []int {
	/* Order doesn't matter ( 1,2,3,4 del[0] -> 4,2,3 ) */
	slice[idx] = slice[len(slice)-1]

	return slice[:len(slice)-1]
}
