package main

import (
	"fmt"

	"github.com/code_go_book_go_programming_language/Part_Three/snip17_gocomptypes_struct_exampleexportableornot/lib"
)

func main() {
	// normalStruct()
	// lowercaseField()
	// nestedStruct()
}

func normalStruct() {
	ep1 := slave.ProfileExportable{
		Name: "Alex",
		Age:  12,
	}
	ep2 := slave.ExportLowercaseProfileNotExportable("Bob", 22)

	fmt.Println(ep1, ep2)
}

func lowercaseField() {

	// The `age` cannot be changed since it wasn't exported (lowercase)
	dog := slave.Dog{
		Name: "Chole",
		// age: 20,
	}

	fmt.Println(dog)
}

func nestedStruct() {
	cat := slave.Cat{
		// If and only if the `animal` struct is being exported,
		// then you can put the values directly in this `{Ani: Ani{.., ..}}`
		Cuteness: 100,
	}
	// Although it was not exported (lowercase initial struct)
	// you CAN access(use) them by assigning values to it directly
	cat.Name = "Lux"
	cat.Age = 2

	// For this one, we're using the exported/exportable `Animal` struct
	cat2 := slave.Cat2{
		Cuteness: 999,
		Animal: slave.Animal{
			Name: "Becca",
			Age:  1,
		},
	}

	fmt.Println(cat)
	fmt.Println(cat2)
}
