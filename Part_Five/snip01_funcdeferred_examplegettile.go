// Fetches an webpage and prints its title. Returns an error if it's not HTML.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	/*
		Two usage examples
		>> go run THISCODE https://google.com
		>> go run THISCODE https://www.bloggingbasics101.com/feed/
	*/

	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/html" && !strings.HasPrefix(contentType, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type $s, not text/html", url, contentType)
	}

	htmlBody, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode &&
			node.Data == "title" &&
			node.FirstChild != nil {
			fmt.Println(node.FirstChild.Data)
		}
	}

	forEachNode(htmlBody, visitNode, nil)
	return nil
}

func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(child, pre, post)
	}

	if post != nil {
		post(node)
	}
}
