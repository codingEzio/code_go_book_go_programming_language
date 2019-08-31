package main

import (
	"fmt"
)

func main() {
	fmt.Println(magic())
}

func magic() (ret string) {
	defer func() {
		recover()
		ret = "hi"
	}()

	panic("oh shit")
}
