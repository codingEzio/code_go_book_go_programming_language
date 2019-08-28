// This package makes request by itself, while invoking others to get the links :)
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"

	"github.com/code_go_book_go_programming_language/Part_Four/snip07_func_examplereturnmultivalue/lib"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	// Request might fail
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Request final result might fail
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// Parse itself might fail
	htmlBody, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// Desired flow
	return libfindhrefs.Visit(nil, htmlBody), nil
}
