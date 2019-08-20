package main

import (
	"fmt"
)

// Note these
// * these are indeed different types, that is, you cannot compare them directly
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	// Do note these named types
	// * when you call `Celsius(x)`, it is "conversion" instead of "function call"!
	// * also, it does NOT change the original value (since it's not a func call right?)
	x := 1
	fmt.Printf("%f %T\n", Celsius(x), Celsius(x))

	// These are REAL conversions 😏
	fmt.Printf("F100 -> C %f %T\n", FToC(Fahrenheit(100)), FToC(Fahrenheit(100)))
	fmt.Printf("C100 -> F %f %T\n", CToF(Celsius(100)), CToF(Celsius(100)))

	// The `%g` was meant to omit any unnecessary digit of a float number
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))

	var CToday Celsius
	var FToday Fahrenheit
	fmt.Printf("%f %g %T\n", CToday, CToday, CToday)
	fmt.Printf("%f %g %T\n", FToday, FToday, FToday)

	// Kinda like the `toString` in Java
	fmt.Println(Celsius(100), Fahrenheit(100))
	fmt.Printf("%v %v\n", Celsius(100), Fahrenheit(100))
	fmt.Printf("%s %s\n", Celsius(100), Fahrenheit(100))
	fmt.Println(Celsius(100).String(), Fahrenheit(100).String())

	// These two won't call the `String` method
	fmt.Printf("%g %g\n", Celsius(100), Fahrenheit(100))
	fmt.Println(float64(Celsius(100)), float64(Fahrenheit(100)))
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// It's kinda like the `toString` method in Java
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
