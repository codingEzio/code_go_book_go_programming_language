// Fetches page and stores as local HTML files, returns its byte size if succeeded
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}

		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}

func fetch2(url string) (filename string, n int64, err error) {
	tmpFolder := "/tmp/"
	pageExt := ".html"

	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index"
	}

	// python.org		=> 	index	( /tmp/ ... .html )
	// python.org/doc	=>	doc		( /tmp/ ... .html )
	f, err := os.Create(tmpFolder + local + pageExt)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		// Close file, but prefer error from `Copy`, if any :)
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	// The file descriptor'll be closed eventually by `defer ...`
	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
