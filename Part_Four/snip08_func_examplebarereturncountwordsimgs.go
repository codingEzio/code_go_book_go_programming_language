// Does request by itself and returns the number of words and images in it :)
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: PROG URL")
	}

	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("words: %d\nimages: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	htmlBody, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
	}

	words, images = countWordsAndImages(htmlBody)
	return
}

func countWordsAndImages(node *html.Node) (words, images int) {
	unvisited := make([]*html.Node, 0)
	unvisited = append(unvisited, node)

	for len(unvisited) > 0 {
		node = unvisited[len(unvisited)-1]
		unvisited = unvisited[:len(unvisited)-1]

		switch node.Type {
		case html.TextNode:
			words += wordCount(node.Data)
		case html.ElementNode:
			if node.Data == "img" {
				images++
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			unvisited = append(unvisited, child)
		}
	}

	// They both have been assigned new values (`words` & `images`)
	// so there's no need to name it explicitly (compiler does it for u)
	return
}

func wordCount(str string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(str))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		n++
	}

	return n
}
