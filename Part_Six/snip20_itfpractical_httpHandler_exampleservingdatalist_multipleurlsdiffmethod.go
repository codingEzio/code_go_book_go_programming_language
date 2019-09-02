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

	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	fmt.Println("Now serving at http://localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

// localhost:8000/listall
// localhost:8000/item
func (db database) list(wr http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(wr, "%s: %s\n", item, price)
	}
}

func (db database) price(wr http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	price, ok := db[item]
	if !ok {
		wr.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(wr, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(wr, "%s\n", price)
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
