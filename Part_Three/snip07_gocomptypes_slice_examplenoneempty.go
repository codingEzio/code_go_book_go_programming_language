// noneEmpty & noneEmptyUsingAppend return a slice holding only the non-empty strings.
// The underlying array is modified during the call (ah, use `A = METHOD(A)` instead)
package main

import "fmt"

func main() {
	data := []string{"One", "", "Three"}
	data2 := []string{"One", "", "Three"}

	// Wrong
	fmt.Println(noneEmpty(data))            // already modified (ret val seems nice but no)
	fmt.Println(noneEmptyUsingAppend(data)) // already modified

	// Right
	data2 = noneEmpty(data2)            // in-place modify
	data2 = noneEmptyUsingAppend(data2) // in-place modify
	fmt.Println(data2)
}

func noneEmpty(strings []string) []string {
	i := 0
	for _, str := range strings {
		if str != "" {
			strings[i] = str // not null? give it to me! (given array => mine)
			i++
		}
	}

	return strings[:i]
}

func noneEmptyUsingAppend(strings []string) []string {
	out := strings[:0]
	for _, str := range strings {
		if str != "" {
			out = append(out, str) // not null? append those to my array!
		}
	}

	return out
}
