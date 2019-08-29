// Finds an HTML element node by id attribute (the output might be a bit jarring, don't worry..)
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: THISCODE HTML_FILE ID")
	}
	filename := os.Args[1]
	htmlID := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	htmlBody, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", ElementByID(htmlBody, htmlID))
}

func ElementByID(node *html.Node, id string) *html.Node {
	pre := func(node *html.Node) bool {
		if node.Type != html.ElementNode {
			return true
		}

		for _, a := range node.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}

		return true
	}

	return forEachElement(node, pre, nil)
}

func forEachElement(node *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	unvisited := make([]*html.Node, 0)
	unvisited = append(unvisited, node)

	for len(unvisited) > 0 {
		node = unvisited[0]
		unvisited = unvisited[1:]

		if pre != nil {
			if !pre(node) {
				return node
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			unvisited = append(unvisited, child)
		}
		if post != nil {
			if !post(node) {
				return node
			}
		}
	}

	return nil
}
