package main

import (
	"fmt"
	"sort"
)

func main() {
	ints_p := []int{1, 4, 4, 1}
	ints_np := []int{532, 4, 4, 1}

	fmt.Println(IsPalindrome(sort.IntSlice(ints_p)))
	fmt.Println(IsPalindrome(sort.IntSlice(ints_np)))
}

func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	max := s.Len() - 1

	for i := 0; i < s.Len()/2; i++ {
		if !equal(i, max-i, s) {
			return false
		}
	}
	return true
}
