// Fetches an webpage and prints its title. Returns an error if it's {not HTML, multi-title}.
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
		How to test this code?
		- create two files, one has single title, the other one has at least two tiles
		- create a simple server: python -m http.server 9000 --bind 127.0.0.1
		- that's all
			>> go run THISCODE http://127.0.0.1:9000/notitle.html
			>> go run THISCODE http://127.0.0.1:9000/onetitle.html
			>> go run THISCODE http://127.0.0.1:9000/twotitles.html
	*/

	for _, arg := range os.Args[1:] {
		if err := title2(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

func title2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// check Content-Type is HTML
	contType := resp.Header.Get("Content-Type")
	if contType != "text/html" && !strings.HasPrefix(contType, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type [%s], not [text/html]", url, contType)
	}

	htmlBody, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	title, err := soloTitle(htmlBody)
	if err != nil {
		return err
	}
	fmt.Println(title)

	return nil
}

func soloTitle(htmlBody *html.Node) (title string, err error) {
	type bailout struct{}

	// don't worry, this statement won't be exec_ed immediately!
	defer func() {
		switch pan := recover(); pan {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(pan) // unexpected panic; carry on panicking
		}
	}()

	// bail out of recursion if we find >=1 non-empty title
	forEachNode2(htmlBody, func(node *html.Node) {
		if node.Type == html.ElementNode &&
			node.Data == "title" &&
			node.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = node.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element!")
	}

	return title, nil
}

func forEachNode2(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode2(child, pre, post)
	}

	if post != nil {
		post(node)
	}
}
