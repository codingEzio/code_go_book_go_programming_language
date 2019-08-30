package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(join(",", "hello", "world"))
	fmt.Println(join(".", "www", "google", "com"))
}

func join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	byt := bytes.Buffer{}
	for _, str := range strs[:len(strs)-1] {
		byt.WriteString(str)
		byt.WriteString(sep)
	}
	byt.WriteString(strs[len(strs)-1])

	return byt.String()
}
