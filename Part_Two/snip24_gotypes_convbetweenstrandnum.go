package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 16
	y := fmt.Sprintf("%d", x)

	// Int -> ASCII
	fmt.Println(strconv.Itoa(x), y)
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y))

	// Int -> Int (convert different base)
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(strconv.FormatUint(uint64(x), 36)) // base <= 36 (1~9 a~z)

	// Int -> Int (formatting not converting)
	fmt.Printf("100000 => %#x (hex)\n", 100000)
	fmt.Printf("1000   => %#o (oct)\n", 10000)
	fmt.Printf("16     => %b  (bin)\n", 16)

	// String(num) -> Int
	a, _ := strconv.Atoi("123")
	b, _ := strconv.ParseInt("123", 10, 8)

	fmt.Println("Converted '123' =>", a, b, "(both int)")
}
