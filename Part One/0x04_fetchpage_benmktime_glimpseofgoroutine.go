// Fetch URLs in parallel and reports their times/sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	/*
		Concurrency & Parallelism
		* concur	do many things at the same time (1 core)(cont switch)
		* parallel	do many things using different cores (not alws fast)(commu)

		How much do you need to know to understand this code?
		* https://golangbot.com/goroutines/
		* https://golangbot.com/channels/
		* https://play.golang.org (run go code online)

		My own "single responsibility code file" principle
		* #TODO more details about 'channel' will be talked about later
		* #TODO more details about "weird" syntax will also be talk about later
	*/

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("\n%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// The `io.Copy` func reads the body of the resp and discards it
	// by writing to the `ioutil.Discard` output stream (args: dest, src)
	// also, the return value of `io.Copy` is exactly what we need in here.
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	// For this file we only need these {time-spent, bytes-size, url},
	// that's why we directly discarded those "bodies" and closed it 🙌
	_ = resp.Body.Close()

	// This one will only execute when there's network error occurred
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
