// Respond URL path if was given
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		You can either use a browser or use the program we prev_ly wrote (0x11_*.go).
	*/
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
