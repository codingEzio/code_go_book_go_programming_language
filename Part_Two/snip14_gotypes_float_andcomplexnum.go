package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	FloatBasics()
	ComplexBasics()
}

func FloatBasics() {
	var fmt_arg float64 = 12345.67890
	fmt.Printf("%%g :: %[1]g, %%e :: %.8[1]e, %%f :: %.4[1]f\n", fmt_arg)

	for x := 0; x < 5; x++ {
		fmt.Printf("x = %d\teˣ = %6.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0, -0, 1/0, -1/0, 0/0

	nan := math.NaN()
	fmt.Println(nan == nan, nan > nan, nan < nan)
}

func ComplexBasics() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	// How to multiply
	// https://www2.clarku.edu/faculty/djoyce/complex/mult.html
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	fmt.Println(1i * 1i)

	fmt.Println(cmplx.Sqrt(-1))
}
