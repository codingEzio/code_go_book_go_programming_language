// Gets the outline of a webpage (basically the same as 'snip10_xx', but writing in a new way)
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline3(url)
	}
}

func outline3(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	htmlBody, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var depth int

	startElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*4, "", node.Data)
			depth++
		}
	}

	endElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*4, "", node.Data)
		}
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
