// Prints flags arguments for different temperature scales (C, F, Kelvin)
package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// *celsiusFlag satisfies the `flag.Value` interface
type celsiusFlag struct{ Celsius }

var temperature = CelsiusFlag("temp", 20.0, "the temperature")

// Copy&Paste: °   °C   °F
func main() {
	flag.Parse()
	fmt.Println(*temperature)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (c Celsius) String() string { return fmt.Sprintf("%.3g°C", c) }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "c", "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "f", "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "k", "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }
