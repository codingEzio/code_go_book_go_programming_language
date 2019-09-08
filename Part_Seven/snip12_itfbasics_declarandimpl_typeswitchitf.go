package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Describer interface {
	Describe()
}

func main() {
	findType("Thor")

	person := Person{
		name: "York Shio",
		age:  40,
	}
	findType(person)
}

func (person Person) Describe() {
	fmt.Printf("%s is %d years old\n", person.name, person.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("Unknown type\n")
	}
}
