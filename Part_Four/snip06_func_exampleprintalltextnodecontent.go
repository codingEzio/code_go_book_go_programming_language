package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	err := printTagText(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func printTagText(reader io.Reader, writer io.Writer) error {
	tokenizer := html.NewTokenizer(os.Stdin)
	var err error
	stack := make([]string, 20)

Tokenize:
	for {
		switch tokenizer.Next() {

		case html.ErrorToken:
			break Tokenize

		case html.StartTagToken:
			b, _ := tokenizer.TagName()
			stack = append(stack, string(b))

		case html.TextToken:
			cur := stack[len(stack)-1]
			if cur == "script" || cur == "style" {
				continue
			}

			text := tokenizer.Text()
			if len(strings.TrimSpace(string(text))) == 0 {
				continue
			}

			_, _ = writer.Write([]byte(fmt.Sprintf("<%s>", cur)))
			_, _ = writer.Write(text)
			if text[len(text)-1] != '\n' {
				_, _ = io.WriteString(writer, "\n")
			}

		case html.EndTagToken:
			stack = stack[:len(stack)-1]
		}
	}

	if err != io.EOF {
		return err
	}

	return nil
}
