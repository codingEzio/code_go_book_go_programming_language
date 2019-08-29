// Crawls the links on a webpage (INFINITELY) (stores the links locally, same-domain only)
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/code_go_book_go_programming_language/Part_Four/snip20_anonymousfunc_examplealgorelatedcrawlinglinksinfinitely_alsomakelocalcopies/findlinks"
)

var origHost string

func main() {
	/*
		Usage examples
		>> go run THISCODE https://golang.org
		>> go run THISCODE https://example.com

		Do remember that
		- only those whom belongs to the same domain will be stored
	*/
	breadthFirst(crawl, os.Args[1:])
}

func save(rawurl string) error {
	url, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}
	if origHost == "" {
		origHost = url.Host
	}
	if origHost != url.Host {
		return nil
	}

	directory := url.Host
	var filename string

	if filepath.Ext(filename) == "" {
		directory = filepath.Join(directory, url.Path)
		filename = filepath.Join(directory, "output.html")
	} else {
		directory = filepath.Join(directory, filepath.Dir(url.Path))
		filename = url.Path
	}

	// Make directory
	err = os.MkdirAll(directory, 0777)
	if err != nil {
		return err
	}

	// Make request
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Make file
	fileObj, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Make copy operation
	_, err = io.Copy(fileObj, resp.Body)
	if err != nil {
		return err
	}

	// Check for delayed write occurs
	err = fileObj.Close()
	if err != nil {
		return err
	}

	return nil
}

func crawl(url string) []string {
	fmt.Println(url)

	err := save(url)
	if err != nil {
		log.Printf(`can't cache "%s": $s`, url, err)
	}

	links, err := findlinks.Extract2(url)
	if err != nil {
		log.Printf(`can't extract links from "%s": %s`, url, err)
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
