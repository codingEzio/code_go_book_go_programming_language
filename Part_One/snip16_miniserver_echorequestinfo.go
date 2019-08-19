package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler__)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handler__(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// We could have written it as (move the 1st part out with its semicolon 😋)
	// >> err := r.ParseForm()
	// >> if err != nil {
	// >>     log.Print(err)
	// >> }
	if err := r.ParseForm(); err != nil {
		// It's a bit different when it comes to the scope of `err`
		// * the shorter one  =>  the end of this `if`
		// * the longer one   =>  still exists until the end of the function
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
