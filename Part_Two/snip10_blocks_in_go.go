package main

import "fmt"

/*
	A tiny comparison between bash-script and Go

	>> func f() {
	>>     v := 1
	>>     // for shell, OK
	>>     // for Go, undefined `v` (in `g()`)
	>>     g()
	>> }
	>>
	>> func g() {
	>>     fmt.Println(v)
	>> }
	>>
	>> func main() {
	>> 	   f()
	>> }
*/

func main() {
	// Everyone to check out the examples inside this book
	// * long & practical
	// * tricky (good examples indeed) (can also be found on StackOverflow!)

	ImplicitBlock()
	BlockScopeWithDeclaration()
	BlockScopeWithAssignment()
}
func ImplicitBlock() {
	/*
		Some of them (a bit more abstract)
		- universe block		all source code
		- package block			all package's code (can be several files in a single dir)
		- file block			file's source code
		- statement block		like `for`, `if`, `switch` etc.
	*/

	for i := 0; i < 5; i++ {
		_ = fmt.Sprint(i)
	} // the end of `i`

	if i := 0; i >= 0 {
		_ = fmt.Sprint(i)
	} // the end of `i`

	switch i := 2; i * 4 {
	case 8:
		_ = fmt.Sprint(i) // each clause in a `switch` stmt acts like a `{}` block
	default:
		_ = fmt.Sprint("hello")
	} // the end of `i`

	// There's also a `select` statement
	// * kinda like `switch`, but serve for a different purposes
	// * it's related to `go routine` & `channels` (no example provided for now)
}

func BlockScopeWithDeclaration() {
	// You can't access either of them at the outermost `{}`
	{
		a := 1
		_ = fmt.Sprintf("%d ", a)
		{
			b := 2
			_ = fmt.Sprintf("%d ", b) // the end of `b`

			// `b` has access 		to `{}` where `a` lives in	{ A { A: exists } }
			// `a` has NO access 	to `{}` where `b` lives in	{ A(B:nope) { B } }
			_ = fmt.Sprintf("%d ", a)
		}
	} // the end of `a`

	t := 1
	{
		y := 1
		{
			_ = fmt.Sprint(t) // totally fine cuz it has the access to outermost `t` & `y`
			_ = fmt.Sprint(y)
			{
				_ = fmt.Sprint(t) // totally fine for the same reason
				_ = fmt.Sprint(y)
			}
		}
		_ = fmt.Sprint(t) // totally fine for the same reason
	}
	{
		_ = fmt.Sprint(t) // it can also access the var `t` ("in access out" still applies)
	}
}

func BlockScopeWithAssignment() {
	v := 1
	{
		v = 2
		{
			v := 3
			_ = fmt.Sprint(v) // reasonable, huh 🤒
		}
		_ = fmt.Sprint(v) // change the val (1=>2) (from the outermost `v`) (assignment😇)
	}
	_ = fmt.Sprint(v) // changed by "assignment" (1=>2) can't be changed by "re-declaration" 😋

	{
		var (
			a = 1
			b = a
		)
		_, _ = fmt.Sprint(a), fmt.Sprint(b)

		// NOT okay for `b` (a doesn't have a value yet 😕)
		// a, b := 1, a
	}
}
