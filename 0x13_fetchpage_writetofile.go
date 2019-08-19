// Fetch pages and report their times, lastly, write to text files
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// From the test results, we came out of this conclusion 😅
	// the response time does reduced, yet the size doesn't change at all
	fetchAndWriteToFile(os.Args[1], os.Args[2])
}

func fetchAndWriteToFile(url string, filename string) {
	start := time.Now()
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)

	// write whole the body
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		panic(err)
	}

	_ = resp.Body.Close()

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs\t%s\n", secs, url)
}
