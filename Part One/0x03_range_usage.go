package main

import "fmt"

func main() {
	for idx, item := range []string{"foo", "bar"} {
		fmt.Println(idx, item)
	}

	for idx, char := range "繁體中文" {
		fmt.Printf("%#U starts at byte pos %d\n", char, idx)
	}

	valToBeMapped := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
	}
	for key, value := range valToBeMapped {
		fmt.Println(key, value)
	}
}
