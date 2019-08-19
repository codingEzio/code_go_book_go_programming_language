// Fetch page & print content (using `ioutil.ReadAll`)
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	prefix := "http://"

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") &&
			!strings.HasPrefix(url, "https://") {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		htmlBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		_ = resp.Body.Close()

		fmt.Printf("\n%s\n%s\n", htmlBody, resp.Status)
	}
}
