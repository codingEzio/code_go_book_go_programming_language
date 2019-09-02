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

// It was declared in the previous package
// func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// localhost:8000/listall
// localhost:8000/item
func (db database) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/listall":
		for item, price := range db {
			fmt.Fprintf(wr, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")

		price, ok := db[item]
		if !ok {
			wr.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(wr, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(wr, "%s\n", price)
	default:
		wr.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(wr, "no such page: %s\n", req.URL)
	}
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
