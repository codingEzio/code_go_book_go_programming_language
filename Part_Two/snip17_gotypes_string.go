package main

import (
	"fmt"
	"html"
	"net/url"
	"unicode/utf8"
)

func main() {
	// StringBasics()
	// StringLiterals()
	// UnicodeCodePoint()
}

func StringBasics() {
	s := "Hello"

	// Index
	fmt.Println(len(s))
	fmt.Println(s[:4], s[1:], s[:]+"-HK!")

	// Formatting
	fmt.Println(s[0])
	fmt.Printf("%[1]d %[1]c %[1]q\n", s[0])

	// Immutability (of strings)
	// * you did NOT modify "the string that `s` originally held"
	// * you only changed the "variable" (not the "value") (now you can modify `s`)
	s_origcopy := s

	// Check more at (variables are always mutable, string itself is not)
	// $ https://stackoverflow.com/a/36721452/6273859
	// $ https://github.com/golang/go/issues/9816
	// $ https://play.golang.org/p/KWccuJsFpj
	s += ", World"
	_ = fmt.Sprintf("s: %s  s_origcopy: %s\n", s, s_origcopy)
}

func StringLiterals() {
	// I'll just point out some of its usages
	// * HTML-templates, JSON-literals, cmd-usage messages(hell yeah!) etc.
	shortLit := `hello \n boys`
	longLit := `line 1
	line 2 with a tab
line 3`

	_ = fmt.Sprint(shortLit)
	_ = fmt.Sprint(longLit)

	exampleHTML := `"Foo's bar" <foobar@example.com>`
	exampleURL := `Foo's Bar?'`

	fmt.Println("Escaped HTML :", html.EscapeString(exampleHTML))
	fmt.Println("Escaped URL  :", url.PathEscape(exampleURL))
}

func UnicodeCodePoint() {
	/*
		Unicode Code Point
		- collects all of the characters and some other things (punc etc.)
		- and assigns each one a standard number called a UCP (in go, a 'rune')
	*/

	// They are kinda the same thing (=> identical output (not always))
	// * a rune whose value is less than 256 can be written as sth like `\x41` (=='A')
	// * but for higher values, a `\u` or `\U` escape MUST be used
	fmt.Println(`世: \xe4\xb8\x96`, `界: \xe7\x95\x8c`)
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")

	greeting := "Hey 你好"

	// for idx := 0; idx < len(greeting); {
	// 	char, size := utf8.DecodeRuneInString(greeting[idx:])
	// 	fmt.Printf("%d\t%c\n", idx, char)
	// 	idx += size
	// }
	for idx, char := range greeting {
		fmt.Printf("%d\t%q\t%d\n", idx, char, char)
	}

	fmt.Println("Actual length  :", len(greeting))                    // 10
	fmt.Println("Length of rune :", utf8.RuneCountInString(greeting)) // 6

	greetingInJp := "こんにちは" // en: Hello, ch: 你好 (音:摳妳雞瓦)
	fmt.Printf("% x\n", greetingInJp)
	fmt.Printf("%x\n", []rune(greetingInJp))

	fmt.Println("\n[Conv by `string()`]")
	fmt.Println(string(45121), string(65))              // 끁 A
	fmt.Println(string("\xe4\xb8\x96"), string(0x4eac)) // 世 京
}
