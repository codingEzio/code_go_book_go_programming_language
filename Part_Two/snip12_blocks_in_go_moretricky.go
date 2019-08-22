package main

import (
	"log"
	"os"
	"runtime"
)

func main() {
	OpenFileUsingIfCatchError()
	SummaryOfColonEqSignInitAndScopeOfBlock() // do check this out
}

func OpenFileUsingIfCatchError() {
	/*
		I personally find this example to be an extremely good one

		>> // The style here might look great (& in other languages)
		>> // but, once the `if` has ended, the `f` became inaccessible immed!
		>>
		>> if f, err := os.Open("what"); err != nil {
		>>     return err
		>> }
		>> // You can add one more block `else` to accomplish this (read & close)
		>> // but it's not a normal practice, check the recommd way down below 😁
		>>
		>> f.ReadByte()		// `f` is beyond the scope of `if` statement
		>> f.Close()		// so there should a `compile error: undefined f`
	*/

	_, filename, _, _ := runtime.Caller(1)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err) // or `panic(err)`
	}

	_ = f.Close()
}

func SummaryOfColonEqSignInitAndScopeOfBlock() {
	/*
		>> var cwd string	// this one is outside of `main` (the style of `init`)
		>>
		>> func init() {
		>>
		>>     // Pitfall #1
		>>	   // 	`cwd` is re-declared in here (global => local)
		>> 	   // 	that means the `main` cannot use this anymore (which is .. bad!)
		>>     //
		>>     // Pitfall #2
		>>     // 	desired flow 	"(1)declare as global var (2)assign value (==use)"
		>>     // 	actual process 	"(1)re-declare (2)never been used ('assign' counts!)
		>>     cwd, err := os.Getwd()
		>>
		>>	   // This version is much much better
		>>     // * initialize a variable of type `error` (in case it doesn't init)
		>>     // * assign the value to `cwd` instead of the mis-use `:=` (shadowing)
		>>     var err error
		>>     cwd, err = os.Getwd()
		>>
		>> 	   if err != nil {
		>>         log.Fatalf("failed: %v", err)
		>>     }
		>> }
	*/
}
