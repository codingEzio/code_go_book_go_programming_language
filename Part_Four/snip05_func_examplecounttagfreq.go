// Counts the frequency of different tags in an HTML document (curl URL | go run THIS)
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	freq, err := tagFreq(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for tag, count := range freq {
		fmt.Printf("%4d %8s\n", count, tag)
	}
}

func tagFreq(reader io.Reader) (map[string]int, error) {
	freq := make(map[string]int, 0)
	tokenizer := html.NewTokenizer(os.Stdin)
	var err error

	for {
		type_ := tokenizer.Next()
		if type_ == html.ErrorToken {
			break
		}

		name, _ := tokenizer.TagName()
		if len(name) > 0 {
			freq[string(name)]++
		}
	}

	if err != io.EOF {
		return freq, err
	}

	return freq, nil
}
