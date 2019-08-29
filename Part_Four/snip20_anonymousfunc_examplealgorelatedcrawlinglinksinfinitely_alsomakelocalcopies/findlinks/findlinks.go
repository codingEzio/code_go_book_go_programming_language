package findlinks

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the responses as HTML, and returns the links in the HTML document.
func Extract2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	htmlBody, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}

				links = append(links, link.String())
			}
		}
	}

	forEachNode(htmlBody, visitNode, nil)
	return links, nil
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
