// curl www.duckduckgo.com | go run THIS_FILE
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int // indent actually

type PrettyPrinter struct {
	writer io.Writer
	err    error
}

func main() {
	htmlBody, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	prettyprinter := NewPrettyPrinter()
	prettyprinter.Pretty(os.Stdout, htmlBody)
}

func NewPrettyPrinter() PrettyPrinter {
	return PrettyPrinter{}
}

func (pp PrettyPrinter) Pretty(writer io.Writer, node *html.Node) error {
	pp.writer = writer
	pp.err = nil
	pp.forEachNode(node, pp.start, pp.end)

	return pp.Err()
}

func (pp PrettyPrinter) Err() error {
	return pp.err
}

func (pp PrettyPrinter) forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}
	if pp.Err() != nil {
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		pp.forEachNode(child, pre, post)
	}

	if post != nil {
		post(node)
	}
	if pp.Err() != nil {
		return
	}
}

func (pp PrettyPrinter) printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(pp.writer, format, args...)
	pp.err = err
}

func (pp PrettyPrinter) startElement(node *html.Node) {
	end := ">"
	if node.FirstChild == nil {
		end = "/>"
	}

	attrs := make([]string, 0, len(node.Attr))
	for _, a := range node.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}

	attrStr := ""
	if len(node.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	name := node.Data

	pp.printf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}

func (pp PrettyPrinter) endElement(node *html.Node) {
	depth--

	if node.FirstChild == nil {
		return
	}
	pp.printf("%*s</%s>\n", depth*2, "", node.Data)
}

func (pp PrettyPrinter) startText(node *html.Node) {
	text := strings.TrimSpace(node.Data)
	if len(text) == 0 {
		return
	}

	pp.printf("%*s%s\n", depth*2, "", node.Data)
}

func (pp PrettyPrinter) startComment(node *html.Node) {
	pp.printf("<!--%s-->\n", node.Data)
}

func (pp PrettyPrinter) start(node *html.Node) {
	switch node.Type {
	case html.ElementNode:
		pp.startElement(node)
	case html.TextNode:
		pp.startText(node)
	case html.CommentNode:
		pp.startComment(node)
	}
}

func (pp PrettyPrinter) end(node *html.Node) {
	switch node.Type {
	case html.ElementNode:
		pp.startElement(node)
	}
}
