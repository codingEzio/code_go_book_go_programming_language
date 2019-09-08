package main

import (
	"fmt"
)

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func main() {
	/*
		This program doesn't show the actual diff usages in cmp with "methods",
		just go check the next example!
	*/

	name := MyString("Eliot Anderson") // i, o, e, o
	var vf VowelsFinder
	vf = name

	fmt.Printf("Vowels are %c\n", vf.FindVowels())
}

func (mstr MyString) FindVowels() []rune {
	var vowels []rune
	for _, rne := range mstr {
		if rne == 'a' || rne == 'e' || rne == 'i' || rne == 'o' || rne == 'u' {
			vowels = append(vowels, rne)
		}
	}
	return vowels
}
