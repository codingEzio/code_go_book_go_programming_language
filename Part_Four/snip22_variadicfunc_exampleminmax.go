package main

import (
	"fmt"
)

func main() {
	aHarmlessArray := []int{10, -1, 20}

	fmt.Println(min(aHarmlessArray...))
	fmt.Println(minAtLeastOneArg(aHarmlessArray[0], aHarmlessArray[1:]...))

	fmt.Println(max(aHarmlessArray...))
	fmt.Println(maxAtLeastOneArg(aHarmlessArray[0], aHarmlessArray[1:]...))
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	temp := nums[0]
	for _, num := range nums {
		if num < temp {
			temp = num // smaller than prev one? replace it!
		}
	}

	return temp
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	temp := nums[0]
	for _, num := range nums {
		if num > temp {
			temp = num // bigger than prev one? replace it!
		}
	}

	return temp
}

func minAtLeastOneArg(first int, nums ...int) int {
	initial := first
	for _, num := range nums {
		if initial > num {
			initial = num
		}
	}

	return initial
}

func maxAtLeastOneArg(first int, nums ...int) int {
	initial := first
	for _, num := range nums {
		if initial < num {
			initial = num
		}
	}

	return initial
}
