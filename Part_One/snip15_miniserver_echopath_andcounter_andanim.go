// Respond URL path if was given and counting visited times (../count)
package main

import (
	"fmt"
	"github.com/code_go_book_go_programming_language/Part_One/lib"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	/*
		Behind the scenes
		* the server runs the handler for each incoming request in a
		* separate goroutine so that it can server multiple req simultaneously

		`Lock` & `Unlock`
		* it was meant to avoid a bug which is called as "race condition"
		* we must ensure that at most one goroutine access the var at a time
		* #TODO might add more details when I fully grasp it
	*/
	http.HandleFunc("/", handler_)
	http.HandleFunc("/count", counter)

	// Optional
	http.HandleFunc("/pattern", func(w http.ResponseWriter, r *http.Request) {
		Part_One.Lissajous(w)
	})

	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

// Echo the Path component of the requested URL
func handler_(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Echo the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
