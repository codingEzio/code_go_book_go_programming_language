// Attempts to contact the server of a URL; retries for several time; reports err if all failed
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	/*
		log.Fatalf("usage: wait url\n")
		|| fmt.Fprintf(os.Stderr, "usage: wait url\n")
		|| os.Exit(1)
	*/

	if len(os.Args) != 2 {
		log.Fatalf("usage: wait url\n")
	}

	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}

		log.SetPrefix("Wait: ")
		log.SetFlags(log.Lshortfile)
		log.Printf("server not responding (%s); retrying...", err)

		time.Sleep(time.Second << uint(tries)) // 1, 2, 4, 8 ... 1min at most
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
