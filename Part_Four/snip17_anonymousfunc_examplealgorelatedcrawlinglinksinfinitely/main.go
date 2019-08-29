// Crawls the links on a webpage (& then links on found webpages, on and on..) (INFINITELY)
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/code_go_book_go_programming_language/Part_Four/snip17_anonymousfunc_examplealgorelatedcrawlinglinksinfinitely/libfindlinks"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)

	links, err := libfindlinks.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return links
}

func breadthFirst(fn func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				worklist = append(worklist, fn(item)...)
			}
		}
	}
}
