package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

/*
	Some restrictions about struct type (not just struct actually)
	> A named struct type S can't declare a filed of the same type S
	> that is, an aggregate value cannot contain it self (recursion maybe?)

	But S may declare a filed of the pointer type *S,
	which let us to make recursive DS like linked-list and trees
*/
func main() {

	vals := []int{4, 1, 2, 3}

	fmt.Println(vals)

	Sort(vals)
	fmt.Println(vals)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	// Equivalent to `return &tree{value: value}`
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// It appends the element of `t` to `values` in order,
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
