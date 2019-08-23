package main

import "fmt"

func main() {
	fmt.Println(IsAnagram("", ""))
	fmt.Println(IsAnagram("a", "a"))
	fmt.Println(IsAnagram("abc", "cba"))
	fmt.Println(IsAnagram("word", "dowr"))
}

func IsAnagram(a, b string) bool {
	/*
		For non-native English speakers:
		* anagram: 相同字母异序词 (word->dorw, secure->rescue)
	*/

	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}

	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}

	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}

	return true
}
