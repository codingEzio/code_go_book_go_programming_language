// Returns the elements which match one of those args being given (local HTML file needed)
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	/*
		Detail usage example
		>> wget -O /tmp/sample.html "https://example.com"
		>> go run snip24_variadicfunc_examplefindelementsbyID.go /tmp/sample.html h1
	*/

	if len(os.Args) < 3 {
		fmt.Println("usage: go run PROGRAM LOCAL_HTML_FILE TAG_NAME")
	}

	filename := os.Args[1]
	tags := os.Args[2:]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	htmlBody, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range ElementsByTag(htmlBody, tags...) {
		fmt.Printf("%+v\n", node)
	}
}

func ElementsByTag(node *html.Node, tags ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	keep := make(map[string]bool, len(tags))

	for _, tag := range tags {
		keep[tag] = true
	}

	pre := func(node *html.Node) bool {
		if node.Type != html.ElementNode {
			return true
		}
		_, ok := keep[node.Data]
		if ok {
			nodes = append(nodes, node)
		}
		return true
	}

	forEachElement(node, pre, nil)
	return nodes
}

func forEachElement(node *html.Node, pre, post func(node *html.Node) bool) *html.Node {
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
