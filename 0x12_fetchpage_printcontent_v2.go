// Fetch page & print content (using `io.Read` instead of `ioutil.ReadAll`)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	/*
		Reference links:
		- https://golang.org/pkg/io/#Copy
		- https://golang.org/pkg/io/ioutil/#ReadAll

		About the one-liner if statement (https://tour.golang.org/flowcontrol/6)
		* vars declared by the stmt are ONLY in scope until the end of `if`
		* in our case we're using `io.Copy` which is totally fine 😁
	*/

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Since we're directly copying the content to `Stdout`
		// there's NO NEED to print it (I do forget about it at first 🙄)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		_ = resp.Body.Close()

		fmt.Printf("\n%s", resp.Status)
	}
}
