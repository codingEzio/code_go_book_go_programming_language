package main

import (
	"fmt"
	"sort"
)

var ages_global_donotchange = map[string]int{
	"Bob":  101,
	"Alex": 20,
}

func main() {
	mapGet()
	mapCreate()
	mapDeleteAndModify()
	mapUnorderedAndOrdered()

	// Our own map-equal function ;)
	fmt.Println(mapEqualOrNot(
		map[string]int{"A": 0},
		map[string]int{"B": 42}))
}

func mapGet() {
	if _, ok := ages_global_donotchange["Dan"]; !ok {
		fmt.Println("We don't have Dan's age!")
	}
}

func mapCreate() {
	agesUsingMake := make(map[string]int)
	agesUsingLiteral := map[string]int{}
	agesUsingLiteral2 := map[string]int{
		"Axis": 31,
	}

	agesUsingMake["Alex"] = 20
	agesUsingLiteral["Alice"] = 42

	_ = fmt.Sprint(agesUsingMake, agesUsingLiteral, agesUsingLiteral2)

	// Though you can use `var` to create a map (len0 nil), better not to do it
	// var agesUsingVar map[string]int // opts on a nil map will cause `panic`
	// agesUsingVar["Alo"] = 15        // this opt is forbidden ofc, so does the others
}

func mapDeleteAndModify() {
	ages := map[string]int{
		"Alex": 20,
		"Bob":  101,
	}

	delete(ages, "Bob")
	delete(ages, "Bob") // => 0

	ages["Alex"]++                  // 21
	ages["Alex"] += 1               // 22
	ages["Alex"] = ages["Alex"] + 1 // 23
}

func mapUnorderedAndOrdered() {

	// The order of map iteration is unspecified, so does map itself :)
	for name, age := range ages_global_donotchange {
		fmt.Println(name, ":", age)
	}

	// Wanna enumerate k/v in order?
	// [1] sort the keys
	// [2] get the val by sorted keys (simple)
	var names_ordered []string
	for name := range ages_global_donotchange {
		names_ordered = append(names_ordered, name)
	}
	sort.Strings(names_ordered)

	for _, name := range names_ordered {
		_ = fmt.Sprint(name, ages_global_donotchange[name])
	}
}

// ----- ----- ----- Helper ----- ----- -----

func mapEqualOrNot(x, y map[string]int) bool {
	/* As with slices, map cannot be compared to each other (we need to write our own) */
	if len(x) != len(y) {
		return false
	}
	for xKey, xVal := range x {
		if yVal, ok := y[xKey]; !ok || yVal != xVal {
			return false
		}
	}
	return true
}
