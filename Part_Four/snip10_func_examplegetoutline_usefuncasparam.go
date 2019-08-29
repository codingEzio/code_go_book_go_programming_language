// Prints the outline (gracefully) (using func as params) (go run THIS URL_LIST)
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline2(url)
	}
}

func outline2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	htmlBody, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(htmlBody, startElement, endElement)

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

// It's actually means "indent" in the output.
// The magic code is `%*s`
// - it specifies how much space to allocate for the string
// - and, you can either write `%5s` (uh) or the `%*s .. 4` 👊
var depth int

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*4, "", node.Data)
		depth++
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*4, "", node.Data)
	}
}
