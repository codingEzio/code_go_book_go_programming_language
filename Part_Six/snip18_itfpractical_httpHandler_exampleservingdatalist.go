package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 5}

	fmt.Println("Now serving at http://localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
// }
func (db database) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(wr, "%s: %s\n", item, price)
	}
}
